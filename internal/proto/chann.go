// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package proto

import (
	"errors"

	"github.com/pion/stun/v3"
)

// ChannelNumber represents CHANNEL-NUMBER attribute.
//
// The CHANNEL-NUMBER attribute contains the number of the channel.
//
// RFC 5766 Section 14.1.
type ChannelNumber uint16 // Encoded as uint16

func (n ChannelNumber) String() string { _ = "STUB: not implemented"; return "" }

// 16 bits of uint + 16 bits of RFFU = 0.
const channelNumberSize = 4

// AddTo adds CHANNEL-NUMBER to message.
func (n ChannelNumber) AddTo(m *stun.Message) error { _ = "STUB: not implemented"; return nil }

// v[2:4] are zeroes (RFFU = 0)

// GetFrom decodes CHANNEL-NUMBER from message.
func (n *ChannelNumber) GetFrom(m *stun.Message) error { _ = "STUB: not implemented"; return nil }

// Asserting length

// v[2:4] is RFFU and equals to 0.

// See https://tools.ietf.org/html/rfc5766#section-11:
//
// 0x4000 through 0x7FFF: These values are the allowed channel
// numbers (16,383 possible values).
const (
	MinChannelNumber = 0x4000
	MaxChannelNumber = 0x7FFF
)

// ErrInvalidChannelNumber means that channel number is not valid as by RFC 5766 Section 11.
var ErrInvalidChannelNumber = errors.New("channel number not in [0x4000, 0x7FFF]")

// isChannelNumberValid returns true if c in [0x4000, 0x7FFF].
func isChannelNumberValid(c uint16) bool { _ = "STUB: not implemented"; return false }

// Valid returns true if channel number has correct value that complies RFC 5766 Section 11 range.
func (n ChannelNumber) Valid() bool { _ = "STUB: not implemented"; return false }
