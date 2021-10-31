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

// CallError reflects an error which has occurred during an external request.
// It contains the cause, an error message, and the appropriate HTTP status code to return.
type CallError struct {
	ResponseCode int
	Message      string
	Cause        error
}

// Error implements the error interface
func (err *CallError) Error() string {
	return err.Message
}

// NewCallError creates a new CallError and prints it to the server log
func NewCallError(responseCode int, message string, err error) *CallError {
	log.Printf("%s; Cause: %s\n", message, err)
	return &CallError{responseCode, message, err}
}

// GetCall executes a GET HTTP call to a given URL, resource, and path parameter.
func GetCall(baseUrl, resourcePath, pathParam string, responseObject interface{}) *CallError {
	requestUrl, err := joinPath(baseUrl, resourcePath, pathParam)
	if err != nil {
		return NewCallError(http.StatusInternalServerError, "Error parsing URL", err)
	}

	log.Printf("Executing GET %s", requestUrl)
	resp, err := http.Get(requestUrl)
	if err != nil {
		return NewCallError(http.StatusInternalServerError, "Error during HTTP call", err)
	}
	log.Printf("Received response with status code %d", resp.StatusCode)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return NewCallError(mapResponseCode(resp.StatusCode), fmt.Sprintf("Response %d from GET %s", resp.StatusCode, requestUrl), nil)
	}

	return readResponse(resp, responseObject)
}

// PostCall executes a POST HTTP call to a given URL, resource, and request body.
func PostCall(baseUrl string, resourcePath string, requestObject interface{}, responseObject interface{}) *CallError {
	requestUrl, err := joinPath(baseUrl, resourcePath)
	if err != nil {
		return NewCallError(http.StatusInternalServerError, "Error parsing URL", err)
	}

	jsonStr, err := json.Marshal(requestObject)
	if err != nil {
		return NewCallError(http.StatusInternalServerError, "Error during request object serialisation", err)
	}

	log.Printf("Executing POST %s; Request Body: %s", requestUrl, jsonStr)
	resp, err := http.Post(requestUrl, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		return NewCallError(http.StatusInternalServerError, "Error during HTTP call", err)
	}
	log.Printf("Received response with status code %d", resp.StatusCode)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return NewCallError(mapResponseCode(resp.StatusCode), fmt.Sprintf("Response %d from POST %s", resp.StatusCode, requestUrl), nil)
	}

	return readResponse(resp, responseObject)
}

// joinPath joins a base URL and multiple path segments
func joinPath(baseUrl string, paths ...string) (string, error) {
	u, err := url.Parse(baseUrl)
	if err != nil {
		return "", err
	}
	allPaths := append([]string{u.Path}, paths...)
	u.Path = path.Join(allPaths...)
	return u.String(), nil
}

// readResponse reads an HTTP response body, parses it into a given struct and validates it
func readResponse(resp *http.Response, object interface{}) *CallError {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return NewCallError(http.StatusBadGateway, "Response body could not be read", err)
	}
	err = json.Unmarshal(body, object)
	if err != nil {
		return NewCallError(http.StatusBadGateway, "Response body could not be parsed into JSON", err)
	}
	log.Println(object)
	_, err = govalidator.ValidateStruct(object)
	if err != nil {
		return NewCallError(http.StatusBadGateway, "Response object is invalid", err)
	}
	return nil
}

// mapResponseCode maps an external call HTTP error response code
// to one to be returned by the endpoint.
// * 404 -> 404 (Not found)
// * 5xx -> 502 (External server error)
// * 4xx -> 500 (External client error)
func mapResponseCode(code int) int {
	if code == http.StatusNotFound {
		return code
	}
	if code < http.StatusInternalServerError {
		return http.StatusInternalServerError
	}
	return http.StatusBadGateway
}
