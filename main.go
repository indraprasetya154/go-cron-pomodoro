package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-co-op/gocron/v2"
)

var (
	workTime       = 25 * time.Minute
	shortBreakTime = 5 * time.Minute
	isWork         = true
	scheduler      gocron.Scheduler
)

func main() {
	var err error
	scheduler, err = gocron.NewScheduler()
	if err != nil {
		panic(err)
	}

	fmt.Println("Pomodoro started!. Press Ctrl+C to stop.")
	startPomodoroCycle()

	scheduler.Start()

	// handle Ctrl+C
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("\nStopping Pomodoro...")
	if err := scheduler.Shutdown(); err != nil {
		panic(err)
	}
}

func startPomodoroCycle() {
	if isWork {
		startWork()
	} else {
		startBreak()
	}
}

func startWork() {
	fmt.Println("🚀 Let's make some penny!")

	// Use time.AfterFunc to schedule the next phase after workTime
	time.AfterFunc(workTime, func() {
		isWork = false
		startPomodoroCycle()
	})
}

func startBreak() {
	fmt.Println("☕️ Let's take a short break!")

	// Use time.AfterFunc to schedule the next phase after shortBreakTime
	time.AfterFunc(shortBreakTime, func() {
		isWork = true
		startPomodoroCycle()
	})
}
