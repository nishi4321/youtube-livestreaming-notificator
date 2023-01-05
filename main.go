package main

import (
	"youtube-livestreaming-notificator/internal/scheduler"
	"youtube-livestreaming-notificator/internal/worker"
)

func main() {
	go scheduler.StartScheduler()
	worker.StartWorker()
}
