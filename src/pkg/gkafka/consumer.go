package gkafka

import (
	"fmt"
	"log"
	"sync"
    "time"
	"os"
    "os/signal"
	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
)

func Consumer(topics []string) error {
    var wg = &sync.WaitGroup{}
    wg.Add(1)
    //广播式消费：消费者1
    go clusterConsumer(wg, Address, topics, "group-1")
    //广播式消费：消费者2
    // go clusterConsumer(wg, Address, topic, "group-2")
 
    wg.Wait()
    return nil
}

func clusterConsumer(wg *sync.WaitGroup, brokers, topics []string, groupId string) error {
	defer wg.Done()
	config := cluster.NewConfig()
	config.Net.MaxOpenRequests = 1
	config.Net.DialTimeout = 2 * time.Second
	config.Net.ReadTimeout = 10 * time.Second
	config.Consumer.Return.Errors = true
    config.Group.Return.Notifications = true
    config.Consumer.Offsets.Initial = sarama.OffsetOldest

    consumer, err := cluster.NewConsumer(brokers, groupId, topics, config)
    if err != nil {
        fmt.Println("cluster.NewConsumer err, message=%s", err)
        return err
    }
    defer consumer.Close()

    // trap SIGINT to trigger a shutdown
    signals := make(chan os.Signal, 1)
    signal.Notify(signals, os.Interrupt)
 
    // consume errors
    go func() {
        for err := range consumer.Errors() {
            log.Printf("%s:Error: %s\n", groupId, err.Error())
        }
    }()
 
    // consume notifications
    go func() {
        for ntf := range consumer.Notifications() {
            log.Printf("%s:Rebalanced: %+v \n", groupId, ntf)
        }
    }()

    var successes int
    Loop:
    for {
        select {
        case msg, ok := <-consumer.Messages():
            if ok {
                fmt.Fprintf(os.Stdout, "%s:%s/%d/%d\t%s\t%s\n", groupId, msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
                consumer.MarkOffset(msg, "")  // mark message as processed
                successes++
            }
        case <-signals:
            break Loop
        }
    }
    fmt.Fprintf(os.Stdout, "%s consume %d messages \n", groupId, successes)

    return nil
}