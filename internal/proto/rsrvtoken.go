// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package proto

import "github.com/pion/stun/v3"

// ReservationToken represents RESERVATION-TOKEN attribute.
//
// The RESERVATION-TOKEN attribute contains a token that uniquely
// identifies a relayed transport address being held in reserve by the
// server. The server includes this attribute in a success response to
// tell the client about the token, and the client includes this
// attribute in a subsequent Allocate request to request the server use
// that relayed transport address for the allocation.
//
// RFC 5766 Section 14.9.
type ReservationToken []byte

const reservationTokenSize = 8 // 8 bytes

// AddTo adds RESERVATION-TOKEN to message.
func (t ReservationToken) AddTo(m *stun.Message) error { _ = "STUB: not implemented"; return nil }

// GetFrom decodes RESERVATION-TOKEN from message.
func (t *ReservationToken) GetFrom(m *stun.Message) error { _ = "STUB: not implemented"; return nil }
