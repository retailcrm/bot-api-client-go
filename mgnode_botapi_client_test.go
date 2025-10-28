package bot_api_client

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	openapi_types "github.com/oapi-codegen/runtime/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListBotsWithResponse(t *testing.T) {
	t.Parallel()

	date := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	since := date.Add(-24 * time.Hour)
	until := date
	active := BooleanTrue
	self := BooleanTrue
	sinceID := SinceID(100)
	untilID := UntilID(200)
	limit := 50
	id := 123
	botRoles := BotRoleQuery{RoleBotRoleResponsible}

	testCases := []struct {
		name          string
		params        *ListBotsParams
		expectedQuery string
	}{
		{
			name: "all parameters",
			params: &ListBotsParams{
				ID:      &id,
				Self:    &self,
				Active:  &active,
				Role:    &botRoles,
				Since:   &since,
				SinceID: &sinceID,
				Until:   &until,
				UntilID: &untilID,
				Limit:   &limit,
			},
			expectedQuery: "active=true&id=123&limit=50&role=responsible&self=true&since=2024-12-31T00%3A00%3A00Z&since_id=100&until=2025-01-01T00%3A00%3A00Z&until_id=200",
		},
		{
			name:          "empty parameters",
			params:        &ListBotsParams{},
			expectedQuery: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			mockDoer := DoerFunc(func(req *http.Request) (*http.Response, error) {
				assert.Equal(t, "GET", req.Method)
				assert.Contains(t, req.URL.Path, "/bots")
				assert.Equal(t, tc.expectedQuery, req.URL.RawQuery)

				body := `{
					"data": [
						{
							"id": 123,
							"client_id": "bot-client-123",
							"name": "Test Bot",
							"is_active": true,
							"is_self": true,
							"is_system": false,
							"roles": ["responsible"],
							"created_at": "2024-12-31T00:00:00.000000Z"
						}
					]
				}`
				return &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(strings.NewReader(body)),
					Header:     make(http.Header),
				}, nil
			})

			client, err := NewClientWithResponses("https://example.com", WithHTTPClient(mockDoer))
			require.NoError(t, err)

			resp, err := client.ListBotsWithResponse(context.Background(), tc.params)
			require.NoError(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, 200, resp.StatusCode())

		})
	}
}

func TestListChannelsWithResponse(t *testing.T) {
	t.Parallel()

	date := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	since := date.Add(-24 * time.Hour)
	until := date
	active := BooleanTrue
	sinceID := SinceID(100)
	untilID := UntilID(200)
	limit := 50
	id := 123
	channelTypes := ChannelTypeQuery{ChannelTypeTelegram, ChannelTypeWhatsapp}

	testCases := []struct {
		name          string
		params        *ListChannelsParams
		expectedQuery string
	}{
		{
			name: "all parameters",
			params: &ListChannelsParams{
				ID:      &id,
				Types:   &channelTypes,
				Active:  &active,
				Since:   &since,
				SinceID: &sinceID,
				Until:   &until,
				UntilID: &untilID,
				Limit:   &limit,
			},
			expectedQuery: "active=true&id=123&limit=50&since=2024-12-31T00%3A00%3A00Z&since_id=100&types=telegram&types=whatsapp&until=2025-01-01T00%3A00%3A00Z&until_id=200",
		},
		{
			name:          "empty parameters",
			params:        &ListChannelsParams{},
			expectedQuery: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			mockDoer := DoerFunc(func(req *http.Request) (*http.Response, error) {
				assert.Equal(t, "GET", req.Method)
				assert.Contains(t, req.URL.Path, "/channels")
				assert.Equal(t, tc.expectedQuery, req.URL.RawQuery)

				body := `{
					"data": [
						{
							"id": 1,
							"type": "telegram",
							"name": "Test Channel",
							"is_active": true,
							"settings": {},
							"activated_at": "2024-12-31T00:00:00.000000Z",
							"created_at": "2024-12-31T00:00:00.000000Z"
						}
					]
				}`
				return &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(strings.NewReader(body)),
					Header:     make(http.Header),
				}, nil
			})

			client, err := NewClientWithResponses("https://example.com", WithHTTPClient(mockDoer))
			require.NoError(t, err)

			resp, err := client.ListChannelsWithResponse(context.Background(), tc.params)
			require.NoError(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, 200, resp.StatusCode())
		})
	}
}

func TestListChatsWithResponse(t *testing.T) {
	t.Parallel()

	date := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	since := date.Add(-24 * time.Hour)
	until := date
	sinceID := SinceID(100)
	untilID := UntilID(200)
	limit := 50
	channelID := 10
	id := 123

	testCases := []struct {
		name          string
		params        *ListChatsParams
		expectedQuery string
	}{
		{
			name: "all parameters",
			params: &ListChatsParams{
				ID:        &id,
				Since:     &since,
				Until:     &until,
				Limit:     &limit,
				SinceID:   &sinceID,
				UntilID:   &untilID,
				ChannelID: &channelID,
			},
			expectedQuery: "channel_id=10&id=123&limit=50&since=2024-12-31T00%3A00%3A00Z&since_id=100&until=2025-01-01T00%3A00%3A00Z&until_id=200",
		},
		{
			name:          "empty parameters",
			params:        &ListChatsParams{},
			expectedQuery: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			mockDoer := DoerFunc(func(req *http.Request) (*http.Response, error) {
				assert.Equal(t, "GET", req.Method)
				assert.Contains(t, req.URL.Path, "/chats")
				assert.Equal(t, tc.expectedQuery, req.URL.RawQuery)

				body := `{
					"data": [
						{
							"id": 1,
							"created_at": "2024-12-31T00:00:00.000000Z",
							"customer": null
						}
					]
				}`
				return &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(strings.NewReader(body)),
					Header:     make(http.Header),
				}, nil
			})

			client, err := NewClientWithResponses("https://example.com", WithHTTPClient(mockDoer))
			require.NoError(t, err)

			resp, err := client.ListChatsWithResponse(context.Background(), tc.params)
			require.NoError(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, 200, resp.StatusCode())
		})
	}
}

func TestListCustomersWithResponse(t *testing.T) {
	t.Parallel()

	date := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	since := date.Add(-24 * time.Hour)
	until := date
	sinceID := SinceID(100)
	untilID := UntilID(200)
	limit := 50
	channelID := 10
	id := 123

	testCases := []struct {
		name          string
		params        *ListCustomersParams
		expectedQuery string
	}{
		{
			name: "all parameters",
			params: &ListCustomersParams{
				ID:        &id,
				Since:     &since,
				Until:     &until,
				SinceID:   &sinceID,
				UntilID:   &untilID,
				Limit:     &limit,
				ChannelID: &channelID,
			},
			expectedQuery: "channel_id=10&id=123&limit=50&since=2024-12-31T00%3A00%3A00Z&since_id=100&until=2025-01-01T00%3A00%3A00Z&until_id=200",
		},
		{
			name:          "empty parameters",
			params:        &ListCustomersParams{},
			expectedQuery: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			mockDoer := DoerFunc(func(req *http.Request) (*http.Response, error) {
				assert.Equal(t, "GET", req.Method)
				assert.Contains(t, req.URL.Path, "/customers")
				assert.Equal(t, tc.expectedQuery, req.URL.RawQuery)

				body := `{
					"data": [
						{
							"id": 1,
							"created_at": "2024-12-31T00:00:00.000000Z",
							"is_blocked": false,
							"revoked_at": null,
							"updated_at": null,
							"utm": null
						}
					]
				}`
				return &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(strings.NewReader(body)),
					Header:     make(http.Header),
				}, nil
			})

			client, err := NewClientWithResponses("https://example.com", WithHTTPClient(mockDoer))
			require.NoError(t, err)

			resp, err := client.ListCustomersWithResponse(context.Background(), tc.params)
			require.NoError(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, 200, resp.StatusCode())
		})
	}
}

func TestListDialogsWithResponse(t *testing.T) {
	t.Parallel()

	date := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	since := date.Add(-24 * time.Hour)
	until := date
	sinceID := SinceID(100)
	untilID := UntilID(200)
	limit := 50
	chatID := 10
	id := 123

	testCases := []struct {
		name          string
		params        *ListDialogsParams
		expectedQuery string
	}{
		{
			name: "all parameters",
			params: &ListDialogsParams{
				ID:      &id,
				Since:   &since,
				Until:   &until,
				SinceID: &sinceID,
				UntilID: &untilID,
				Limit:   &limit,
				ChatID:  &chatID,
			},
			expectedQuery: "chat_id=10&id=123&limit=50&since=2024-12-31T00%3A00%3A00Z&since_id=100&until=2025-01-01T00%3A00%3A00Z&until_id=200",
		},
		{
			name:          "empty parameters",
			params:        &ListDialogsParams{},
			expectedQuery: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			mockDoer := DoerFunc(func(req *http.Request) (*http.Response, error) {
				assert.Equal(t, "GET", req.Method)
				assert.Contains(t, req.URL.Path, "/dialogs")
				assert.Equal(t, tc.expectedQuery, req.URL.RawQuery)

				body := `{
					"data": [
						{
							"id": 1,
							"chat_id": 123,
							"is_active": true,
							"is_assigned": false,
							"created_at": "2024-12-31T00:00:00.000000Z"
						}
					]
				}`
				return &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(strings.NewReader(body)),
					Header:     make(http.Header),
				}, nil
			})

			client, err := NewClientWithResponses("https://example.com", WithHTTPClient(mockDoer))
			require.NoError(t, err)

			resp, err := client.ListDialogsWithResponse(context.Background(), tc.params)
			require.NoError(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, 200, resp.StatusCode())
		})
	}
}

func TestListMembersWithResponse(t *testing.T) {
	t.Parallel()

	date := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	since := date.Add(-24 * time.Hour)
	until := date
	sinceID := SinceID(100)
	untilID := UntilID(200)
	limit := 50
	chatID := 10
	userID := 20
	state := ListMembersParamsState("active")
	id := 123

	testCases := []struct {
		name          string
		params        *ListMembersParams
		expectedQuery string
	}{
		{
			name: "all parameters",
			params: &ListMembersParams{
				ID:      &id,
				Since:   &since,
				Until:   &until,
				SinceID: &sinceID,
				UntilID: &untilID,
				Limit:   &limit,
				ChatID:  &chatID,
				UserID:  &userID,
				State:   &state,
			},
			expectedQuery: "chat_id=10&id=123&limit=50&since=2024-12-31T00%3A00%3A00Z&since_id=100&state=active&until=2025-01-01T00%3A00%3A00Z&until_id=200&user_id=20",
		},
		{
			name:          "empty parameters",
			params:        &ListMembersParams{},
			expectedQuery: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			mockDoer := DoerFunc(func(req *http.Request) (*http.Response, error) {
				assert.Equal(t, "GET", req.Method)
				assert.Contains(t, req.URL.Path, "/members")
				assert.Equal(t, tc.expectedQuery, req.URL.RawQuery)

				body := `{
					"data": [
						{
							"id": 1,
							"chat_id": 123,
							"user_id": 456,
							"is_author": false,
							"state": "active",
							"created_at": "2024-12-31T00:00:00.000000Z"
						}
					]
				}`
				return &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(strings.NewReader(body)),
					Header:     make(http.Header),
				}, nil
			})

			client, err := NewClientWithResponses("https://example.com", WithHTTPClient(mockDoer))
			require.NoError(t, err)

			resp, err := client.ListMembersWithResponse(context.Background(), tc.params)
			require.NoError(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, 200, resp.StatusCode())
		})
	}
}

func TestListMessagesWithResponse(t *testing.T) {
	t.Parallel()

	date := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	since := date.Add(-24 * time.Hour)
	until := date
	sinceID := SinceID(100)
	untilID := UntilID(200)
	limit := 50
	chatID := 10
	userID := 20
	channelType := ChannelTypeTelegram
	messageType := MessageTypeText
	scope := ListMessagesParamsScopePrivate
	includeMass := BooleanTrue
	messageIds := []int64{1, 2, 3}

	testCases := []struct {
		name          string
		params        *ListMessagesParams
		expectedQuery string
	}{
		{
			name: "all parameters",
			params: &ListMessagesParams{
				Since:                    &since,
				Until:                    &until,
				SinceID:                  &sinceID,
				UntilID:                  &untilID,
				Limit:                    &limit,
				MessageID:                &messageIds,
				ChatID:                   &chatID,
				UserID:                   &userID,
				ChannelType:              &channelType,
				Type:                     &messageType,
				IncludeMassCommunication: &includeMass,
				Scope:                    &scope,
			},
			expectedQuery: "channel_type=telegram&chat_id=10&id=1&id=2&id=3&include_mass_communication=true&limit=50&scope=private&since=2024-12-31T00%3A00%3A00Z&since_id=100&type=text&until=2025-01-01T00%3A00%3A00Z&until_id=200&user_id=20",
		},
		{
			name:          "empty parameters",
			params:        &ListMessagesParams{},
			expectedQuery: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			mockDoer := DoerFunc(func(req *http.Request) (*http.Response, error) {
				assert.Equal(t, "GET", req.Method)
				assert.Contains(t, req.URL.Path, "/messages")
				assert.Equal(t, tc.expectedQuery, req.URL.RawQuery)

				body := `{
					"data": [
						{
							"id": 1,
							"chat_id": 123,
							"is_edit": false,
							"is_read": false,
							"note": "",
							"created_at": "2024-12-31T00:00:00.000000Z",
							"channel_sent_at": null,
							"actions": []
						}
					]
				}`
				return &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(strings.NewReader(body)),
					Header:     make(http.Header),
				}, nil
			})

			client, err := NewClientWithResponses("https://example.com", WithHTTPClient(mockDoer))
			require.NoError(t, err)

			resp, err := client.ListMessagesWithResponse(context.Background(), tc.params)
			require.NoError(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, 200, resp.StatusCode())
		})
	}
}

func TestListCommandsWithResponse(t *testing.T) {
	t.Parallel()

	date := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	since := date.Add(-24 * time.Hour)
	until := date
	sinceID := SinceID(100)
	untilID := UntilID(200)
	limit := 50
	id := 123
	name := "start"

	testCases := []struct {
		name          string
		params        *ListCommandsParams
		expectedQuery string
	}{
		{
			name: "all parameters",
			params: &ListCommandsParams{
				ID:      &id,
				Limit:   &limit,
				SinceID: &sinceID,
				Since:   &since,
				UntilID: &untilID,
				Until:   &until,
				Name:    &name,
			},
			expectedQuery: "id=123&limit=50&name=start&since=2024-12-31T00%3A00%3A00Z&since_id=100&until=2025-01-01T00%3A00%3A00Z&until_id=200",
		},
		{
			name:          "empty parameters",
			params:        &ListCommandsParams{},
			expectedQuery: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			mockDoer := DoerFunc(func(req *http.Request) (*http.Response, error) {
				assert.Equal(t, "GET", req.Method)
				assert.Contains(t, req.URL.Path, "/my/commands")
				assert.Equal(t, tc.expectedQuery, req.URL.RawQuery)

				body := `{
					"data": [
						{
							"id": 1,
							"name": "start",
							"description": "Start command",
							"created_at": "2024-12-31T00:00:00.000000Z",
							"updated_at": null
						}
					]
				}`
				return &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(strings.NewReader(body)),
					Header:     make(http.Header),
				}, nil
			})

			client, err := NewClientWithResponses("https://example.com", WithHTTPClient(mockDoer))
			require.NoError(t, err)

			resp, err := client.ListCommandsWithResponse(context.Background(), tc.params)
			require.NoError(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, 200, resp.StatusCode())
		})
	}
}

func TestListUsersWithResponse(t *testing.T) {
	t.Parallel()

	date := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	since := date.Add(-24 * time.Hour)
	until := date
	sinceID := SinceID(100)
	untilID := UntilID(200)
	limit := 50
	active := BooleanTrue
	id := 123

	testCases := []struct {
		name          string
		params        *ListUsersParams
		expectedQuery string
	}{
		{
			name: "all parameters",
			params: &ListUsersParams{
				ID:      &id,
				Since:   &since,
				Until:   &until,
				SinceID: &sinceID,
				UntilID: &untilID,
				Limit:   &limit,
				Active:  &active,
			},
			expectedQuery: "active=true&id=123&limit=50&since=2024-12-31T00%3A00%3A00Z&since_id=100&until=2025-01-01T00%3A00%3A00Z&until_id=200",
		},
		{
			name:          "empty parameters",
			params:        &ListUsersParams{},
			expectedQuery: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			mockDoer := DoerFunc(func(req *http.Request) (*http.Response, error) {
				assert.Equal(t, "GET", req.Method)
				assert.Contains(t, req.URL.Path, "/users")
				assert.Equal(t, tc.expectedQuery, req.URL.RawQuery)

				body := `{
					"data": [
						{
							"id": 1,
							"available": true,
							"connected": true,
							"is_active": true,
							"is_online": false,
							"is_technical_account": false,
							"created_at": "2024-12-31T00:00:00.000000Z",
							"revoked_at": null,
							"updated_at": null
						}
					]
				}`
				return &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(strings.NewReader(body)),
					Header:     make(http.Header),
				}, nil
			})

			client, err := NewClientWithResponses("https://example.com", WithHTTPClient(mockDoer))
			require.NoError(t, err)

			resp, err := client.ListUsersWithResponse(context.Background(), tc.params)
			require.NoError(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, 200, resp.StatusCode())
		})
	}
}

func TestCreateDialogWithResponse(t *testing.T) {
	t.Parallel()

	mockDoer := DoerFunc(func(req *http.Request) (*http.Response, error) {
		assert.Equal(t, "POST", req.Method)
		assert.Contains(t, req.URL.Path, "/chats/1/dialogs")

		body := `{"data": {"id": 123}}`
		return &http.Response{
			StatusCode: 201,
			Body:       io.NopCloser(strings.NewReader(body)),
			Header:     make(http.Header),
		}, nil
	})

	client, err := NewClientWithResponses("https://example.com", WithHTTPClient(mockDoer))
	require.NoError(t, err)

	resp, err := client.CreateDialogWithResponse(context.Background(), 1, CreateDialogJSONRequestBody{})
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode())
}

func TestCreateDialogWithBodyWithResponse(t *testing.T) {
	t.Parallel()

	mockDoer := DoerFunc(func(req *http.Request) (*http.Response, error) {
		assert.Equal(t, "POST", req.Method)
		assert.Contains(t, req.URL.Path, "/chats/1/dialogs")

		body := `{"data": {"id": 123}}`
		return &http.Response{
			StatusCode: 201,
			Body:       io.NopCloser(strings.NewReader(body)),
			Header:     make(http.Header),
		}, nil
	})

	client, err := NewClientWithResponses("https://example.com", WithHTTPClient(mockDoer))
	require.NoError(t, err)

	resp, err := client.CreateDialogWithBodyWithResponse(context.Background(), 1, "application/json", strings.NewReader("{}"))
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode())
}

func TestAssignDialogResponsibleWithResponse(t *testing.T) {
	t.Parallel()

	mockDoer := DoerFunc(func(req *http.Request) (*http.Response, error) {
		assert.Equal(t, "PATCH", req.Method)
		assert.Contains(t, req.URL.Path, "/dialogs/123/assign")

		body := `{"data": {"success": true}}`
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader(body)),
			Header:     make(http.Header),
		}, nil
	})

	client, err := NewClientWithResponses("https://example.com", WithHTTPClient(mockDoer))
	require.NoError(t, err)

	resp, err := client.AssignDialogResponsibleWithResponse(context.Background(), 123, AssignDialogResponsibleJSONRequestBody{
		UserID: int64(1),
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnassignDialogResponsibleWithResponse(t *testing.T) {
	t.Parallel()

	mockDoer := DoerFunc(func(req *http.Request) (*http.Response, error) {
		assert.Equal(t, "PATCH", req.Method)
		assert.Contains(t, req.URL.Path, "/dialogs/123/unassign")

		body := `{"data": {"success": true}}`
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader(body)),
			Header:     make(http.Header),
		}, nil
	})

	client, err := NewClientWithResponses("https://example.com", WithHTTPClient(mockDoer))
	require.NoError(t, err)

	resp, err := client.UnassignDialogResponsibleWithResponse(context.Background(), 123)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestCloseDialogWithResponse(t *testing.T) {
	t.Parallel()

	mockDoer := DoerFunc(func(req *http.Request) (*http.Response, error) {
		assert.Equal(t, "DELETE", req.Method)
		assert.Contains(t, req.URL.Path, "/dialogs/123/close")

		body := `{"data": {"success": true}}`
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader(body)),
			Header:     make(http.Header),
		}, nil
	})

	client, err := NewClientWithResponses("https://example.com", WithHTTPClient(mockDoer))
	require.NoError(t, err)

	resp, err := client.CloseDialogWithResponse(context.Background(), 123)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestDialogAddTagsWithResponse(t *testing.T) {
	t.Parallel()

	mockDoer := DoerFunc(func(req *http.Request) (*http.Response, error) {
		assert.Equal(t, "PATCH", req.Method)
		assert.Contains(t, req.URL.Path, "/dialogs/123/tags/add")

		body := `{"data": {"success": true}}`
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader(body)),
			Header:     make(http.Header),
		}, nil
	})

	client, err := NewClientWithResponses("https://example.com", WithHTTPClient(mockDoer))
	require.NoError(t, err)

	resp, err := client.DialogAddTagsWithResponse(context.Background(), 123, DialogAddTagsJSONRequestBody{
		Tags: []struct {
			ColorCode *ColorCode `binding:"omitempty,enum-valid" json:"color_code,omitempty"`
			Name      string     `binding:"required,min=1,max=255" json:"name"`
		}{
			{Name: "tag1"},
			{Name: "tag2"},
		},
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestDialogDeleteTagsWithResponse(t *testing.T) {
	t.Parallel()

	mockDoer := DoerFunc(func(req *http.Request) (*http.Response, error) {
		assert.Equal(t, "PATCH", req.Method)
		assert.Contains(t, req.URL.Path, "/dialogs/123/tags/delete")

		body := `{"data": {"success": true}}`
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader(body)),
			Header:     make(http.Header),
		}, nil
	})

	client, err := NewClientWithResponses("https://example.com", WithHTTPClient(mockDoer))
	require.NoError(t, err)

	resp, err := client.DialogDeleteTagsWithResponse(context.Background(), 123, DialogDeleteTagsJSONRequestBody{
		Tags: []struct {
			Name string `binding:"required,min=1,max=255" json:"name"`
		}{
			{Name: "tag1"},
		},
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestSendMessageWithResponse(t *testing.T) {
	t.Parallel()

	mockDoer := DoerFunc(func(req *http.Request) (*http.Response, error) {
		assert.Equal(t, "POST", req.Method)
		assert.Contains(t, req.URL.Path, "/messages")

		body := `{"data": {"id": 123}}`
		return &http.Response{
			StatusCode: 201,
			Body:       io.NopCloser(strings.NewReader(body)),
			Header:     make(http.Header),
		}, nil
	})

	client, err := NewClientWithResponses("https://example.com", WithHTTPClient(mockDoer))
	require.NoError(t, err)

	resp, err := client.SendMessageWithResponse(context.Background(), SendMessageJSONRequestBody{
		ChatID:  int64(1),
		Content: strPtr("test message"),
		Scope:   MessageScopePrivate,
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode())
}

func TestSendMessageWithBodyWithResponse(t *testing.T) {
	t.Parallel()

	mockDoer := DoerFunc(func(req *http.Request) (*http.Response, error) {
		assert.Equal(t, "POST", req.Method)
		assert.Contains(t, req.URL.Path, "/messages")

		body := `{"data": {"id": 123}}`
		return &http.Response{
			StatusCode: 201,
			Body:       io.NopCloser(strings.NewReader(body)),
			Header:     make(http.Header),
		}, nil
	})

	client, err := NewClientWithResponses("https://example.com", WithHTTPClient(mockDoer))
	require.NoError(t, err)

	resp, err := client.SendMessageWithBodyWithResponse(context.Background(), "application/json", strings.NewReader(`{"chat_id": 1, "scope": "private"}`))
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode())
}

func TestEditMessageWithResponse(t *testing.T) {
	t.Parallel()

	mockDoer := DoerFunc(func(req *http.Request) (*http.Response, error) {
		assert.Equal(t, "PATCH", req.Method)
		assert.Contains(t, req.URL.Path, "/messages/123")

		body := `{"data": {"success": true}}`
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader(body)),
			Header:     make(http.Header),
		}, nil
	})

	client, err := NewClientWithResponses("https://example.com", WithHTTPClient(mockDoer))
	require.NoError(t, err)

	resp, err := client.EditMessageWithResponse(context.Background(), 123, EditMessageJSONRequestBody{
		Content: strPtr("edited message"),
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestDeleteMessageWithResponse(t *testing.T) {
	t.Parallel()

	mockDoer := DoerFunc(func(req *http.Request) (*http.Response, error) {
		assert.Equal(t, "DELETE", req.Method)
		assert.Contains(t, req.URL.Path, "/messages/123")

		body := `{"data": {"success": true}}`
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader(body)),
			Header:     make(http.Header),
		}, nil
	})

	client, err := NewClientWithResponses("https://example.com", WithHTTPClient(mockDoer))
	require.NoError(t, err)

	resp, err := client.DeleteMessageWithResponse(context.Background(), 123)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUploadFileWithBodyWithResponse(t *testing.T) {
	t.Parallel()

	mockDoer := DoerFunc(func(req *http.Request) (*http.Response, error) {
		assert.Equal(t, "POST", req.Method)
		assert.Contains(t, req.URL.Path, "/files/upload")

		body := `{"data": {"id": "file-123"}}`
		return &http.Response{
			StatusCode: 201,
			Body:       io.NopCloser(strings.NewReader(body)),
			Header:     make(http.Header),
		}, nil
	})

	client, err := NewClientWithResponses("https://example.com", WithHTTPClient(mockDoer))
	require.NoError(t, err)

	resp, err := client.UploadFileWithBodyWithResponse(context.Background(), "multipart/form-data", strings.NewReader("file data"))
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode())
}

func TestUploadFileByUrlWithResponse(t *testing.T) {
	t.Parallel()

	mockDoer := DoerFunc(func(req *http.Request) (*http.Response, error) {
		assert.Equal(t, "POST", req.Method)
		assert.Contains(t, req.URL.Path, "/files/upload_by_url")

		body := `{"data": {"id": "file-123"}}`
		return &http.Response{
			StatusCode: 201,
			Body:       io.NopCloser(strings.NewReader(body)),
			Header:     make(http.Header),
		}, nil
	})

	client, err := NewClientWithResponses("https://example.com", WithHTTPClient(mockDoer))
	require.NoError(t, err)

	resp, err := client.UploadFileByUrlWithResponse(context.Background(), UploadFileByUrlJSONRequestBody{
		Url: "https://example.com/file.png",
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode())
}

func TestGetFileUrlWithResponse(t *testing.T) {
	t.Parallel()

	mockDoer := DoerFunc(func(req *http.Request) (*http.Response, error) {
		assert.Equal(t, "GET", req.Method)
		assert.Contains(t, req.URL.Path, "/files/")

		body := `{"data": {"url": "https://example.com/file.png"}}`
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader(body)),
			Header:     make(http.Header),
		}, nil
	})

	client, err := NewClientWithResponses("https://example.com", WithHTTPClient(mockDoer))
	require.NoError(t, err)

	fileID := openapi_types.UUID{}
	resp, err := client.GetFileUrlWithResponse(context.Background(), fileID)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUpdateFileMetadataWithResponse(t *testing.T) {
	t.Parallel()

	mockDoer := DoerFunc(func(req *http.Request) (*http.Response, error) {
		assert.Equal(t, "PUT", req.Method)
		assert.Contains(t, req.URL.Path, "/files/")

		body := `{"data": {"success": true}}`
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader(body)),
			Header:     make(http.Header),
		}, nil
	})

	client, err := NewClientWithResponses("https://example.com", WithHTTPClient(mockDoer))
	require.NoError(t, err)

	fileID := openapi_types.UUID{}
	resp, err := client.UpdateFileMetadataWithResponse(context.Background(), fileID, UpdateFileMetadataJSONRequestBody{})
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestCreateOrUpdateCommandWithResponse(t *testing.T) {
	t.Parallel()

	mockDoer := DoerFunc(func(req *http.Request) (*http.Response, error) {
		assert.Equal(t, "PUT", req.Method)
		assert.Contains(t, req.URL.Path, "/my/commands/test")

		body := `{"data": {"success": true}}`
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader(body)),
			Header:     make(http.Header),
		}, nil
	})

	client, err := NewClientWithResponses("https://example.com", WithHTTPClient(mockDoer))
	require.NoError(t, err)

	resp, err := client.CreateOrUpdateCommandWithResponse(context.Background(), "test", CreateOrUpdateCommandJSONRequestBody{})
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestDeleteCommandWithResponse(t *testing.T) {
	t.Parallel()

	mockDoer := DoerFunc(func(req *http.Request) (*http.Response, error) {
		assert.Equal(t, "DELETE", req.Method)
		assert.Contains(t, req.URL.Path, "/my/commands/test")

		body := `{"data": {"success": true}}`
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader(body)),
			Header:     make(http.Header),
		}, nil
	})

	client, err := NewClientWithResponses("https://example.com", WithHTTPClient(mockDoer))
	require.NoError(t, err)

	resp, err := client.DeleteCommandWithResponse(context.Background(), "test")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUpdateBotWithResponse(t *testing.T) {
	t.Parallel()

	mockDoer := DoerFunc(func(req *http.Request) (*http.Response, error) {
		assert.Equal(t, "PATCH", req.Method)
		assert.Contains(t, req.URL.Path, "/my/info")

		body := `{"data": {"success": true}}`
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader(body)),
			Header:     make(http.Header),
		}, nil
	})

	client, err := NewClientWithResponses("https://example.com", WithHTTPClient(mockDoer))
	require.NoError(t, err)

	resp, err := client.UpdateBotWithResponse(context.Background(), UpdateBotJSONRequestBody{})
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func strPtr(s string) *string {
	return &s
}
