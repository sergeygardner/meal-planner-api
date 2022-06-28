package handler

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/sergeygardner/meal-planner-api/application/event"
	ApplicationServicePassword "github.com/sergeygardner/meal-planner-api/application/service/password"
	"github.com/sergeygardner/meal-planner-api/application/service/update"
	"github.com/sergeygardner/meal-planner-api/domain/dto"
	"github.com/sergeygardner/meal-planner-api/domain/entity"
	"github.com/sergeygardner/meal-planner-api/domain/model"
	"github.com/sergeygardner/meal-planner-api/domain/response"
	InfrastructureService "github.com/sergeygardner/meal-planner-api/infrastructure/service"
	ApplicationMiddlewareJWT "github.com/sergeygardner/meal-planner-api/infrastructure/service/jwt"
	"github.com/sergeygardner/meal-planner-api/infrastructure/service/repository"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

var (
	errorUserNotFound            = errors.New("user is not found by credentials")
	errorUserRegister            = errors.New("user has not registered by credentials")
	errorAuthConfirmationNotSent = errors.New("the server hasn't been sent the confirmation")
)

func AuthCredentials(authCredentialsDTO dto.UserCredentialsDTO) (*response.AuthConfirmation, *entity.UserConfirmation, error) {
	userRepository := repository.GetFactoryRepository().GetUserRepository()
	user, errorUser := userRepository.FindOne(userRepository.GetCriteriaByUsername(authCredentialsDTO.Username))
	passwordOk := ApplicationServicePassword.CheckPassword(authCredentialsDTO.Password, user.Password)

	if errorUser != nil || user == nil || !passwordOk {
		return nil, nil, errorUserNotFound
	} else {
		userConfirmationRepository := repository.GetFactoryRepository().GetUserConfirmationRepository()
		userConfirmation, errorConfirmationFindOne := userConfirmationRepository.FindOne(userConfirmationRepository.GetCriteriaByUserIdAndActive(user))

		if errorConfirmationFindOne != nil {
			userConfirmationInsertOne, errorConfirmationInsertOne := userConfirmationRepository.InsertOne(prepareUserConfirmationRepositoryInsert(*user))
			userConfirmation = userConfirmationInsertOne

			if errorConfirmationInsertOne != nil {
				return nil, nil, errors.Wrapf(errorAuthConfirmationNotSent, "an error occurred while inserting a confirmation in the database by provided data %s", authCredentialsDTO)
			}
		}

		messageBusService := InfrastructureService.GetMessageBusService()
		errorMessageBusServiceAddEventListener := messageBusService.AddAuthConfirmationEvent()

		if errorMessageBusServiceAddEventListener != nil {
			return nil, nil, errors.Wrap(errorAuthConfirmationNotSent, "an error occurred while adding a confirmation event")
		} else {
			defer func(messageBusService InfrastructureService.MessageBusServiceInterface, topic string, data interface{}) {
				errorMessageBusServicePublish := messageBusService.Publish(topic, data)

				if errorMessageBusServicePublish != nil {
					log.Error(errors.Wrapf(errorMessageBusServicePublish, "an error occurred while publising a confirmation by provided data topic=%s,data=%s", topic, data))
					return
				}

				errorMessageBusServiceRemoveEventListener := messageBusService.RemoveAuthConfirmationEvent()
				if errorMessageBusServiceRemoveEventListener != nil {
					log.Error(errors.Wrap(errorMessageBusServiceRemoveEventListener, "an error occurred while removing a confirmation event"))
					return
				}
			}(messageBusService, event.AuthConfirmationTopicName, userConfirmation)
		}

		return &response.AuthConfirmation{Message: "The server has been sent the confirmation", Status: http.StatusOK}, userConfirmation, nil
	}
}

func AuthConfirmation(authConfirmationDTO dto.AuthConfirmationDTO) (*response.AuthToken, error) {
	userRepository := repository.GetFactoryRepository().GetUserRepository()
	userConfirmationRepository := repository.GetFactoryRepository().GetUserConfirmationRepository()
	user, errorUserFindOne := userRepository.FindOne(userRepository.GetCriteriaByUsername(authConfirmationDTO.UserCredentialsDTO.Username))
	passwordOk := ApplicationServicePassword.CheckPassword(authConfirmationDTO.Password, user.Password)

	if errorUserFindOne != nil {
		return nil, errors.Wrapf(errorUserFindOne, "an error occurred while getting a user by provided data %s", authConfirmationDTO)
	} else if !passwordOk {
		return nil, errors.Wrapf(errorUserNotFound, "an error occurred while checking a password for a user by provided data password=%s", authConfirmationDTO.Password)
	} else {
		userConfirmation, errorUserConfirmationFindOne := userConfirmationRepository.FindOne(userConfirmationRepository.GetCriteriaByUserIdAndActive(user))

		if errorUserConfirmationFindOne != nil || userConfirmation.Value != authConfirmationDTO.Code {
			return nil, errorUserNotFound
		} else {
			errorSetUserConfirmationInActive := update.SetUserConfirmationInActive(userConfirmation)

			if errorSetUserConfirmationInActive != nil {
				return nil, errors.Wrapf(errorSetUserConfirmationInActive, "an error occurred while setting a confirmation to the inactive status a user by provided data %v", userConfirmation)
			}

			return AuthToken(user)
		}
	}
}

func AuthToken(user *entity.User) (*response.AuthToken, error) {
	authToken, errorMakeTokens := makeTokens(user)

	if errorMakeTokens != nil {
		return nil, errors.Wrapf(errorMakeTokens, "an error occurred while making tokens by provided data %v", user)
	} else {
		return authToken, nil
	}
}

func AuthTokenByUserId(userId *uuid.UUID) (*response.AuthToken, error) {
	userRepository := repository.GetFactoryRepository().GetUserRepository()
	user, errorUserFindOne := userRepository.FindOne(userRepository.GetCriteriaByUserId(userId))

	if errorUserFindOne != nil {
		return nil, errors.Wrapf(errorUserFindOne, "an error occurred while getting a user by provided data userId=%s", userId)
	} else {
		return AuthToken(user)
	}
}

func AuthRegister(userRegisterDTO dto.UserRegisterDTO) (*entity.User, error) {
	userRepository := repository.GetFactoryRepository().GetUserRepository()

	_, errorUserFindOne := userRepository.FindOne(userRepository.GetCriteriaByUsername(userRegisterDTO.UserCredentialsDTO.Username))

	if errorUserFindOne == nil {
		return nil, errorUserRegister
	} else {
		password, errorCastPassword := ApplicationServicePassword.CastPassword(userRegisterDTO.Password)

		if errorCastPassword != nil {
			return nil, errors.Wrapf(errorCastPassword, "an error occurred while casting a password for a user by provided data password=%s", userRegisterDTO.Password)
		} else {
			user, errorUserInsertOne := userRepository.InsertOne(prepareUserRepositoryInsert(userRegisterDTO, string(password)))

			if errorUserInsertOne != nil {
				return nil, errors.Wrapf(errorUserInsertOne, "an error occurred while creating a user in the database by privided data %s", userRegisterDTO)
			} else {
				return user, nil
			}
		}
	}
}

func getNewTokenSignedString(user *entity.User, expiresAt time.Duration) (string, error) {
	accessToken := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		model.Token{
			UserId:    user.Id,
			Username:  user.Username,
			UserRoles: user.Roles,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(expiresAt)),
			},
		},
	)

	return accessToken.SignedString(ApplicationMiddlewareJWT.GetJwtKey())
}

func makeTokens(user *entity.User) (*response.AuthToken, error) {
	accessTokenSignedString, errorAccessTokenSignedString := getNewTokenSignedString(user, model.JWTAccessTokenExpire)
	refreshTokenSignedString, errorRefreshTokenSignedString := getNewTokenSignedString(user, model.JWTRefreshTokenExpire)

	if errorAccessTokenSignedString != nil {
		return nil, errors.Wrapf(errorAccessTokenSignedString, "an error occurred while creating an access token for a user by privided data %v", user)
	} else if errorRefreshTokenSignedString != nil {
		return nil, errors.Wrapf(errorRefreshTokenSignedString, "an error occurred while creating an refresh token for a user by privided data %v", user)
	} else {
		return &response.AuthToken{
			AccessToken:  accessTokenSignedString,
			RefreshToken: refreshTokenSignedString,
		}, nil
	}
}
