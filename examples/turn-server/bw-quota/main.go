// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

// Package main implements a TURN server with per-user bandwidth quotas.
//
// This example demonstrates how to implement bandwidth rate limiting on a per-user basis
// by wrapping the relay PacketConn with a rate-limited connection. Each user (identified
// by username+realm) gets their own rate limiter that caps their total bandwidth usage.
package main

import (
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"

	"github.com/pion/turn/v5"
	"golang.org/x/time/rate"
)

// userRateLimiters manages per-user rate limiters.
// The key is "userid:realm" to support different quotas per realm.
//
// Note: Inactive rate limiters should be periodically checked and removed, this is not implemented
// here for simplifity.
type userRateLimiters struct {
	limiters sync.Map // map[string]*rate.Limiter
	limit    rate.Limit
	burst    int
}

func newUserRateLimiters(bytesPerSec int) *userRateLimiters { _ = "STUB: not implemented"; return nil }

// Allow burst up to 1 second worth of data.

func (u *userRateLimiters) getLimiter(userID, realm string) *rate.Limiter {
	_ = "STUB: not implemented"
	return nil
}

//nolint:forcetypeassert

// Create new limiter for this user.

//nolint:forcetypeassert

// rateLimitedConn wraps a net.PacketConn with rate limiting.
// Both ReadFrom and WriteTo consume tokens from the same limiter, implementing
// a shared bandwidth cap for all traffic (upload + download combined).
type rateLimitedConn struct {
	net.PacketConn
	limiter *rate.Limiter
}

func (c *rateLimitedConn) ReadFrom(p []byte) (n int, addr net.Addr, err error) {
	_ = "STUB: not implemented"
	return 0, *new(net.Addr), nil
}

func (c *rateLimitedConn) WriteTo(p []byte, addr net.Addr) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Silently drop packet

// bwQuotaGenerator wraps a RelayAddressGenerator to add bandwidth quota enforcement.
type bwQuotaGenerator struct {
	turn.RelayAddressGenerator
	rateLimiters *userRateLimiters
}

func (g *bwQuotaGenerator) AllocatePacketConn(conf turn.AllocateListenerConfig) (net.PacketConn, net.Addr, error) {
	_ = "STUB: not implemented"
	return *new(net.PacketConn), *new(net.Addr), nil
}

func main() { //nolint:cyclop
	publicIP := flag.String("public-ip", "", "IP Address that TURN can be contacted by.")
	port := flag.Int("port", 3478, "Listening port.")
	users := flag.String("users", "", "List of username and password (e.g. \"user=pass,user=pass\")")
	realm := flag.String("realm", "pion.ly", "Realm (defaults to \"pion.ly\")")
	bwLimit := flag.Int("bw-limit", 100*1024/8,
		"Bandwidth limit per user in bytes/sec (default: 100 Kbps = 12500 bytes/sec)")
	testMode := flag.Bool("test", false, "Start a TURN client and UDP echo server for testing")
	testPort := flag.String("test-port", "5000", "UDP port for the TURN client proxy")
	peerAddr := flag.String("peer-addr", "127.0.0.1:5001", "Peer address for testing")
	flag.Parse()

	if len(*publicIP) == 0 {
		log.Fatalf("'public-ip' is required")
	} else if len(*users) == 0 {
		log.Fatalf("'users' is required")
	}

	rateLimiters := newUserRateLimiters(*bwLimit)
	log.Printf("Bandwidth limit per user: %d bytes/sec (%.1f Kbps)", *bwLimit, float64(*bwLimit)*8/1024)

	udpListener, err := net.ListenPacket("udp4", net.JoinHostPort("0.0.0.0", strconv.Itoa(*port))) //nolint:noctx
	if err != nil {
		log.Panicf("Failed to create TURN server listener: %s", err)
	}

	// Parse users.
	usersMap := map[string][]byte{}
	for userPass := range strings.SplitSeq(*users, ",") {
		parts := strings.SplitN(userPass, "=", 2)
		if len(parts) != 2 {
			log.Fatalf("Invalid user credential format '%s': expected 'username=password'", userPass)
		}
		usersMap[parts[0]] = turn.GenerateAuthKey(parts[0], *realm, parts[1])
	}

	bwGenerator := &bwQuotaGenerator{
		RelayAddressGenerator: &turn.RelayAddressGeneratorStatic{
			RelayAddress: net.ParseIP(*publicIP),
			Address:      "0.0.0.0",
		},
		rateLimiters: rateLimiters,
	}

	server, err := turn.NewServer(turn.ServerConfig{
		Realm: *realm,
		AuthHandler: func(ra *turn.RequestAttributes) (string, []byte, bool) {
			if key, ok := usersMap[ra.Username]; ok {
				return ra.Username, key, true
			}

			return "", nil, false
		},
		PacketConnConfigs: []turn.PacketConnConfig{
			{
				PacketConn:            udpListener,
				RelayAddressGenerator: bwGenerator,
			},
		},
	})
	if err != nil {
		log.Panic(err)
	}

	log.Printf("TURN server started on %s:%d", *publicIP, *port)

	if *testMode {
		go runTestClient(*publicIP, *port, *realm, *users, *peerAddr, *testPort)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	log.Println("Shutting down...")

	if err = server.Close(); err != nil {
		log.Panic(err)
	}
}

// runTestClient starts a TURN client, allocates a relay, and exposes a UDP echo server.
// This allows testing the bandwidth quota with tools like iperf.
func runTestClient(publicIP string, port int, realm, users string, peer, testPort string) {
	_ = "STUB: not implemented" //nolint:cyclop
	// Parse first user credentials.
	return
}

// Small delay to let server start.

//nolint:noctx

// Create a UDP echo server that forwards traffic through the relay.
//nolint:noctx

//nolint:forcetypeassert

// Identify client and drop packet if client hasn'y shown up yet.

// Store client address
