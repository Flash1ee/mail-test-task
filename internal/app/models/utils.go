package models

func SliceSum(slices ...[]byte) []byte {
	var res []byte
	for _, el := range slices {
		res = append(res, el...)
	}
	return res
}
