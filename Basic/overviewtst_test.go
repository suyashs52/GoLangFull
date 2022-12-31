package main

import "testing"

var tests = []struct {
	name     string
	divided  float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},

	{"invalid-data", 00.0, 0.0, 0.0, true},
}

// go test -v //verbose
//go test -cover
//go test -coverprofile=coverage.out && go tool cover -html=coverage.out

func TestDivide(t *testing.T) {
	_, err := divide(10.0, 1.0)

	if err != nil {
		t.Error("Got an error when we should not have")
	}
}

func TestBadDivide(t *testing.T) {
	_, err := divide(10.0, 0)
	if err == nil {
		t.Error("Did not get an error when we should have")
	}
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.divided, tt.divisor)

		if tt.isErr {
			if err == nil {
				t.Error("expected an error but did not get one")
			}
		} else {
			if err != nil {
				t.Error("did not expect one but got one ", err.Error())
			}
		}

		if got != tt.expected {
			t.Errorf("expected %f bug got %f", tt.expected, got)
		}
	}
}
