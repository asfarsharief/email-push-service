package listners

import (
	emailprocessor "email-push-service/internal/emailProcessor"
	"email-push-service/pkg/logger"
	"fmt"

	"github.com/nats-io/nats.go"
)

type NatsObject struct {
	url   string
	conn  *nats.Conn
	topic string
}

func NewNatsListner(url string) ListnerInterface {
	if url == "default" {
		url = nats.DefaultURL
	}
	return &NatsObject{
		url: url,
	}
}

func (no *NatsObject) InitializeListner() error {
	logger.Info("Starting connnection for NATS queue")
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		logger.Error("Starting connnection failed", err)
		return err
	}
	no.conn = nc
	return nil
}

func (no *NatsObject) Listen(topic string) {
	fmt.Println("Starting listening")
	_, err := no.conn.Subscribe("jobs", func(msg *nats.Msg) {
		logger.Info("Received job:", string(msg.Data))
		emailprocessor.ProcessRequest(msg.Data)
	})

	if err != nil {
		logger.Error("Error subscribing:", err)
		return
	}

	fmt.Println("Listening for jobs... Press Ctrl+C to exit.")
	select {}
}
