package main

func main() {
	Map := make(map[int]int)
	for i := 0; i < 100000; i++ {
		go writeMap(Map, i, i)
		go readMap(Map, i)
	}
}

func readMap(Map map[int]int, key int) int {
	return Map[key]
}

func writeMap(Map map[int]int, key int, value int) {
	Map[key] = value
}

//大致意思就是说，并发访问map是不安全的，会出现未定义行为，导致程序退出。所以如果希望在多协程中并发访问map，必须提供某种同步机制，
//一般情况下通过读写锁sync.RWMutex实现对map的并发访问控制，将map和sync.RWMutex封装一下，可以实现对map的安全并发访问。
//同时Go1.9版本之后，已经支持，sync包下：
//func (m *Map) Delete(key interface{})
//func (m *Map) Load(key interface{})(value interface{},ok bool)
//func (m *Map) LoadOrStore(key,value interface{}) (actual interface{},loaded bool)
//func (m *Map) Range(f func(key,value interface{}) bool)
//func (m *Map) Store(key,value interface{})
//
//第三方包：go-concurrentMap也是可以实现并发安全的
//地址：github.com/Golangltd/go-concurrentMap
