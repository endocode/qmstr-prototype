package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/QMSTR/qmstr-prototype/pkg/master"
	"github.com/QMSTR/qmstr-prototype/pkg/model"
	"github.com/spf13/pflag"
)

func main() {
	var verbose bool
	var printVersion bool
	pflag.BoolVarP(&verbose, "verbose", "v", false, "enable verbose log output")
	pflag.BoolVarP(&printVersion, "version", "V", false, "print version and exit")
	pflag.Parse()

	var infoWriter io.Writer
	if verbose {
		infoWriter = os.Stdout
	} else {
		infoWriter = ioutil.Discard
	}

	info := log.New(infoWriter, "INFO: ", log.Ldate|log.Ltime)
	logr := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	info.Printf("quartermaster master process starting")
	defer info.Printf("quartermaster master process exiting")
	if printVersion {
		// --version prints the version and then exits:
		fmt.Println("Quartermaster master 0.0.1")
		return
	}
	// default: run the master server until a quit requests comes in
	modl := model.NewModel()
	// TODO: also react to a SIGTERM/SIGKILL
	logr.Printf(<-master.StartHTTPServer(info, logr, modl))
}
