# Redis cache library for Golang [![Build Status](https://travis-ci.org/go-redis/cache.svg?branch=v3)](https://travis-ci.org/go-redis/cache)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fmatfax%2Fgo-redis-wrapper.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fmatfax%2Fgo-redis-wrapper?ref=badge_shield)

## Installation

    go get github.com/go-redis/cache

## Quickstart


```go
package cache_test

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/vmihailenco/msgpack"

	"github.com/go-redis/cache"
)

type Object struct {
	Str string
	Num int
}

func Example_basicUsage() {
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server1": ":6379",
			"server2": ":6380",
		},
	})

	codec := &cache.Codec{
		Redis: ring,

		Marshal: func(v interface{}) ([]byte, error) {
			return msgpack.Marshal(v)
		},
		Unmarshal: func(b []byte, v interface{}) error {
			return msgpack.Unmarshal(b, v)
		},
	}

	key := "mykey"
	obj := &Object{
		Str: "mystring",
		Num: 42,
	}

	codec.Set(&cache.Item{
		Key:        key,
		Object:     obj,
		Expiration: time.Hour,
	})

	var wanted Object
	if err := codec.Get(key, &wanted); err == nil {
		fmt.Println(wanted)
	}

	// Output: {mystring 42}
}

func Example_advancedUsage() {
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server1": ":6379",
			"server2": ":6380",
		},
	})

	codec := &cache.Codec{
		Redis: ring,

		Marshal: func(v interface{}) ([]byte, error) {
			return msgpack.Marshal(v)
		},
		Unmarshal: func(b []byte, v interface{}) error {
			return msgpack.Unmarshal(b, v)
		},
	}

	obj := new(Object)
	err := codec.Once(&cache.Item{
		Key:    "mykey",
		Object: obj, // destination
		Func: func() (interface{}, error) {
			return &Object{
				Str: "mystring",
				Num: 42,
			}, nil
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(obj)
	// Output: &{mystring 42}
}
```


## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fmatfax%2Fgo-redis-wrapper.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fmatfax%2Fgo-redis-wrapper?ref=badge_large)