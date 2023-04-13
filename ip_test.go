package ip

import (
	"net/http"
	"testing"
)

func TestLocal(t *testing.T) {
	localIP, err := Local()
	if err != nil {
		t.Errorf("Failed to get local IP address: %v", err)
	}
	if localIP == nil {
		t.Error("Local IP address is nil")
	}
}

func TestRemote(t *testing.T) {
	// Create a mock HTTP request
	req, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		t.Errorf("Failed to create HTTP request: %v", err)
	}
	req.RemoteAddr = "192.168.1.1:1234"

	remoteIP, err := Remote(req)
	if err != nil {
		t.Errorf("Failed to get remote IP address: %v", err)
	}
	if remoteIP.String() != "192.168.1.1" {
		t.Errorf("Expected remote IP address to be 192.168.1.1, got %v", remoteIP)
	}

	// Test with invalid IP address
	req.RemoteAddr = "invalid-ip"
	_, err = Remote(req)
	if err == nil {
		t.Error("Expected error for invalid IP address")
	}
}

