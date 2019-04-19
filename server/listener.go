package server

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func Start() int {
	channel := make(chan os.Signal)
	signal.Notify(channel, os.Interrupt, os.Kill)

	fmt.Println("server start,press ctrl + c quit")

	<-channel

	return 0;
}
func TimeOutFn() {
	timeout := make (chan bool, 1)
	go func() {
		time.Sleep(1e9) // sleep one second
		timeout <- true
	}()
	ch := make (chan int)
	select {
	case <- ch:
	case <- timeout:
		fmt.Println("timeout!")
	}
}
//func ListenActiveMQ()  {
//	connect ,err := stomp.Dial("tcp", "localhost:61613");
//	if err != nil {
//		fmt.Println(err)
//	}
//}

type Listener_working struct {
	listen_port int

}