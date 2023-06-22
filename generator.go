package uuid7

import (
	"math/rand"
	"sync"

	"github.com/GoWebProd/gip/fasttime"
)

type Generator struct {
	counter uint32
	mu      sync.Mutex
	rnd     rand.Source
}

func New() *Generator {
	return &Generator{
		rnd: rand.NewSource(fasttime.Now()),
	}
}

func (u *Generator) Next() UUID {
	ts := fasttime.NowNano() / 1_000_000

	return u.NextWithTimestamp(ts)
}

// Timestamp in milliseconds
func (u *Generator) NextWithTimestamp(ts int64) UUID {
	u.mu.Lock()

	u.counter += 1

	cnt := u.counter
	rnd1 := uint64(u.rnd.Int63())
	rnd2 := uint64(u.rnd.Int63())

	u.mu.Unlock()

	return From(cnt, rnd1, rnd2, ts)
}
