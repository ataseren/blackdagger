package cmd

import (
	"os"
	"testing"
	"time"
)

func TestSchedulerCommand(t *testing.T) {
	tmpDir, _, _ := setupTest(t)
	defer func() {
		_ = os.RemoveAll(tmpDir)
	}()

	go func() {
		testRunCommand(t, schedulerCmd(), cmdTest{
			args:        []string{"scheduler"},
			expectedOut: []string{"starting blackdagger scheduler"},
		})
	}()

	time.Sleep(time.Millisecond * 500)
}
