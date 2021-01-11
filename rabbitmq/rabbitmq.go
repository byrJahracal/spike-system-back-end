package rabbitmq

import (
	"pku-class/market/backstage/handler"
	eh "pku-class/market/error-handler"

	"github.com/streadway/amqp"
)

var Connect *amqp.Connection
var Queue amqp.Queue
var Channel *amqp.Channel

func init() {
	var user string = "wyh"
	var pwd string = "123456"
	var host string = "182.92.69.9" //182.92.69.9
	var port string = "5672"
	url := "amqp://" + user + ":" + pwd + "@" + host + ":" + port + "/"
	var err error
	Connect, err = amqp.Dial(url)
	eh.ErrorHandler(err, "消息队列服务器连接失败！", "")
	Channel, err = Connect.Channel()
	eh.ErrorHandler(err, "打开消息队列失败！", "")

	Queue, err = Channel.QueueDeclare(
		"market",
		false,
		false,
		false,
		false,
		nil,
	)
	eh.ErrorHandler(err, "初始化队列失败！", "消息队列创建成功！")
}

func ConsumerMode() {
	msgs, err := Channel.Consume(
		Queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	eh.ErrorHandler(err, "创建消费者失败！", "")
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			handler.FlashPostHandler(d.Body)
		}
	}()
	<-forever
}

func Publish(body []byte) {
	err := Channel.Publish(
		"",
		Queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	eh.ErrorHandler(err, "发布消息失败！", "")
}
