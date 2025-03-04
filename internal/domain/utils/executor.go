package utils

import "context"

// Executor defines an interface for executing commands.
// It provides a common abstraction for command execution across different environments.
//
// The Execute method runs a command with given arguments and returns the output as a string,
// along with any error that occurred during execution.
// It accepts a context.Context to allow for timeout control, cancellation, and other request-scoped values.
type Executor interface {
	Execute(ctx context.Context, cmd string, args ...string) (string, error)
}
