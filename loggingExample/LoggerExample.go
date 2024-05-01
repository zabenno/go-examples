package main

import (
	"fmt"
	"log/slog"
	"loggingExample/utils"
	"sync"
)

func logloop(logger *slog.Logger, message string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 20; i++ {
		logger.Info(message + fmt.Sprintf(" This used logger:%p", &logger))
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	logger1 := utils.SetupLogger("DEBUG")
	logger2 := utils.SetupLogger("DEBUG")

	logger1.Info(fmt.Sprintf("Logger 1 at memory address: %p", &logger1))
	logger1.Info(fmt.Sprintf("Logger 2 at memory address: %p", &logger2))

	go logloop(&logger1, "This is logger 1", &wg)
	go logloop(&logger2, "This is logger 2", &wg)

	wg.Wait()

}
