package main

import (
	"log"
    "ansq_demo/tasks"
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

    srv := asynq.NewServer(
        redisConnOpt,
        asynq.Config{Concurrency: 10},
    )

    mux := asynq.NewServeMux()
    mux.HandleFunc(tasks.TypeWelcomeEmail, tasks.HandleWelcomeEmailTask)
    mux.HandleFunc(tasks.TypeReminderEmail, tasks.HandleReminderEmailTask)

    if err := srv.Run(mux); err != nil {
        log.Fatal(err)
    }
}