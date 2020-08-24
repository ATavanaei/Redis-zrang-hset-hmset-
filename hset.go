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
  
	// set datat in hset
	_, err = c.Do("HMSET", "album:3", "title", "james", "artist", "besting", "price", 24.5, "likes", 12000)
	if err != nil {
		panic(err)
	}

	// get data from hset with key title of value
  title, err := redis.String(c.Do("HGET", "album:3", "price"))
	if err != nil {
		panic(err)
	}else {
		fmt.Println(title)
	}

	// get all set data of key and scan
	type Album struct {
		Title  string  `redis:"title"`
		Artist string  `redis:"artist"`
		Price  float64 `redis:"price"`
		Likes  int     `redis:"likes"`
	}
	var album Album
	values, err := redis.Values(c.Do("HGETALL", "album:2"))
	if err != nil {
		panic(err)
	}else {
		err = redis.ScanStruct(values, &album)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%+v", album)
	}

	// exist value by key
	v, err := c.Do("HEXISTS", "album:2", "title")
	if err != nil {
		panic(err)
	}else{
		fmt.Println(v)
	}
}
