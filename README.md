# bloomfilter-go
![Go](https://github.com/jdxyw/bloomfilter-go/workflows/Go/badge.svg?branch=main)
[![codecov](https://codecov.io/gh/jdxyw/bloomfilter-go/branch/main/graph/badge.svg?token=ARRCE62BDV)](https://codecov.io/gh/jdxyw/bloomfilter-go)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/jdxyw/bloomfilter-go/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/jdxyw/bloomfilter-go)](https://goreportcard.com/report/github.com/jdxyw/bloomfilter-go)

A Bloom filter is a space-efficient probabilistic data structure, conceived by Burton Howard Bloom in 1970, that is used to test whether an element is a member of a set. False positive matches are possible, but false negatives are not – in other words, a query returns either "possibly in set" or "definitely not in set".  (from wikipedia)

In short, if a bloom filter return `false` for a specific element, then the element is `definitely not` in this bloom filter. 

If a bloom filter return `true` for a specific element, then it's possible that the element is not in this bloom filter. The probability (error rate) is based on the number of the hash functions, the total number of bit, and the bit per element.

`NewBloomFilter` provides two parameters. The first one is `entries` that indicate your expected max number of element.
The second one is `err` that indicate the false positive rate (error rate) you allows. Based on these two parameters, this package could find the optimal number for hash function and bit per element.

## Install

Install this package through `go get`.

```bash
go get github.com/jdxyw/bloomfilter-go
```

## Usage

The usage is quite simple.

```go
package main

import (
	"fmt"
	bf "github.com/jdxyw/bloomfilter-go"
)

func main() {
    // The first parameter is your expected elements or expected max number of elements.
    // The second parameter is your expected collision error rate you can accept.
	b, _ := bf.NewBloomFilter(100000, 0.01)

	b.Set([]byte("Java"))
	b.Set([]byte("Python"))
	b.Set([]byte("Go"))
	b.Set([]byte("C++"))

	if b.Check([]byte("Python")) == true {
		fmt.Println("The Python is in this bloomfilter.")
	}
}
```

## Benchmark

```bash
go test -bench=. ./benchmark/
```

```
BenchmarkEntries1000000Err005-8          8743465               155 ns/op
BenchmarkEntries2000000Err005-8          8170458               152 ns/op
BenchmarkEntries5000000Err005-8          6824294               183 ns/op
BenchmarkEntries5000000Err001-8          5236119               237 ns/op
BenchmarkEntries10000000Err005-8         5822541               218 ns/op
BenchmarkEntries50000000Err001-8         3174663               369 ns/op
```