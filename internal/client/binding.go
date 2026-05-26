// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package client

import (
	"net"
	"sync"
	"time"
)

// Channel number:
//
//	0x4000 through 0x7FFF: These values are the allowed channel
//	numbers (16,383 possible values).
const (
	minChannelNumber uint16 = 0x4000
	maxChannelNumber uint16 = 0x7fff
)

type bindingState int32

const (
	bindingStateIdle bindingState = iota
	bindingStateRequest
	bindingStateReady
	bindingStateRefresh
	bindingStateFailed
)

type binding struct {
	number       uint16          // Read-only
	st           bindingState    // Thread-safe (atomic op)
	addr         net.Addr        // Read-only
	mgr          *bindingManager // Read-only
	muBind       sync.Mutex      // Thread-safe, for ChannelBind ops
	_refreshedAt time.Time       // Protected by mutex
	mutex        sync.RWMutex    // Thread-safe
}

func (b *binding) setState(state bindingState) { _ = "STUB: not implemented"; return }

func (b *binding) state() bindingState { _ = "STUB: not implemented"; return *new(bindingState) }

func (b *binding) setRefreshedAt(at time.Time) { _ = "STUB: not implemented"; return }

func (b *binding) refreshedAt() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (b *binding) ok() bool { _ = "STUB: not implemented"; return false }

// Thread-safe binding map.
type bindingManager struct {
	chanMap map[uint16]*binding
	addrMap map[string]*binding
	next    uint16
	mutex   sync.RWMutex
}

func newBindingManager() *bindingManager { _ = "STUB: not implemented"; return nil }

func (mgr *bindingManager) assignChannelNumber() uint16 { _ = "STUB: not implemented"; return 0 }

func (mgr *bindingManager) create(addr net.Addr) *binding { _ = "STUB: not implemented"; return nil }

func (mgr *bindingManager) findByAddr(addr net.Addr) (*binding, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (mgr *bindingManager) findByNumber(number uint16) (*binding, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (mgr *bindingManager) deleteByAddr(addr net.Addr) bool {
	_ = "STUB: not implemented"
	return false
}

func (mgr *bindingManager) deleteByNumber(number uint16) bool {
	_ = "STUB: not implemented"
	return false
}

func (mgr *bindingManager) size() int { _ = "STUB: not implemented"; return 0 }

func (mgr *bindingManager) all() []*binding { _ = "STUB: not implemented"; return nil }
