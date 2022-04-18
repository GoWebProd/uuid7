package uuid7

var hexArray = []byte("0123456789abcdef")

func byteToHex(dest []byte, b byte) {
	dest[0] = hexArray[byte((b>>4)&0x0f)]
	dest[1] = hexArray[byte(b&0x0f)]
}
