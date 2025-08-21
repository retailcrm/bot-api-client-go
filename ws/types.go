package ws

import (
	"encoding/json"
	"fmt"
	guuid "github.com/google/uuid"
)

type UUID = guuid.UUID

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
	case "message_new":
		fallthrough
	case "message_updated":
		fallthrough
	case "message_deleted":
		fallthrough
	case "message_restore":
		var body MessageDataSchema
		if err := json.Unmarshal(raw.Data, &body); err != nil {
			return fmt.Errorf("failed to unmarshal data: %w", err)
		}
		e.Data = body

	case "chat_created":
		fallthrough
	case "chat_updated":
		var body ChatDataSchema
		if err := json.Unmarshal(raw.Data, &body); err != nil {
			return fmt.Errorf("failed to unmarshal data: %w", err)
		}
		e.Data = body

	case "chats_deleted":
		var body ChatsDeletedDataSchema
		if err := json.Unmarshal(raw.Data, &body); err != nil {
			return fmt.Errorf("failed to unmarshal data: %w", err)
		}
		e.Data = body

	case "dialog_closed":
		fallthrough
	case "dialog_opened":
		var body DialogDataSchema
		if err := json.Unmarshal(raw.Data, &body); err != nil {
			return fmt.Errorf("failed to unmarshal data: %w", err)
		}
		e.Data = body

	case "dialog_assign":
		var body DialogAssignDataSchema
		if err := json.Unmarshal(raw.Data, &body); err != nil {
			return fmt.Errorf("failed to unmarshal data: %w", err)
		}
		e.Data = body

	case "user_online_updated":
		var body UserOnlineUpdatedDataSchema
		if err := json.Unmarshal(raw.Data, &body); err != nil {
			return fmt.Errorf("failed to unmarshal data: %w", err)
		}
		e.Data = body
	case "user_joined_chat":
		var body UserJoinedChatDataSchema
		if err := json.Unmarshal(raw.Data, &body); err != nil {
			return fmt.Errorf("failed to unmarshal data: %w", err)
		}
		e.Data = body
	case "user_left_chat":
		var body UserLeftChatDataSchema
		if err := json.Unmarshal(raw.Data, &body); err != nil {
			return fmt.Errorf("failed to unmarshal data: %w", err)
		}
		e.Data = body
	case "user_updated":
		var body UserUpdatedDataSchema
		if err := json.Unmarshal(raw.Data, &body); err != nil {
			return fmt.Errorf("failed to unmarshal data: %w", err)
		}
		e.Data = body

	case "customer_updated":
		var body CustomerUpdatedDataSchema
		if err := json.Unmarshal(raw.Data, &body); err != nil {
			return fmt.Errorf("failed to unmarshal data: %w", err)
		}
		e.Data = body

	case "bot_updated":
		var body BotUpdatedDataSchema
		if err := json.Unmarshal(raw.Data, &body); err != nil {
			return fmt.Errorf("failed to unmarshal data: %w", err)
		}
		e.Data = body

	case "channel_updated":
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
