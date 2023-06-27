package events

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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

type TestEventHandler struct {
	ID int
}

func (h *TestEventHandler) Handle(event EventInterface) {}

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
	s.handler1 = TestEventHandler{
		ID: 1,
	}
	s.handler2 = TestEventHandler{
		ID: 2,
	}
	s.handler3 = TestEventHandler{
		ID: 3,
	}
	s.event1 = TestEvent{Name: "test1", Payload: "payload1"}
	s.event2 = TestEvent{Name: "test2", Payload: "payload2"}
}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	err := s.eventDispatcher.Register(s.event1.GetName(), &s.handler1)
	s.Nil(err)
	s.Equal(1, len(s.eventDispatcher.handlers[s.event1.GetName()]))

	err = s.eventDispatcher.Register(s.event1.GetName(), &s.handler2)
	s.Nil(err)
	s.Equal(2, len(s.eventDispatcher.handlers[s.event1.GetName()]))

	assert.Equal(s.T(), &s.handler1, s.eventDispatcher.handlers[s.event1.GetName()][0])
	assert.Equal(s.T(), &s.handler2, s.eventDispatcher.handlers[s.event1.GetName()][1])
}

func (s *EventDispatcherTestSuite) TestEventDispatcher_Register_WithSameHandler() {
	err := s.eventDispatcher.Register(s.event1.GetName(), &s.handler1)
	s.Nil(err)
	s.Equal(1, len(s.eventDispatcher.handlers[s.event1.GetName()]))

	err = s.eventDispatcher.Register(s.event1.GetName(), &s.handler1)
	s.Equal(ErrHandlerAlreadyRegistered, err)
	s.Equal(1, len(s.eventDispatcher.handlers[s.event1.GetName()]))
}
