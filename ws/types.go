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
	EventTypeMessageRestored   EventTypeSchema = "message_restored"
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

const (
	WaitingLevelWarning WaitingLevelSchema = "warning"
	WaitingLevelDanger  WaitingLevelSchema = "danger"
	WaitingLevelNone    WaitingLevelSchema = "none"
)

const (
	MessageTypeText    MessageTypeSchema = "text"
	MessageTypeSystem  MessageTypeSchema = "system"
	MessageTypeCommand MessageTypeSchema = "command"
	MessageTypeOrder   MessageTypeSchema = "order"
	MessageTypeProduct MessageTypeSchema = "product"
	MessageTypeFile    MessageTypeSchema = "file"
	MessageTypeImage   MessageTypeSchema = "image"
	MessageTypeAudio   MessageTypeSchema = "audio"
)

const (
	MessageScopeUndefined MessageScopeSchema = "undefined"
	MessageScopePublic    MessageScopeSchema = "public"
	MessageScopePrivate   MessageScopeSchema = "private"
)

const (
	MessageStatusUndefined MessageStatusSchema = "undefined"
	MessageStatusReceived  MessageStatusSchema = "received"
	MessageStatusSending   MessageStatusSchema = "sending"
	MessageStatusSent      MessageStatusSchema = "sent"
	MessageStatusFailed    MessageStatusSchema = "failed"
	MessageStatusSeen      MessageStatusSchema = "seen"
)

const (
	ResponsibleTypeUser ResponsibleTypeSchema = "user"
	ResponsibleTypeBot  ResponsibleTypeSchema = "bot"
)

const (
	UserTypeUser     UserTypeSchema = "user"
	UserTypeBot      UserTypeSchema = "bot"
	UserTypeCustomer UserTypeSchema = "customer"
	UserTypeChannel  UserTypeSchema = "channel"
)

const (
	MessageOrderStatusCodeNew        MessageOrderStatusCodeSchema = "new"
	MessageOrderStatusCodeApproval   MessageOrderStatusCodeSchema = "approval"
	MessageOrderStatusCodeAssembling MessageOrderStatusCodeSchema = "assembling"
	MessageOrderStatusCodeDelivery   MessageOrderStatusCodeSchema = "delivery"
	MessageOrderStatusCodeComplete   MessageOrderStatusCodeSchema = "complete"
	MessageOrderStatusCodeCancel     MessageOrderStatusCodeSchema = "cancel"
)

const (
	MessageFileKindNone  MessageFileKindSchema = "none"
	MessageFileKindImage MessageFileKindSchema = "image"
	MessageFileKindVideo MessageFileKindSchema = "video"
	MessageFileKindFile  MessageFileKindSchema = "file"
	MessageFileKindAudio MessageFileKindSchema = "audio"
)

const (
	MessageErrorCodeUnknown           MessageErrorCodeSchema = "unknown"
	MessageErrorCodeNetworkError      MessageErrorCodeSchema = "network_error"
	MessageErrorCodeMalformedResponse MessageErrorCodeSchema = "malformed_response"
	MessageErrorCodeAsyncSendTimeout  MessageErrorCodeSchema = "async_send_timeout"
	MessageErrorCodeGeneral           MessageErrorCodeSchema = "general"
	MessageErrorCodeCustomerNotExists MessageErrorCodeSchema = "customer_not_exists"
	MessageErrorCodeReplyTimedOut     MessageErrorCodeSchema = "reply_timed_out"
	MessageErrorCodeSpamSuspicion     MessageErrorCodeSchema = "spam_suspicion"
	MessageErrorCodeAccessRestricted  MessageErrorCodeSchema = "access_restricted"
)

const (
	MessageActionDialogOpened      MessageActionSchema = "dialog_opened"
	MessageActionDialogClosed      MessageActionSchema = "dialog_closed"
	MessageActionUserJoined        MessageActionSchema = "user_joined"
	MessageActionUserLeft          MessageActionSchema = "user_left"
	MessageActionDialogAssign      MessageActionSchema = "dialog_assign"
	MessageActionCustomerBlocked   MessageActionSchema = "customer_blocked"
	MessageActionCustomerUnblocked MessageActionSchema = "customer_unblocked"
	MessageActionDialogUnassign    MessageActionSchema = "dialog_unassign"
)

const (
	SuggestionTypeText  SuggestionTypeSchema = "text"
	SuggestionTypeEmail SuggestionTypeSchema = "email"
	SuggestionTypePhone SuggestionTypeSchema = "phone"
	SuggestionTypeURL   SuggestionTypeSchema = "url"
)

const (
	ChannelTypeTelegram      ChannelTypeSchema = "telegram"
	ChannelTypeFbmessenger   ChannelTypeSchema = "fbmessenger"
	ChannelTypeViber         ChannelTypeSchema = "viber"
	ChannelTypeWhatsapp      ChannelTypeSchema = "whatsapp"
	ChannelTypeSkype         ChannelTypeSchema = "skype"
	ChannelTypeVk            ChannelTypeSchema = "vk"
	ChannelTypeInstagram     ChannelTypeSchema = "instagram"
	ChannelTypeConsultant    ChannelTypeSchema = "consultant"
	ChannelTypeYandexChat    ChannelTypeSchema = "yandex_chat"
	ChannelTypeOdnoklassniki ChannelTypeSchema = "odnoklassniki"
	ChannelTypeMax           ChannelTypeSchema = "max"
	ChannelTypeOzon          ChannelTypeSchema = "ozon"
	ChannelTypeWildberries   ChannelTypeSchema = "wildberries"
	ChannelTypeYandexMarket  ChannelTypeSchema = "yandex_market"
	ChannelTypeMegaMarket    ChannelTypeSchema = "mega_market"
	ChannelTypeAvito         ChannelTypeSchema = "avito"
	ChannelTypeDrom          ChannelTypeSchema = "drom"
	ChannelTypeYoula         ChannelTypeSchema = "youla"
	ChannelTypeCustom        ChannelTypeSchema = "custom"
)

const (
	ChannelFeatureNone    ChannelFeatureSchema = "none"
	ChannelFeatureReceive ChannelFeatureSchema = "receive"
	ChannelFeatureSend    ChannelFeatureSchema = "send"
	ChannelFeatureBoth    ChannelFeatureSchema = "both"
)

const (
	CustomerExternalIdAny   CustomerExternalIdSchema = "any"
	CustomerExternalIdPhone CustomerExternalIdSchema = "phone"
)

const (
	SendingPolicyAfterReplyTimeoutNo       SendingPolicyAfterReplyTimeoutSchema = "no"
	SendingPolicyAfterReplyTimeoutTemplate SendingPolicyAfterReplyTimeoutSchema = "template"
)

const (
	SendingPolicyNewCustomerNo       SendingPolicyNewCustomerSchema = "no"
	SendingPolicyNewCustomerTemplate SendingPolicyNewCustomerSchema = "template"
	SendingPolicyNewCustomerText     SendingPolicyNewCustomerSchema = "text"
)

const (
	SendingPolicyOutgoingAllowed    SendingPolicyOutgoingSchema = "allowed"
	SendingPolicyOutgoingRestricted SendingPolicyOutgoingSchema = "restricted"
)

const (
	WAChannelQualityHigh   WAChannelQualitySchema = "high"
	WAChannelQualityMedium WAChannelQualitySchema = "medium"
	WAChannelQualityLow    WAChannelQualitySchema = "low"
)

const (
	WAChannelStatusConnected  WAChannelStatusSchema = "connected"
	WAChannelStatusFlagged    WAChannelStatusSchema = "flagged"
	WAChannelStatusOffline    WAChannelStatusSchema = "offline"
	WAChannelStatusPending    WAChannelStatusSchema = "pending"
	WAChannelStatusRestricted WAChannelStatusSchema = "restricted"
)

const (
	MemberStateActive    MemberStateSchema = "active"
	MemberStateKicked    MemberStateSchema = "kicked"
	MemberStateLeaved    MemberStateSchema = "leaved"
	MemberStateUndefined MemberStateSchema = "undefined"
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
	case EventTypeMessageRestored:
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
