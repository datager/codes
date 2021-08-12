package clients

import (
	"time"

	"codes/gocodes/dg/utils/log"

	"github.com/Shopify/sarama"
	"github.com/bsm/sarama-cluster"
	"github.com/golang/glog"
	"github.com/pkg/errors"
)

type KafkaProducerImpl struct {
	Brokers  []string
	producer sarama.SyncProducer
	topic    string
}

func NewKafkaProducer(topic string, brokers []string) (KafkaProducer, error) {
	instance := &KafkaProducerImpl{}
	instance.Brokers = brokers
	instance.topic = topic

	config := sarama.NewConfig()
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Compression = sarama.CompressionNone
	config.Producer.MaxMessageBytes = 104857600
	config.Consumer.Fetch.Max = 104857600
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	var err error
	instance.producer, err = sarama.NewSyncProducer(instance.Brokers, config)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return instance, nil
}

func (this *KafkaProducerImpl) Send(data []byte) error {

	_, _, err := this.producer.SendMessage(&sarama.ProducerMessage{
		Topic: this.topic,
		Value: sarama.ByteEncoder(data),
	})

	if err != nil {
		return errors.WithStack(err)
	}

	log.Infoln("send success")

	return nil
}

func (this *KafkaProducerImpl) Close() error {
	return this.producer.Close()
}

type KafkaConsumerImpl struct {
	Brokers     []string
	topic       string
	consumer    *cluster.Consumer
	messageChan chan []byte
	quit        chan int
}

func NewKafkaConsumer(topic string, group string, brokers []string) (KafkaConsumer, error) {
	instance := &KafkaConsumerImpl{
		messageChan: make(chan []byte, 1000),
		quit:        make(chan int),
	}
	instance.Brokers = brokers
	instance.topic = topic

	// init (custom) config, enable errors and notifications
	kafkaConfig := cluster.NewConfig()
	kafkaConfig.Version = sarama.V1_0_0_0
	kafkaConfig.Consumer.Return.Errors = true
	kafkaConfig.Group.Return.Notifications = true
	kafkaConfig.Consumer.Fetch.Max = 104857600

	consumer, err := cluster.NewConsumer(brokers, group, []string{topic}, kafkaConfig)
	if err != nil {
		glog.Errorln(" create kafka consumer err:", err)
		return nil, err
	}

	go func(consumer *cluster.Consumer) {

		glog.Infoln("kafka face inferry input start ")
		defer func() {
			glog.Infoln("face inferry input closed ")
		}()
		for {
			select {
			case <-instance.quit:
				consumer.Close()
				return
			case err := <-consumer.Errors():
				glog.Warningln("kafka err:", err)
				<-time.After(time.Second)
				break
			case info := <-consumer.Notifications():
				glog.Infoln("kafka info:", info)
				break
			case kafkaMessage := <-consumer.Messages():
				instance.messageChan <- kafkaMessage.Value
				break

			}
		}
	}(consumer)

	instance.consumer = consumer
	return instance, nil
}

func (this *KafkaConsumerImpl) GetMessage() ([]byte, error) {
	return <-this.messageChan, nil
}

func (this *KafkaConsumerImpl) Close() error {

	this.quit <- 1

	return nil
}
