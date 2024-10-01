package main

import (
	"fmt"

	"github.com/bits-and-blooms/bloom/v3"
)

func main() {
	// using go lib
	fmt.Println("--- using lib ---")
	filter := bloom.NewWithEstimates(1000000, 0.0001)
	filter.AddString("lemon")
	filter.AddString("orange")

	fmt.Printf("check exist lemon: %v \n", filter.TestString("lemon"))
	fmt.Printf("check exist orange: %v \n", filter.TestString("orange"))
	fmt.Printf("check exist non-exist: %v \n", filter.TestString("non-exist"))
	fmt.Println("--- end using lib ---")

	// using pure implement
	fmt.Println("--- using pure implement---")
	pureFilter := NewBloomFilter(1000000, hashFuncsBL)
	pureFilter.AddString("lemon")
	pureFilter.AddString("orange")
	fmt.Printf("check exist lemon: %v \n", pureFilter.TestString("lemon"))
	fmt.Printf("check exist orange: %v \n", pureFilter.TestString("orange"))
	fmt.Printf("check exist non-exist: %v \n", pureFilter.TestString("non-exist"))
	fmt.Println("--- end using pure implement ---")
}
