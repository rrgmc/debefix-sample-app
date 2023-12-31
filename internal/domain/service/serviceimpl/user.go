package serviceimpl

import (
	"context"

	"github.com/google/uuid"
	"github.com/rrgmc/debefix-sample-app/internal/domain/entity"
	"github.com/rrgmc/debefix-sample-app/internal/domain/repository"
	"github.com/rrgmc/debefix-sample-app/internal/domain/service"
	"github.com/rrgmc/debefix-sample-app/internal/domain/validator"
)

type userService struct {
	rctx           repository.Context
	userRepository repository.UserRepository
	userValidator  validator.UserValidator
}

func NewUserService(rctx repository.Context, userRepository repository.UserRepository,
	countryService service.CountryService) service.UserService {
	return &userService{
		rctx:           rctx,
		userRepository: userRepository,
		userValidator:  validator.NewUserValidator(countryService),
	}
}

func (d userService) GetUserList(ctx context.Context, filter entity.UserFilter) ([]entity.User, error) {
	err := d.userValidator.ValidateUserFilter(ctx, filter)
	if err != nil {
		return nil, err
	}
	return d.userRepository.GetUserList(ctx, d.rctx, filter)
}

func (d userService) GetUserByID(ctx context.Context, userID uuid.UUID) (entity.User, error) {
	return d.userRepository.GetUserByID(ctx, d.rctx, userID)
}

func (d userService) ExistsUserByID(ctx context.Context, userID uuid.UUID) (bool, error) {
	return d.userRepository.ExistsUserByID(ctx, d.rctx, userID)
}

func (d userService) AddUser(ctx context.Context, user entity.UserAdd) (entity.User, error) {
	err := d.userValidator.ValidateUserAdd(ctx, user)
	if err != nil {
		return entity.User{}, err
	}
	return d.userRepository.AddUser(ctx, d.rctx, user)
}

func (d userService) UpdateUserByID(ctx context.Context, userID uuid.UUID, user entity.UserUpdate) (entity.User, error) {
	err := d.userValidator.ValidateUserUpdate(ctx, user)
	if err != nil {
		return entity.User{}, err
	}
	return d.userRepository.UpdateUserByID(ctx, d.rctx, userID, user)
}

func (d userService) DeleteUserByID(ctx context.Context, userID uuid.UUID) error {
	return d.userRepository.DeleteUserByID(ctx, d.rctx, userID)
}
