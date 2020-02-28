package main

import (
	"fmt"
	"sync"

	"github.com/shadyzoz/goemitter"
)

type myEmitter struct {
	goemitter.EventEmitter
}

func main() {
	emitter := myEmitter{}
	var wg sync.WaitGroup
	emitter.On("hello", func(args ...interface{}) {
		fmt.Println("hello", args[0])
		wg := args[1].(*sync.WaitGroup)
		wg.Done()
	})
	wg.Add(1)
	emitter.Emit("hello", "goemitter", &wg)
	wg.Wait()
}
