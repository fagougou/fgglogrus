package fgglogrus

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func parentPath(path string) string {
	index := strings.LastIndex(path, "/")
	return path[:index]
}

func AppName() string {
	modDirPath, _ := filepath.Abs(".")
	file, err := os.Open(modDirPath + "/go.mod")
	for err != nil && len(modDirPath) > 1 {
		modDirPath = parentPath(modDirPath)
		file, err = os.Open(modDirPath + "/go.mod")
	}

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var appName string
	for scanner.Scan() {
		line := scanner.Text()
		if index := strings.LastIndex(line, "module"); index != -1 {
			appName = line[index+7:]
			break
		}
	}

	return strings.Trim(appName, " ")
}

func Filename(appName string) string {
	dt := time.Now()
	prefix := ""
	if os.Getenv("ENV") == "production" {
		prefix = "/app/"
	}
	return prefix + "logs/" + appName + "-" + dt.Format("2006-01-02") + ".log"
}
