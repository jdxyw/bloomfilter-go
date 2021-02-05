package bloomfilter

import (
	"errors"
	"math"
	"sync"
)

const (
	ln2s2 = 0.480453013918201 // ln(2)^2
	ln2   = 0.693147180559945 // ln(2)
	seed  = 0x9747b28c
)

var (
	errNegativeError   = errors.New("error should be between 0 and 1")
	errEntriesTooSmall = errors.New("entries should be larger than 10000")
)

// BloomFilter is the base struct for bloom filter.
type BloomFilter struct {
	bitmap  []byte
	h       int     // The number of the has function
	entries int64   // The expected number of elements this filter should supports.
	err     float64 // The expected probability of collision.
	bits    int     // The total bits this bloom filter supports
	mutex sync.RWMutex
}

// NewBloomFilter returns a pointer of a BloomFilter struct.
func NewBloomFilter(entires int64, err float64) (*BloomFilter, error) {
	if err < 0 || err > 1 {
		return nil, errNegativeError
	}

	if entires < 10000 {
		return nil, errEntriesTooSmall
	}
	// bits per element
	// The optimal bits each entries need based on the expected error rate.
	bpe := -math.Log(err) / ln2s2

	totalBits := int(float64(entires) * bpe)

	totalBytes := 0

	if totalBits%8 == 0 {
		totalBytes = totalBits / 8
	} else {
		totalBytes = totalBits/8 + 1
	}

	return &BloomFilter{
		bitmap:  make([]byte, totalBytes, totalBytes),
		entries: entires,
		err:     err,
		h:       int(ln2 * bpe),
		bits:    totalBits,
	}, nil
}

// Check returns true if the element exists in this bloom filter.
func (b *BloomFilter) Check(data []byte) bool {
	h1 := MurmurHash2(data, seed)
	h2 := MurmurHash2(data, h1)

	b.mutex.RLock()
	defer b.mutex.RUnlock()

	for i := 0; i < b.h; i++ {
		x := (h1 + uint32(i)*h2) % uint32(b.bits)
		if b.isBitSet(x) == false {
			return false
		}
	}

	return true
}

// Set adds data to this bloom filter
func (b *BloomFilter) Set(data []byte) {
	h1 := MurmurHash2(data, seed)
	h2 := MurmurHash2(data, h1)

	b.mutex.Lock()
	defer b.mutex.Unlock()

	for i := 0; i < b.h; i++ {
		x := (h1 + uint32(i)*h2) % uint32(b.bits)
		b.setBit(x)
	}
}

func (b *BloomFilter) isBitSet(bidx uint32) bool {
	var bytesIdx, bitIdx uint32

	bytesIdx = bidx / uint32(8)
	if bidx%8 == 0 {
		bitIdx = uint32(7)
	} else {
		bitIdx = bidx % uint32(8)
	}

	c := b.bitmap[int(bytesIdx)]
	if (c&(1<<bitIdx))>>bitIdx == 0x01 {
		return true
	}

	return false
}

func (b *BloomFilter) setBit(bidx uint32) {
	var bytesIdx, bitIdx uint32

	bytesIdx = bidx / uint32(8)
	if bidx%8 == 0 {
		bitIdx = uint32(7)
	} else {
		bitIdx = bidx % uint32(8)
	}

	b.bitmap[int(bytesIdx)] = b.bitmap[int(bytesIdx)] | (1 << bitIdx)
}
