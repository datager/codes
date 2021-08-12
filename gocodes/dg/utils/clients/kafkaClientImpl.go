package clients

// import (
// 	"github.com/Shopify/sarama"
// 	"codes/gocodes/dg/utils/log"
// 	"github.com/pkg/errors"
// )

// type KafkaProducerImpl struct {
// 	Brokers  []string
// 	producer sarama.SyncProducer
// 	topic    string
// }

// func NewKafkaProducer(topic string, brokers []string) (KafkaProducer, error) {
// 	instance := &KafkaProducerImpl{}
// 	instance.Brokers = brokers
// 	instance.topic = topic

// 	config := sarama.NewConfig()
// 	config.Producer.Retry.Max = 5
// 	config.Producer.RequiredAcks = sarama.WaitForLocal
// 	config.Producer.Compression = sarama.CompressionNone
// 	config.Producer.MaxMessageBytes = 104857600
// 	config.Consumer.Fetch.Max = 104857600
// 	config.Producer.Return.Successes = true
// 	config.Producer.Return.Errors = true

// 	var err error
// 	instance.producer, err = sarama.NewSyncProducer(instance.Brokers, config)
// 	if err != nil {
// 		return nil, errors.WithStack(err)
// 	}

// 	return instance, nil
// }

// func (this *KafkaProducerImpl) Send(data []byte) error {

// 	_, _, err := this.producer.SendMessage(&sarama.ProducerMessage{
// 		Topic: this.topic,
// 		Value: sarama.ByteEncoder(data),
// 	})

// 	if err != nil {
// 		return errors.WithStack(err)
// 	}

// 	log.Infoln("send success")

// 	return nil
// }

// func (this *KafkaProducerImpl) Close() error {
// 	return this.producer.Close()
// }

// type KafkaConsumerImpl struct {
// 	Brokers   []string
// 	topic     string
// 	consumer  sarama.Consumer
// 	partition sarama.PartitionConsumer
// }

// func NewKafkaConsumer(topic string, brokers []string) (KafkaConsumer, error) {
// 	instance := &KafkaConsumerImpl{}
// 	instance.Brokers = brokers

// 	var err error
// 	config := sarama.NewConfig()
// 	config.Consumer.Return.Errors = true
// 	instance.consumer, err = sarama.NewConsumer(instance.Brokers, config)
// 	if err != nil {
// 		return nil, errors.WithStack(err)
// 	}

// 	instance.partition, err = instance.consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)

// 	return instance, nil
// }

// func (this *KafkaConsumerImpl) GetMessage() ([]byte, error) {

// 	select {
// 	case msg := <-this.partition.Messages():
// 		return msg.Value, nil
// 	case err := <-this.partition.Errors():
// 		return nil, errors.WithStack(err.Err)
// 	}
// }

// func (this *KafkaConsumerImpl) Close() error {
// 	err := this.partition.Close()
// 	if err != nil {
// 		return errors.WithStack(err)
// 	}

// 	err = this.consumer.Close()
// 	if err != nil {
// 		return errors.WithStack(err)
// 	}

// 	return nil
// }
