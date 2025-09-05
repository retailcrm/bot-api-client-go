# Bot API Client

A Go client library for interacting with the Message Gateway Bot API. This client provides a comprehensive set of methods for managing bots, channels, chats, dialogs, messages, and more.

## Installation

```bash
go get github.com/retailcrm/bot-api-client-go
```

## Usage

### Initializing the Client

```go
package main

import (
    "github.com/retailcrm/bot-api-client-go"
    "log"
)

func main() {
    client, err := bot_api_client.NewClientWithResponses(
        "https://mg-s1.retailcrm.pro/api/bot/v1/",
        bot_api_client.WithBotToken("BOT_TOKEN"),
    )

    if err != nil {
        log.Fatalf("Error creating client: %v", err)

        return
    }
}
```

### REST API Examples

#### Sending a Message

```go
response, err := client.SendMessageWithResponse(
    context.Background(),
    bot_api_client.SendMessageJSONRequestBody{},
)

if err != nil {
    log.Fatalf("Error sending message: %v", err)
}

if response.JSONDefault != nil {
    log.Printf("Error: %s", response.JSONDefault.Errors[0])
}

if response.JSON200 != nil {
    log.Printf("Message id: %d", response.JSON200.MessageId)
}
```

### WebSocket Support

```go
package main

import (
    "context"
    "github.com/retailcrm/bot-api-client-go/ws"
    "log"
)

func main() {
    controller, err := ws.NewController(
        "wss://mg-s1.retailcrm.pro/api/bot/v1/ws",
        "BOT_TOKEN",
    )

    if err != nil {
        log.Fatalf("Error creating client: %v", err)

        return
    }

    err = controller.SubscribeToReceiveEventsOperation(
        context.Background(),
        ws.EventsChannelParameters{Events: "message_new"},
        func(ctx context.Context, msg ws.EventMessageFromEventsChannel) error {
            switch data := msg.Payload.Data.(type) {
            case ws.MessageDataSchema:
                log.Printf("New event `%s` content: `%s`", msg.Payload.Type, *data.Message.Content)
            }
            return nil
        },
    )

    if err != nil {
        log.Fatalf("Error subscribing: %v", err)

        return
    }

    select {}
}
```

