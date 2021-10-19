package models

import (
	"bytes"
	"encoding/binary"
)

type ResponseBody interface {
	Encode() ([]byte, error)
	Decode() error
}

type Response struct {
	ReturnCode int32
	Body       ResponseBody
}

type ResponseOk struct {
	ClientId   string
	ClientType int32
	UserName   string
	ExpiresIn  int32
	UserId     int64
}
type ResponseError struct {
	ErrorString string
}

func (r *ResponseOk) Encode() ([]byte, error) {
	data := new(bytes.Buffer)
	if err := binary.Write(data, binary.LittleEndian, r); err != nil {
		return nil, err
	}
	return data.Bytes(), nil
}
func (r *ResponseOk) Decode(binData []byte) error {
	data := bytes.NewReader(binData)
	if err := binary.Read(data, binary.LittleEndian, r); err != nil {
		return err
	}
	return nil
}
func (r *ResponseError) Encode() ([]byte, error) {
	data := new(bytes.Buffer)
	if err := binary.Write(data, binary.LittleEndian, r); err != nil {
		return nil, err
	}
	return data.Bytes(), nil
}
func (r *ResponseError) Decode(binData []byte) error {
	data := bytes.NewReader(binData)
	if err := binary.Read(data, binary.LittleEndian, r); err != nil {
		return err
	}
	return nil
}

func (r *Response) Encode() ([]byte, error) {
	dataBin := new(bytes.Buffer)
	if err := binary.Write(dataBin, binary.LittleEndian, r.ReturnCode); err != nil {
		return nil, err
	}
	data, err := r.Body.Encode()
	if err != nil {
		return nil, err
	}
	return SliceSum(dataBin.Bytes(), data), nil
}
func (r *Response) Decode(binData []byte) error {
	data := bytes.NewReader(binData)

	if err := binary.Read(data, binary.LittleEndian, r.ReturnCode); err != nil {
		return err
	}
	err := r.Body.Decode()
	if err != nil {
		return err
	}

	return nil
}
