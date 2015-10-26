package main

import (
	"fmt"
	"github.com/juragan360/concurrency/fibo"
	"sync"
	"time"
)

func main() {
	iteration := 100
	fmt.Println("T3 - WaitGroup")

	tstart := time.Now()
	wg := new(sync.WaitGroup)
	for i := 0; i < iteration; i++ {
		wg.Add(1)
		go func(i int, wg *sync.WaitGroup) {
			f := fibo.F{}
			f.Seed = i
			f.Calc()
			fmt.Printf("Fibo of %d is %d \n", f.Seed, f.Result)
			time.Sleep(100 * time.Millisecond)
			wg.Done()
		}(i, wg)
	}
	wg.Wait()

	fmt.Printf("Done. Completed in %v \n", time.Since(tstart))
}
