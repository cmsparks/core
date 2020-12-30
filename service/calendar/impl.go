package calendar

import (
  "fmt"

  gcal "google.golang.org/api/calendar/v3"

	"github.com/acm-uiuc/core/config"
)

type calendarImpl struct {
  calendar *gcal.Service
}

func (service *calendarImpl) GetEvents() (*gcal.Events, error) {
	id, err := config.GetConfigValue("GOOGLE_CALENDAR_ID")
	if err != nil {
		return nil, fmt.Errorf("failed to get config value: %w", err)
	}

  res, err := service.calendar.Events.List(id).Do()
  /*if err != nil {
		return false, fmt.Errorf("Unable to retrieve calendar events list: %v", err)
  }*/
  
  return res, nil
}
