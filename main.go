package main

import (
	"sync"

	"github.com/luisgs7/golang-get-started/controllers"
)

func main() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7}

	var wg sync.WaitGroup

	for _, value := range numbers {
		wg.Add(1)
		go func(i int) {
			controllers.Hello(i)
			wg.Done()
		}(value)
	}

	wg.Wait()
}
