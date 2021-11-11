# log24go
Log24go is a basic logging library for Golang.

Inspired by the Golang logger standard library.

Project is mainly used to increase knowledge in Golang.

**WORK IS IN PROGRESS**

## Integration

To integrate EliasCha/log24go into your own Golang application, please see below example.

```go
package main

import (
	"os"

	"github.com/EliasCha/log24go"
)

func App() {
    log := &log24go.Logger{Out: os.Stdout}
    log.Info("Info")
    log.Warn("Warn")
    log.Error("Error")
    log.Debug("Debug")
}
```

## Supported Levels

Currently there is only four levels supported:
* Info - Used to inform the user about the application.
* Warn - Used to warn the user about the health of the application.
* Error - Used to show Error in the application that crashes the application functionality.
* Debug - Used for debugging purposes.
