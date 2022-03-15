package bucket_api

import (
	"fmt"
	"context"
	"io"
	"net/http"
	"encoding/gob"
	"log"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/S3"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/segmentio/kafka-go"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type kafkaMsg struct {
	topic	string
	partition	int
	key, value	[]byte
	Offset		int
}
type retrieve struct {
	pipeReader		io.PipeReader
	pipeWriter		io.PipeWriter
}

func DecodeAndSend(reader *bufio.Reader) error {
	// Calling Pipe method
	ret := retrieve{}
    ret.pipeReader, ret.pipeWriter = io.Pipe()
  
    go func() {
		defer pipeWriter.Close()
		//remember to cancel the context also 
        
		decoder := gob.NewDecoder(reader)
		interruptChan := make(chan os.Signal, 1)
		signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

		for {
			select {
			case <- interruptChan:
				log.Info("DecodeAndSend got an interrupt")
				break
			case <- context.Done():
				log.Info("DecodeAndSend got an interrupt")
				break
			default:
				msg := kafkaMsg{}
				err = decoder.Decode(&msg)
				if err != nil {
					log.Info("DecodeAndSend Got an error while decoding msg, %v", err)
					break
				}
				//fmt.Fprint(pipeWriter, msg)
				ret.pipeWriter.Write([]byte(msg))
				log.Info("DecodeAndSend: Msg written to pipe, msg length %v", len(msg.Value))
			}
		}
		decoder = nil  
    }()
}

func DownlaodVideoFromKafka() {

	kafkaBrokerUrl := ""
	brokers := strings.Split(kafkaBrokerUrl, ",")	
	kafkaconfig := kafka.ReaderConfig{
		Brokers:         brokers,
		GroupID:         kafkaClientId,
		Topic:           kafkaTopic,
		MinBytes:        10e3,            // 10KB
		MaxBytes:        10e6,            // 10MB
		MaxWait:         1 * time.Second, // Maximum amount of time to wait for new data to come when fetching batches of messages from kafka.
		ReadLagInterval: -1,
	}

	//reader := kafka.NewReader(kafkaconfig)
	//defer reader.Close()
	pipeReader, pipeWriter := io.Pipe()
	defer pipeWriter.close()
	vidReader := bufio.NewReader(pipeReader)
	encoder := gob.NewEncoder(pipeWriter)

	go func() {
		if err := DownloadAndSend(vidReader); err != nil {
			log.Infof("Could not download and send video, %v", err)
		}
	}

	interruptChan := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <- interruptChan:
			log.Info("DownlaodVideoFromKafka got an interrupt")
			err = nil
			break
		
		case <- context.Done():
			log.Info("DownlaodVideoFromKafka got a context closure")
			err = nil
			break
		default:
			m, err := reader.ReadMessage(context.Background())
			if err != nil {
				log.Error().Msgf("error while receiving message: %s", err.Error())
				continue
			} else if msg != nil {		
				log.Infof("message at topic/partition/offset %v/%v/%v: %s\n", m.Topic, m.Partition, m.Offset, string(m.value))

				if err = encoder.Encode(m.value); err != nil {
					err = fmt.Errorf("Error while encoding with gob: %v", err)
					break
					}
			} else {
				log.Infof("Messages are not coming")
			}
		}
	}

	if err := reader.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
}





