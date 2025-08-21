package bot_api_client

import (
	"context"
	"net/http"
)

func WithBotToken(token string) ClientOption {
	return WithRequestEditorFn(func(_ context.Context, req *http.Request) error {
		req.Header.Set("X-Bot-Token", token)

		return nil
	})
}
