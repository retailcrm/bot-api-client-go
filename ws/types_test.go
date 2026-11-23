package ws

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDialogOpenedContext(t *testing.T) {
	var event EventSchema
	require.NoError(t, json.Unmarshal([]byte(`{"type":"dialog_opened","data":{"dialog":{"id":2,"created_at":"2026-10-09T10:00:00Z"},"preceding_message":{"message_id":1,"created_at":"2026-10-09T09:00:00Z","author_id":3}}}`), &event))
	opened, ok := event.Data.(DialogOpenedDataSchema)
	require.True(t, ok)
	require.NotNil(t, opened.PrecedingMessage)
	require.Equal(t, int64(1), opened.PrecedingMessage.MessageId)
	require.False(t, opened.PrecedingMessage.CreatedAt.IsZero())
	require.NoError(t, json.Unmarshal([]byte(`{"type":"dialog_closed","data":{"dialog":{"id":2,"created_at":"2026-10-09T10:00:00Z"}}}`), &event))
	require.IsType(t, DialogDataSchema{}, event.Data)
}
