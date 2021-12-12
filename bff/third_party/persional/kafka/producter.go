package lib_kafka

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	kLog "github.com/go-kratos/kratos/v2/log"
)

var Address = []string{"10.130.138.164:9092", "10.130.138.164:9093", "10.130.138.164:9094"}

// var _ kLog.Logger = (*KafkaLogger)(nil)

type KafkaLogger struct {
	log  *log.Logger
	pool *sync.Pool
	p    *sarama.SyncProducer
}

// NewKafkaLogger new a logger with writer.
func NewKafkaLogger(w io.Writer) kLog.Logger {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second
	p, err := sarama.NewSyncProducer(Address, config)
	if err != nil {
		p = nil
		log.Printf("sarama.NewSyncProducer err, message=%s \n", err)
	}
	return &kafkaLogger{
		log: log.New(w, "", 0),
		pool: &sync.Pool{
			New: func() interface{} {
				return new(bytes.Buffer)
			},
		},
		p: &p,
	}
}

// Log print the kv pairs log.
func (l *KafkaLogger) Log(level kLog.Level, keyvals ...interface{}) error {
	if len(keyvals) == 0 {
		return nil
	}
	if (len(keyvals) & 1) == 1 {
		keyvals = append(keyvals, "KEYVALS UNPAIRED")
	}
	buf := l.pool.Get().(*bytes.Buffer)
	buf.WriteString(level.String())
	for i := 0; i < len(keyvals); i += 2 {
		_, _ = fmt.Fprintf(buf, " %s=%v", keyvals[i], keyvals[i+1])
	}
	if l.p != nil {
		err := l.send(l.p, buf.String())
		if err != nil {
			_ = l.log.Output(4, buf.String()) //nolint:gomnd
		}
	} else {
		_ = l.log.Output(4, buf.String()) //nolint:gomnd
	}
	buf.Reset()
	l.pool.Put(buf)
	return nil
}

func (l *KafkaLogger) Close() {
	if l.p != nil {
		(*l.p).Close()
	}
}

func (l *KafkaLogger) send(p *sarama.SyncProducer, value string) error {
	topic := "log"
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(value),
	}
	part, offset, err := (*p).SendMessage(msg)
	if err == nil {
		fmt.Fprintf(os.Stdout, value+"发送成功，partition=%d, offset=%d \n", part, offset)
	}
	return err
}
