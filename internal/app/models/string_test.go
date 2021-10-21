package models

import "testing"

func cmpProtoString(str String, str2 String) bool {
	if str.Len != str2.Len || len(str.Str) != len(str.Str) {
		return false
	}

	for i := range str.Str {
		if str.Str[i] != str2.Str[i] {
			return false
		}
	}

	return true
}

func TestString_GetBytesLength_OrdinaryString(t *testing.T) {
	testString := "mailisvk"
	expectedLen := len(testString) + INT32_SIZE

	protoLen := ConvertToProtoString(testString).GetBytesLength()
	if protoLen != expectedLen {
		t.Fatal("incorrect len")
	}
}
func TestString_GetBytesLength_EmptyString(t *testing.T) {
	testString := ""
	expectedLen := len(testString) + INT32_SIZE

	protoLen := ConvertToProtoString(testString).GetBytesLength()
	if protoLen != expectedLen {
		t.Fatal("incorrect len")
	}
}
func TestString_Encode_Decode(t *testing.T) {
	testString := "mailisvk"
	expectedProtoStr := ConvertToProtoString(testString)

	encoded, err := expectedProtoStr.Encode()
	if err != nil {
		t.Fatal("invalid encoded protoStr")
	}

	var decodedProtoStr String
	if err := decodedProtoStr.Decode(encoded); err != nil {
		t.Fatal("invalid decoded protoStr")
	}

	if !cmpProtoString(expectedProtoStr, decodedProtoStr) {
		t.Fatalf("invalid compare: %v not equal %v", expectedProtoStr, decodedProtoStr)
	}
}
func TestString_Encode_Decode_Empty(t *testing.T) {
	testString := ""
	expectedProtoStr := ConvertToProtoString(testString)

	encoded, err := expectedProtoStr.Encode()
	if err != nil {
		t.Fatal("invalid encoded protoStr")
	}

	var decodedProtoStr String
	if err = decodedProtoStr.Decode(encoded); err != nil {
		t.Fatal("invalid decoded protoStr")
	}

	if !cmpProtoString(expectedProtoStr, decodedProtoStr) {
		t.Fatalf("invalid compare: %v not equal %v", expectedProtoStr, decodedProtoStr)
	}
}
func TestString_ToString(t *testing.T) {
	testString := "mailisvk"

	protoString := ConvertToProtoString(testString)
	res, err := protoString.ToString()
	if err != nil || res != testString {
		t.Fatalf("invalid convert from protoStr to str\nexpected: %v received: %v",
			testString, res)
	}
}
func TestString_ToString_Empty(t *testing.T) {
	testString := ""

	protoString := ConvertToProtoString(testString)
	res, err := protoString.ToString()
	if err != nil || res != testString {
		t.Fatalf("invalid convert from protoStr to str\nexpected: %v received: %v",
			testString, res)
	}
}
