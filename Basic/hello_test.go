package main

import "testing"

func TestHello(t *testing.T) {

	input := Hello()

	checkWithValue := "hello, world"

	if input != checkWithValue {
		t.Errorf("input is mismatching with value %s", input)
	}
}

func TestHelloRAvali(t *testing.T) {

	input := Hello()

	checkWithValue := "hello, world"

	if input != checkWithValue {
		t.Errorf("input is mismatching with value %s", input)
	}

}
