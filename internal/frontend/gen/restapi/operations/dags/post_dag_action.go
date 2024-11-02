// Code generated by go-swagger; DO NOT EDIT.

package dags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostDagActionHandlerFunc turns a function with the right signature into a post dag action handler
type PostDagActionHandlerFunc func(PostDagActionParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostDagActionHandlerFunc) Handle(params PostDagActionParams) middleware.Responder {
	return fn(params)
}

// PostDagActionHandler interface for that can handle valid post dag action params
type PostDagActionHandler interface {
	Handle(PostDagActionParams) middleware.Responder
}

// NewPostDagAction creates a new http.Handler for the post dag action operation
func NewPostDagAction(ctx *middleware.Context, handler PostDagActionHandler) *PostDagAction {
	return &PostDagAction{Context: ctx, Handler: handler}
}

/*
	PostDagAction swagger:route POST /dags/{dagId} dags postDagAction

Performs an action on a DAG.
*/
type PostDagAction struct {
	Context *middleware.Context
	Handler PostDagActionHandler
}

func (o *PostDagAction) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostDagActionParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostDagActionBody post dag action body
//
// swagger:model PostDagActionBody
type PostDagActionBody struct {

	// action
	// Required: true
	// Enum: ["start","suspend","stop","retry","mark-success","mark-failed","save","rename"]
	Action *string `json:"action"`

	// params
	Params string `json:"params,omitempty"`

	// request Id
	RequestID string `json:"requestId,omitempty"`

	// step
	Step string `json:"step,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this post dag action body
func (o *PostDagActionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var postDagActionBodyTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["start","suspend","stop","retry","mark-success","mark-failed","save","rename"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postDagActionBodyTypeActionPropEnum = append(postDagActionBodyTypeActionPropEnum, v)
	}
}

const (

	// PostDagActionBodyActionStart captures enum value "start"
	PostDagActionBodyActionStart string = "start"

	// PostDagActionBodyActionSuspend captures enum value "suspend"
	PostDagActionBodyActionSuspend string = "suspend"

	// PostDagActionBodyActionStop captures enum value "stop"
	PostDagActionBodyActionStop string = "stop"

	// PostDagActionBodyActionRetry captures enum value "retry"
	PostDagActionBodyActionRetry string = "retry"

	// PostDagActionBodyActionMarkDashSuccess captures enum value "mark-success"
	PostDagActionBodyActionMarkDashSuccess string = "mark-success"

	// PostDagActionBodyActionMarkDashFailed captures enum value "mark-failed"
	PostDagActionBodyActionMarkDashFailed string = "mark-failed"

	// PostDagActionBodyActionSave captures enum value "save"
	PostDagActionBodyActionSave string = "save"

	// PostDagActionBodyActionRename captures enum value "rename"
	PostDagActionBodyActionRename string = "rename"
)

// prop value enum
func (o *PostDagActionBody) validateActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, postDagActionBodyTypeActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *PostDagActionBody) validateAction(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"action", "body", o.Action); err != nil {
		return err
	}

	// value enum
	if err := o.validateActionEnum("body"+"."+"action", "body", *o.Action); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post dag action body based on context it is used
func (o *PostDagActionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostDagActionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostDagActionBody) UnmarshalBinary(b []byte) error {
	var res PostDagActionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
