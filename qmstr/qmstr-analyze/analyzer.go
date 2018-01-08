package analyze

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	// Logger is the default logger.
	Logger *log.Logger
)

func initLogging(debug bool) {
	var infoWriter io.Writer
	if debug {
		infoWriter = os.Stdout
	} else {
		infoWriter = ioutil.Discard
	}
	Logger = log.New(infoWriter, "", log.Ldate|log.Ltime)
}