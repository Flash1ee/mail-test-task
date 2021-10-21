package models

import (
	"bytes"
	"encoding/binary"
	"strings"
)

const INT32_SIZE = 4

type String struct {
	Str []int8
	Len int32
}

func (s *String) Encode() ([]byte, error) {
	data := new(bytes.Buffer)
	if err := binary.Write(data, binary.LittleEndian, s.Len); err != nil {
		return nil, InvalidEncode
	}

	for i := int32(0); i < s.Len; i++ {
		if err := data.WriteByte(byte(s.Str[i])); err != nil {
			return nil, InvalidEncode
		}
	}

	return data.Bytes(), nil
}
func (s *String) Decode(binData []byte) error {
	data := bytes.NewBuffer(binData)

	if err := binary.Read(data, binary.LittleEndian, &s.Len); err != nil {
		return InvalidDecode
	}

	s.Str = make([]int8, s.Len, s.Len)
	for i := int32(0); i < s.Len; i++ {
		if err := binary.Read(data, binary.LittleEndian, &s.Str[i]); err != nil {
			return InvalidDecode
		}
	}

	return nil
}
func (s String) ToString() (string, error) {
	decodeString := new(strings.Builder)

	for _, char := range s.Str {
		if err := decodeString.WriteByte(byte(char)); err != nil {
			return "", InvalidDecode
		}
	}

	return decodeString.String(), nil
}
func (s String) GetBytesLength() int {
	return len(s.Str) + INT32_SIZE
}
