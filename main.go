package mian

import (
	"os"
	"os/signal"
	"syscall"
)

func init() {

}

func main() {

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

}
