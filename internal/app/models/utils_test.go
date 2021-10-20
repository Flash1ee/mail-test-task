package models

import "testing"

func cmpSliceByte(sb []byte, sb1 []byte) bool {
	if len(sb) != len(sb1) {
		return false
	}
	for i := range sb {
		if sb[i] != sb1[i] {
			return false
		}
	}
	return true
}
func TestSliceSum_Empty(t *testing.T) {
	var src [][]byte
	var exp []byte
	res := SliceSum(src...)
	if !cmpSliceByte(exp, res) {
		t.Fatalf("receive slice %v not equal expected slice %v ", exp, res)
	}
}
func TestSliceSum_OneSlice(t *testing.T) {
	var src = [][]byte{
		{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40},
	}
	exp := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40}
	res := SliceSum(src...)
	if !cmpSliceByte(exp, res) {
		t.Fatalf("receive slice %v not equal expected slice %v ", exp, res)
	}
}
func TestSliceSum_FourSlice(t *testing.T) {
	var src = [][]byte{
		{0x18, 0x2d},
		{0x44, 0x54},
		{0xfb, 0x21},
		{0x09, 0x40},
	}
	exp := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40}
	res := SliceSum(src...)
	if !cmpSliceByte(exp, res) {
		t.Fatalf("receive slice %v not equal expected slice %v ", exp, res)
	}
}
func TestGetIprotoString(t *testing.T) {
	testString := "string"
	exp := ConvertToProtoString(testString)
	res := GetIprotoString(testString)
	if !cmpProtoString(exp, res) {
		t.Fatalf("receive string %v not equal expected %v", res, exp)
	}
}
func TestGetIprotoStringEmpty(t *testing.T) {
	testString := ""
	exp := ConvertToProtoString(testString)
	res := GetIprotoString(testString)
	if !cmpProtoString(exp, res) {
		t.Fatalf("receive string %v not equal expected %v", res, exp)
	}
}
