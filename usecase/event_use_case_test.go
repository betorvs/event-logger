package usecase

import (
	"encoding/json"
	"testing"

	"github.com/betorvs/event-logger/appcontext"
	"github.com/betorvs/event-logger/domain"
	localtest "github.com/betorvs/event-logger/test"
	"github.com/stretchr/testify/assert"
)

func TestLogEvent(t *testing.T) {

	appcontext.Current.Add(appcontext.Logger, localtest.InitMockLogger)

	event := new(domain.Event)
	rawJSON := []byte(localtest.TestValidJSON)

	err := json.Unmarshal(rawJSON, &event)
	assert.NoError(t, err)
	err = LogEvent(event)
	assert.NoError(t, err)

	event1 := new(domain.Event)
	rawJSON1 := []byte(localtest.TestInvalidJSON)

	err1 := json.Unmarshal(rawJSON1, &event1)
	assert.NoError(t, err1)
	err1 = LogEvent(event1)
	assert.Error(t, err1)
}
