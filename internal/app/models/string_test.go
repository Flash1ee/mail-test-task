package models

import "testing"

func convertToProtoString(str string) String {
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

func TestString_GetBytesLength_OrdinaryString(t *testing.T) {
	testString := "mailisvk"
	expectedLen := len(testString) + INT32_SIZE

	protoLen := convertToProtoString(testString).GetBytesLength()
	if protoLen != expectedLen {
		t.Fatal("incorrect len")
	}
}
func TestString_GetBytesLength_EmptyString(t *testing.T) {
	testString := ""
	expectedLen := len(testString) + INT32_SIZE

	protoLen := convertToProtoString(testString).GetBytesLength()
	if protoLen != expectedLen {
		t.Fatal("incorrect len")
	}
}
