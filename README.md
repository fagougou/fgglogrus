# fgglogrus
a [logrus](https://github.com/sirupsen/logrus) wrapper with some configs

# Feature
- Add fields: `app`, `env`
- Log to file when `os.Getenv("ENV") === production`
- Logfiles sperated by date.

# Usage

```go
package main

import (
	. "github.com/fagougou/fgglogrus"
	"github.com/sirupsen/logrus"
)

func main() {

	FggLogrus.Info("A group of walrus emerges from the ocean")

	FggLogrus.WithFields(logrus.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	FggLogrus.Error("The ice breaks!")
	FggLogrus.Error("The ice errrrrrr!")

}

```
