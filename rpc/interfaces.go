package rpc

import (
	"github.com/cppforlife/bosh-cpi-go/apiv1"
)

//go:generate counterfeiter -o rpcfakes/fake_action_factory.go . ActionFactory

type ActionFactory interface {
	Create(string, apiv1.CallContext, apiv1.ApiVersions) (interface{}, error)
}

type Dispatcher interface {
	// Dispatch interprets request bytes, executes request,
	// captures response and return response bytes.
	// It panics if built-in errors fail to serialize.
	Dispatch([]byte) []byte
}

type Caller interface {
	Call(interface{}, []interface{}) (interface{}, error)
}

type CloudError interface {
	Error() string
	Type() string
}

type RetryableError interface {
	Error() string
	CanRetry() bool
}
