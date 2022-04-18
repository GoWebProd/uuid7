package uuid7_test

import (
	"testing"

	"github.com/GoWebProd/uuid7"
)

func BenchmarkNext(b *testing.B) {
	u := uuid7.New()

	for i := 0; i < b.N; i++ {
		_ = u.Next()
	}
}
