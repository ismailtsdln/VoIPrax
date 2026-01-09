package sip

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Message represents a SIP message (Request or Response)
type Message struct {
	IsResponse bool
	Method     string
	RequestURI string
	StatusCode int
	Reason     string
	Version    string
	Headers    map[string][]string
	Body       []byte
}

// Parse parses a raw SIP message string
func Parse(raw []byte) (*Message, error) {
	msg := &Message{
		Headers: make(map[string][]string),
	}

	reader := bufio.NewReader(bytes.NewReader(raw))

	// Read Start Line
	startLine, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("failed to read start line: %w", err)
	}
	startLine = strings.TrimSpace(startLine)

	parts := strings.Split(startLine, " ")
	if len(parts) < 3 {
		return nil, fmt.Errorf("invalid start line: %s", startLine)
	}

	if strings.HasPrefix(parts[0], "SIP/") {
		// Response: SIP/2.0 200 OK
		msg.IsResponse = true
		msg.Version = parts[0]
		code, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("invalid status code: %s", parts[1])
		}
		msg.StatusCode = code
		msg.Reason = strings.Join(parts[2:], " ")
	} else {
		// Request: INVITE sip:user@domain SIP/2.0
		msg.IsResponse = false
		msg.Method = parts[0]
		msg.RequestURI = parts[1]
		msg.Version = parts[2]
	}

	// Read Headers
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimSpace(line)
		if line == "" {
			break // End of headers
		}

		headerParts := strings.SplitN(line, ":", 2)
		if len(headerParts) != 2 {
			continue // Skip malformed headers
		}

		key := http.CanonicalHeaderKey(strings.TrimSpace(headerParts[0]))
		val := strings.TrimSpace(headerParts[1])
		msg.Headers[key] = append(msg.Headers[key], val)
	}

	// Read Body based on Content-Length
	if clHeaders, ok := msg.Headers["Content-Length"]; ok && len(clHeaders) > 0 {
		cl, err := strconv.Atoi(clHeaders[0])
		if err == nil && cl > 0 {
			body := make([]byte, cl)
			_, err = reader.Read(body)
			if err == nil {
				msg.Body = body
			}
		}
	}

	return msg, nil
}

// String returns the raw string representation of the SIP message
func (m *Message) String() string {
	var b strings.Builder
	if m.IsResponse {
		b.WriteString(fmt.Sprintf("%s %d %s\r\n", m.Version, m.StatusCode, m.Reason))
	} else {
		b.WriteString(fmt.Sprintf("%s %s %s\r\n", m.Method, m.RequestURI, m.Version))
	}

	for k, vs := range m.Headers {
		for _, v := range vs {
			b.WriteString(fmt.Sprintf("%s: %s\r\n", k, v))
		}
	}
	b.WriteString("\r\n")
	b.Write(m.Body)

	return b.String()
}
