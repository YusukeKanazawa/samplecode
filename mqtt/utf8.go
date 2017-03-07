package mqtt

import ()

func decodeUTF8(b []byte) (string, int) {
	l := int(b[0]<<8 | b[1])
	str := b[2 : l+2]
	return string(str), l + 2
}
func encodeUTF8(utf8str string) []byte {
	str := []byte(utf8str)
	l := len(str)
	return append([]byte{byte(l >> 8), byte(l & 0xFF)}, str...)
}
