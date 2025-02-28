package main

func GetBinaryDigit(b byte, idx int) int {
	return int(b >> idx & 0b0000_0001)
}
