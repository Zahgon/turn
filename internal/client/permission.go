// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package client

import (
	"net"
	"sync"
)

type permState int32

const (
	permStateIdle permState = iota
	permStatePermitted
)

type permission struct {
	addr  net.Addr
	st    permState    // Thread-safe (atomic op)
	mutex sync.RWMutex // Thread-safe
}

func (p *permission) setState(state permState) { _ = "STUB: not implemented"; return }

func (p *permission) state() permState { _ = "STUB: not implemented"; return *new(permState) }

// Thread-safe permission map.
type permissionMap struct {
	permMap map[string]*permission
	mutex   sync.RWMutex
}

func (m *permissionMap) insert(addr net.Addr, p *permission) bool {
	_ = "STUB: not implemented"
	return false
}

func (m *permissionMap) find(addr net.Addr) (*permission, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (m *permissionMap) delete(addr net.Addr) { _ = "STUB: not implemented"; return }

func (m *permissionMap) addrs() []net.Addr { _ = "STUB: not implemented"; return nil }

func newPermissionMap() *permissionMap { _ = "STUB: not implemented"; return nil }
