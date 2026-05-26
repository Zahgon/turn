// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package allocation

import (
	"net"
	"time"

	"github.com/pion/logging"
	"github.com/pion/turn/v5/internal/proto"
)

// ChannelBind represents a TURN Channel
// See: https://tools.ietf.org/html/rfc5766#section-2.5
type ChannelBind struct {
	Peer   net.Addr
	Number proto.ChannelNumber

	allocation    *Allocation
	lifetimeTimer *time.Timer
	log           logging.LeveledLogger
}

// NewChannelBind creates a new ChannelBind.
func NewChannelBind(number proto.ChannelNumber, peer net.Addr, log logging.LeveledLogger) *ChannelBind {
	_ = "STUB: not implemented"
	return nil
}

func (c *ChannelBind) start(lifetime time.Duration) { _ = "STUB: not implemented"; return }

func (c *ChannelBind) refresh(lifetime time.Duration) { _ = "STUB: not implemented"; return }
