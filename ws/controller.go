package ws

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/lerenn/asyncapi-codegen/pkg/extensions"
	"net/http"
)

type Controller struct {
	url     string
	headers http.Header
	options []ControllerOption
}

type Option func(controller *Controller) error

func NewController(url, token string, options ...Option) (*AppController, error) {
	ctrl := &Controller{
		url: url,
		headers: http.Header{
			"X-Bot-Token": []string{token},
		},
		options: make([]ControllerOption, 0),
	}

	for _, option := range options {
		if err := option(ctrl); err != nil {
			return nil, fmt.Errorf("error while applying option: %w", err)
		}
	}

	appController, err := NewAppController(ctrl, ctrl.options...)
	if err != nil {
		panic(err)
	}

	return appController, nil
}

func (c Controller) Publish(ctx context.Context, channel string, mw extensions.BrokerMessage) error {
	return nil
}

func (c Controller) Subscribe(ctx context.Context, channel string) (extensions.BrokerChannelSubscription, error) {
	conn, _, err := websocket.DefaultDialer.Dial(
		fmt.Sprintf("%s?%s", c.url, channel),
		c.headers,
	)
	if err != nil {
		return extensions.BrokerChannelSubscription{}, err
	}

	messages := make(chan extensions.AcknowledgeableBrokerMessage)
	cancel := make(chan any)

	go func() {
		defer close(cancel)
		defer close(messages)

		for {
			_, message, readErr := conn.ReadMessage()

			if readErr != nil {
				return
			}

			messages <- extensions.NewAcknowledgeableBrokerMessage(
				extensions.BrokerMessage{Payload: message},
				NoopAcknowledgementHandler{},
			)
		}
	}()

	sub := extensions.NewBrokerChannelSubscription(messages, cancel)

	sub.WaitForCancellationAsync(func() {
		if err := conn.Close(); err != nil {
			panic(err)
		}
	})

	return sub, nil
}

type NoopAcknowledgementHandler struct {
}

func (k NoopAcknowledgementHandler) AckMessage() {
}

func (k NoopAcknowledgementHandler) NakMessage() {
}
