// Package util implements some utility functions.
package util

import (
	"fmt"
	"runtime/debug"

	"github.com/sainnhe/go-common/pkg/log"
)

// Recover allow the program to recover from panic and print logs using [log.GetDefault].
//
// NOTE: It should be used via defer, otherwise panics can't be captured.
func Recover() {
	if err := recover(); err != nil {
		// We must use [fmt.Sprintf] here otherwise [debug.Stack] will be printed in a single line.
		log.GetGlobalLogger().Error(
			fmt.Sprintf("Recovered from panic: %+v\n%s", err, string(debug.Stack())),
		)
	}
}

// ToStr converts a variable to a string.
// It's basically a convenient wrapper of [fmt.Sprintf] that uses %+v as placeholder.
func ToStr(v any) string {
	return fmt.Sprintf("%+v", v)
}
