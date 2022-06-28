package handler

import (
	"errors"
	"github.com/google/uuid"
	"github.com/sergeygardner/meal-planner-api/application/handler"
	"github.com/sergeygardner/meal-planner-api/domain/dto"
	DomainEntity "github.com/sergeygardner/meal-planner-api/domain/entity"
	"github.com/sergeygardner/meal-planner-api/domain/kind"
	"reflect"
	"time"
)

var (
	userDTO                 *dto.UserDTO
	userId                  *uuid.UUID
	statusUserDeleteSuccess = "the recipe user has been deleted successful"
	statusUserDeleteError   = errors.New("the recipe user has not been deleted")
)

func userInfo(message string) (int, error) {
	var (
		userIdValue uuid.UUID
		errorUserId error
	)

	if message == "UserInfo" {
		showDialogMessage("input id for User")

		return StatusContinue, nil
	}

	if userId == nil {
		userIdValue, errorUserId = uuid.Parse(message)

		userId = &userIdValue
	} else {
		errorUserId = nil
	}

	if errorUserId != nil {
		return StatusError, errorUserId
	} else {
		user, errorUser := handler.UserInfo(userId)

		userId = nil

		if errorUser != nil {
			return StatusError, errorUser
		} else {
			printTable("User", []*DomainEntity.User{user}, DomainEntity.User{})

			return StatusOk, nil
		}
	}
}

func userUpdate(message string) (int, error) {
	var (
		userIdValue uuid.UUID
		errorUserId error
	)

	if message == "UserUpdate" {
		showDialogMessage("input id for User")

		return StatusContinue, nil
	}

	if userId == nil {
		userIdValue, errorUserId = uuid.Parse(message)

		userId = &userIdValue
	} else {
		errorUserId = nil
	}

	if errorUserId != nil {
		return StatusError, errorUserId
	} else if userDTO == nil {
		userDTO = &dto.UserDTO{}
		userDTO.DateUpdate = time.Now().UTC()
		showDialogMessage("your password")
	} else if userDTO.Password == "" {
		userDTO.Password = message
		showDialogMessage("your name")
	} else if userDTO.Name == "" {
		userDTO.Name = message
		showDialogMessage("your surname")
	} else if userDTO.Surname == "" {
		userDTO.Surname = message
		showDialogMessage("your middle name")
	} else if userDTO.MiddleName == "" {
		userDTO.MiddleName = message
		showDialogMessage("your birthday (YYYY-MM-DDT00:00:00Z) [RFC3339]")

	} else if reflect.ValueOf(userDTO.Birthday).IsZero() {
		parsedDate, errorParsedDate := time.Parse(time.RFC3339, message)

		if errorParsedDate != nil {
			return StatusError, errorParsedDate
		}

		userRegisterDTO.Birthday = parsedDate

		showDialogMessage("input status for User. choose from (%v,%v,%v)", kind.UserStatusDisabled, kind.UserStatusRegister, kind.UserStatusNeedConfirmation)
	} else if userDTO.Status == "" {
		userDTO.Status = kind.UserStatus(message)

		user, errorUser := handler.UserUpdate(userId, userDTO)

		userId = nil
		userDTO = nil

		if errorUser != nil {
			return StatusError, errorUser
		} else {
			printTable("User", []*DomainEntity.User{user}, DomainEntity.User{})

			return StatusOk, nil
		}
	}

	return StatusContinue, nil
}

func userDelete(message string) (int, error) {
	var (
		userIdValue uuid.UUID
		errorUserId error
	)

	if message == "UserDelete" {
		showDialogMessage("input id for User")

		return StatusContinue, nil
	}

	if userId == nil {
		userIdValue, errorUserId = uuid.Parse(message)

		userId = &userIdValue
	} else {
		errorUserId = nil
	}

	if errorUserId != nil {
		return StatusError, errorUserId
	} else {
		userDeleteStatus, errorUserDeleteStatus := handler.UserDelete(userId)

		userId = nil

		if errorUserDeleteStatus != nil {
			return StatusError, errorUserDeleteStatus
		} else if userDeleteStatus {
			showInfoMessage(statusUserDeleteSuccess)

			return StatusOk, nil
		} else {
			return StatusError, statusUserDeleteError
		}
	}
}
