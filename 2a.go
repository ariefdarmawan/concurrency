package main

import (
	"fmt"
	"github.com/juragan360/concurrency/fibo"
	"time"
)

func main() {
	iteration := 100
	fmt.Println("T2 - WaitGroup")

	crcvd := make(chan int)
	cfibo := make(chan fibo.F)
	cdone := make(chan bool)

	go func() {
		for {
			select {
			case i := <-crcvd:
				go func(i int) {
					cfibo <- GetFibo(i)
				}(i)

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
		/*
			go func(i int) {
				cfibo <- GetFibo(i)
			}(i)
		*/
		crcvd <- i
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
