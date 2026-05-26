// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package client

import (
	"net"
	"sync"
	"time"

	"github.com/pion/logging"
	"github.com/pion/stun/v3"
	"github.com/pion/transport/v4"
)

// AllocationConfig is a set of configuration params use by NewUDPConn and NewTCPAllocation.
type AllocationConfig struct {
	Client                    Client
	RelayedAddr               net.Addr
	ServerAddr                net.Addr
	Integrity                 stun.MessageIntegrity
	Nonce                     stun.Nonce
	Username                  stun.Username
	Realm                     stun.Realm
	Lifetime                  time.Duration
	Net                       transport.Net
	Log                       logging.LeveledLogger
	PermissionRefreshInterval time.Duration
	BindingRefreshInterval    time.Duration
	BindingCheckInterval      time.Duration
}

type allocation struct {
	client            Client                // Read-only
	relayedAddr       net.Addr              // Read-only
	serverAddr        net.Addr              // Read-only
	permMap           *permissionMap        // Thread-safe
	integrity         stun.MessageIntegrity // Read-only
	username          stun.Username         // Read-only
	realm             stun.Realm            // Read-only
	_nonce            stun.Nonce            // Needs mutex x
	_lifetime         time.Duration         // Needs mutex x
	net               transport.Net         // Thread-safe
	refreshAllocTimer *PeriodicTimer        // Thread-safe
	refreshPermsTimer *PeriodicTimer        // Thread-safe
	readTimer         *time.Timer           // Thread-safe
	mutex             sync.RWMutex          // Thread-safe
	log               logging.LeveledLogger // Read-only
}

func (a *allocation) setNonceFromMsg(msg *stun.Message) {
	_ = "STUB: not implemented"
	// Update nonce
	return
}

func (a *allocation) refreshAllocation(lifetime time.Duration, dontWait bool) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:err113

// Getting lifetime from response

func (a *allocation) refreshPermissions() error { _ = "STUB: not implemented"; return nil }

func (a *allocation) onRefreshTimers(id int) { _ = "STUB: not implemented"; return }

// Limit the max retries on errTryAgain to 3
// when stale nonce returns, sencond retry should succeed

func (a *allocation) nonce() stun.Nonce { _ = "STUB: not implemented"; return *new(stun.Nonce) }

func (a *allocation) setNonce(nonce stun.Nonce) { _ = "STUB: not implemented"; return }

func (a *allocation) lifetime() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (a *allocation) setLifetime(lifetime time.Duration) { _ = "STUB: not implemented"; return }
