package mapper

import (
	"pirate-lang-go/modules/account/dto"
	"pirate-lang-go/modules/account/entity"
)

func ToUserEntity(user *dto.CreateAccountRequest) *entity.User {

	if user == nil {
		return nil
	}

	return &entity.User{
		UserName: user.Username,
		Password: user.Password,
		Email:    user.Email,
	}
}
func ToPaginatedUsersResponse(users *entity.PaginatedUsers) *dto.PaginatedUsersResponse {
	if users == nil {
		return nil
	}

	userDtos := make([]*dto.UserResponse, 0, len(users.Items))
	for _, user := range users.Items {
		userDtos = append(userDtos, &dto.UserResponse{
			Id:            user.ID,
			Username:      user.UserName,
			Email:         user.Email,
			IsSocialLogin: user.IsSocialLogin,
			IsLocked:      user.IsLocked,
		})
	}

	return &dto.PaginatedUsersResponse{
		Items:       userDtos,
		TotalItems:  users.TotalItems,
		TotalPages:  users.TotalPages,
		CurrentPage: users.CurrentPage,
		PageSize:    users.PageSize,
	}
}
