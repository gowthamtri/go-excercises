package main

import (
	"fmt"
	"sync"
)

type ChopS struct{
	sync.Mutex
}

type Philo struct{
	leftCS, rightCS *ChopS
}

var wg sync.WaitGroup
var host = make(chan int, 2)
	
func(p Philo) eat(id int){
	for i:=0; i<3; i++{
		host <- 1
		
		p.leftCS.Lock()
		p.rightCS.Lock()
		
		fmt.Println("Philosopher", id+1,"is Eating")
		
		p.rightCS.Unlock()
		p.leftCS.Unlock()
		
		fmt.Println("Philosopher", id+1,"Finished Eating")
		
		<- host
	}
	wg.Done()
}

func main () {
	
	CSticks := make([] *ChopS, 5)
	
	for i:=0; i<5; i++ {
		CSticks[i] = new (ChopS)
	}

	philos := make([] *Philo, 5)

	for i:=0; i<5; i++ {
		philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5]}
	}

	wg.Add(5)
	for i:=0; i<5; i++ {
		go philos[i].eat(i)
	}
	wg.Wait()
}