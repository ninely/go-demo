package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"math"
	"strconv"
	"time"
)

type byteDataGenerate struct {
	data []byte
}

func NewByteDataGenerate(size int) *byteDataGenerate {
	data := make([]byte, size)
	for i := 0; i < size; i++ {
		data[i] = math.MaxUint8
	}
	return &byteDataGenerate{
		data: data,
	}
}

func (b *byteDataGenerate) GetNextByteData() []byte {
	for i := len(b.data) - 1; i > 0; i-- {
		if b.data[i] > 0 {
			b.data[i] -= 1
			return b.data
		}
		b.data[i] = math.MaxUint8
	}

	fmt.Println("----- out of total size -----")

	for i := 0; i < len(b.data); i++ {
		b.data[i] = math.MaxUint8
	}
	return b.data
}

type redisTest struct {
	rdb *redis.Client
}

func NewRedisTest() *redisTest {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	return &redisTest{
		rdb: rdb,
	}
}

func (r *redisTest) WriteValue(ctx context.Context, size, batch int) {
	bData := NewByteDataGenerate(size)
	for i := 1; i <= batch; i++ {
		key := strconv.Itoa(i)
		value := bData.GetNextByteData()
		err := r.rdb.Set(ctx, key, value, 200*time.Second).Err()
		if err != nil {
			panic(err)
		}
	}
}

func (r *redisTest) InfoMemory(ctx context.Context) {
	val, err := r.rdb.Info(ctx, "memory").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
}

func (r *redisTest) InfoKeyspace(ctx context.Context) {
	val, err := r.rdb.Info(ctx, "keyspace").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
}
