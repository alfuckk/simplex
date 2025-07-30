package svc

import (
	"context"
	v1 "simplex/app/user/api/v1"
	"simplex/app/user/internal/repo"
)

type UserService interface {
	GetProfile(ctx context.Context, userId string) (*v1.GetProfileResponseData, error)
	UpdateProfile(ctx context.Context, userId string, req *v1.UpdateProfileRequest) error
}

func NewUserService(
	service *Service,
	userRepo repo.UserRepository,
) UserService {
	return &userService{
		userRepo: userRepo,
		Service:  service,
	}
}

type userService struct {
	userRepo repo.UserRepository
	*Service
}

func (s *userService) GetProfile(ctx context.Context, userId string) (*v1.GetProfileResponseData, error) {
	user, err := s.userRepo.GetByID(ctx, userId)
	if err != nil {
		return nil, err
	}

	return &v1.GetProfileResponseData{
		UserId:   user.UserId,
		Nickname: user.Nickname,
	}, nil
}

func (s *userService) UpdateProfile(ctx context.Context, userId string, req *v1.UpdateProfileRequest) error {
	user, err := s.userRepo.GetByID(ctx, userId)
	if err != nil {
		return err
	}

	user.Email = req.Email
	user.Nickname = req.Nickname

	if err = s.userRepo.Update(ctx, user); err != nil {
		return err
	}

	return nil
}
