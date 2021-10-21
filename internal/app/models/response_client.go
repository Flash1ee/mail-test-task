package models

type ResponseClientOk struct {
	ReturnCode int32
	ClientId   string
	ClientType int32
	UserName   string
	ExpiresIn  int32
	UserId     int64
}

type ResponseClientError struct {
	ReturnCode  int32
	ErrorString string
}

func ConvertToClientResponse(body Response) (interface{}, error) {
	var err error
	if body.ReturnCode < 0 {
		return nil, InvalidErrCode
	}

	if body.Body == nil {
		return nil, EmptyBodyErr
	}

	if body.ReturnCode != 0 {
		ret := ResponseClientError{ReturnCode: body.ReturnCode}
		if ret.ErrorString, err = body.Body.(*ResponseError).ErrorString.ToString(); err != nil {
			return nil, err
		}
		return ret, nil
	}

	stringClientId, err := body.Body.(*ResponseOk).ClientId.ToString()
	if err != nil {
		return nil, err
	}

	stringUserName, err := body.Body.(*ResponseOk).UserName.ToString()
	if err != nil {
		return nil, err
	}

	return ResponseClientOk{
		ReturnCode: body.ReturnCode,
		ClientId:   stringClientId,
		ClientType: body.Body.(*ResponseOk).ClientType,
		UserName:   stringUserName,
		ExpiresIn:  body.Body.(*ResponseOk).ExpiresIn,
		UserId:     body.Body.(*ResponseOk).UserId,
	}, nil

}
