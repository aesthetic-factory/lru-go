package main

import (
	"fmt"
	"root/lru"
)

func main() {

	cache := new(lru.LRU[int32, int])
	cache.Init(3)
	cache.Insert(1, 100)
	cache.Insert(2, 200)
	cache.Insert(3, 300)
	cache.Insert(1, 101)
	cache.Insert(4, 400)
	cache.Show()

	fmt.Println("Get 2")
	get2Res, get2Err := cache.Get(2)
	fmt.Println("Get 2: ", get2Err, " value: ", get2Res)
	cache.Show()

	fmt.Println("Get 3")
	get3Res, get3Err := cache.Get(3)
	fmt.Println("Get 3: ", get3Err, " value: ", get3Res)
	cache.Show()

	fmt.Println("Remove 4")
	cache.Remove(4)
	cache.Show()
}
