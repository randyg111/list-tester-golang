package main

import (
	"fmt"
	"time"

	ll "github.com/randyg111/linked-list-golang"
)

func main() {
	var list ll.List[int]
	n := 40000
	for i := 0; i < n; i++ {
		list.Add(i)
	}
	list.Shuffle()
	// Compare binary search and linear search
	start := time.Now()
	for i := 0; i < n; i++ {
		list.IndexOf(i)
	}

	fmt.Println("Linear Search:", time.Since(start)/time.Duration(n))
	list.Sort()
	start = time.Now()
	for i := 0; i < n; i++ {
		list.Search(i)
	}
	fmt.Println("Binary Search:", time.Since(start)/time.Duration(n))

	list.Shuffle()
	// Compare bogo sort and merge sort
	start = time.Now()
	list.Sort()
	fmt.Println("Merge Sort:", time.Since(start)/time.Duration(n))
	list.Clear()
	m := 5
	for i := 0; i < m; i++ {
		list.Add(i)
	}
	list.Shuffle()
	start = time.Now()
	list.Bogo()
	fmt.Println("Bogo Sort:", time.Since(start)/time.Duration(m))
}
