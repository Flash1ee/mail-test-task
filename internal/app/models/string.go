package models

import (
	"bytes"
	"encoding/binary"
	"strings"
)

// не очень хорошо, но использовать пакет unsafe не советуют из-за непереносимости..
const INT32_SIZE = 4

type String struct {
	Str []int8
	Len int32
}

func (s *String) Encode() ([]byte, error) {
	data := new(bytes.Buffer)
	if err := binary.Write(data, binary.LittleEndian, s.Len); err != nil {
		return nil, err
	}
	for i := int32(0); i < s.Len; i++ {
		if err := data.WriteByte(byte(s.Str[i])); err != nil {
			return nil, err
		}
	}
	return data.Bytes(), nil
}
func (s *String) Decode(binData []byte) error {
	data := bytes.NewBuffer(binData)
	if err := binary.Read(data, binary.LittleEndian, s.Len); err != nil {
		return err
	}

	s.Str = make([]int8, 0, s.Len)
	for i := int32(0); i < s.Len; i++ {
		if err := binary.Read(data, binary.LittleEndian, &s.Str[i]); err != nil {
			return err
		}
	}

	return nil
}
func (s String) ToString() (string, error) {
	decodeString := new(strings.Builder)
	for _, char := range s.Str {
		if err := decodeString.WriteByte(byte(char)); err != nil {
			return "", err
		}
	}
	return decodeString.String(), nil
}
func (s String) GetBytesLength() int {
	return len(s.Str) + INT32_SIZE
}
