package events

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type TestEvent struct {
	Name    string
	Payload interface{}
}

func (e *TestEvent) GetName() string {
	return e.Name
}

func (e *TestEvent) GetPayload() interface{} {
	return e.Payload
}

func (e *TestEvent) GetDateTime() time.Time {
	return time.Now()
}

type TestEventHandler struct{}

func (h *TestEventHandler) Handle(event EventInterface) error {
	return nil
}

type EventDispatcherTestSuite struct {
	suite.Suite
	event1          TestEvent
	event2          TestEvent
	handler1        TestEventHandler
	handler2        TestEventHandler
	handler3        TestEventHandler
	eventDispatcher EventDispatcher
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}

func (s *EventDispatcherTestSuite) SetupTest() {
	s.eventDispatcher = *NewEventDispatcher()
	s.handler1 = TestEventHandler{}
	s.handler2 = TestEventHandler{}
	s.handler3 = TestEventHandler{}
	s.event1 = TestEvent{Name: "test1", Payload: "payload1"}
	s.event2 = TestEvent{Name: "test2", Payload: "payload2"}
}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Register() {}
