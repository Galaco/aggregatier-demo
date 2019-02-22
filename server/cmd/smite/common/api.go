package common

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	apiParamLanguageEN        = 1
	apiParamLanguageGER       = 2
	apiParamLanguageFR        = 3
	apiParamLanguageSPA       = 7
	apiParamLanguageSPA_LATAM = 9
	apiParamLanguagePORTU     = 10
	apiParamLanguageRU        = 11
	apiParamLanguagePOL       = 12
	apiParamLanguageTURK      = 13
)

const (
	apiEndpointBaseUri        = "http://api.smitegame.com/smiteapi.svc/"
	apiEndpointRequestSession = "createSession"
	apiEndpointRequestGods    = "getgods"
)

type Api struct {
	devId, authKey string
}

func (api *Api) CreateSession() (*Session, error) {
	timestamp := api.getTimestampFormatted()

	hash := api.createHash(apiEndpointRequestSession, timestamp)
	url := apiEndpointBaseUri +
		apiEndpointRequestSession +
		"Json/" +
		api.devId +
		"/" +
		hash +
		"/" +
		timestamp

	session := &Session{}
	if err := json.Unmarshal([]byte(api.getRequest(url)), session); err != nil {
		return nil, err
	}
	return session, nil
}

func (api *Api) GetGods(sessionId string) ([]God, error) {
	log.Println(sessionId)
	timestamp := api.getTimestampFormatted()

	hash := api.createHash(apiEndpointRequestGods, timestamp)
	url := apiEndpointBaseUri +
		apiEndpointRequestGods +
		"Json/" +
		api.devId +
		"/" +
		hash +
		"/" +
		sessionId +
		"/" +
		timestamp +
		"/" +
		strconv.FormatInt(int64(apiParamLanguageEN), 10)

	gods := make([]God, 0)
	if err := json.Unmarshal([]byte(api.getRequest(url)), &gods); err != nil {
		return nil, err
	}
	return gods, nil
}

func (api *Api) getTimestamp() time.Time {
	return time.Now().UTC()
}

func (api *Api) getTimestampFormatted() string {
	return api.getTimestamp().Format("20060102150405") //YmdHis
}

func (api *Api) createHash(route string, timestamp string) string {
	hasher := md5.New()
	hasher.Write([]byte(api.devId + strings.ToLower(route) + api.authKey + timestamp))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (api *Api) getRequest(url string) string {
	response, err := http.Get(url)
	if err != nil {
		return "[]"
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return "[]"
		}
		return string(contents)
	}
}

func NewApi(devId string, authKey string) *Api {
	return &Api{
		devId:   devId,
		authKey: authKey,
	}
}
