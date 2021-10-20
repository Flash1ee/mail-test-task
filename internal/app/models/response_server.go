package models

import (
	"bytes"
	"encoding/binary"
)

type ResponseBody interface {
	Encode() ([]byte, error)
	Decode([]byte) error
}

type Response struct {
	ReturnCode int32
	Body       ResponseBody
}

type ResponseOk struct {
	ClientId   String
	ClientType int32
	UserName   String
	ExpiresIn  int32
	UserId     int64
}
type ResponseError struct {
	ErrorString String
}

func (r *ResponseOk) Encode() ([]byte, error) {
	data := new(bytes.Buffer)
	binClientId, err := r.ClientId.Encode()
	if err != nil {
		return nil, err
	}
	if _, err := data.Write(binClientId); err != nil {
		return nil, err
	}
	if err := binary.Write(data, binary.LittleEndian, r.ClientType); err != nil {
		return nil, err
	}
	binUserName, err := r.UserName.Encode()
	if err != nil {
		return nil, err
	}
	if _, err := data.Write(binUserName); err != nil {
		return nil, err
	}
	if err := binary.Write(data, binary.LittleEndian, r.ExpiresIn); err != nil {
		return nil, err
	}
	if err := binary.Write(data, binary.LittleEndian, r.UserId); err != nil {
		return nil, err
	}

	return data.Bytes(), nil
}
func (r *ResponseOk) Decode(binData []byte) error {
	data := bytes.NewBuffer(binData)
	if err := r.ClientId.Decode(data.Bytes()); err != nil {
		return err
	}
	data.Next(r.ClientId.GetBytesLength())
	if err := binary.Read(data, binary.LittleEndian, r.ClientType); err != nil {
		return err
	}
	if err := r.UserName.Decode(data.Bytes()); err != nil {
		return err
	}
	data.Next(r.UserName.GetBytesLength())
	if err := binary.Read(data, binary.LittleEndian, r.ExpiresIn); err != nil {
		return err
	}
	if err := binary.Read(data, binary.LittleEndian, r.UserId); err != nil {
		return err
	}
	return nil
}
func (r *ResponseError) Encode() ([]byte, error) {
	binErrString, err := r.ErrorString.Encode()
	if err != nil {
		return nil, err
	}
	return binErrString, nil
}
func (r *ResponseError) Decode(binData []byte) error {
	data := bytes.NewBuffer(binData)
	if err := r.ErrorString.Decode(data.Bytes()); err != nil {
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
	data := bytes.NewBuffer(binData)

	if err := binary.Read(data, binary.LittleEndian, r.ReturnCode); err != nil {
		return err
	}
	if r.ReturnCode != 0 {
		r.Body = &ResponseError{}
	} else {
		r.Body = &ResponseOk{}
	}
	err := r.Body.Decode(data.Bytes())
	if err != nil {
		return err
	}
	return nil
}
