package fgglogrus

import (
	"os"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

var oldPointers []*os.File

func TestWriteFile(t *testing.T) {
	log := logrus.New()
	appName1 := "test1"

	initLoggerToFile(log, appName1)

	assert.Contains(t, allFileName[0], appName1)

	for _, f := range allFilePointer {
		oldPointers = append(oldPointers, f)
		_, err := f.WriteString("writes\n")
		assert.Nil(t, err)
	}
}
func TestCloseFile(t *testing.T) {
	log := logrus.New()
	appName2 := "test2"

	initLoggerToFile(log, appName2)

	time.Sleep(time.Millisecond * 500)

	for _, oldf := range oldPointers {
		_, err := oldf.WriteString("writes\n")
		assert.NotNil(t, err)
	}

	// remove file created by unit test
	defer os.RemoveAll("logs")

}
