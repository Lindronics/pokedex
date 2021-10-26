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

func GetCall(baseUrl, resourcePath, pathParam string) ([]byte, error) {
	requestUrl := joinPath(baseUrl, resourcePath, pathParam)

	log.Printf("Executing GET %s", requestUrl)
	resp, err := http.Get(requestUrl)
	if err != nil {
		log.Printf("Error during HTTP call")
		return nil, err
	}
	log.Printf("Received response with status code %d", resp.StatusCode)

	defer resp.Body.Close()
	return readResponse(resp)
}

func PostCall(baseUrl string, resourcePath string, requestBody interface{}) ([]byte, error) {
	jsonStr, err := json.Marshal(requestBody)
	if err != nil {
		log.Printf("Error during request object serialisation")
		return nil, err
	}

	requestUrl := joinPath(baseUrl, resourcePath)

	log.Printf("Executing POST %s; Request Body: %s", requestUrl, jsonStr)
	resp, err := http.Post(requestUrl, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Printf("Error during HTTP call")
		return nil, err
	}
	log.Printf("Received response with status code %d", resp.StatusCode)

	defer resp.Body.Close()
	return readResponse(resp)
}

func joinPath(baseUrl string, paths ...string) string {
	u, err := url.Parse(baseUrl)
	if err != nil {
		log.Printf("Error parsing URL")
	}
	allPaths := append([]string{u.Path}, paths...)
	u.Path = path.Join(allPaths...)
	return u.String()
}

func readResponse(resp *http.Response) ([]byte, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Response body could not be read")
	}
	return body, err
}
