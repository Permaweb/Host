// Encoding and decoding on gob objects.

package badger

import (
	"bytes"
	"encoding/gob"

	"github.com/Permaweb/host"
)

// String

func encodeString(s string) (b []byte, err error) {
	buffer := &bytes.Buffer{}
	err = gob.NewEncoder(buffer).Encode(s)
	b = buffer.Bytes()
	return
}

func decodeString(b []byte) (s string, err error) {
	buffer := bytes.NewBuffer(b)
	err = gob.NewDecoder(buffer).Decode(&s)
	return
}

// Repo

func encodeRepo(repo host.Repo) (b []byte, err error) {
	buffer := &bytes.Buffer{}
	err = gob.NewEncoder(buffer).Encode(repo)
	b = buffer.Bytes()
	return
}

func decodeRepo(b []byte) (repo host.Repo, err error) {
	buffer := bytes.NewBuffer(b)
	err = gob.NewDecoder(buffer).Decode(&repo)
	return
}
