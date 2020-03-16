package main

import (
	"fmt"
	"sync"
	"time"
)

type timer struct {
	id       int
	isActive bool
	duration time.Duration
}

type Timers struct {
	timers []*timer
	sync.Mutex
}

func (timers *Timers) AddLoopedTimer(
	id int, duration time.Duration, callback func(iteration int),
) {
	timers.Lock()

	timer := &timer{}
	timer.id = id
	timer.isActive = true
	timer.duration = duration
	timers.timers = append(timers.timers, timer)

	fmt.Printf("%d: new timer\n", timer.id)

	timers.Unlock()

	timer.startLooped(callback)
}

func (timers *Timers) DisableTimer(id int) {
	timers.Lock()
	defer timers.Unlock()

	for index, timer := range timers.timers {
		if timer.id == id {
			timers.timers[index].isActive = false
			timers.timers = append(
				timers.timers[:index], timers.timers[index+1:]...,
			)
		}
	}
}

func (timer *timer) startLooped(callback func(iteration int)) {
	iteration := 0
	for {
		if !timer.isActive {
			fmt.Printf("%d: deactivated\n", timer.id)
			return
		}

		time.Sleep(timer.duration)
		iteration++

		fmt.Printf("%d: tick\n", timer.id)

		if timer.isActive {
			callback(iteration)
		}
	}
}
