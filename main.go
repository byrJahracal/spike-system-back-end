package main

import (
	"flag"
	"pku-class/market/rabbitmq"
	"pku-class/market/router"
)

var mode = flag.Int("mode", 1, "input your mode") //0生产者，1消费者
var port = flag.String("port", "8081", "input your mode")

func main() {
	flag.Parse()
	defer rabbitmq.Connect.Close()
	defer rabbitmq.Channel.Close()
	if *mode == 0 {
		router := router.SetRouter()
		router.Run("0.0.0.0:" + *port)
	} else if *mode == 1 {
		rabbitmq.ConsumerMode()
	}
}
