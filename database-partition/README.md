# Databas partition

Experiment to simulate log events storing partionated by month

## Run

```sh
➜  docker compose up --build
➜  cd app && go test ./...
```

## To enjoy 

Execute the random log generator and explore how database persist these logs on created partitions
```
➜  cd app && go run ./...
```