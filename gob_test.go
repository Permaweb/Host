// Encoding and decoding on gob objects.

package main

import "testing"

func TestGobString(t *testing.T) {
	const test = "Hello World ❤"

	b, err := encodeString(test)
	if err != nil {
		t.Errorf("Couldn't encode a string to gob: %s.", err.Error())
	}

	decoded, err := decodeString(b)
	if err != nil {
		t.Errorf("Couldn't decode a string from gob: %s.", err.Error())
	}

	if decoded != test {
		t.Errorf("Gob+String was incorrect, got: %s, want: %s.", decoded, test)
	}
}
