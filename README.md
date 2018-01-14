# gocraft-work-example
An example CLI program for showcasing [gocraft/work](https://github.com/gocraft/work)

### Requirements
You will need Redis and godep installed

### Get it running
```
dep ensure
go build

./gocraft-work-example enqueue
Enqueued: send_email with Paylod: map[address:test@example.com subject:hello world customer_id:4]
Enqueued: upload_s3 with Paylod: map[bucket:my-s3-bucket]

./gocraft-work-example process
Starting job: upload_s3
Upload: my-s3-bucket
Starting job: send_email
Send email to: test@example.com with subject: hello world and customer id: 4
```
