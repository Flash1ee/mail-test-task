package models

import "testing"

func TestString_GetBytesLength_OrdinaryString(t *testing.T) {
	testString := "mailisvk"
	expectedLen := len(testString) + INT32_SIZE

	protoLen := GetIprotoString(testString).GetBytesLength()
	if protoLen != expectedLen {
		t.Fatal("incorrect len")
	}
}
func TestString_GetBytesLength_EmptyString(t *testing.T) {
	testString := ""
	expectedLen := len(testString) + INT32_SIZE

	protoLen := GetIprotoString(testString).GetBytesLength()
	if protoLen != expectedLen {
		t.Fatal("incorrect len")
	}
}
