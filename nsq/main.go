package main

import (
	"fmt"
	"time"

	"github.com/nsqio/go-nsq"
)

/*
对于输出我作如下理解，因为初次启动 nsq 相关程序，ConsumerA[test/test-channel-a] 查询 nsqlookupd 主题为 test，返回错误，主题不存在。ConsumerB[test/test-channel-b] 也执行上面的动作。这个时候应该不会创建两个 channel，test-channel-a 和 test-channel-b，也不会创建主题。接下来 Producer 成功连接 nsqd，这个时候会创建 test 主题。等待了一会后 ConsumerB 尝试查询主题成功，进而连接 nsqd，成功建立 test-channel-b，消费已被生产出的 15 条消息，因为 test-channel-a 还未被创建，所以目前已有的消息是不会被复制分发的。接着 ConsumerA 尝试查询主题成功，进而连接 nsqd，成功建立 test-channel-a，接下来的消息都是被复制分发的，两个消费者都能收到

两个 channel 都指定为 test-channel-a 将得到如下输出，可以确定的是多个消费者守在同一个 channel 中，同一条消息将只会被一个消费者处理
*/

// ConsumerHandler 消费者处理者
type ConsumerHandler struct{}

// HandleMessage 处理消息
func (*ConsumerHandler) HandleMessage(msg *nsq.Message) error {
	fmt.Println(string(msg.Body))
	return nil
}

// Producer 生产者
func Producer() {
	producer, err := nsq.NewProducer(":4150", nsq.NewConfig())
	if err != nil {
		fmt.Println("NewProducer", err)
		panic(err)
	}

	i := 1
	for {
		if err := producer.Publish("test1", []byte(fmt.Sprintf("Welcome %d", i))); err != nil {
			fmt.Println("Publish", err)
			panic(err)
		}
		fmt.Printf("PUBLISH(message): Welcome %d\n", i)
		time.Sleep(time.Second * 5)
		i++
	}
}

// ConsumerA 消费者
func ConsumerA() {
	consumer, err := nsq.NewConsumer("test1", "test-channel-a", nsq.NewConfig())
	if err != nil {
		fmt.Println("NewConsumer", err)
		panic(err)
	}

	consumer.AddHandler(&ConsumerHandler{})

	if err := consumer.ConnectToNSQLookupd(":4161"); err != nil {
		fmt.Println("ConnectToNSQLookupd", err)
		panic(err)
	}
}

// ConsumerB 消费者
func ConsumerB() {
	consumer, err := nsq.NewConsumer("test1", "test-channel-b", nsq.NewConfig())
	if err != nil {
		fmt.Println("NewConsumer", err)
		panic(err)
	}

	consumer.AddHandler(&ConsumerHandler{})
	if err := consumer.ConnectToNSQLookupd(":4161"); err != nil {
		fmt.Println("ConnectToNSQLookupd", err)
		panic(err)
	}
}

func main() {
	ConsumerA()
	ConsumerB()
	Producer()
}
