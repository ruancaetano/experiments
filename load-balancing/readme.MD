# Load Balancing

Simple experiments about web application horizontal scalability, to observe the behavior of different types of load balancing provided by nginx;


> To load balacing be effective, the applicaitons must be into different machines, but to simplify, all instances of the experiments was executed in the same machine.

## Run

```sh
➜  docker compose up --build
```

## To enjoy 

Update the number of instances in nginx.conf upstream directive, and observe on application logs the number of requests distributed for each instance using benchmark tools like wrt:

```
➜  wrk -t 10 -c 50 -d 30s http://localhost:80 --latency
```