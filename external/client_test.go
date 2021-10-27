// Tests the external call client.
// I decided not to include unit tests for GetCall and PostCall, as they are
// indirectly covered by the service tests.
// In a production environment, it would make sense to test them separately too.
package external

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

type testStruct struct {
	ValueOne string           `json:"value_one" valid:"required"`
	ValueTwo nestedTestStruct `json:"value_two" valid:"required"`
}

type nestedTestStruct struct {
	ValueOne string `json:"value_one" valid:"required"`
}

// TestReadResponse tests whether readResponse throws an error upon invalid response bodies
func TestReadResponse(t *testing.T) {
	tables := []struct {
		responseBody  string
		errorExpected bool
	}{
		{`invalid json`, true},
		{`{}`, true},
		{`{"value_one": "test"}`, true},
		{`{"value_one": "test", "value_two": 1}`, true},
		{`{"value_one": 5, "value_two": {"value_one": "test"}}`, true},
		{`{"value_one": "test", "value_two": {"value_one": "test"}}`, false},
	}
	for _, table := range tables {
		body := ioutil.NopCloser(strings.NewReader(table.responseBody))
		resp := http.Response{Body: body}
		var object testStruct
		err := readResponse(&resp, &object)
		if (err != nil) != table.errorExpected {
			t.Errorf("Should have thrown error: %t", table.errorExpected)
		}
	}

}

// TestJoinPath tests whether joinPath forms valid URLs, assuming correct input for baseUrl
func TestJoinPath(t *testing.T) {
	tables := []struct {
		base          string
		paths         []string
		expected      string
		errorExpected bool
	}{
		{"https://test.api/v1", []string{"a", "b", "c/d"}, "https://test.api/v1/a/b/c/d", false},
		{"https://test.api/v1", []string{"a", "//b"}, "https://test.api/v1/a/b", false},
	}
	for _, table := range tables {
		url, err := joinPath(table.base, table.paths...)
		if (err != nil) != table.errorExpected {
			t.Errorf("Should have thrown error: %t", table.errorExpected)
		}
		if url != table.expected {
			t.Errorf("URL %s did not match expected %s", url, table.expected)
		}
	}
}

// TestMapResponseCode tests whether mapResponseCode returns the correct status codes
func TestMapResponseCode(t *testing.T) {
	tables := []struct {
		in  int
		out int
	}{
		{404, 404},
		{400, 500},
		{403, 500},
		{500, 502},
	}
	for _, table := range tables {
		out := mapResponseCode(table.in)
		if out != table.out {
			t.Errorf("Expected %d but got %d", table.out, out)
		}
	}
}
