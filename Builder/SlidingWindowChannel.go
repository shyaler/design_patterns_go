package main

import "fmt"

type SlidingWindowChannel struct{
	windowSize int
}

func (ch *SlidingWindowChannel) setUpChan(size int){
	fmt.Printf("Setting up sliding window channel of window size %d\n", size)
	ch.windowSize = size
}