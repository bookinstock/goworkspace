package redisx

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func RunExample1() {

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}

var ctx = context.Background()

var rdb *redis.Client

func Init() {
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

func Run() {

	Init()

	set("k", "v", time.Minute)

	v := get("k")
	fmt.Println("v=", v)
}

func set(k string, v interface{}, ttl time.Duration) {
	rdb.Set(ctx, k, v, ttl).Err()
}

func get(k string) string {
	v, _ := rdb.Get(ctx, k).Result()
	return v
}

// func foo() {
// 	rdb := redis.NewClient(&redis.Options{
// 		Addr:     "localhost:6379",
// 		Password: "",
// 		DB:       0,
// 	})

// 	// SET key value EX 10 NX
// 	set, err := rdb.SetNX(ctx, "key", "value", 10*time.Second).Result()

// 	// SET key value keepttl NX
// 	set, err = rdb.SetNX(ctx, "key", "value", redis.KeepTTL).Result()

// 	// SORT list LIMIT 0 2 ASC
// 	vals, err := rdb.Sort(ctx, "list", &redis.Sort{Offset: 0, Count: 2, Order: "ASC"}).Result()

// 	// // ZRANGEBYSCORE zset -inf +inf WITHSCORES LIMIT 0 2
// 	// vals, err = rdb.ZRangeByScoreWithScores(ctx, "zset", &redis.ZRangeBy{
// 	// 	Min:    "-inf",
// 	// 	Max:    "+inf",
// 	// 	Offset: 0,
// 	// 	Count:  2,
// 	// }).Result()

// 	// // ZINTERSTORE out 2 zset1 zset2 WEIGHTS 2 3 AGGREGATE SUM
// 	// vals, err = rdb.ZInterStore(ctx, "out", &redis.ZStore{
// 	// 	Keys:    []string{"zset1", "zset2"},
// 	// 	Weights: []int64{2, 3},
// 	// }).Result()

// 	// // EVAL "return {KEYS[1],ARGV[1]}" 1 "key" "hello"
// 	// vals, err := rdb.Eval(ctx, "return {KEYS[1],ARGV[1]}", []string{"key"}, "hello").Result()

// 	// custom command
// 	res, err := rdb.Do(ctx, "set", "key", "value").Result()

// 	fmt.Println("set=", set, "vals", vals, "res=", res, "err=", err)
// }

// // redis mock

// func redis_mock() {

// }

// func NewsInfoForCache(redisDB *redis.Client, newsID int) (info string, err error) {
// 	cacheKey := fmt.Sprintf("news_redis_cache_%d", newsID)
// 	info, err = redisDB.Get(ctx, cacheKey).Result()
// 	if err == redis.Nil {
// 		// info, err = call api()
// 		info = "test"
// 		err = redisDB.Set(ctx, cacheKey, info, 30*time.Minute).Err()
// 	}
// 	return
// }

// func TestNewsInfoForCache(t *testing.T) {
// 	db, mock := redismock.NewClientMock()

// 	newsID := 123456789
// 	key := fmt.Sprintf("news_redis_cache_%d", newsID)

// 	// mock ignoring `call api()`

// 	mock.ExpectGet(key).RedisNil()
// 	mock.Regexp().ExpectSet(key, `[a-z]+`, 30*time.Minute).SetErr(errors.New("FAIL"))

// 	_, err := NewsInfoForCache(db, newsID)
// 	if err == nil || err.Error() != "FAIL" {
// 		t.Error("wrong error")
// 	}

// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Error(err)
// 	}
// }

// // distirbuted lock

// func main() {
// 	// Connect to redis.
// 	client := redis.NewClient(&redis.Options{
// 		Network: "tcp",
// 		Addr:    "127.0.0.1:6379",
// 	})
// 	defer client.Close()

// 	// Create a new lock client.
// 	locker := redislock.New(client)

// 	ctx := context.Background()

// 	// Try to obtain lock.
// 	lock, err := locker.Obtain(ctx, "my-key", 100*time.Millisecond, nil)
// 	if err == redislock.ErrNotObtained {
// 		fmt.Println("Could not obtain lock!")
// 	} else if err != nil {
// 		log.Fatalln(err)
// 	}

// 	// Don't forget to defer Release.
// 	defer lock.Release(ctx)
// 	fmt.Println("I have a lock!")

// 	// Sleep and check the remaining TTL.
// 	time.Sleep(50 * time.Millisecond)
// 	if ttl, err := lock.TTL(ctx); err != nil {
// 		log.Fatalln(err)
// 	} else if ttl > 0 {
// 		fmt.Println("Yay, I still have my lock!")
// 	}

// 	// Extend my lock.
// 	if err := lock.Refresh(ctx, 100*time.Millisecond, nil); err != nil {
// 		log.Fatalln(err)
// 	}

// 	// Sleep a little longer, then check.
// 	time.Sleep(100 * time.Millisecond)
// 	if ttl, err := lock.TTL(ctx); err != nil {
// 		log.Fatalln(err)
// 	} else if ttl == 0 {
// 		fmt.Println("Now, my lock has expired!")
// 	}

// }

// // cache

// type Object struct {
// 	Str string
// 	Num int
// }

// func Example_basicUsage() {
// 	ring := redis.NewRing(&redis.RingOptions{
// 		Addrs: map[string]string{
// 			"server1": ":6379",
// 			"server2": ":6380",
// 		},
// 	})

// 	mycache := cache.New(&cache.Options{
// 		Redis:      ring,
// 		LocalCache: cache.NewTinyLFU(1000, time.Minute),
// 	})

// 	ctx := context.TODO()
// 	key := "mykey"
// 	obj := &Object{
// 		Str: "mystring",
// 		Num: 42,
// 	}

// 	if err := mycache.Set(&cache.Item{
// 		Ctx:   ctx,
// 		Key:   key,
// 		Value: obj,
// 		TTL:   time.Hour,
// 	}); err != nil {
// 		panic(err)
// 	}

// 	var wanted Object
// 	if err := mycache.Get(ctx, key, &wanted); err == nil {
// 		fmt.Println(wanted)
// 	}

// 	// Output: {mystring 42}
// }

// func Example_advancedUsage() {
// 	ring := redis.NewRing(&redis.RingOptions{
// 		Addrs: map[string]string{
// 			"server1": ":6379",
// 			"server2": ":6380",
// 		},
// 	})

// 	mycache := cache.New(&cache.Options{
// 		Redis:      ring,
// 		LocalCache: cache.NewTinyLFU(1000, time.Minute),
// 	})

// 	obj := new(Object)
// 	err := mycache.Once(&cache.Item{
// 		Key:   "mykey",
// 		Value: obj, // destination
// 		Do: func(*cache.Item) (interface{}, error) {
// 			return &Object{
// 				Str: "mystring",
// 				Num: 42,
// 			}, nil
// 		},
// 	})
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(obj)
// 	// Output: &{mystring 42}
// }

// // rate limit
// func ExampleNewLimiter() {
// 	ctx := context.Background()
// 	rdb := redis.NewClient(&redis.Options{
// 		Addr: "localhost:6379",
// 	})
// 	// redis option params
// 	// redis option pool params
// 	_ = rdb.FlushDB(ctx).Err()

// 	limiter := redis_rate.NewLimiter(rdb)
// 	res, err := limiter.Allow(ctx, "project:123", redis_rate.PerSecond(10))
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("allowed", res.Allowed, "remaining", res.Remaining)
// 	// Output: allowed 1 remaining 9
// }
