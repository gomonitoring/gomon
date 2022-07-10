package tasks

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gomonitoring/gomon/internal/model"
	log "github.com/sirupsen/logrus"
)

func CallUrl(url string, id uint, threshhold int, resetTime int64) (string, error) {
	var statusCode int
	resp, err := http.Get(url)
	if err != nil {
		statusCode = -1
	} else {
		statusCode = resp.StatusCode
	}
	result := model.CallUrlResult{
		Id:         id,
		StatusCode: statusCode,
		Threshhold: threshhold,
		ResetTime:  resetTime,
		Time:       time.Now().Unix(),
	}
	log.Infoln("call", url)
	return encodeCallResult(&result), nil
}

func encodeCallResult(result *model.CallUrlResult) string {
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
