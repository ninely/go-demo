package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"math"
	"strconv"
	"strings"
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
	keyFormat := getKeyFormat(batch)
	bData := NewByteDataGenerate(size)
	for i := 1; i <= batch; i++ {
		key := fmt.Sprintf(keyFormat, strconv.Itoa(i))
		value := bData.GetNextByteData()
		err := r.rdb.Set(ctx, key, value, 250*time.Second).Err()
		if err != nil {
			panic(err)
		}
	}
}

func getKeyFormat(batch int) string {
	b := batch
	strSize := 0
	for b > 0 {
		b /= 10
		strSize += 1
	}
	return "%0" + strconv.Itoa(strSize) + "s"
}

func (r *redisTest) InfoMemory(ctx context.Context) (string, string) {
	val, err := r.rdb.Info(ctx, "memory").Result()
	if err != nil {
		panic(err)
	}
	rows := strings.Split(val, "\r\n")
	usedMemory := rows[1]
	usedMemoryTerms := strings.Split(usedMemory, ":")
	used := usedMemoryTerms[1]

	usedMemoryHuman := rows[2]
	usedMemoryHumanTerms := strings.Split(usedMemoryHuman, ":")
	usedHuman := usedMemoryHumanTerms[1]
	return used, usedHuman
}

func (r *redisTest) InfoKeyspace(ctx context.Context) {
	val, err := r.rdb.Info(ctx, "keyspace").Result()
	if err != nil {
		panic(err)
	}
	rows := strings.Split(val, "\r\n")
	fmt.Println(rows[1])
}
