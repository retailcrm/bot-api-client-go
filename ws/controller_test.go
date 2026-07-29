package ws

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
	bot_api_client "github.com/retailcrm/bot-api-client-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserAgent(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		options  []Option
		expected string
	}{
		{
			name:     "default",
			expected: bot_api_client.DefaultUserAgent(),
		},
		{
			name:     "custom",
			options:  []Option{WithUserAgent("custom-client/1.0.0")},
			expected: "custom-client/1.0.0",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			headers := make(chan http.Header, 1)
			upgrader := websocket.Upgrader{}
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				headers <- req.Header
				conn, err := upgrader.Upgrade(w, req, nil)
				if err == nil {
					_ = conn.Close()
				}
			}))
			defer server.Close()

			controller, err := NewController(
				strings.Replace(server.URL, "http://", "ws://", 1),
				"token",
				tc.options...,
			)
			require.NoError(t, err)

			broker, ok := controller.broker.(*Controller)
			require.True(t, ok)

			conn, _, err := websocket.DefaultDialer.Dial(
				strings.Replace(server.URL, "http://", "ws://", 1),
				broker.headers,
			)
			require.NoError(t, err)
			require.NoError(t, conn.Close())
			assert.Equal(t, tc.expected, (<-headers).Get("User-Agent"))
		})
	}
}
