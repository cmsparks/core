package calendar

import (
	"fmt"
	"context"
	
  "golang.org/x/oauth2/google"
  gcal "google.golang.org/api/calendar/v3"
)

type CalendarService interface {
	GetEvents() (*gcal.Events, error)
}

const scope = "https://www.googleapis.com/auth/calendar"

func New() (CalendarService, error) {
	ctx := context.Background()
  client, err := google.DefaultClient(ctx, scope)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %w", err)
	}

  svc, err := gcal.New(client)
  if err != nil {
		return nil, fmt.Errorf("Unable to create Calendar service: %v", err)
	}

  return &calendarImpl{
	  calendar: svc,
	}, nil
}
