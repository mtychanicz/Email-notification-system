package sender

import (
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-redisstream/pkg/redisstream"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
	"github.com/redis/go-redis/v9"
)

type sendNotificationMessagePayload struct {
	Message    string   `json:"message"`
	Sender     string   `json:"sender"`
	Recipients []string `json:"recipients"`
}
type retryMessagePayload struct {
	Message    string   `json:"message"`
	Sender     string   `json:"sender"`
	Recipients []string `json:"recipients"`
	Status     string   `json:"status"` //value shloud be either success or failure
	Cause      string   `json:"cause"`  //only populated if status == "failure"
}
type EmailNotificationHandler struct {
}

func (h *EmailNotificationHandler) Handle(msg *message.Message) ([]*message.Message, error) {
	retryMsgPayload := retryMessagePayload{}
	err := json.Unmarshal(msg.Payload, &retryMsgPayload)
	status := "success"
	cause := ""
	if err != nil {
		status = "failure"
		cause = err.Error()
	} else if retryMsgPayload.Status != status {
		// err := //Send email here you doofus
		// if err != nil {
		// 	status = "failure"
		// 	cause = err.Error()
		// }
	}
	retryMsgPayload.Status = status
	retryMsgPayload.Cause = cause
	if marshaled, err := json.Marshal(retryMsgPayload); err == nil {
		return message.Messages{message.NewMessage(msg.UUID, marshaled)}, nil
	} else {
		return nil, err
	}
}

func GetRouterConfig(
	redisUrl, messageTopic, retryTopic, consumerGroup string,
	logger watermill.LoggerAdapter,
) (*message.Router, error) {
	if logger == nil {
		logger = watermill.NewStdLogger(false, false)
	}
	router, err := message.NewRouter(*&message.RouterConfig{}, logger)

	router.AddPlugin(plugin.SignalsHandler)

	router.AddMiddleware(
		middleware.CorrelationID,

		middleware.Recoverer,
	)

	emailNotificationHandler := EmailNotificationHandler{}
	sub, err := newRedisSubscriber(newRedisClient(redisUrl), consumerGroup, logger)
	if err != nil {
		panic(err)
	}
	pub, err := newRedisPublisher(newRedisClient(redisUrl), logger)
	if err != nil {
		panic(err)
	}
	router.AddHandler(
		"socket_notification_handler",
		messageTopic,
		sub,
		retryTopic,
		pub,
		emailNotificationHandler.Handle,
	)
	return router, err
}
func newRedisSubscriber(client *redis.Client, consumerGroup string, logger watermill.LoggerAdapter) (*redisstream.Subscriber, error) {
	return redisstream.NewSubscriber(
		redisstream.SubscriberConfig{
			Client:        client,
			ConsumerGroup: consumerGroup,
		}, logger,
	)
}
func newRedisPublisher(client *redis.Client, logger watermill.LoggerAdapter) (*redisstream.Publisher, error) {
	return redisstream.NewPublisher(
		redisstream.PublisherConfig{
			Client: client,
		}, logger,
	)
}
func newRedisClient(redisUrl string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: redisUrl,
	})
}
