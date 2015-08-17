package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

var logloc string = "/home/ubuntu/.NetworkObserver/Logs"

func init() {
	// Check if the log folder exists
	e, err := exists(logloc)

	if err != nil {
		os.Exit(-1)
	}

	if !e {
		fmt.Println("Creating the log folder at \"" + logloc + "\"")
		os.Mkdir(logloc, 0777)
	}
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)

	// The file/folder exists
	if err == nil {
		return true, nil
	}

	// the file/folder does not exist
	if os.IsNotExist(err) {
		return false, nil
	}

	// There was an error
	return true, err
}

func WriteString(s string) {
	filename := logloc + "/" + time.Now().Format("01-02-2006") + ".log"

	file, err := os.Create(filename)
	defer file.Close()

	if err != nil {
		fmt.Println("There was an error opening the logging file.")
		fmt.Println(err)
		return
	}

	log.SetOutput(file)
	log.Println(s)
}
