package benchmark

import (
	"math/rand"
	"bloomfilter-go"
	"testing"
)

const fixedByetesLen = 6

func benchmark(entries int64, err float64, n int) {
	b, _ := bloomfilter.NewBloomFilter(entries, err)

	rand.Seed(6)
	keys := make([][]byte, 0)
	for i := 0; i < n; i++ {
		key := make([]byte, fixedByetesLen)
		rand.Read(key)
		b.Set(key)

		if rand.Float32() < 0.05 {
			keys = append(keys, key)
		}
	}

	for _, key := range keys {
		b.Check(key)
	}
}
func BenchmarkEntries1000000Err005(b *testing.B) {
	benchmark(1000000, 0.05, b.N)
}

func BenchmarkEntries2000000Err005(b *testing.B) {
	benchmark(2000000, 0.05, b.N)
}

func BenchmarkEntries5000000Err005(b *testing.B) {
	benchmark(5000000, 0.05, b.N)
}

func BenchmarkEntries5000000Err001(b *testing.B) {
	benchmark(5000000, 0.01, b.N)
}

func BenchmarkEntries10000000Err005(b *testing.B) {
	benchmark(10000000, 0.05, b.N)
}

func BenchmarkEntries50000000Err001(b *testing.B) {
	benchmark(50000000, 0.01, b.N)
}