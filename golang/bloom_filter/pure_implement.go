package main

import (
	"fmt"
	"hash/fnv"
)

type BloomFilter struct {
	bitSet        []bool
	hashFunctions []func(data string) int
}

func NewBloomFilter(size int, hashFuncs ...func(data string) int) *BloomFilter {
	return &BloomFilter{
		bitSet:        make([]bool, size),
		hashFunctions: hashFuncs,
	}
}

func hashFuncsBL(data string) int {
	result, err := fnv.New64a().Write([]byte(data))
	if err != nil {
		fmt.Printf("error hash func: %v", err)
	}

	return result
}

func (bf *BloomFilter) AddString(data string) {
	for _, hashFunc := range bf.hashFunctions {
		index := hashFunc(data) % int(len(bf.bitSet))
		bf.bitSet[index] = true
	}
}

func (bf *BloomFilter) TestString(data string) bool {
	for _, hashFunc := range bf.hashFunctions {
		index := hashFunc(data) % int(len(bf.bitSet))
		if !bf.bitSet[index] {
			return false
		}
	}

	return true
}
