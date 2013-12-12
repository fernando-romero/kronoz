package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

const (
	ADDRESS = "127.0.0.1:6379"
)

var (
	c, err = redis.Dial("tcp", ADDRESS)
)

func main() {

	c.Do("ZADD", "test_sorted_set", 1, 999)
	c.Do("ZADD", "test_sorted_set", 2, 888)
	c.Do("ZADD", "test_sorted_set", 3, 777)
	c.Do("ZADD", "test_sorted_set", 2, 222)
	c.Do("ZADD", "test_sorted_set", 3, 333)
	c.Do("ZADD", "test_sorted_set", 4, 666)
	c.Do("ZADD", "test_sorted_set", 5, 555)

	reply, err := c.Do("ZRANGEBYSCORE", "test_sorted_set", 2, 3, "WITHSCORES")

	result, err := redis.Strings(reply, nil)
	fmt.Println(result)

	if err != nil {
		fmt.Println(err)
	}
}
