package secrets

import (
	"encoding/json"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/google/uuid"
	events "soeren.cloud-events/pkg"
)

type RebootData struct {
	SecretName string `json:"secretName"`
	Path       string `json:"path"`
	Status     string `json:"status"`
}

func NewSystemRebootedEvent(source string, data RebootData) cloudevents.Event {
	event := cloudevents.NewEvent()
	event.SetSource(source)
	event.SetType(string(events.TypeSystemRebooted))
	event.SetID(uuid.New().String())
	event.SetTime(time.Now())
	event.SetSubject("os-reboot")
	event.SetDataContentType(cloudevents.ApplicationJSON)
	event.SetData(cloudevents.ApplicationJSON, data)
	return event
}

func NewMarshalledSystemRebootEvent(source string, data RebootData) []byte {
	m, _ := json.Marshal(NewSystemRebootedEvent(source, data))
	return m
}
