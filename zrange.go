package main

import(
	"github.com/garyburd/redigo/redis"
	"fmt"
)


func main(){

	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("connect to redis err", err.Error())
		return
	}

	/// add zrange data
	n,err := c.Do("zadd", "curbike", 5, "dsfsdfsdf")
	if err != nil {
		panic(err)
	}else{
		fmt.Println(n)
	}

  // get count of datat ion zrange by min and max setting
	num, err := c.Do("zcard", "curbike")
	if err != nil {
		fmt.Println("zcard failed", err.Error())
	} else {
		fmt.Printf("curbike's size is %s:", num)
	}

	// get count of datat in zrange by min and max setting
	num, err := c.Do("zcount", "curbike", 1, 100)
	if err != nil {
		fmt.Println("zcount failed ", err.Error())
	} else {
		fmt.Println("zcount num is :", num)
	}

	/// get value of zrange with key by min and max setting
  result, err := redis.Values(c.Do("zrange", "curbike", 0, 100))
	if err != nil {
		fmt.Println("interstore failed", err.Error())
	} else {
		fmt.Printf("interstore newset elsements are:")
		for _, v := range result {
			fmt.Printf("%s ", v.([]byte))
		}
		fmt.Println()
	}
}
