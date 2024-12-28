package models

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID                    uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name                  string     `json:"name"`
	Email                 string     `json:"email"`
	Password              string     `json:"password"`
	Telephone             NullString `json:"telephone"`
	Instagram             NullString `json:"instagram"`
	ProfileImage          NullString `json:"profileimage"`
	SubscriptionLevel     NullString `json:"subscriptionlevel"`
	Level                 NullString `json:"level"`
	TeamName              NullString `json:"teamname"`
	LastQuestionnaireDate NullTime   `json:"lastquestionnairedate"`
}

func (User) TableName() string {
	return "msuser"
}

func GetUserByID(db *gorm.DB, id uuid.UUID) (User, error) {
	var user User
	result := db.Where("id = ?", id).First(&user)

	if result.Error != nil {
		return User{}, result.Error
	}

	return user, nil
}

func GetAllUser(db *gorm.DB) ([]User, error) {
	var users []User

	result := db.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	fmt.Println(result)

	return users, nil
}

func GetUserByEmail(db *gorm.DB, email string) (User, error) {
	var user User
	result := db.Where("email = ?", email).First(&user)

	if result.Error != nil {
		return User{}, result.Error
	}

	return user, nil
}

func CheckUserExistsByEmail(db *gorm.DB, email string) (bool, error) {
	var user = User{ID: uuid.Nil}
	result := db.Where("email = ?", email).First(&user)

	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return false, result.Error
	}

	if user.ID == uuid.Nil {
		return false, nil
	}

	return true, nil
}

func CreateUser(db *gorm.DB, name, email, password string) (User, error) {
	var user = User{ID: uuid.New(), Name: name, Email: email, Password: password}
	result := db.Create(&user)

	if result.Error != nil {
		return User{}, result.Error
	}

	return user, nil
}

func UpdateUser(db *gorm.DB, user User) (User, error) {
	result := db.Save(&user)

	if result.Error != nil {
		return User{}, result.Error
	}

	return user, nil
}
