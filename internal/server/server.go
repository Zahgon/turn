// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

// Package server implements the private API to implement a TURN server
package server

import (
	"crypto/tls"
	"net"
	"time"

	"github.com/pion/logging"
	"github.com/pion/stun/v3"
	"github.com/pion/turn/v5/internal/allocation"
	"github.com/pion/turn/v5/internal/auth"
)

// Request contains all the state needed to process a single incoming datagram.
type Request struct {
	// Current Request State
	Conn    net.PacketConn
	SrcAddr net.Addr
	Buff    []byte
	TLS     *tls.ConnectionState

	// Server State
	AllocationManager *allocation.Manager
	NonceHash         NonceManager

	// User Configuration
	AuthHandler auth.AuthHandler

	// Quota Handler
	QuotaHandler func(username string, realm string, srcAddr net.Addr) (ok bool)

	Log   logging.LeveledLogger
	Realm string

	ChannelBindTimeout time.Duration
	PermissionTimeout  time.Duration
	AllocationLifetime time.Duration
}

// HandleRequest processes the give Request.
func HandleRequest(r Request) error { _ = "STUB: not implemented"; return nil }

func handleDataPacket(req Request) error { _ = "STUB: not implemented"; return nil }

//nolint:errorlint

//nolint:errorlint

func handleTURNPacket(req Request) error { _ = "STUB: not implemented"; return nil }

// nolint:errorlint

// nolint:errorlint

// nolint:errorlint

func getMessageHandler(class stun.MessageClass, method stun.Method) ( // nolint:cyclop
	func(req Request, stunMsg *stun.Message) error,
	error,
) {
	_ = "STUB: not implemented"
	return nil, nil
}
