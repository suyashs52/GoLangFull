package main

import "testing"

// go test -coverprofile=coverage.out && go tool cover -html=coverage.out
// go test -v ./...
func TestRun(t *testing.T) {
	err := run()
	if err != nil {
		t.Errorf("failed run()")
	}
}
