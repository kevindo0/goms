package gkafka

import (
	"fmt"
	"time"
	"log"
	"os"
	"github.com/Shopify/sarama"
)

var Address = []string{"10.6.124.21:9092"}

func SyncProducer(topic string) error {
	config := sarama.NewConfig()
	config.Net.MaxOpenRequests = 3
	config.Net.DialTimeout = 5 * time.Second
	config.Net.ReadTimeout = 10 * time.Second
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Timeout = 5 * time.Second
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(Address, config)
	if err != nil {
        fmt.Println("sarama.NewSyncProducer err, message=%s", err)
        return err
    }
    defer producer.Close()
    srcValue := "sync: this is a message. index=%d"
    for i := 0; i < 10; i++ {
        value := fmt.Sprintf(srcValue, i)
        msg := &sarama.ProducerMessage{
            Topic: topic,
            Value: sarama.StringEncoder(value),
        }
        part, offset, err := producer.SendMessage(msg)
        if err != nil {
            log.Printf("send message(%s) err=%s \n", value, err)
        }else {
            fmt.Fprintf(os.Stdout, value + "发送成功，partition=%d, offset=%d \n", part, offset)
        }
        time.Sleep(time.Second)
    }
    return nil
}