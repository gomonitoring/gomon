package tasks

import (
	"encoding/base64"
	"encoding/json"
)

type CallUrlResult struct {
	Id         uint
	StatusCode int
	Threshhold int
	ResetTime  int64
	Time       int64
}

func CallUrl(url string, id uint, threshhold int, resetTime int64) (string, error) {
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

func encodeCallResult(result CallUrlResult) string {
	resJSON, _ := json.Marshal(result)
	return base64.StdEncoding.EncodeToString(resJSON)
}

func decodeCallResult(b64data string, output interface{}) error {
	decodedstg, err := base64.StdEncoding.DecodeString(b64data)
	if err != nil {
		return err
	}
	data := []byte(decodedstg)
	err = json.Unmarshal(data, output)
	return err
}
