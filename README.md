# logger

## Как использовать?

```
package main

import (
	"your_project/logger"
)

func main() {
	log := logger.New()

	log.Debug("Это debug сообщение")
	log.Info("Это info сообщение")
	log.Warn("Это warn сообщение")
	log.Error("Это error сообщение")
}
```
