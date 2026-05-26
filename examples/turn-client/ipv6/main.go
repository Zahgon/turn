// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

// Package main implements a TURN client using UDP with IPv6 support (RFC 6156)
package main

import (
	"flag"
	"log"
	"net"
	"strconv"
	"strings"

	"github.com/pion/logging"
	"github.com/pion/turn/v5"
)

func main() { //nolint:cyclop
	host := flag.String("host", "", "TURN Server IPv6 address.")
	port := flag.Int("port", 3478, "Listening port.")
	user := flag.String("user", "", "A pair of username and password (e.g. \"user=pass\")")
	realm := flag.String("realm", "pion.ly", "Realm (defaults to \"pion.ly\")")
	ping := flag.Bool("ping", false, "Run ping test")
	flag.Parse()

	if len(*host) == 0 {
		log.Fatalf("'host' is required")
	}

	if len(*user) == 0 {
		log.Fatalf("'user' is required")
	}

	cred := strings.SplitN(*user, "=", 2)

	// TURN client won't create a local listening socket by itself.
	// Using udp6 for IPv6 support
	conn, err := net.ListenPacket("udp6", "[::]:0") // nolint: noctx
	if err != nil {
		log.Panicf("Failed to listen: %s", err)
	}
	defer func() {
		if closeErr := conn.Close(); closeErr != nil {
			log.Panicf("Failed to close connection: %s", closeErr)
		}
	}()

	// Format server address (net.JoinHostPort handles IPv6 brackets automatically)
	turnServerAddr := net.JoinHostPort(*host, strconv.Itoa(*port))

	cfg := &turn.ClientConfig{
		STUNServerAddr: turnServerAddr,
		TURNServerAddr: turnServerAddr,
		Conn:           conn,
		Username:       cred[0],
		Password:       cred[1],
		Realm:          *realm,
		LoggerFactory:  logging.NewDefaultLoggerFactory(),
	}

	client, err := turn.NewClient(cfg)
	if err != nil {
		log.Panicf("Failed to create TURN client: %s", err)
	}
	defer client.Close()

	// Start listening on the conn provided.
	err = client.Listen()
	if err != nil {
		log.Panicf("Failed to listen: %s", err)
	}

	// Allocate a relay socket on the TURN server. On success, it
	// will return a net.PacketConn which represents the remote
	// socket. With IPv6, this will allocate an IPv6 relay address.
	relayConn, err := client.Allocate()
	if err != nil {
		log.Panicf("Failed to allocate: %s", err)
	}
	defer func() {
		if closeErr := relayConn.Close(); closeErr != nil {
			log.Panicf("Failed to close connection: %s", closeErr)
		}
	}()

	// The relayConn's local address is actually the transport
	// address assigned on the TURN server (IPv6 address).
	log.Printf("relayed-address=%s (IPv6)", relayConn.LocalAddr().String())

	// If you provided `-ping`, perform a ping test against the
	// relayConn we have just allocated.
	if *ping {
		err = doPingTest(client, relayConn)
		if err != nil {
			log.Panicf("Failed to ping: %s", err)
		}
	}
}

func doPingTest(client *turn.Client, relayConn net.PacketConn) error {
	_ = "STUB: not implemented" //nolint:cyclop
	// Send BindingRequest to learn our external IPv6 address
	return nil
}

// Set up pinger socket (pingerConn)
// nolint: noctx

// Punch a UDP hole for the relayConn by sending a data to the mappedAddr.
// This will trigger a TURN client to generate a permission request to the
// TURN server. After this, packets from the IPv6 address will be accepted by
// the TURN server.

// Start read-loop on pingerConn

// Start read-loop on relayConn

// Echo back

// Send 10 packets from relayConn to the echo server

// For simplicity, this example does not wait for the pong (reply).
// Instead, sleep 1 second.
