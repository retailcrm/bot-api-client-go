package bot_api_client

import (
	"errors"
)

type Err interface {
	Error() error
}

func ExtractError(resp Err, err error) error {
	if err != nil {
		return err
	}

	if resp != nil && resp.Error() != nil {
		return resp.Error()
	}

	return nil
}

func (r ListBotsResp) Error() error {
	if r.JSONDefault != nil {
		return errors.New(r.JSONDefault.Errors[0])
	}

	return nil
}

func (r ListChannelsResp) Error() error {
	if r.JSONDefault != nil {
		return errors.New(r.JSONDefault.Errors[0])
	}

	return nil
}

func (r ListChatsResp) Error() error {
	if r.JSONDefault != nil {
		return errors.New(r.JSONDefault.Errors[0])
	}

	return nil
}

func (r CreateDialogResp) Error() error {
	if r.JSONDefault != nil {
		return errors.New(r.JSONDefault.Errors[0])
	}

	return nil
}

func (r ListCustomersResp) Error() error {
	if r.JSONDefault != nil {
		return errors.New(r.JSONDefault.Errors[0])
	}

	return nil
}

func (r ListDialogsResp) Error() error {
	if r.JSONDefault != nil {
		return errors.New(r.JSONDefault.Errors[0])
	}

	return nil
}

func (r AssignDialogResponsibleResp) Error() error {
	if r.JSONDefault != nil {
		return errors.New(r.JSONDefault.Errors[0])
	}

	return nil
}

func (r CloseDialogResp) Error() error {
	if r.JSONDefault != nil {
		return errors.New(r.JSONDefault.Errors[0])
	}

	return nil
}

func (r DialogAddTagsResp) Error() error {
	if r.JSONDefault != nil {
		return errors.New(r.JSONDefault.Errors[0])
	}

	return nil
}

func (r DialogDeleteTagsResp) Error() error {
	if r.JSONDefault != nil {
		return errors.New(r.JSONDefault.Errors[0])
	}

	return nil
}

func (r UnassignDialogResponsibleResp) Error() error {
	if r.JSONDefault != nil {
		return errors.New(r.JSONDefault.Errors[0])
	}

	return nil
}

func (r UploadFileResp) Error() error {
	if r.JSONDefault != nil {
		return errors.New(r.JSONDefault.Errors[0])
	}

	return nil
}

func (r UploadFileByUrlResp) Error() error {
	if r.JSONDefault != nil {
		return errors.New(r.JSONDefault.Errors[0])
	}

	return nil
}

func (r GetFileUrlResp) Error() error {
	if r.JSONDefault != nil {
		return errors.New(r.JSONDefault.Errors[0])
	}

	return nil
}

func (r UpdateFileMetadataResp) Error() error {
	if r.JSONDefault != nil {
		return errors.New(r.JSONDefault.Errors[0])
	}

	return nil
}

func (r ListMembersResp) Error() error {
	if r.JSONDefault != nil {
		return errors.New(r.JSONDefault.Errors[0])
	}

	return nil
}

func (r ListMessagesResp) Error() error {
	if r.JSONDefault != nil {
		return errors.New(r.JSONDefault.Errors[0])
	}

	return nil
}

func (r SendMessageResp) Error() error {
	if r.JSONDefault != nil {
		return errors.New(r.JSONDefault.Errors[0])
	}

	return nil
}

func (r DeleteMessageResp) Error() error {
	if r.JSONDefault != nil {
		return errors.New(r.JSONDefault.Errors[0])
	}

	return nil
}

func (r EditMessageResp) Error() error {
	if r.JSONDefault != nil {
		return errors.New(r.JSONDefault.Errors[0])
	}

	return nil
}

func (r ListCommandsResp) Error() error {
	if r.JSONDefault != nil {
		return errors.New(r.JSONDefault.Errors[0])
	}

	return nil
}

func (r DeleteCommandResp) Error() error {
	if r.JSONDefault != nil {
		return errors.New(r.JSONDefault.Errors[0])
	}

	return nil
}

func (r CreateOrUpdateCommandResp) Error() error {
	if r.JSONDefault != nil {
		return errors.New(r.JSONDefault.Errors[0])
	}

	return nil
}

func (r UpdateBotResp) Error() error {
	if r.JSONDefault != nil {
		return errors.New(r.JSONDefault.Errors[0])
	}

	return nil
}

func (r ListUsersResp) Error() error {
	if r.JSONDefault != nil {
		return errors.New(r.JSONDefault.Errors[0])
	}

	return nil
}

func (r WebSocketConnectionResp) Error() error {
	if r.JSONDefault != nil {
		return errors.New(r.JSONDefault.Errors[0])
	}

	return nil
}
