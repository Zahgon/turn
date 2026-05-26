// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package client

import (
	"sync"
	"time"
)

// PeriodicTimerTimeoutHandler is a handler called on timeout.
type PeriodicTimerTimeoutHandler func(timerID int)

// PeriodicTimer is a periodic timer.
type PeriodicTimer struct {
	id             int
	interval       time.Duration
	timeoutHandler PeriodicTimerTimeoutHandler
	stopFunc       func()
	mutex          sync.RWMutex
}

// NewPeriodicTimer create a new timer.
func NewPeriodicTimer(id int, timeoutHandler PeriodicTimerTimeoutHandler, interval time.Duration) *PeriodicTimer {
	_ = "STUB: not implemented"
	return nil
}

// Start starts the timer.
func (t *PeriodicTimer) Start() bool { _ = "STUB: not implemented"; return false }

// This is a noop if the timer is always running

// Stop stops the timer.
func (t *PeriodicTimer) Stop() { _ = "STUB: not implemented"; return }

// IsRunning tests if the timer is running.
// Debug purpose only.
func (t *PeriodicTimer) IsRunning() bool { _ = "STUB: not implemented"; return false }
