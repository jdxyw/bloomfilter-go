package bloomfilter

const (
	// m and r are mixing constants generated offline. Taken from the original implementation in C.
	m = 0x5bd1e995
	r = 24
)

// MurmurHash2 returns a 32-bit hash.
// The original algorithm is developed by Austin Appleby.
// MurmurHash is a non-cryptographic hash function suitable for general hash-based lookup.
// http://sites.google.com/site/murmurhash/
// https://en.wikipedia.org/wiki/MurmurHash
func MurmurHash2(b []byte, seed uint32) uint32 {
	var k uint32

	h := seed ^ uint32(len(b))

	for i := len(b); i >= 4; i -= 4 {
		k = uint32(b[0])
		k |= uint32(b[1]) << 8
		k |= uint32(b[2]) << 16
		k |= uint32(b[3]) << 24

		k *= m
		k ^= k >> r
		k *= m
		h *= m
		h ^= k

		b = b[4:]
	}

	switch len(b) {
	case 3:
		h ^= uint32(b[2]) << 16
		fallthrough
	case 2:
		h ^= uint32(b[1]) << 8
		fallthrough
	case 1:
		h ^= uint32(b[0])
		h *= m
	}

	h ^= h >> 13
	h *= m
	h ^= h >> 15

	return h
}
