package notifications

import (
	"github.com/Shopify/sarama"
	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
	"time"
)

type ProducerConfig interface {
	GetAddr() string
	GetTopic() string
	GetMaxRetry() int
	GetIsSuccess() bool
	GetRetryBackoffSec() int

}

func InitProducer(conf ProducerConfig) *NotificationProducer{
	return &NotificationProducer{
		conf: conf,
		Nots: make(chan *Notifications),
	}
}

type NotificationProducer struct {
	conf ProducerConfig
	Nots chan *Notifications
}

func (np *NotificationProducer) PushNotifications(nots []*Notification){
	np.Nots <- &Notifications{Nots: nots}
}

func (np *NotificationProducer) Run(){
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = np.conf.GetMaxRetry()
	config.Producer.Return.Successes = np.conf.GetIsSuccess()
	config.Metadata.Retry.Backoff = time.Second * time.Duration(np.conf.GetRetryBackoffSec())
	producer, err := sarama.NewSyncProducer([]string{np.conf.GetAddr()}, config)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			panic(err)
		}
	}()

	// Produce messages to topic (asynchronously)
	topic := np.conf.GetTopic()
	for {
		nots := <- np.Nots
		notsProto, err := proto.Marshal(nots)
		if err != nil {

			log.WithFields(log.Fields{
				"err": err,
			}).Error("Cannot marshal notification in proto3")
			continue
		}
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.ByteEncoder(notsProto),
		}
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			panic(err)
		}
		log.WithFields(log.Fields{
			"topic": err,
			"partition": partition,
			"offset": offset,
		}).Info("Notification is stored in topic")
	}
}


