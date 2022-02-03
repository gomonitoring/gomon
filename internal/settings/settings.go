package settings

import (
	"fmt"
	"os"
	"strconv"
)

var DefaultResetTime = getStrEnv("DEFAUTL_RESET_TIME", "5m")

var MachineryBroker = getStrEnv("MACHINERY_BROKER", "redis://localhost:6379")
var MachineryResultBackend = getStrEnv("MACHINERY_RESULTS_BACKEND", "redis://localhost:6379")
var LocalWorkerConcurrency = getIntEnv("LOCAL_WORKER_CONCURRENCY", 1)
var MonitoringWorkerConcurrency = getIntEnv("MONITORING_SERVER_CONCURRENCY", 10)
var RedisHost = getStrEnv("REDIS_HOST", "localhost:6379")
var RedisPassword = getStrEnv("REDIS_PASSWORD", "")
var CallUrlsSchedule = getStrEnv("CALL_URL_SCHEDULE", "* * * * *")

func getStrEnv(key string, def string) string {
	val := os.Getenv(key)
	if val == "" {
		return def
	}
	return val
}

func getIntEnv(key string, def int) int {
	val := os.Getenv(key)
	if val == "" {
		return def
	} else {
		ret, err := strconv.Atoi(val)
		if err != nil {
			panic(fmt.Sprintf("Environment variable: %s must be an integer", key))
		}
		return ret
	}
}
