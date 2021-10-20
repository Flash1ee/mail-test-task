package models

import (
	"bytes"
	"encoding/binary"
)

type Request struct {
	Header Header
	SvcMsg int32
	Token  String
	Scope  String
}

func (r *Request) Encode() ([]byte, error) {
	headerBin, err := r.Header.Encode()
	if err != nil {
		return nil, err
	}

	svcBin := new(bytes.Buffer)
	if err := binary.Write(svcBin, binary.LittleEndian, r.SvcMsg); err != nil {
		return nil, err
	}

	tokenBin, err := r.Token.Encode()
	if err != nil {
		return nil, err
	}

	scopeBin, err := r.Scope.Encode()
	if err != nil {
		return nil, err
	}

	return SliceSum(headerBin, svcBin.Bytes(), tokenBin, scopeBin), nil
}
func (r *Request) Decode(binData []byte) error {
	data := bytes.NewBuffer(binData)
	if err := r.Header.Decode(data.Bytes()); err != nil {
		return err
	}
	data.Next(r.Header.HeaderSize())
	if err := binary.Read(data, binary.LittleEndian, &r.SvcMsg); err != nil {
		return err
	}
	if err := r.Token.Decode(data.Bytes()); err != nil {
		return err
	}
	data.Next(r.Token.GetBytesLength())

	if err := r.Scope.Decode(data.Bytes()); err != nil {
		return err
	}
	return nil
}
