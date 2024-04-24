package dto

import "github.com/dna-technology/dna-task-livecoding-golang/internal/pkg/db/user"

type UserDto struct {
	UserId   string
	FullName string
	Email    string
}

func (u *UserDto) FromEntity(entity user.User) UserDto {
	return UserDto{
		UserId:   entity.UserId,
		FullName: entity.FullName,
		Email:    entity.Email,
	}
}

func (u *UserDto) ToEntity() user.User {
	return user.User{
		UserId:   u.UserId,
		FullName: u.FullName,
		Email:    u.Email,
	}
}
