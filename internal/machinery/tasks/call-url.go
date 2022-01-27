package tasks

import (
	"encoding/json"
)

type CallUrlResult struct {
	Id         uint
	StatusCode int
	Threshhold int
	ResetTime  int64
	Time       int64
}

func CallUrl(url string, id uint, threshhold int, resetTime int64) ([]byte, error) {
	statusCode := 200 // TODO implement sned request and return results
	var time int64
	time = 1
	result := CallUrlResult{
		Id:         id,
		StatusCode: statusCode,
		Threshhold: threshhold,
		ResetTime:  resetTime,
		Time:       time,
	}
	return encodeCallResult(result), nil
}

func encodeCallResult(result CallUrlResult) []byte {
	reqJSON, _ := json.Marshal(result)
	return reqJSON
}

func decodeCallResult(data []byte, output interface{}) error {
	error := json.Unmarshal(data, output)
	return error
}
