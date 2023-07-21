package service

import (
	"fmt"

	"github.com/Pramessh-P/weather-api/internal/helper"
	"github.com/Pramessh-P/weather-api/internal/model"
	"github.com/Pramessh-P/weather-api/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Signup(user *model.User)error
	Login(user *helper.UserLogin)(error)
	GetRecentActivities(user string)error
}

type userService struct{
	userRepo repository.UserRepository
}

func NewuserService(userRepository repository.UserRepository)*userService{

	return &userService{
		userRepo: userRepository,
	}
}
func(u *userService)Signup(user *model.User)error{
	
	
	hash,err:=bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
	if err !=nil{
		return err
	}
	user.Password=string(hash)
	err=u.userRepo.StoreUser(user)
	
	if err!=nil{
		fmt.Println(err)
		return err
	}
	return nil
}

func (u *userService)Login(user *helper.UserLogin)(error){
	hashedPassword,err:=u.userRepo.LoginRepo(user)
	if err!=nil{
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(user.Password));err!=nil{
		fmt.Println(err)
		return err
	}

	return nil
}
func (u *userService)GetRecentActivities(user string)error{
	err:=u.userRepo.UserActivity(user)
	return err
}
