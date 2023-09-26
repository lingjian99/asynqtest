package main

import (
	"ansq_demo/tasks"
	"log"
	"time"

	"github.com/hibiken/asynq"
)


func main() {

    redisConnOpt := asynq.RedisClientOpt{
        Addr: "192.168.3.184:6379",
        //Username: "admin",
        // Omit if no password is required
        Password: "XinYun123",
        // Use a dedicated db number for asynq.
        // By default, Redis offers 16 databases (0..15)
        DB: 0,
    }
    client := asynq.NewClient(redisConnOpt)

    t1, err := tasks.NewWelcomeEmailTask(42)
    if err != nil {
        log.Fatal(err)
    }

    t2, err := tasks.NewReminderEmailTask(42)
    if err != nil {
        log.Fatal(err)
    }

    // Process the task immediately.
    info, err := client.Enqueue(t1)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf(" [*] Successfully enqueued task: %+v", info)

    // Process the task 24 hours later.
    info, err = client.Enqueue(t2, asynq.ProcessIn(24*time.Hour))
    if err != nil {
        log.Fatal(err)
    }
    log.Printf(" [*] Successfully enqueued task: %+v", info)
}
