package uuid7

import (
	"math/rand"
	"sync/atomic"
	"time"
	"unsafe"

	"github.com/GoWebProd/gip/fasttime"
)

type Generator struct {
	lastMS  int64
	counter uint32
	rnd     rand.Source
}

func New() *Generator {
	return &Generator{
		rnd: rand.New(rand.NewSource(time.Now().Unix())),
	}
}

func (u *Generator) Next() string {
	ts := fasttime.NowNano() / 1_000_000
	cnt := atomic.AddUint32(&u.counter, 1)

	var val [16]byte

	val1 := (*uint64)(unsafe.Pointer(&val[0]))
	val2 := (*uint64)(unsafe.Pointer(&val[8]))

	*val1 = (2 << 62) | ((uint64(cnt) & 0xFFF) << 50) | (uint64(u.rnd.Int63()) & 0xFFFFFFFFFFFFF)
	*val2 = (uint64(ts) << 16) + (7 << 12) + uint64(u.rnd.Int63())&0xFFF

	var buf [36]byte

	buf[8] = '-'
	buf[13] = '-'
	buf[18] = '-'
	buf[23] = '-'

	byteToHex(buf[34:36], val[0])
	byteToHex(buf[32:34], val[1])
	byteToHex(buf[30:32], val[2])
	byteToHex(buf[28:30], val[3])
	byteToHex(buf[26:28], val[4])
	byteToHex(buf[24:26], val[5])
	byteToHex(buf[21:23], val[6])
	byteToHex(buf[19:21], val[7])
	byteToHex(buf[16:18], val[8])
	byteToHex(buf[14:16], val[9])
	byteToHex(buf[11:13], val[10])
	byteToHex(buf[9:11], val[11])
	byteToHex(buf[6:8], val[12])
	byteToHex(buf[4:6], val[13])
	byteToHex(buf[2:4], val[14])
	byteToHex(buf[0:2], val[15])

	return string(buf[:])
}
