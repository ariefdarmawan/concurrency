package main

import (
	"fmt"
	"github.com/juragan360/concurrency/fibo"
	"time"
)

func main() {
	iteration := 100
	fmt.Println("T1 - Go Routine")

	tstart := time.Now()
	for i := 0; i < iteration; i++ {
		go func() {
			f := fibo.F{}
			f.Seed = i
			f.Calc()
			fmt.Printf("Fibo of %d is %d \n", f.Seed, f.Result)
			time.Sleep(100 * time.Millisecond)
		}()
	}

	fmt.Printf("Done. Completed in %v \n", time.Since(tstart))
}
