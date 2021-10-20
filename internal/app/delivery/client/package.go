package client

import "mail-test-task/internal/app/models"

const (
	svc_id  int32 = 0x00000002
	svc_msg int32 = 0x00000001
)

func packingHeader(body []byte) ([]byte, error) {
	bodyLength := int32(len(body))
	var reqID int32 = 1

	header, err := (&models.Header{
		SvcId:     svc_id,
		BodyLen:   bodyLength,
		RequestId: reqID,
	}).Encode()

	if err != nil {
		return nil, err
	}
	return header, nil
}

func packingBody(token, scope string) ([]byte, error) {
	protoToken := models.GetIprotoString(token)
	protoScope := models.GetIprotoString(scope)

	body, err := (&models.Request{
		SvcMsg: svc_msg,
		Token:  protoToken,
		Scope:  protoScope,
	}).Encode()

	if err != nil {
		return nil, err
	}
	return body, nil
}

func getPackage(token string, scope string) ([]byte, error) {
	body, err := packingBody(token, scope)
	if err != nil {
		return nil, err
	}
	header, err := packingHeader(body)
	if err != nil {
		return nil, err
	}
	return models.SliceSum(header, body), nil
}
