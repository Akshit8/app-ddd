package graceful

import (
	"os"
	"os/signal"
	"syscall"
)

func Run() {
	interupt := make(chan os.Signal, 1)
	defer close(interupt)

	signal.Notify(interupt, syscall.SIGTERM, syscall.SIGINT)

	<-interupt
}
