package tasks

import (
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/go-redis/redis"
	"github.com/gomonitoring/gomon/internal/machinery/worker"
	"github.com/gomonitoring/gomon/internal/model"
	"github.com/gomonitoring/gomon/internal/settings"
	"github.com/gomonitoring/gomon/internal/storage"
)

func CollectResults(results ...string) error {
	calls := make([]model.CallUrlResult, len(results))
	for i, data := range results {
		result := model.CallUrlResult{}
		decodeCallResult(data, &result)
		calls[i] = result
		if result.StatusCode > 499 || result.StatusCode < 100 {
			handleFaliure(result.Id, result.Threshhold, result.Time, result.ResetTime, result.StatusCode)
		}
	}
	var st storage.LocalWorker = worker.GetLocalWorkerStorage()
	err := st.SaveCallResults(calls)
	if err != nil {
		log.Fatalln("could not collect url call results", err)
	}
	log.Infoln("url call results collected")
	return nil
}

func handleFaliure(id uint, threshhold int, ts int64, resetTime int64, statusCode int) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     settings.RedisHost,
		Password: settings.RedisPassword,
		DB:       0,
	})
	prefix_base := 36
	url_key := strconv.FormatUint(uint64(id), prefix_base)
	if countKeysWithPrefix(redisClient, url_key+"_") >= threshhold {
		iter := redisClient.Scan(0, url_key+"_*", 0).Iterator()
		for iter.Next() {
			redisClient.Del(iter.Val())
		}
		sig := tasks.Signature{
			Name:       "create_alert",
			RoutingKey: "local",
			Args: []tasks.Arg{
				{
					Type:  "uint",
					Value: id,
				}, {
					Type:  "int64",
					Value: ts,
				},
			},
		}
		_, err := GetMachineryServer().SendTask(&sig)
		if err != nil {
			log.Fatal("could not push create alert task to queue.")
		}
	} else {
		event_key := url_key + "_" + strconv.FormatInt(ts, prefix_base)
		redisClient.Set(event_key, statusCode, time.Duration(resetTime))
	}
}

func countKeysWithPrefix(client *redis.Client, prefix string) int {
	var cursor uint64 = 0
	var n int
	for {
		var keys []string
		var err error
		keys, cursor, err = client.Scan(cursor, prefix+"*", 0).Result()
		if err != nil {
			panic(err)
		}
		n += len(keys)
		if cursor == 0 {
			break
		}
	}
	return n
}
