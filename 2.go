package main

import (
	"fmt"
	"github.com/juragan360/concurrency/fibo"
	"time"
)

func main() {
	iteration := 100
	fmt.Println("T2 - WaitGroup")

	cfibo := make(chan fibo.F)
	cdone := make(chan bool)

	go func() {
		for {
			select {
			case f := <-cfibo:
				fmt.Printf("Fibo of %d is %d \n", f.Seed, f.Result)
				if f.Seed == 99 {
					cdone <- true
				}
			}
		}
	}()

	tstart := time.Now()
	for i := 0; i < iteration; i++ {
		go func(i int) {
			cfibo <- GetFibo(i)
		}(i)
	}

	wait := true
	for wait {
		select {
		case <-cdone:
			wait = false
		}
	}

	fmt.Printf("Done. Completed in %v \n", time.Since(tstart))
}

func GetFibo(i int) fibo.F {
	f := fibo.F{}
	f.Seed = i
	f.Calc()
	time.Sleep(100 * time.Millisecond)
	return f
}
