package dyndns

import (
	"encoding/json"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/google/uuid"
	"github.com/soerenschneider/soeren.cloud-events/pkg"
)

type NewIpAppliedEvent struct {
	IPv4     string
	IPv6     string
	Hostname string
	Source   string
}

func NewNewIpAppliedEvent(source string, data *NewIpAppliedEvent) cloudevents.Event {
	event := cloudevents.NewEvent()
	event.SetSource(source)
	event.SetType(string(pkg.TypeDyndnsNewIpApplied))
	event.SetID(uuid.New().String())
	event.SetTime(time.Now())
	event.SetSubject("new-ip-applied")
	if data != nil {
		event.SetDataContentType(cloudevents.ApplicationJSON)
		_ = event.SetData(cloudevents.ApplicationJSON, data)
	}
	return event
}

func NewMarshalledNewIpAppliedEvent(source string, data *NewIpAppliedEvent) []byte {
	m, _ := json.Marshal(NewNewIpAppliedEvent(source, data))
	return m
}
