package ip_util

import  (
	"testing"
)

// TestExternalIP verifies that TestExternalIP functions without returning an error
func TestExternalIP(t *testing.T) {
	_, err := ExternalIP()
	if err != nil {
		t.Errorf("ExternalIP failed : %w", err)
	}
}