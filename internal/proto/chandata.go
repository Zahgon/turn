// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package proto

import (
	"errors"
)

// ChannelData represents The ChannelData Message.
//
// See RFC 5766 Section 11.4.
type ChannelData struct {
	Data   []byte // Can be sub slice of Raw
	Length int    // Ignored while encoding, len(Data) is used
	Number ChannelNumber
	Raw    []byte
}

// Equal returns true if compareTo == c.
func (c *ChannelData) Equal(compareTo *ChannelData) bool { _ = "STUB: not implemented"; return false }

// Grow ensures that internal buffer will fit v more bytes and
// increases it capacity if necessary.
//
// Similar to stun.Message.grow method.
func (c *ChannelData) grow(v int) { _ = "STUB: not implemented"; return }

// Reset resets Length, Data and Raw length.
func (c *ChannelData) Reset() { _ = "STUB: not implemented"; return }

// Encode encodes ChannelData Message to Raw.
func (c *ChannelData) Encode() { _ = "STUB: not implemented"; return }

const padding = 4

func nearestPaddedValueLength(l int) int { _ = "STUB: not implemented"; return 0 }

// WriteHeader writes channel number and length.
func (c *ChannelData) WriteHeader() { _ = "STUB: not implemented"; return }

// Making WriteHeader call valid even when c.Raw
// is nil or len(c.Raw) is less than needed for header.

// Early bounds check to guarantee safety of writes below.

// nolint:gosec // G115

// ErrBadChannelDataLength means that channel data length is not equal
// to actual data length.
var ErrBadChannelDataLength = errors.New("channelData length != len(Data)")

// Decode decodes The ChannelData Message from Raw.
func (c *ChannelData) Decode() error { _ = "STUB: not implemented"; return nil }

const (
	channelDataLengthSize = 2
	channelDataNumberSize = channelDataLengthSize
	channelDataHeaderSize = channelDataLengthSize + channelDataNumberSize
)

// IsChannelData returns true if buf looks like the ChannelData Message.
func IsChannelData(buf []byte) bool { _ = "STUB: not implemented"; return false }

// Quick check for channel number.
