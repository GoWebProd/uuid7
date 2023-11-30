package uuid7

import (
	"bytes"
	"encoding/binary"
	"unsafe"

	"github.com/pkg/errors"
)

type UUID [16]byte

var (
	ErrBadUUID    = errors.New("bad UUID")
	ErrBadVersion = errors.New("bad UUID version")
)

func From(cnt uint32, rnd1 uint64, rnd2 uint64, ts int64) UUID {
	var val [16]byte

	binary.LittleEndian.PutUint64(val[0:8], (2<<62)|((uint64(cnt)&0xFFF)<<50)|(rnd1&0xFFFFFFFFFFFFF))
	binary.LittleEndian.PutUint64(val[8:16], (uint64(ts)<<16)+(7<<12)+rnd2&0xFFF)

	return val
}

func Parse(uuid string) (UUID, error) {
	const uuidLen = 36
	var u UUID

	bytesPtr := *(*[]byte)(unsafe.Pointer(&uuid))
	if uuidLen != len(bytesPtr) {
		return u, ErrBadUUID
	}
	s := *(*[]byte)(unsafe.Pointer(&struct {
		data uintptr
		len  int
		cap  int
	}{uintptr(unsafe.Pointer(&bytesPtr[0])), uuidLen, uuidLen}))

	if s[8] != '-' || s[13] != '-' || s[18] != '-' || s[23] != '-' {
		return u, ErrBadUUID
	}

	if !isHex(s[34]) || !isHex(s[32]) || !isHex(s[30]) || !isHex(s[28]) ||
		!isHex(s[26]) || !isHex(s[24]) || !isHex(s[21]) || !isHex(s[19]) ||
		!isHex(s[16]) || !isHex(s[14]) || !isHex(s[11]) || !isHex(s[9]) ||
		!isHex(s[6]) || !isHex(s[4]) || !isHex(s[2]) || !isHex(s[0]) {
		return u, ErrBadUUID
	}

	u[0] = hexToByte(s[34:36])
	u[1] = hexToByte(s[32:34])
	u[2] = hexToByte(s[30:32])
	u[3] = hexToByte(s[28:30])
	u[4] = hexToByte(s[26:28])
	u[5] = hexToByte(s[24:26])
	u[6] = hexToByte(s[21:23])
	u[7] = hexToByte(s[19:21])
	u[8] = hexToByte(s[16:18])
	u[9] = hexToByte(s[14:16])
	u[10] = hexToByte(s[11:13])
	u[11] = hexToByte(s[9:11])
	u[12] = hexToByte(s[6:8])
	u[13] = hexToByte(s[4:6])
	u[14] = hexToByte(s[2:4])
	u[15] = hexToByte(s[0:2])

	if v := u.version(); v != 7 {
		return u, errors.Wrapf(ErrBadVersion, "expected version 7, got %d", v)
	}

	return u, nil
}

func (u UUID) version() uint32 {
	return uint32(u[9] >> 4)
}

func (u UUID) Timestamp() uint64 {
	return binary.LittleEndian.Uint64(u[8:16]) >> 16
}

func (u UUID) String() string {
	var buf [36]byte

	buf[8] = '-'
	buf[13] = '-'
	buf[18] = '-'
	buf[23] = '-'

	byteToHex(buf[34:36], u[0])
	byteToHex(buf[32:34], u[1])
	byteToHex(buf[30:32], u[2])
	byteToHex(buf[28:30], u[3])
	byteToHex(buf[26:28], u[4])
	byteToHex(buf[24:26], u[5])
	byteToHex(buf[21:23], u[6])
	byteToHex(buf[19:21], u[7])
	byteToHex(buf[16:18], u[8])
	byteToHex(buf[14:16], u[9])
	byteToHex(buf[11:13], u[10])
	byteToHex(buf[9:11], u[11])
	byteToHex(buf[6:8], u[12])
	byteToHex(buf[4:6], u[13])
	byteToHex(buf[2:4], u[14])
	byteToHex(buf[0:2], u[15])

	return string(buf[:])
}

var emptyUUID [16]byte

func (u UUID) Empty() bool {
	return bytes.Equal(u[:], emptyUUID[:])
}
