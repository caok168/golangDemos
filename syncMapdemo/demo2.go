package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	fmt.Println(time.Now().Hour())

	var mapInt = new(sync.Map)

	// add elem
	mapInt.Store(1, 1)
	mapInt.Store(2, 2)
	mapInt.Store(3, 3)
	fmt.Println("before delete key:")
	// iterator
	mapInt.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)

		return true
	})

	// del
	mapInt.Delete(2)
	fmt.Println("after delete key:")
	mapInt.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})

	// query
	v, ok := mapInt.Load(1)
	if ok {
		fmt.Println(v)
	}

	// load or store
	v, ok = mapInt.LoadOrStore(2, 10)
	fmt.Println("load or store:", v, ",ok:", ok)

	v, ok = mapInt.LoadOrStore(3, 10)
	fmt.Println("load or store:", v, ",ok:", ok)

	mapInt.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}
