package wrapper_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/matfax/go-redis-wrapper"
)

func BenchmarkOnce(b *testing.B) {
	codec := newCodec()
	codec.UseLocalCache(1000, time.Minute)

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var n int
			err := codec.Once(&wrapper.Item{
				Key:    "bench-once",
				Object: &n,
				Func: func() (interface{}, error) {
					return 42, nil
				},
			})
			if err != nil {
				panic(err)
			}
			if n != 42 {
				panic(fmt.Sprintf("%d != 42", n))
			}
		}
	})
}
