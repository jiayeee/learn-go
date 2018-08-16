package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const urlPrefix string = "/list/"

type uError string

func (error uError) Error() string {
	return error.Message()
}

func (error uError) Message() string {
	return string(error)
}

func HandleFileListing(writer http.ResponseWriter, request *http.Request) error {

	if strings.Index(request.URL.Path, urlPrefix) != 0 {
		return uError("path must start with " + urlPrefix)
	}

	path := request.URL.Path[len(urlPrefix):]
	file, e := os.Open(path)
	if e != nil {
		//http.Error(writer, e.Error(), http.StatusInternalServerError)
		return e
	}
	defer file.Close()

	bytes, e := ioutil.ReadAll(file)
	if e != nil {
		//panic(e)
		return e
	}

	writer.Write(bytes)

	return nil
}
