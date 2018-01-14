package main

import (
	"fmt"
	"log"

	"github.com/garyburd/redigo/redis"
	"github.com/gocraft/work"
)

var redisPool = &redis.Pool{
	MaxActive: 5,
	MaxIdle:   5,
	Wait:      true,
	Dial: func() (redis.Conn, error) {
		return redis.Dial("tcp", ":6379")
	},
}

var enqueuer = work.NewEnqueuer("my_app_namespace", redisPool)

func enqueueJob(job string, payload work.Q) {
	_, err := enqueuer.Enqueue(job, payload)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Enqueued:", job, "with Paylod:", payload)
}

func enqueueEmail() {
	enqueueJob(
		"send_email",
		work.Q{"address": "test@example.com", "subject": "hello world", "customer_id": 4},
	)
}

func enqueueS3() {
	enqueueJob(
		"upload_s3",
		work.Q{"bucket": "my-s3-bucket"},
	)
}
