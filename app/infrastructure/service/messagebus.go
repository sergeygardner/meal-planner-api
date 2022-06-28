package service

import (
	"fmt"
	"github.com/sergeygardner/meal-planner-api/application/event"
	"github.com/vardius/message-bus"
	"reflect"
	"runtime"
	"sync"
)

var (
	messageBusService MessageBusServiceInterface
)

type MessageBusServiceInterface interface {
	addEventListener(topic string, handlerName string, handler interface{}) error
	removeEventListener(topic string, handlerName string) error
	Publish(topic string, data interface{}) error
	AddAuthConfirmationEvent() error
	RemoveAuthConfirmationEvent() error
}

type MessageBusService struct {
	driver           messagebus.MessageBus
	handlers         map[string]map[string]interface{}
	syncWaitingGroup sync.WaitGroup
	MessageBusServiceInterface
}

func (mbs *MessageBusService) addEventListener(topic string, handlerName string, handler interface{}) error {
	if _, ok := mbs.handlers[topic]; !ok {
		mbs.handlers[topic] = map[string]interface{}{}
	}

	if _, ok := mbs.handlers[topic][handlerName]; !ok {
		mbs.handlers[topic][handlerName] = mbs.prepareEvent(handler)
	}

	errorSubscribe := mbs.driver.Subscribe(
		topic,
		mbs.handlers[topic][handlerName],
	)

	if errorSubscribe != nil {
		return errorSubscribe
	}

	return nil
}

func (mbs *MessageBusService) removeEventListener(topic string, handlerName string) error {
	if _, ok := mbs.handlers[topic][handlerName]; !ok {
		return fmt.Errorf("a handler's name '%s' does not exist in a topic '%s'", handlerName, topic)
	}

	errorUnsubscribe := mbs.driver.Unsubscribe(
		topic,
		mbs.handlers[topic][handlerName],
	)

	if errorUnsubscribe != nil {
		return errorUnsubscribe
	}

	return nil
}

func (mbs *MessageBusService) prepareEvent(handler interface{}) interface{} {
	return func(data interface{}) {
		defer mbs.syncWaitingGroup.Done()

		reflect.ValueOf(handler).Call([]reflect.Value{reflect.ValueOf(data)})
		return
	}
}

func (mbs *MessageBusService) Publish(topic string, data interface{}) error {
	if _, ok := mbs.handlers[topic]; !ok {
		return fmt.Errorf("not enough handlers in the topic = %s", topic)
	}

	mbs.syncWaitingGroup.Add(len(mbs.handlers[topic]))
	mbs.driver.Publish(topic, data)
	mbs.syncWaitingGroup.Wait()

	return nil
}

func (mbs *MessageBusService) AddAuthConfirmationEvent() error {
	return mbs.addEventListener(
		event.AuthConfirmationTopicName,
		event.UserConfirmationEventName,
		event.UserConfirmationEvent,
	)
}

func (mbs *MessageBusService) RemoveAuthConfirmationEvent() error {
	return mbs.removeEventListener(
		event.AuthConfirmationTopicName,
		event.UserConfirmationEventName,
	)
}

func GetMessageBusService() MessageBusServiceInterface {
	if messageBusService == nil {
		messageBusService = &MessageBusService{
			driver:           messagebus.New(runtime.NumCPU()),
			syncWaitingGroup: sync.WaitGroup{},
			handlers:         make(map[string]map[string]interface{}),
		}
	}

	return messageBusService
}
