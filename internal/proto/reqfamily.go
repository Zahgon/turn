// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package proto

import (
	"errors"

	"github.com/pion/stun/v3"
)

// RequestedAddressFamily represents the REQUESTED-ADDRESS-FAMILY Attribute as
// defined in RFC 6156 Section 4.1.1.
type RequestedAddressFamily byte

const requestedFamilySize = 4

var errInvalidRequestedFamilyValue = errors.New("invalid value for requested family attribute")

// GetFrom decodes REQUESTED-ADDRESS-FAMILY from message.
func (f *RequestedAddressFamily) GetFrom(m *stun.Message) error {
	_ = "STUB: not implemented"
	return nil
}

func (f RequestedAddressFamily) String() string { _ = "STUB: not implemented"; return "" }

// AddTo adds REQUESTED-ADDRESS-FAMILY to message.
func (f RequestedAddressFamily) AddTo(m *stun.Message) error { _ = "STUB: not implemented"; return nil }

// b[1:4] is RFFU = 0.
// The RFFU field MUST be set to zero on transmission and MUST be
// ignored on reception. It is reserved for future uses.

// Values for RequestedAddressFamily as defined in RFC 6156 Section 4.1.1.
const (
	RequestedFamilyIPv4 RequestedAddressFamily = 0x01
	RequestedFamilyIPv6 RequestedAddressFamily = 0x02
)
