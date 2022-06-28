package service

import (
	"encoding/json"
	"github.com/sergeygardner/meal-planner-api/domain/dto"
	"io"
)

func CreateDTOFromAuthCredentials(data io.Reader) (dto.UserCredentialsDTO, error) {
	authCredentialsDTO := &dto.UserCredentialsDTO{}
	errorDTO := json.NewDecoder(data).Decode(&authCredentialsDTO)

	return *authCredentialsDTO, errorDTO
}

func CreateDTOFromAuthConfirmation(data io.Reader) (dto.AuthConfirmationDTO, error) {
	authConfirmationDTO := &dto.AuthConfirmationDTO{}
	errorDTO := json.NewDecoder(data).Decode(&authConfirmationDTO)

	return *authConfirmationDTO, errorDTO
}

func CreateDTOFromUserRegister(data io.Reader) (dto.UserRegisterDTO, error) {
	userRegisterDTO := &dto.UserRegisterDTO{}
	errorDTO := json.NewDecoder(data).Decode(&userRegisterDTO)

	return *userRegisterDTO, errorDTO
}

func CreateDTOFromUserUpdate(data io.Reader) (dto.UserDTO, error) {
	userDTO := &dto.UserDTO{}
	errorDTO := json.NewDecoder(data).Decode(&userDTO)

	return *userDTO, errorDTO
}
