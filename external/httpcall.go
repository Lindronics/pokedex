package external

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
)

func GetCall(baseUrl string, resourcePath string) []byte {
	resp, err := http.Get(joinPath(baseUrl, resourcePath))
	if err != nil {
		log.Fatal("Response error")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Response body corrupted")
	}
	return body
}

func PostCall(baseUrl string, resourcePath string, requestBody interface{}) []byte {
	jsonStr, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatal("Error during serialisation")
	}
	resp, err := http.Post(joinPath(baseUrl, resourcePath), "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Fatal("Response error")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Response body corrupted")
	}
	return body
}

func joinPath(baseUrl, resourcePath string) string {
	u, err := url.Parse(baseUrl)
	if err != nil {
		log.Fatal("Error parsing URL")
	}
	u.Path = path.Join(u.Path, resourcePath)
	return u.String()
}
