package notification

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"fmt"
)

type Service struct {
	Client *firebase.App
}

func NewService() (*Service, error) {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}

	return &Service{app}, nil
}
