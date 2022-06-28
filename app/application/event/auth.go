package event

import (
	DomainEntity "github.com/sergeygardner/meal-planner-api/domain/entity"
	log "github.com/sirupsen/logrus"
)

const (
	AuthConfirmationTopicName = "auth:confirmation"
	UserConfirmationEventName = "UserConfirmationEvent"
)

var userConfirmationInActivePanicMessage = "an error occurred while executing UserConfirmationEvent. The UserConfirmation is inactive."

// UserConfirmationEvent sends confirmation to the user. It is Panic while UserConfirmation.Active is false
func UserConfirmationEvent(userConfirmation *DomainEntity.UserConfirmation) {
	if !userConfirmation.Active {
		panic(userConfirmationInActivePanicMessage)
	} else {
		log.Infof("I have sent the confirmation with code '%s'\n", userConfirmation.Value)
	}
}
