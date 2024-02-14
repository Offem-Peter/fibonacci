package main

import (
	"context"
	"log"
	"math/big"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

type FibonacciService struct {
	current   *big.Int
	prev      *big.Int
	client    *redis.Client
	rwmu      sync.RWMutex
	writeSync time.Duration
}

var ctx = context.Background()

func NewFibonacciService(redisAddr string) *FibonacciService {
	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddr,
		DB:   0,
	})

	pingResult, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic("cannot connect to redis")
	}

	log.Println("connected to redis", pingResult)

	// Check Redis for existing values.
	currentStr, err := rdb.Get(ctx, "current").Result()
	if err == redis.Nil {
		// If not present, start with the default values.
		currentStr = "0"
		rdb.Set(ctx, "current", currentStr, 0)
	}
	if err != nil && err != redis.Nil {
		log.Fatalf("Error getting current from Redis: %v", err)
	}

	prevStr, err := rdb.Get(ctx, "previous").Result()
	if err == redis.Nil {
		// If not present, start with the default values.
		prevStr = "1"
		rdb.Set(ctx, "previous", prevStr, 0)
	}
	if err != nil && err != redis.Nil {
		log.Fatalf("Error getting previous from Redis: %v", err)
	}

	// Parse strings to *big.Int.
	current := new(big.Int)
	current.SetString(currentStr, 10)

	prev := new(big.Int)
	prev.SetString(prevStr, 10)

	// Initialize the FibonacciService with the starting values of the sequence.
	return &FibonacciService{
		current: current,
		prev:    prev,
		client:  rdb,
	}
}

func (f *FibonacciService) syncWithRedis(ctx context.Context) {
	for {
		<-time.After(f.writeSync)

		f.rwmu.RLock()
		current, prev := f.current.String(), f.prev.String()
		f.rwmu.RUnlock()

		// Use a pipeline to reduce round trips to Redis.
		pipe := f.client.Pipeline()
		pipe.Set(ctx, "current", current, 0)
		pipe.Set(ctx, "previous", prev, 0)
		_, err := pipe.Exec(ctx)
		if err != nil {
			log.Printf("Error syncing with Redis: %v", err)
		}
		log.Println("synced with redis")
	}
}

func (f *FibonacciService) Current(ctx context.Context) *big.Int {
	f.rwmu.Lock()
	defer f.rwmu.Unlock()

	current, err := f.client.Get(ctx, "current").Result()
	if err == redis.Nil {
		f.client.Set(ctx, "current", f.current.String(), 0)
	} else if err != nil {
		log.Println(err)
	}

	// If current is already set in Redis, parse it to the big.Int.
	f.current.SetString(current, 10)
	return new(big.Int).Set(f.current)
}

func (f *FibonacciService) Next(ctx context.Context) *big.Int {
	f.rwmu.Lock() // Lock for write operation
	defer f.rwmu.Unlock()

	// Calculate next number
	next := new(big.Int).Add(f.current, f.prev)
	f.prev.Set(f.current)
	f.current.Set(next)

	return new(big.Int).Set(f.current)
}

func (f *FibonacciService) Previous(ctx context.Context) *big.Int {
	f.rwmu.Lock() // Lock for write operation
	defer f.rwmu.Unlock()

	// Calculate previous number if current is not the first two Fibonacci numbers
	if f.current.Cmp(big.NewInt(0)) == 0 {
		return new(big.Int).Set(f.current) // Return 0 if the current number is 0
	} else if f.current.Cmp(big.NewInt(1)) == 0 {
		f.current.Set(big.NewInt(0)) // Set to 0 if the current number is 1
		f.prev.Set(big.NewInt(1))    // The previous number before 1 is 0 in the Fibonacci sequence
	} else {
		// Calculate the previous Fibonacci number
		prev := new(big.Int).Sub(f.current, f.prev)
		f.current.Set(f.prev)
		f.prev.Set(prev)
	}

	return new(big.Int).Set(f.current)
}
