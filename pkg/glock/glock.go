// -------------------------------------------------------------------------------------------
// Copyright (c) Team Sorghum. All rights reserved.
// Licensed under the GPL v3 License. See LICENSE in the project root for license information.
// -------------------------------------------------------------------------------------------

// Package glock implements goroutine lock.
package glock

import (
	"github.com/teamsorghum/go-common/pkg/concurrent"
	"github.com/teamsorghum/go-common/pkg/log"
)

// wg is the wait group used to implement goroutine lock.
var wg = &concurrent.WaitGroup{
	Name:   "glock",
	Logger: log.Global(),
}

// Lock locks goroutine to ensure that the task won't be interrupted.
func Lock() {
	wg.Add(1)
}

// Unlock unlocks a goroutine lock.
//
// NOTE: This function must be used via defer to avoid panic in the middle and causing the lock to not be released.
func Unlock() {
	wg.Done()
}

// Wait waits for all goroutine locks to be released.
func Wait() {
	if count := wg.GetCount(); count > 0 {
		wg.Logger.Info("[glock] Waiting for goroutine locks to be released...", "count", count)
		wg.Wait()
	}
}
