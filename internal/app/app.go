package app

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/core-go/health"
	_ "github.com/core-go/mongo"
	"github.com/core-go/mq"
	"github.com/core-go/mq/kafka"
	"github.com/core-go/mq/log"
	"github.com/core-go/mq/pubsub"
	"github.com/core-go/mq/validator"
	"reflect"

	v "github.com/go-playground/validator/v10"
)

type ApplicationContext struct {
	HealthHandler *health.Handler
	Receive       func(ctx context.Context, handler func(context.Context, *mq.Message, error) error)
	Handler       *mq.Handler
}

type KafkaCustomWriter struct {
	Send func(ctx context.Context, data []byte, attributes map[string]string) (string, error)
}

func NewKafkaCustomWriter(send func(ctx context.Context, data []byte, attributes map[string]string) (string, error)) *KafkaCustomWriter {
	return &KafkaCustomWriter{Send: send}
}

func NewUserValidator() validator.Validator {
	val := validator.NewDefaultValidator()
	val.CustomValidateList = append(val.CustomValidateList, validator.CustomValidate{Fn: CheckActive, Tag: "active"})
	return val
}

func CheckActive(fl v.FieldLevel) bool {
	return fl.Field().Bool()
}

func (w *KafkaCustomWriter) Write(ctx context.Context, model interface{}) error {
	data := new(bytes.Buffer)
	err := json.NewEncoder(data).Encode(model.(*User))
	if err != nil {
		return err
	}
	_, err = w.Send(ctx, data.Bytes(), nil)
	return err
}

func NewApp(ctx context.Context, root Root) (*ApplicationContext, error) {
	log.Initialize(root.Log)
	//db,er1 := mongo.Setup(ctx, root.Mongo)
	//if er1 != nil {
	//	log.Error(ctx, "Cannot connect to MongoDB: Error: " + er1.Error())
	//	return nil, er1
	//}

	logError := log.ErrorMsg
	var logInfo func(context.Context, string)
	if log.IsErrorEnable() {
		logInfo = log.InfoMsg
	}

	receive, er2 := pubsub.NewSubscriberByConfig(ctx, root.Sub.Subscriber, true)
	if er2 != nil {
		log.Error(ctx, "Cannot create a new receiver. Error: "+er2.Error())
		return nil, er2
	}

	kafkaWriter, er3 := kafka.NewWriterByConfig(root.KafkaWriter)
	if er3 != nil {
		log.Error(ctx, "Cannot new a new sender. Error:"+er3.Error())
		return nil, er3
	}

	writer := NewKafkaCustomWriter(kafkaWriter.Write)
	var userType = reflect.TypeOf(User{})
	checker := validator.NewErrorChecker(NewUserValidator().Validate)
	val := mq.NewValidator(userType, checker.Check)

	kafkaChecker := kafka.NewKafkaHealthChecker(root.KafkaWriter.Brokers, "kafka_producer")
	receiveChecker := pubsub.NewPubHealthChecker("pubsub_subcriber", receive.Client, root.Sub.Subscriber.SubscriptionId)
	var healthHandler *health.Handler
	var handler *mq.Handler
	if root.Pub != nil {
		sender, er3 := pubsub.NewPublisherByConfig(ctx, *root.Pub)
		if er3 != nil {
			log.Error(ctx, "Cannot new a new sender. Error:"+er3.Error())
			return nil, er3
		}
		retryService := mq.NewRetryService(sender.Publish, logError, logInfo)
		handler = mq.NewHandlerByConfig(root.Sub.Config, writer.Write, &userType, retryService.Retry, val.Validate, nil, logError, logInfo)
		senderChecker := pubsub.NewPubHealthChecker("pubsub_publisher", sender.Client, root.Pub.TopicId)
		healthHandler = health.NewHandler(kafkaChecker, receiveChecker, senderChecker)
	} else {
		healthHandler = health.NewHandler(kafkaChecker, receiveChecker)
		handler = mq.NewHandlerWithRetryConfig(writer.Write, &userType, val.Validate, root.Retry, true, logError, logInfo)
	}
	return &ApplicationContext{
		HealthHandler: healthHandler,
		Receive:       receive.Subscribe,
		Handler:       handler,
	}, nil
}
