package repository

import (
	// "fmt"

	// "github.com/Pramessh-P/weather-api/internal/database"
	"fmt"

	"github.com/Pramessh-P/weather-api/internal/database"
	"github.com/Pramessh-P/weather-api/internal/helper"
	"github.com/Pramessh-P/weather-api/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	StoreUser(user *model.User) error
	LoginRepo(user *helper.UserLogin)(string,error)
	UserActivity(user string)error
}

type UserInMemoryRepository struct {
	userData map[string]*model.User
}

func NewUserRepository() *UserInMemoryRepository {
	return &UserInMemoryRepository{
		userData: make(map[string]*model.User),
	}
}

func (r *UserInMemoryRepository) StoreUser(user *model.User)  error {
	r.userData[user.Email]=user
	// var userData *model.User

	result :=database.DB.Create(user)
	if result.Error != nil {
        return result.Error
    }
	fmt.Println("users added database to successfully")
	return  nil
}
func (r *UserInMemoryRepository)LoginRepo(user *helper.UserLogin)(string,error){
	userEmail:=user.Email
	hashedPass,err:=getPasswordByEmail(userEmail)
	return hashedPass,err
}

func getPasswordByEmail(email string) (string, error) {
    var user model.User

    result := database.DB.Where("email = ?", email).First(&user)
    if result.Error != nil {
        return "", result.Error
    }

    return user.Password, nil
}
func (r *UserInMemoryRepository)UserActivity(user string)error{
	var users model.User
    result := database.DB.Where("email = ?", user).First(&users)
	if result.Error!=nil{
		if result.Error==gorm.ErrRecordNotFound{
			return result.Error
		}
		return result.Error
	}
	return nil
}