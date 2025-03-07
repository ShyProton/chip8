package utils

func GetBinaryDigit(b byte, idx int) int {
	const LastDigitMask = 0b0000_0001

	return int(b >> idx & LastDigitMask)
}
