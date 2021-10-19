package models

import (
	"bytes"
	"encoding/binary"
)

type Header struct {
	SvcId     int32
	BodyLen   int32
	RequestId int32
}

func (h *Header) Encode() ([]byte, error) {
	data := new(bytes.Buffer)
	if err := binary.Write(data, binary.LittleEndian, h); err != nil {
		return nil, err
	}
	return data.Bytes(), nil
}

func (h *Header) Decode(binData []byte) error {
	data := bytes.NewBuffer(binData)
	if err := binary.Read(data, binary.LittleEndian, h); err != nil {
		return err
	}
	return nil
}
