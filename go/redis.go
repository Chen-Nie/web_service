package src

import (
    "github.com/gomodule/redigo/redis"
)

func redisConnection() {
    con, err := redis.Dial("tcp", "localhost:6379")
    if err != nil {log.Fatal(err)}
    return con
}
