package models

func SliceSum(slices ...[]byte) []byte {
	var res []byte
	for _, el := range slices {
		res = append(res, el...)
	}
	return res
}
func GetIprotoString(str string) String {
	length := int32(len(str))
	binStr := make([]int8, 0, length)

	for _, val := range str {
		binStr = append(binStr, int8(val))
	}

	return String{Str: binStr, Len: length}
}
func ConvertToProtoString(str string) String {
	length := len(str)
	b := make([]int8, 0, length)
	for _, val := range str {
		b = append(b, int8(val))
	}
	return String{
		Str: b,
		Len: int32(length),
	}
}
