package bot_api_client

import (
	"context"
	"net/http"
)

const (
	botTokenHeader = "X-Bot-Token"
)

func WithBotToken(token string) ClientOption {
	return WithRequestEditorFn(func(_ context.Context, req *http.Request) error {
		req.Header.Set(botTokenHeader, token)

		return nil
	})
}
