package dyndns

import (
	"encoding/json"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/google/uuid"
	"github.com/soerenschneider/soeren.cloud-events/pkg"
)

type NewIpDetectedEvent struct {
	IPv4     string
	IPv6     string
	Hostname string
}

func NewNewIpDetectedEvent(source string, data *NewIpDetectedEvent) cloudevents.Event {
	event := cloudevents.NewEvent()
	event.SetSource(source)
	event.SetType(string(pkg.TypeDyndnsNewIpDetected))
	event.SetID(uuid.New().String())
	event.SetTime(time.Now())
	event.SetSubject("new-ip-detected")
	if data != nil {
		event.SetDataContentType(cloudevents.ApplicationJSON)
		_ = event.SetData(cloudevents.ApplicationJSON, data)
	}
	return event
}

func NewMarshalledNewIpDetectedEvent(source string, data *NewIpDetectedEvent) []byte {
	m, _ := json.Marshal(NewNewIpDetectedEvent(source, data))
	return m
}
