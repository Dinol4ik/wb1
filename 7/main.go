package main

import (
	"fmt"
	"sync"
)

type ConcurrencyMap struct {
	m map[string]int
	sync.RWMutex
}

func NewMap() *ConcurrencyMap {
	return &ConcurrencyMap{
		m: make(map[string]int),
	}
}

// Запись в мапу не потокобезопасна, используем mutex

func (cMap *ConcurrencyMap) Set(value int) {
	key := fmt.Sprintf("worker%d", value)
	cMap.Lock()
	cMap.m[key] = value
	cMap.Unlock()
	fmt.Printf("value %d added in map with key %s\n", value, key)
}

//удаление из map используя mutex.lock

func (cMap *ConcurrencyMap) Remove(key string) {
	cMap.Lock()
	defer cMap.Unlock()
	delete(cMap.m, key)
}

// используем RMutex для чтения из map
// при mutex.lock у пишущей горутины - произойдет лок читающей горутины

func (cMap *ConcurrencyMap) Get(key string) (int, bool) {
	cMap.RLock()
	value, ok := cMap.m[key]
	cMap.RUnlock()
	return value, ok
}

func main() {
	const workerPul = 3
	var text string
	wg := &sync.WaitGroup{}
	myMap := NewMap()
	wg.Add(workerPul)

	// конкурентно запишем в мапу значение
	for i := 0; i < workerPul; i++ {
		go func(i int, concurrencyMap *ConcurrencyMap) {
			defer wg.Done()
			concurrencyMap.Set(i)
		}(i, myMap)
	}
	wg.Wait()

	// конкурентно считаем из мапы значение
	wg.Add(workerPul)
	for i := 0; i < workerPul; i++ {
		text = fmt.Sprintf("worker%d", i)
		go func(text string, concurrencyMap *ConcurrencyMap) {
			defer wg.Done()
			result, ok := concurrencyMap.Get(text)
			if ok {
				fmt.Println(result)
			}
		}(text, myMap)
	}
	wg.Wait()
	fmt.Println(myMap.m)
}
