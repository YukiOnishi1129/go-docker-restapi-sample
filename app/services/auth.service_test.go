package services_test

import (
	"myapp/models"
	"myapp/repositories"
	"myapp/utils/logic"
	"myapp/utils/validation"
)

type userRepoMock struct {
	repositories.UserRepository
	FakeGetUserByEmail func(user *models.User, email string) error
	FakeGetAllUserByEmail func(users *[]models.User, email string) error
	FakeCreateUser func(createUsers *models.User) error
}

type authLogicMock struct {
	logic.AuthLogic
}

type userLogicMock struct {
	logic.UserLogic
}

type responseLogicMock struct {
	logic.ResponseLogic
}

type jwtLogicMock struct {
	logic.JWTLogic
}

type authValidation struct {
	validation.AuthValidation
}