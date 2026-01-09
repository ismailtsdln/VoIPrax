package sip

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/ismailtsdln/VoIPrax/internal/logger"
)

// TransportType represents the SIP transport protocol
type TransportType string

const (
	UDP TransportType = "UDP"
	TCP TransportType = "TCP"
)

// Stack handles sending and receiving SIP messages
type Stack struct {
	logger *logger.Logger
	conn   net.PacketConn
	tcpLis net.Listener
	mu     sync.RWMutex
}

// NewStack creates a new SIP stack instance
func NewStack(log *logger.Logger) *Stack {
	return &Stack{
		logger: log,
	}
}

// ListenUDP starts listening for SIP messages over UDP
func (s *Stack) ListenUDP(addr string) error {
	conn, err := net.ListenPacket("udp", addr)
	if err != nil {
		return fmt.Errorf("failed to listen on UDP %s: %w", addr, err)
	}
	s.conn = conn
	s.logger.Info().Str("addr", addr).Msg("SIP Stack listening on UDP")
	return nil
}

// SendUDP sends a SIP message to the specified address over UDP
func (s *Stack) SendUDP(addr string, msg *Message) error {
	raddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return fmt.Errorf("failed to resolve UDP address %s: %w", addr, err)
	}

	raw := []byte(msg.String())
	_, err = s.conn.WriteTo(raw, raddr)
	if err != nil {
		return fmt.Errorf("failed to send UDP message: %w", err)
	}

	s.logger.Debug().
		Str("to", addr).
		Str("method", msg.Method).
		Int("status", msg.StatusCode).
		Msg("SIP message sent via UDP")
	return nil
}

// ReceiveUDP waits for a SIP message on the UDP connection
func (s *Stack) ReceiveUDP(ctx context.Context, timeout time.Duration) (*Message, string, error) {
	buf := make([]byte, 65535)

	if timeout > 0 {
		s.conn.SetReadDeadline(time.Now().Add(timeout))
	} else {
		s.conn.SetReadDeadline(time.Time{})
	}

	n, addr, err := s.conn.ReadFrom(buf)
	if err != nil {
		return nil, "", err
	}

	msg, err := Parse(buf[:n])
	if err != nil {
		return nil, addr.String(), fmt.Errorf("failed to parse SIP message: %w", err)
	}

	s.logger.Debug().
		Str("from", addr.String()).
		Str("method", msg.Method).
		Int("status", msg.StatusCode).
		Msg("SIP message received via UDP")

	return msg, addr.String(), nil
}

// Close closes the SIP stack connections
func (s *Stack) Close() error {
	if s.conn != nil {
		return s.conn.Close()
	}
	return nil
}
