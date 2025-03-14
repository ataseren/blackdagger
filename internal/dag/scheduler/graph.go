package scheduler

import (
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/ErdemOzgen/blackdagger/internal/dag"
	"github.com/ErdemOzgen/blackdagger/internal/logger"
)

// ExecutionGraph represents a graph of steps.
type ExecutionGraph struct {
	startedAt       time.Time
	finishedAt      time.Time
	outputVariables *dag.SyncMap
	dict            map[int]*Node
	nodes           []*Node
	from            map[int][]int
	to              map[int][]int
	mu              sync.RWMutex
	logger          logger.Logger
}

// NewExecutionGraph creates a new execution graph with the given steps.
func NewExecutionGraph(lg logger.Logger, steps ...dag.Step) (*ExecutionGraph, error) {
	graph := &ExecutionGraph{
		outputVariables: &dag.SyncMap{},
		dict:            make(map[int]*Node),
		from:            make(map[int][]int),
		to:              make(map[int][]int),
		nodes:           []*Node{},
		logger:          lg,
	}
	if graph.logger == nil {
		graph.logger = logger.Default
	}
	for _, step := range steps {
		step.OutputVariables = graph.outputVariables
		node := &Node{data: NodeData{Step: step}}
		node.init()
		graph.dict[node.id] = node
		graph.nodes = append(graph.nodes, node)
	}
	if err := graph.setup(); err != nil {
		return nil, err
	}
	return graph, nil
}

// NewExecutionGraphForRetry creates a new execution graph for retry with
// given nodes.
func NewExecutionGraphForRetry(lg logger.Logger, nodes ...*Node) (*ExecutionGraph, error) {
	graph := &ExecutionGraph{
		outputVariables: &dag.SyncMap{},
		dict:            make(map[int]*Node),
		from:            make(map[int][]int),
		to:              make(map[int][]int),
		nodes:           []*Node{},
		logger:          lg,
	}
	for _, node := range nodes {
		if node.data.Step.OutputVariables != nil {
			node.data.Step.OutputVariables.Range(func(key, value any) bool {
				k, ok := key.(string)
				if !ok {
					return false
				}
				v, ok := value.(string)
				if !ok {
					return false
				}

				graph.outputVariables.Store(key, value)
				err := os.Setenv(k, v[len(key.(string))+1:])
				if err != nil {
					graph.logger.Error("Failed to set env", "error", err)
				}
				return true
			})
		}
		node.data.Step.OutputVariables = graph.outputVariables
		node.init()
		graph.dict[node.id] = node
		graph.nodes = append(graph.nodes, node)
	}
	if err := graph.setup(); err != nil {
		return nil, err
	}
	if err := graph.setupRetry(); err != nil {
		return nil, err
	}
	return graph, nil
}

// Duration returns the duration of the execution.
func (g *ExecutionGraph) Duration() time.Duration {
	g.mu.RLock()
	defer g.mu.RUnlock()
	if g.finishedAt.IsZero() {
		return time.Since(g.startedAt)
	}
	return g.finishedAt.Sub(g.startedAt)
}

func (g *ExecutionGraph) IsStarted() bool {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return !g.startedAt.IsZero()
}

func (g *ExecutionGraph) IsFinished() bool {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return !g.finishedAt.IsZero()
}

func (g *ExecutionGraph) StartAt() time.Time {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return g.startedAt
}

func (g *ExecutionGraph) IsRunning() bool {
	g.mu.RLock()
	defer g.mu.RUnlock()
	for _, node := range g.Nodes() {
		if node.State().Status == NodeStatusRunning {
			return true
		}
	}
	return false
}

func (g *ExecutionGraph) FinishAt() time.Time {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return g.finishedAt
}

func (g *ExecutionGraph) Finish() {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.finishedAt = time.Now()
}

func (g *ExecutionGraph) Start() {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.startedAt = time.Now()
}

// Nodes returns the nodes of the execution graph.
func (g *ExecutionGraph) Nodes() []*Node {
	return g.nodes
}

func (g *ExecutionGraph) NodeData() []NodeData {
	g.mu.Lock()
	defer g.mu.Unlock()

	var ret []NodeData
	for _, node := range g.nodes {
		node.mu.Lock()
		ret = append(ret, node.data)
		node.mu.Unlock()
	}

	return ret
}

func (g *ExecutionGraph) node(id int) *Node {
	return g.dict[id]
}

func (g *ExecutionGraph) setupRetry() error {
	dict := map[int]NodeStatus{}
	retry := map[int]bool{}
	for _, node := range g.nodes {
		dict[node.id] = node.data.State.Status
		retry[node.id] = false
	}
	var frontier []int
	for _, node := range g.nodes {
		if len(node.data.Step.Depends) == 0 {
			frontier = append(frontier, node.id)
		}
	}
	for len(frontier) > 0 {
		var next []int
		for _, u := range frontier {
			if retry[u] || dict[u] == NodeStatusError ||
				dict[u] == NodeStatusCancel {
				g.logger.Info("clear node state", "step", g.dict[u].data.Step.Name)
				g.dict[u].clearState()
				retry[u] = true
			}
			for _, v := range g.from[u] {
				if retry[u] {
					retry[v] = true
				}
				next = append(next, v)
			}
		}
		frontier = next
	}
	return nil
}

func (g *ExecutionGraph) setup() error {
	for _, node := range g.nodes {
		for _, dep := range node.data.Step.Depends {
			depStep, err := g.findStep(dep)
			if err != nil {
				return err
			}
			g.addEdge(depStep, node)
		}
	}

	if g.hasCycle() {
		return errCycleDetected
	}

	return nil
}

func (g *ExecutionGraph) hasCycle() bool {
	var inDegrees = make(map[int]int)
	for node, depends := range g.to {
		inDegrees[node] = len(depends)
	}

	var q []int
	for _, node := range g.nodes {
		if inDegrees[node.id] != 0 {
			continue
		}
		q = append(q, node.id)
	}

	for len(q) > 0 {
		var f = q[0]
		q = q[1:]

		var tos = g.from[f]
		for _, to := range tos {
			inDegrees[to]--
			if inDegrees[to] == 0 {
				q = append(q, to)
			}
		}
	}

	for _, degree := range inDegrees {
		if degree > 0 {
			return true
		}
	}

	return false
}

func (g *ExecutionGraph) addEdge(from, to *Node) {
	g.from[from.id] = append(g.from[from.id], to.id)
	g.to[to.id] = append(g.to[to.id], from.id)
}

func (g *ExecutionGraph) findStep(name string) (*Node, error) {
	for _, n := range g.dict {
		if n.data.Step.Name == name {
			return n, nil
		}
	}
	return nil, fmt.Errorf("%w: %s", errStepNotFound, name)
}

var (
	errCycleDetected = errors.New("cycle detected")
	errStepNotFound  = errors.New("step not found")
)
