package fgglogrus

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func AppName() string {
	modPath, _ := filepath.Abs("./go.mod")
	file, err := os.Open(modPath)
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

func Filename() string {
	dt := time.Now()
	return dt.Format("2006-01-02") + ".log"
}
