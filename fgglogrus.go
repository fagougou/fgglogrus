package fgglogrus

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	FggLogrus      *logrus.Logger
	allFilePointer []*os.File
	allFileName    []string
)

type formatter struct {
	fields logrus.Fields
	lf     logrus.Formatter
}

// Format satisfies the logrus.Formatter interface.
func (f *formatter) Format(e *logrus.Entry) ([]byte, error) {
	for k, v := range f.fields {
		e.Data[k] = v
	}
	return f.lf.Format(e)
}

func contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func closeFiles(files []*os.File, log *logrus.Logger) {
	var wg sync.WaitGroup

	for _, file := range files {
		wg.Add(1)
		go func() {
			err := file.Close()
			if err != nil {
				log.Fatal(err)
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func initLoggerToFile(log *logrus.Logger, appName string) {
	// open file
	os.Mkdir("logs", 0755)

	filename := Filename(appName)

	LOG_FILE, _ := filepath.Abs(filename)

	logFile, err := os.OpenFile(LOG_FILE, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)

	if err != nil {
		log.Fatal(err)
	}

	if !contains(allFileName, filename) {
		allFileName = allFileName[:0]
		closeFiles(allFilePointer, log)

		allFilePointer = append(allFilePointer, logFile)
		allFileName = append(allFileName, filename)
	}
	// defer logFile.Close()
	log.SetOutput(logFile)
}

func init() {
	log := logrus.New()

	// get app name
	appName := AppName()
	environment := os.Getenv("ENV")

	log.SetFormatter(&formatter{
		fields: logrus.Fields{
			"app": appName,
			"env": environment,
		},
		lf: &logrus.JSONFormatter{},
	})

	if environment == "production" {
		initLoggerToFile(log, appName)
	} else {
		log.SetOutput(os.Stdout)
	}

	FggLogrus = log
}
