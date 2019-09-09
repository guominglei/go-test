package main

import "os"
import "log"
import "io"
import "io/ioutil"

var Trace *log.Logger
var Info *log.Logger
var Warning *log.Logger
var Error *log.Logger

func init() {
	file, err := os.OpenFile("errors.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file", err)
	}

	Trace = log.New(ioutil.Discard, "Trace: ", log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Llongfile)

	Warning = log.New(os.Stdout, "Warning: ", log.Ldate|log.LUTC|log.Llongfile)

	Error = log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {

	Trace.Println("tranceback stack info")
	Info.Println("this is info ")
	Warning.Println("Warning  info aaa a ")
	Error.Println("ERRRRRRRRRR")
}
