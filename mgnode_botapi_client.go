package bot_api_client

import (
	"errors"
	"fmt"
	"net/http"
)

type Err interface {
	Error() error
}

type HTTPResponseCarrier interface {
	Response() *http.Response
	ResponseBody() []byte
}

func ExtractError(resp Err, err error) error {
	if err != nil {
		return err
	}

	if resp != nil {
		if response, body, ok := responseDetails(resp); ok && response != nil && response.StatusCode == http.StatusTooManyRequests {
			cause := resp.Error()
			if cause == nil {
				cause = fmt.Errorf("unexpected HTTP status: %d", response.StatusCode)
			}
			return NewRateLimitError(response, body, cause)
		}
		if resp.Error() != nil {
			return resp.Error()
		}
		if response, _, ok := responseDetails(resp); ok && response != nil && response.StatusCode >= http.StatusBadRequest {
			return fmt.Errorf("unexpected HTTP status: %d", response.StatusCode)
		}
	}

	return nil
}

func defaultResponseError(response *ErrorResponse) error {
	if response == nil || len(response.Errors) == 0 {
		return nil
	}
	return errors.New(response.Errors[0])
}

func (r ListBotsResp) Error() error {
	return defaultResponseError(r.JSONDefault)
}

func (r ListChannelsResp) Error() error {
	return defaultResponseError(r.JSONDefault)
}

func (r ListChatsResp) Error() error {
	return defaultResponseError(r.JSONDefault)
}

func (r CreateDialogResp) Error() error {
	return defaultResponseError(r.JSONDefault)
}

func (r ListCustomersResp) Error() error {
	return defaultResponseError(r.JSONDefault)
}

func (r ListDialogsResp) Error() error {
	return defaultResponseError(r.JSONDefault)
}

func (r AssignDialogResponsibleResp) Error() error {
	return defaultResponseError(r.JSONDefault)
}

func (r CloseDialogResp) Error() error {
	return defaultResponseError(r.JSONDefault)
}

func (r DialogAddTagsResp) Error() error {
	return defaultResponseError(r.JSONDefault)
}

func (r DialogDeleteTagsResp) Error() error {
	return defaultResponseError(r.JSONDefault)
}

func (r UnassignDialogResponsibleResp) Error() error {
	return defaultResponseError(r.JSONDefault)
}

func (r UploadFileResp) Error() error {
	return defaultResponseError(r.JSONDefault)
}

func (r UploadFileByUrlResp) Error() error {
	return defaultResponseError(r.JSONDefault)
}

func (r GetFileUrlResp) Error() error {
	return defaultResponseError(r.JSONDefault)
}

func (r UpdateFileMetadataResp) Error() error {
	return defaultResponseError(r.JSONDefault)
}

func (r ListMembersResp) Error() error {
	return defaultResponseError(r.JSONDefault)
}

func (r ListMessagesResp) Error() error {
	return defaultResponseError(r.JSONDefault)
}

func (r SendMessageResp) Error() error {
	return defaultResponseError(r.JSONDefault)
}

func (r DeleteMessageResp) Error() error {
	return defaultResponseError(r.JSONDefault)
}

func (r EditMessageResp) Error() error {
	return defaultResponseError(r.JSONDefault)
}

func (r ListCommandsResp) Error() error {
	return defaultResponseError(r.JSONDefault)
}

func (r DeleteCommandResp) Error() error {
	return defaultResponseError(r.JSONDefault)
}

func (r CreateOrUpdateCommandResp) Error() error {
	return defaultResponseError(r.JSONDefault)
}

func (r UpdateBotResp) Error() error {
	return defaultResponseError(r.JSONDefault)
}

func (r ListUsersResp) Error() error {
	return defaultResponseError(r.JSONDefault)
}

func (r WebSocketConnectionResp) Error() error {
	return defaultResponseError(r.JSONDefault)
}

func (r ListBotsResp) Response() *http.Response { return r.HTTPResponse }
func (r ListBotsResp) ResponseBody() []byte     { return r.Body }

func (r ListChannelsResp) Response() *http.Response { return r.HTTPResponse }
func (r ListChannelsResp) ResponseBody() []byte     { return r.Body }

func (r ListChatsResp) Response() *http.Response { return r.HTTPResponse }
func (r ListChatsResp) ResponseBody() []byte     { return r.Body }

func (r CreateDialogResp) Response() *http.Response { return r.HTTPResponse }
func (r CreateDialogResp) ResponseBody() []byte     { return r.Body }

func (r ListCustomersResp) Response() *http.Response { return r.HTTPResponse }
func (r ListCustomersResp) ResponseBody() []byte     { return r.Body }

func (r ListDialogsResp) Response() *http.Response { return r.HTTPResponse }
func (r ListDialogsResp) ResponseBody() []byte     { return r.Body }

func (r AssignDialogResponsibleResp) Response() *http.Response { return r.HTTPResponse }
func (r AssignDialogResponsibleResp) ResponseBody() []byte     { return r.Body }

func (r CloseDialogResp) Response() *http.Response { return r.HTTPResponse }
func (r CloseDialogResp) ResponseBody() []byte     { return r.Body }

func (r DialogAddTagsResp) Response() *http.Response { return r.HTTPResponse }
func (r DialogAddTagsResp) ResponseBody() []byte     { return r.Body }

func (r DialogDeleteTagsResp) Response() *http.Response { return r.HTTPResponse }
func (r DialogDeleteTagsResp) ResponseBody() []byte     { return r.Body }

func (r UnassignDialogResponsibleResp) Response() *http.Response { return r.HTTPResponse }
func (r UnassignDialogResponsibleResp) ResponseBody() []byte     { return r.Body }

func (r UploadFileResp) Response() *http.Response { return r.HTTPResponse }
func (r UploadFileResp) ResponseBody() []byte     { return r.Body }

func (r UploadFileByUrlResp) Response() *http.Response { return r.HTTPResponse }
func (r UploadFileByUrlResp) ResponseBody() []byte     { return r.Body }

func (r GetFileUrlResp) Response() *http.Response { return r.HTTPResponse }
func (r GetFileUrlResp) ResponseBody() []byte     { return r.Body }

func (r UpdateFileMetadataResp) Response() *http.Response { return r.HTTPResponse }
func (r UpdateFileMetadataResp) ResponseBody() []byte     { return r.Body }

func (r ListMembersResp) Response() *http.Response { return r.HTTPResponse }
func (r ListMembersResp) ResponseBody() []byte     { return r.Body }

func (r ListMessagesResp) Response() *http.Response { return r.HTTPResponse }
func (r ListMessagesResp) ResponseBody() []byte     { return r.Body }

func (r SendMessageResp) Response() *http.Response { return r.HTTPResponse }
func (r SendMessageResp) ResponseBody() []byte     { return r.Body }

func (r DeleteMessageResp) Response() *http.Response { return r.HTTPResponse }
func (r DeleteMessageResp) ResponseBody() []byte     { return r.Body }

func (r EditMessageResp) Response() *http.Response { return r.HTTPResponse }
func (r EditMessageResp) ResponseBody() []byte     { return r.Body }

func (r ListCommandsResp) Response() *http.Response { return r.HTTPResponse }
func (r ListCommandsResp) ResponseBody() []byte     { return r.Body }

func (r DeleteCommandResp) Response() *http.Response { return r.HTTPResponse }
func (r DeleteCommandResp) ResponseBody() []byte     { return r.Body }

func (r CreateOrUpdateCommandResp) Response() *http.Response { return r.HTTPResponse }
func (r CreateOrUpdateCommandResp) ResponseBody() []byte     { return r.Body }

func (r UpdateBotResp) Response() *http.Response { return r.HTTPResponse }
func (r UpdateBotResp) ResponseBody() []byte     { return r.Body }

func (r ListUsersResp) Response() *http.Response { return r.HTTPResponse }
func (r ListUsersResp) ResponseBody() []byte     { return r.Body }

func (r WebSocketConnectionResp) Response() *http.Response { return r.HTTPResponse }
func (r WebSocketConnectionResp) ResponseBody() []byte     { return r.Body }
