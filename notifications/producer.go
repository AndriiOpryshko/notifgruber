package notifications

import (
	"github.com/Shopify/sarama"
	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
	"time"
)

func InitProducer(addr string, topic string) *NotificationProducer{
	return &NotificationProducer{
		addr: addr,
		topic: topic,
		maxRetry: 5,
		success: true,
		Nots: make(chan *Notifications),
	}
}

type NotificationProducer struct {
	addr string
	topic string
	maxRetry int
	success bool
	Nots chan *Notifications
}

func (np *NotificationProducer) PushNotifications(nots []*Notification){
	np.Nots <- &Notifications{Nots: nots}
}

func (np *NotificationProducer) Run(){
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = np.maxRetry
	config.Producer.Return.Successes = np.success
	config.Metadata.Retry.Backoff = 2*time.Second
	producer, err := sarama.NewSyncProducer([]string{np.addr}, config)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			panic(err)
		}
	}()

	// Produce messages to topic (asynchronously)
	topic := np.topic
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


