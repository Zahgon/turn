// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package allocation

import (
	"net"
	"time"

	"github.com/pion/logging"
)

const DefaultPermissionTimeout = time.Duration(5) * time.Minute

// Permission represents a TURN permission. TURN permissions mimic the address-restricted
// filtering mechanism of NATs that comply with [RFC4787].
// See: https://tools.ietf.org/html/rfc5766#section-2.3
type Permission struct {
	Addr          net.Addr
	allocation    *Allocation
	timeout       time.Duration
	lifetimeTimer *time.Timer
	log           logging.LeveledLogger
}

// NewPermission create a new Permission.
func NewPermission(addr net.Addr, log logging.LeveledLogger, timeout time.Duration) *Permission {
	_ = "STUB: not implemented"
	return nil
}

func (p *Permission) start(lifetime time.Duration) { _ = "STUB: not implemented"; return }

func (p *Permission) refresh(lifetime time.Duration) { _ = "STUB: not implemented"; return }
