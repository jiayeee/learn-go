package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

type testingUserError string

func (error testingUserError) Error() string {
	return error.Message()
}

func (error testingUserError) Message() string {
	return string(error)
}

func errorPanic(writer http.ResponseWriter, request *http.Request) error {
	panic(123)
}

func errorUserError(writer http.ResponseWriter, request *http.Request) error {
	return testingUserError("user error")
}

func errorNotFound(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrNotExist
}

func errorNoPermission(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrPermission
}

func errorUnknow(writer http.ResponseWriter, request *http.Request) error {
	return errors.New("unknow error")
}

func noError(writer http.ResponseWriter, request *http.Request) error {
	fmt.Fprintln(writer, "no error")
	return nil
}

var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	{errorPanic, 500, "Internal Server Error"},
	{errorUserError, 400, "user error"},
	{errorNotFound, 404, "Not Found"},
	{errorNoPermission, 403, "Forbidden"},
	{errorUnknow, 500, "Internal Server Error"},
	{noError, 200, "no error"},
}

/**
 * 模拟 http 请求
 */
func TestErrorWrapper(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)

		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "http://www.imooc.com", nil)
		f(response, request)

		verifyResposne(response.Result(), tt.code, tt.message, t)
	}
}

/**
 * 起服务器处理 http 请求
 */
func TestErrorWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)

		server := httptest.NewServer(http.HandlerFunc(f))
		resp, _ := http.Get(server.URL)

		verifyResposne(resp, tt.code, tt.message, t)
	}
}

func verifyResposne(resp *http.Response, expectedCode int, expectedMessage string, t *testing.T) {

	bytes, _ := ioutil.ReadAll(resp.Body)
	body := strings.Trim(string(bytes), "\n")

	if resp.StatusCode != expectedCode || body != expectedMessage {
		t.Errorf("expect(%d, %s), got (%d, %s)\n", expectedCode, expectedMessage, resp.StatusCode, body)
	}
}
