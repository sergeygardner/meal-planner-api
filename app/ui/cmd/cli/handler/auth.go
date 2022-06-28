package handler

import (
	"context"
	"github.com/go-chi/jwtauth/v5"
	"github.com/pkg/errors"
	ApplicationHandler "github.com/sergeygardner/meal-planner-api/application/handler"
	"github.com/sergeygardner/meal-planner-api/domain/dto"
	"github.com/sergeygardner/meal-planner-api/domain/model"
	"github.com/sergeygardner/meal-planner-api/infrastructure/service/repository"
	UiService "github.com/sergeygardner/meal-planner-api/ui/service"
	"reflect"
	"time"
)

var (
	userCredentialsDTO    *dto.UserCredentialsDTO
	authConfirmationDTO   *dto.AuthConfirmationDTO
	userRegisterDTO       *dto.UserRegisterDTO
	errorAuthConfirmation error
	errorWeirdBehaviour   = errors.New("an error occurred while running command. Weird behaviour!")
	errorAuthentication   = errors.New("an error occurred while running command. You are not authenticated!")
)

//func AuthCheck(w http.ResponseWriter, _ *http.Request) {
//	w.WriteHeader(http.StatusOK)
//}

func auth(_ string) (int, error) {
	userRepository := repository.GetFactoryRepository().GetUserRepository()
	user, errorUser := userRepository.FindOne(userRepository.GetCriteriaByUsername("username"))

	if errorUser != nil {
		return StatusError, errorUser
	}

	newAuthToken, errorAuthToken := ApplicationHandler.AuthTokenByUserId(&user.Id)

	if errorAuthToken != nil {
		return StatusError, errorAuthToken
	} else {
		authToken = newAuthToken

		showInfoMessage("You are authenticated", "")

		return StatusOk, nil
	}
}

func authCredentials(message string) (int, error) {
	if userCredentialsDTO == nil {
		userCredentialsDTO = &dto.UserCredentialsDTO{}
		showDialogMessage("your username")
	} else if userCredentialsDTO.Username == "" {
		userCredentialsDTO.Username = message
		showDialogMessage("your password")
	} else if userCredentialsDTO.Password == "" {
		userCredentialsDTO.Password = message

		responseAuthConfirmation, _, errorAuthCredentials := ApplicationHandler.AuthCredentials(*userCredentialsDTO)

		if errorAuthCredentials != nil {
			return StatusError, errorAuthCredentials
		} else {
			showInfoMessage(responseAuthConfirmation.Message, "")

			return StatusOk, nil
		}
	} else {
		return StatusError, errorWeirdBehaviour
	}

	return StatusContinue, nil
}

func authConfirmation(message string) (int, error) {
	if authConfirmationDTO == nil {
		authToken = nil
		authConfirmationDTO = &dto.AuthConfirmationDTO{UserCredentialsDTO: *userCredentialsDTO}
		showDialogMessage("your confirmation code")
	} else if authConfirmationDTO.Code == "" {
		authConfirmationDTO.Code = message

		authToken, errorAuthConfirmation = ApplicationHandler.AuthConfirmation(*authConfirmationDTO)

		if errorAuthConfirmation != nil {
			return StatusError, errorAuthConfirmation
		} else {
			authConfirmationDTO = nil
			userCredentialsDTO = nil

			showInfoMessage("You are authenticated", "")

			return StatusOk, nil
		}
	} else {
		return StatusError, errorWeirdBehaviour
	}

	return StatusContinue, nil
}

func authRefresh(_ string) (int, error) {
	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}
	newAuthToken, errorAuthToken := ApplicationHandler.AuthTokenByUserId(&token.UserId)

	if errorAuthToken != nil {
		return StatusError, errorAuthToken
	} else {
		authToken = newAuthToken

		return StatusOk, nil
	}
}

func authRegister(message string) (int, error) {
	if userRegisterDTO == nil {
		userRegisterDTO = &dto.UserRegisterDTO{}
		showDialogMessage("your username")
	} else if userRegisterDTO.Username == "" {
		userRegisterDTO.Username = message
		showDialogMessage("your password")
	} else if userRegisterDTO.Password == "" {
		userRegisterDTO.Password = message
		showDialogMessage("your name")
	} else if userRegisterDTO.Name == "" {
		userRegisterDTO.Name = message
		showDialogMessage("your surname")
	} else if userRegisterDTO.Surname == "" {
		userRegisterDTO.Surname = message
		showDialogMessage("your middle name")
	} else if userRegisterDTO.MiddleName == "" {
		userRegisterDTO.MiddleName = message
		showDialogMessage("your birthday (YYYY-MM-DDT00:00:00Z) [RFC3339]")

	} else if reflect.ValueOf(userRegisterDTO.Birthday).IsZero() {
		parsedDate, errorParsedDate := time.Parse(time.RFC3339, message)

		if errorParsedDate != nil {
			return StatusError, errorParsedDate
		}

		userRegisterDTO.Birthday = parsedDate

		_, errorAuthRegister := ApplicationHandler.AuthRegister(*userRegisterDTO)

		if errorAuthRegister != nil {
			return StatusError, errorAuthRegister
		} else {
			showInfoMessage("You are registered", "")
			return StatusOk, nil
		}
	}

	return StatusContinue, nil
}

func tokenFromContext() (*model.Token, error) {
	if authToken == nil {
		return nil, errorAuthentication
	}

	jwtToken, errorVerifyToken := jwtauth.VerifyToken(jwtAuth, authToken.RefreshToken)

	if errorVerifyToken != nil {
		return nil, errorVerifyToken
	}

	contextWithJwt := jwtauth.NewContext(context.TODO(), jwtToken, errorVerifyToken)

	return UiService.ExtractClaimsFromContext(contextWithJwt)
}
