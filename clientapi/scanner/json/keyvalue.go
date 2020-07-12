package json

import (
	"encoding/json"
	"io"
)

// KeyValueScanner scans json keys and values from stream (io.Reader)
// Supports only json objects ({}) and arrays ([]) as value
type KeyValueScanner struct {
	r io.Reader

	bufferSize int
	bufferItr  int
	buffer     []byte

	key   string
	value []byte

	err error
}

func NewKeyValueScanner(r io.Reader, bufSize int) *KeyValueScanner {
	return &KeyValueScanner{
		r: r,

		bufferSize: 0,
		bufferItr:  0,
		buffer:     make([]byte, bufSize),
	}
}

func (s *KeyValueScanner) Scan() bool {
	s.key = ""
	s.value = nil

	s.key, s.err = s.findKey()
	if s.err != nil {
		return false
	}

	s.value, s.err = s.findValue()
	if s.err != nil {
		return false
	}

	return true
}

func (s *KeyValueScanner) Key() string {
	return s.key
}

func (s *KeyValueScanner) Value() json.RawMessage {
	return s.value
}

// TODO: needs better error handling
func (s *KeyValueScanner) Error() error {
	return s.err
}

func (s *KeyValueScanner) findKey() (string, error) {
	var (
		b              byte
		err            error
		beginningFound bool
		keyBuffer      []byte
	)

	for {
		// TODO: it's better to work with the buffer directly
		b, err = s.nextByte()
		if err != nil {
			return "", err
		}

		if !beginningFound {
			if b == '"' {
				beginningFound = true
			}

			continue
		}

		if b == '"' {
			return string(keyBuffer), nil
		}

		keyBuffer = append(keyBuffer, b)
	}

}

func (s *KeyValueScanner) findValue() ([]byte, error) {
	var (
		b              byte
		err            error
		closuresNeeded int
		valueBuffer    []byte
	)

	for {
		b, err = s.nextByte()
		if err != nil {
			return nil, err
		}

		if b == '{' || b == '[' {
			closuresNeeded++
		} else if b == '}' || b == ']' {
			closuresNeeded--
			if closuresNeeded <= 0 {
				valueBuffer = append(valueBuffer, b)
				break
			}
		}

		if closuresNeeded > 0 {
			valueBuffer = append(valueBuffer, b)
		}
	}

	return valueBuffer, nil
}

func (s *KeyValueScanner) nextByte() (byte, error) {
	if s.bufferSize == s.bufferItr {
		n, err := s.r.Read(s.buffer)
		if err != nil {
			return byte(0), err
		}
		s.bufferSize = n
		s.bufferItr = 0
	}

	result := s.buffer[s.bufferItr]
	s.bufferItr++
	return result, nil
}
