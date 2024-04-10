package blockless

import (
	"github.com/Bruce960205/b7s/models/execute"
)

type Executor interface {
	ExecuteFunction(requestID string, request execute.Request) (execute.Result, error)
}
