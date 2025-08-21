package bot_api_client

import (
	"time"
)

const RFC3339Micro string = "2006-01-02T15:04:05.999999Z07:00"

type DateTimeRFC3339 struct {
	Time time.Time `json:"time"`
}

type DateTimeRFC3339Micro struct {
	Time time.Time `json:"time"`
}

func (t DateTimeRFC3339) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.Time.Format(time.RFC3339) + `"`), nil
}

func (t *DateTimeRFC3339) UnmarshalJSON(data []byte) error {
	if len(data) == 2 && string(data) == `""` {
		return nil
	}

	parsedTime, err := time.Parse(`"`+time.RFC3339+`"`, string(data))
	if err != nil {
		return err
	}
	t.Time = parsedTime

	return nil
}

func (t DateTimeRFC3339Micro) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.Time.Format(RFC3339Micro) + `"`), nil
}

func (t *DateTimeRFC3339Micro) UnmarshalJSON(data []byte) error {
	if len(data) == 2 && string(data) == `""` {
		return nil
	}

	parsedTime, err := time.Parse(`"`+RFC3339Micro+`"`, string(data))
	if err != nil {
		return err
	}
	t.Time = parsedTime

	return nil
}
