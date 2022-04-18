package uuid7

var hexArray = []byte("0123456789abcdef")

func byteToHex(dest []byte, b byte) {
	dest[0] = hexArray[byte((b>>4)&0x0f)]
	dest[1] = hexArray[byte(b&0x0f)]
}

func hexToByte(data []byte) byte {
	var b byte

	switch {
	case data[0] >= '0' && data[0] <= '9':
		b = byte(data[0] - '0')
	case data[0] >= 'a' && data[0] <= 'f':
		b = byte(data[0] - 'a' + 10)
	case data[0] >= 'A' && data[0] <= 'F':
		b = byte(data[0] - 'A' + 10)
	}

	b <<= 4

	switch {
	case data[1] >= '0' && data[1] <= '9':
		b |= byte(data[1] - '0')
	case data[1] >= 'a' && data[1] <= 'f':
		b |= byte(data[1] - 'a' + 10)
	case data[1] >= 'A' && data[1] <= 'F':
		b |= byte(data[1] - 'A' + 10)
	}

	return b
}

func isHex(data byte) bool {
	return (data >= '0' && data <= '9') || (data >= 'a' && data <= 'f') || (data >= 'A' && data <= 'F')
}
