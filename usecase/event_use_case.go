package usecase

import (
	"fmt"

	"github.com/betorvs/event-logger/config"
	"github.com/betorvs/event-logger/domain"
)

// LogEvent func print alerts received from am
func LogEvent(event *domain.Event) (err error) {
	if event.Name == "" || event.Kind == "" {
		err = fmt.Errorf("should contain name and/or kind")
		return err
	}
	logger := config.GetLogger
	defer logger().Sync()
	status := 0
	if event.Status != 0 {
		status = event.Status
	}
	logger().Infow(event.Message,
		"name", event.Name,
		"kind", event.Kind,
		"labels", event.Labels,
		"source", event.Source,
		"status", status,
	)

	return nil
}
