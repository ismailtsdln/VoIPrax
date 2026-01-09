package sip

import (
	"testing"
)

func TestParseInvite(t *testing.T) {
	raw := []byte("INVITE sip:alice@atlanta.com SIP/2.0\r\n" +
		"Via: SIP/2.0/UDP pc33.atlanta.com;branch=z9hG4bK776asdh1\r\n" +
		"Max-Forwards: 70\r\n" +
		"To: Alice <sip:alice@atlanta.com>\r\n" +
		"From: Bob <sip:bob@biloxi.com>;tag=192837465\r\n" +
		"Call-ID: a84b4c76e66710@pc33.atlanta.com\r\n" +
		"CSeq: 314159 INVITE\r\n" +
		"Contact: <sip:bob@pc33.atlanta.com>\r\n" +
		"Content-Type: application/sdp\r\n" +
		"Content-Length: 0\r\n\r\n")

	msg, err := Parse(raw)
	if err != nil {
		t.Fatalf("Failed to parse INVITE: %v", err)
	}

	if msg.IsResponse {
		t.Error("Expected Request, got Response")
	}

	if msg.Method != "INVITE" {
		t.Errorf("Expected Method INVITE, got %s", msg.Method)
	}

	if msg.RequestURI != "sip:alice@atlanta.com" {
		t.Errorf("Expected URI sip:alice@atlanta.com, got %s", msg.RequestURI)
	}

	if len(msg.Headers["Via"]) == 0 {
		t.Error("Expected Via header, got none")
	}
}

func TestParseResponse(t *testing.T) {
	raw := []byte("SIP/2.0 200 OK\r\n" +
		"Via: SIP/2.0/UDP pc33.atlanta.com;branch=z9hG4bK776asdh1;received=192.0.2.1\r\n" +
		"To: Alice <sip:alice@atlanta.com>;tag=a6c859\r\n" +
		"From: Bob <sip:bob@biloxi.com>;tag=192837465\r\n" +
		"Call-ID: a84b4c76e66710@pc33.atlanta.com\r\n" +
		"CSeq: 314159 INVITE\r\n" +
		"Contact: <sip:alice@pc33.atlanta.com>\r\n" +
		"Content-Type: application/sdp\r\n" +
		"Content-Length: 0\r\n\r\n")

	msg, err := Parse(raw)
	if err != nil {
		t.Fatalf("Failed to parse Response: %v", err)
	}

	if !msg.IsResponse {
		t.Error("Expected Response, got Request")
	}

	if msg.StatusCode != 200 {
		t.Errorf("Expected StatusCode 200, got %d", msg.StatusCode)
	}

	if msg.Reason != "OK" {
		t.Errorf("Expected Reason OK, got %s", msg.Reason)
	}
}
