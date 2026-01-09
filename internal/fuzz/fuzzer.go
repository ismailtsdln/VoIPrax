package fuzz

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/ismailtsdln/VoIPrax/internal/sip"
)

// Fuzzer provides methods for SIP message fuzzing
type Fuzzer struct {
	payloads []string
}

// NewFuzzer creates a new Fuzzer instance with default payloads
func NewFuzzer() *Fuzzer {
	return &Fuzzer{
		payloads: []string{
			strings.Repeat("A", 1024),
			strings.Repeat("A", 4096),
			"%s%s%s%s%s%s%s%s%s%s",
			"'; DROP TABLE users; --",
			"<script>alert(1)</script>",
			"../../../../etc/passwd",
			"\x00\x01\x02\x03",
			"!!!!!@@@@@#####$$$$$",
		},
	}
}

// FuzzHeader replaces a header value with a fuzzed payload
func (f *Fuzzer) FuzzHeader(msg *sip.Message, headerName string) *sip.Message {
	if _, ok := msg.Headers[headerName]; !ok {
		return msg
	}

	payload := f.payloads[rand.Intn(len(f.payloads))]
	msg.Headers[headerName] = []string{payload}
	return msg
}

// FuzzMethod replaces the SIP method with a fuzzed payload
func (f *Fuzzer) FuzzMethod(msg *sip.Message) *sip.Message {
	payload := f.payloads[rand.Intn(len(f.payloads))]
	msg.Method = payload
	return msg
}

// GenerateInviteTemplate creates a basic INVITE message template
func GenerateInviteTemplate(targetURI, fromURI, toURI string) *sip.Message {
	return &sip.Message{
		Method:     "INVITE",
		RequestURI: targetURI,
		Version:    "SIP/2.0",
		Headers: map[string][]string{
			"Via":            {"SIP/2.0/UDP 127.0.0.1:5060;branch=z9hG4bK" + fmt.Sprint(rand.Int())},
			"From":           {fmt.Sprintf("<%s>;tag=%d", fromURI, rand.Int())},
			"To":             {fmt.Sprintf("<%s>", toURI)},
			"Call-ID":        {fmt.Sprintf("%d@127.0.0.1", rand.Int())},
			"CSeq":           {fmt.Sprintf("%d INVITE", rand.Intn(1000))},
			"Contact":        {fmt.Sprintf("<sip:voiprax@127.0.0.1:5060>")},
			"Max-Forwards":   {"70"},
			"Content-Type":   {"application/sdp"},
			"Content-Length": {"0"},
		},
	}
}
