package external

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"

	"github.com/asaskevich/govalidator"
)

type HttpError struct {
	ResponseCode int
	Message      string
	Cause        error
}

func (err *HttpError) Error() string {
	return err.Message
}

func NewServiceError(responseCode int, message string, err error) *HttpError {
	log.Printf("%s; Cause: %s\n", message, err)
	return &HttpError{responseCode, message, err}
}

func GetCall(baseUrl, resourcePath, pathParam string, responseObject interface{}) *HttpError {
	requestUrl, err := joinPath(baseUrl, resourcePath, pathParam)
	if err != nil {
		return NewServiceError(500, "Error parsing URL", err)
	}

	log.Printf("Executing GET %s", requestUrl)
	resp, err := http.Get(requestUrl)
	if err != nil {
		log.Printf("Error during HTTP call")
		return &HttpError{500, "Error during HTTP call", err}
	}
	log.Printf("Received response with status code %d", resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		return NewServiceError(mapResponseCode(resp.StatusCode), fmt.Sprintf("Response %d from GET %s", resp.StatusCode, requestUrl), nil)
	}

	defer resp.Body.Close()
	return readResponse(resp, responseObject)
}

func PostCall(baseUrl string, resourcePath string, requestObject interface{}, responseObject interface{}) *HttpError {
	requestUrl, err := joinPath(baseUrl, resourcePath)
	if err != nil {
		return NewServiceError(500, "Error parsing URL", err)
	}

	jsonStr, err := json.Marshal(requestObject)
	if err != nil {
		return NewServiceError(500, "Error during request object serialisation", err)
	}

	log.Printf("Executing POST %s; Request Body: %s", requestUrl, jsonStr)
	resp, err := http.Post(requestUrl, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		return NewServiceError(500, "Error during HTTP call", err)
	}
	log.Printf("Received response with status code %d", resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		return NewServiceError(mapResponseCode(resp.StatusCode), fmt.Sprintf("Response %d from POST %s", resp.StatusCode, requestUrl), nil)
	}

	defer resp.Body.Close()
	return readResponse(resp, responseObject)
}

func joinPath(baseUrl string, paths ...string) (string, error) {
	u, err := url.Parse(baseUrl)
	if err != nil {
		return "", err
	}
	allPaths := append([]string{u.Path}, paths...)
	u.Path = path.Join(allPaths...)
	return u.String(), nil
}

func readResponse(resp *http.Response, object interface{}) *HttpError {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return NewServiceError(http.StatusBadGateway, "Response body could not be read", err)
	}
	err = json.Unmarshal(body, object)
	if err != nil {
		return NewServiceError(http.StatusBadGateway, "Response body could not be read", err)
	}
	log.Println(object)
	_, err = govalidator.ValidateStruct(object)
	if err != nil {
		return NewServiceError(http.StatusBadGateway, "Response object is invalid", err)
	}
	return nil
}

func mapResponseCode(code int) int {
	if code == http.StatusNotFound {
		return code
	}
	if code < 500 {
		return http.StatusInternalServerError
	}
	return http.StatusBadGateway
}
