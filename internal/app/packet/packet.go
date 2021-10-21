package packet

import "mail-test-task/internal/app/models"

const (
	svc_id  int32 = 0x00000002
	svc_msg int32 = 0x00000001
)

func packingHeader(body []byte) ([]byte, error) {
	bodyLength := int32(len(body))
	var reqID int32 = 1

	header := models.Header{
		SvcId:     svc_id,
		BodyLen:   bodyLength,
		RequestId: reqID,
	}
	encoded, err := header.Encode()

	if err != nil {
		return nil, InvalidPackingHeader
	}
	return encoded, nil
}

func packingBody(token, scope string) ([]byte, error) {
	protoToken := models.GetIprotoString(token)
	protoScope := models.GetIprotoString(scope)

	body := models.Request{
		SvcMsg: svc_msg,
		Token:  protoToken,
		Scope:  protoScope,
	}
	encoded, err := body.Encode()
	if err != nil {
		return nil, InvalidPackingBody
	}
	return encoded, nil
}

func GetPacket(token string, scope string) ([]byte, error) {
	body, err := packingBody(token, scope)
	if err != nil {
		return nil, err
	}
	header, err := packingHeader(body)
	if err != nil {
		return nil, err
	}
	headerLen := models.Header{}.HeaderSize()
	return models.SliceSum(header, body[headerLen:]), nil
}
