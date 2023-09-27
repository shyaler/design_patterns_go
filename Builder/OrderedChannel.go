package main

import "fmt"

type OrderedChannel struct {
	size int
}

func (ch *OrderedChannel) setUpChan(size int) {
	fmt.Printf("Setting up ordered channel of size %d\n", size)
	ch.size = size
}
