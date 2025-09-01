package ws

import (
	"encoding/json"
	"fmt"
	guuid "github.com/google/uuid"
)

type UUID = guuid.UUID

const (
	EventTypeMessageNew        EventTypeSchema = "message_new"
	EventTypeMessageUpdated    EventTypeSchema = "message_updated"
	EventTypeMessageDeleted    EventTypeSchema = "message_deleted"
	EventTypeMessageRestore    EventTypeSchema = "message_restore"
	EventTypeChatCrated        EventTypeSchema = "chat_created"
	EventTypeChatUpdated       EventTypeSchema = "chat_updated"
	EventTypeChatsDeleted      EventTypeSchema = "chats_deleted"
	EventTypeDialogClosed      EventTypeSchema = "dialog_closed"
	EventTypeDialogOpened      EventTypeSchema = "dialog_opened"
	EventTypeDialogAssign      EventTypeSchema = "dialog_assign"
	EventTypeUserOnlineUpdated EventTypeSchema = "user_online_updated"
	EventTypeUserJoinedChat    EventTypeSchema = "user_joined_chat"
	EventTypeUserLeftChat      EventTypeSchema = "user_left_chat"
	EventTypeUserUpdated       EventTypeSchema = "user_updated"
	EventTypeCustomerUpdated   EventTypeSchema = "customer_updated"
	EventTypeBotUpdated        EventTypeSchema = "bot_updated"
	EventTypeChannelUpdated    EventTypeSchema = "channel_updated"
)

func (e *EventSchema) UnmarshalJSON(data []byte) error {
	type EventSchemaRaw struct {
		Data json.RawMessage `json:"data"`
		Meta MetaSchema      `json:"meta"`
		Type string          `json:"type"`
	}

	var raw EventSchemaRaw
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	e.Type = EventTypeSchema(raw.Type)
	e.Meta = raw.Meta

	switch e.Type {
	case EventTypeMessageNew:
		fallthrough
	case EventTypeMessageUpdated:
		fallthrough
	case EventTypeMessageDeleted:
		fallthrough
	case EventTypeMessageRestore:
		var body MessageDataSchema
		if err := json.Unmarshal(raw.Data, &body); err != nil {
			return fmt.Errorf("failed to unmarshal data: %w", err)
		}
		e.Data = body

	case EventTypeChatCrated:
		fallthrough
	case EventTypeChatUpdated:
		var body ChatDataSchema
		if err := json.Unmarshal(raw.Data, &body); err != nil {
			return fmt.Errorf("failed to unmarshal data: %w", err)
		}
		e.Data = body

	case EventTypeChatsDeleted:
		var body ChatsDeletedDataSchema
		if err := json.Unmarshal(raw.Data, &body); err != nil {
			return fmt.Errorf("failed to unmarshal data: %w", err)
		}
		e.Data = body

	case EventTypeDialogClosed:
		fallthrough
	case EventTypeDialogOpened:
		var body DialogDataSchema
		if err := json.Unmarshal(raw.Data, &body); err != nil {
			return fmt.Errorf("failed to unmarshal data: %w", err)
		}
		e.Data = body

	case EventTypeDialogAssign:
		var body DialogAssignDataSchema
		if err := json.Unmarshal(raw.Data, &body); err != nil {
			return fmt.Errorf("failed to unmarshal data: %w", err)
		}
		e.Data = body

	case EventTypeUserOnlineUpdated:
		var body UserOnlineUpdatedDataSchema
		if err := json.Unmarshal(raw.Data, &body); err != nil {
			return fmt.Errorf("failed to unmarshal data: %w", err)
		}
		e.Data = body
	case EventTypeUserJoinedChat:
		var body UserJoinedChatDataSchema
		if err := json.Unmarshal(raw.Data, &body); err != nil {
			return fmt.Errorf("failed to unmarshal data: %w", err)
		}
		e.Data = body
	case EventTypeUserLeftChat:
		var body UserLeftChatDataSchema
		if err := json.Unmarshal(raw.Data, &body); err != nil {
			return fmt.Errorf("failed to unmarshal data: %w", err)
		}
		e.Data = body
	case EventTypeUserUpdated:
		var body UserUpdatedDataSchema
		if err := json.Unmarshal(raw.Data, &body); err != nil {
			return fmt.Errorf("failed to unmarshal data: %w", err)
		}
		e.Data = body

	case EventTypeCustomerUpdated:
		var body CustomerUpdatedDataSchema
		if err := json.Unmarshal(raw.Data, &body); err != nil {
			return fmt.Errorf("failed to unmarshal data: %w", err)
		}
		e.Data = body

	case EventTypeBotUpdated:
		var body BotUpdatedDataSchema
		if err := json.Unmarshal(raw.Data, &body); err != nil {
			return fmt.Errorf("failed to unmarshal data: %w", err)
		}
		e.Data = body

	case EventTypeChannelUpdated:
		var body ChannelUpdatedDataSchema
		if err := json.Unmarshal(raw.Data, &body); err != nil {
			return fmt.Errorf("failed to unmarshal data: %w", err)
		}
		e.Data = body

	default:
		return fmt.Errorf("unknown event type: %s", e.Type)
	}

	return nil
}
