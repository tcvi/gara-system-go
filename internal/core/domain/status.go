package domain

import "github.com/pkg/errors"

type Status string

var (
	New        Status = "New"
	Processing Status = "Processing"
	Done       Status = "Done"
)

func ParseStatus(status string) (*Status, error) {
	switch status {
	case string(New):
		return &New, nil
	case string(Processing):
		return &Processing, nil
	case string(Done):
		return &Done, nil
	}

	return nil, errors.New("Invalid Status")
}
