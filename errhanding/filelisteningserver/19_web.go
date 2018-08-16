package main

import (
	"github.com/gopm/modules/log"
	"learn-go/errhanding/filelisteningserver/filelisting"
	"net/http"
	_ "net/http/pprof"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

/**
 * 统一错误处理逻辑
 */
func errWrapper(handler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {

		defer func() {
			if r := recover(); r != nil {
				log.Error("panic: %v\n", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)

		if userError, ok := err.(userError); ok {
			message := userError.Message()
			http.Error(writer, message, http.StatusBadRequest)
			return
		}

		if err != nil {
			log.Warn("Error handling request : %s", err.Error())

			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}

			http.Error(writer, http.StatusText(code), code)
		}
	}
}

type userError interface {
	error
	Message() string
}

func main() {

	http.HandleFunc("/", errWrapper(filelisting.HandleFileListing))

	e := http.ListenAndServe(":8888", nil)
	if e != nil {
		panic(e)
	}
}
