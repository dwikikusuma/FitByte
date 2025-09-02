package service

import (
	customErrors "FitByte/internal/errors"
	"FitByte/internal/models"
	"FitByte/internal/repositories"
	"FitByte/pkg/log"
	"context"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(ctx context.Context, user models.User) error
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (u *userService) Register(ctx context.Context, user models.User) error {
	isUserExist, err := u.userRepo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		log.Logger.Error().Err(err).Msg("error occurred on Register(ctx context.Context, user models.User)")
		return err
	}

	if isUserExist != nil {
		log.Logger.Warn().Str("email", user.Email).Msg("user already exists")
		return customErrors.ErrUserAlreadyExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Logger.Error().Err(err).Msg("error occurred on Register(ctx context.Context, user models.User)")
		return err
	}

	user.Password = string(hashedPassword)
	err = u.userRepo.CreateUser(ctx, user)
	if err != nil {
		log.Logger.Error().Err(err).Msg("error occurred on Register(ctx context.Context, user models.User)")
		return err
	}

	return nil
}
