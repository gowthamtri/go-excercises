package main

import (
	"fmt"
	"sync"
)

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	left, right *ChopS
	id int
}

var wg sync.WaitGroup

func (p Philo) Eat(host chan int) {
	defer wg.Done()

	<-host

	p.left.Lock()
	p.right.Lock()

	fmt.Printf("\nPhilosopher %d is starting to eat.", p.id)
	
	fmt.Printf("\nPhilosopher %d is finishing eating.", p.id)

	p.right.Unlock()
	p.left.Unlock()

	host <- 1
}

func main() {
	chopSticks := make([]*ChopS, 5)
	for i:=0;i<5;i++ {
		chopSticks[i] = new(ChopS)
	}

	philosophers := make([]*Philo, 5)
	for i:=0;i<5;i++ {
		philosophers[i] = &Philo { chopSticks[i], chopSticks[(i + 1) % 5], i + 1}
	}

	// Channel to synchronize eating
	host := make(chan int, 2)

	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ { 
			wg.Add(1)
			go philosophers[i].Eat(host)
		}
	}

	host <- 1
	host <- 1

	wg.Wait()
}