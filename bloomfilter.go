package bloomfilter

import (
	"errors"
	"math"
)

const (
	ln2s2 = 0.480453013918201 // ln(2)^2
	ln2   = 0.693147180559945 // ln(2)
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
		totalBytes = totalBits%8 + 1
	} else {
		totalBytes = totalBits / 8
	}

	return &BloomFilter{
		bitmap:  make([]byte, totalBytes, totalBytes),
		entries: entires,
		err:     err,
		h:       int(ln2 * bpe),
	}, nil
}
