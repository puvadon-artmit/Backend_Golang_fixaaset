package AuthServices

import (
	"errors"

	"github.com/google/uuid"
	AuthDto "github.com/puvadon-artmit/gofiber-template/api/auth/entitys/request"
	authResponse "github.com/puvadon-artmit/gofiber-template/api/auth/entitys/response"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"github.com/puvadon-artmit/gofiber-template/utils"
)

// func GetProfileById(id string) (Result *authResponse.ResultGetProfile, Error error) {
// 	db := database.DB
// 	result := authResponse.ResultGetProfile{}
// 	tx := db.Table("users").Select("user_id, username, firstname, lastname").Where("user_id=?", id).Scan(&result)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}
// 	return &result, nil
// }

func GetProfileById(userID string) (result []*model.User, err error) {
	db := database.DB
	var records []*model.User
	tx := db.Where("user_id = ?", userID).Preload("Roles.PermissionGroup.Permission").Preload("Roles.PermissionGroup.PermissionComponent").Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func SignIn(dto AuthDto.LoginDto) (token *string, Error error) {
	db := database.DB
	var user []model.User
	db.Find(&user, "username = ?", dto.Username)
	if len(user) == 0 {
		return nil, errors.New("username not found")
	}
	if !utils.Compare(dto.Password, user[0].Password) {
		return nil, errors.New("password does not match")
	}
	tokenString, err := utils.GenerateJwt(user[0])
	if err != nil {
		return nil, errors.New("failed to generate token, please try again")
	}
	return &tokenString, nil
}

// func SignUp(user model.User) (result *model.User, Error error) {
// 	db := database.DB
// 	var userDB []model.User
// 	db.Find(&userDB, "username = ?", user.Username)
// 	if len(userDB) != 0 {
// 		return nil, errors.New("User Already has!")
// 	}
// 	var password = user.Password
// 	user.Password = utils.Encode(password)
// 	user.UserID = uuid.New().String()
// 	err := db.Create(&user).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return nil, nil
// }

func SignUp(user model.User) (result *model.User, Error error) {
	db := database.DB
	var userDB []model.User
	db.Find(&userDB, "username = ?", user.Username)
	if len(userDB) != 0 {
		return nil, errors.New("user already exists")
	}
	var password = user.Password
	user.Password = utils.Encode(password)
	user.UserID = uuid.New().String()
	err := db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// func SignIn(dto AuthDto.LoginDto) (token *string, Error error) {
// 	db := database.DB
// 	var user []model.User
// 	db.Find(&user, "username = ?", dto.Username)
// 	if len(user) == 0 {
// 		return nil, errors.New("Username not found")
// 	}
// 	if !utils.Compare(dto.Password, user[0].Password) {
// 		return nil, errors.New("Password is not compare!")
// 	}
// 	tokenString, err := utils.GenerateJwt(user[0])
// 	if err != nil {
// 		return nil, errors.New("Generate token fail, try again.")
// 	}
// 	return &tokenString, nil
// }

func GetAll() (value *[]authResponse.ResultGetProfile, err error) {
	db := database.DB
	var result []authResponse.ResultGetProfile
	tx := db.Table("users").Select("user_id, username, firstname, lastname").Scan(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &result, nil
}

func CreateNewUser(user model.User, roles []model.Role) (value *model.User, Error error) {
	db := database.DB
	user.UserID = uuid.New().String()

	// Hash the password
	user.Password = utils.Encode(user.Password)

	tx := db.Create(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	for i := range roles {
		userroles := model.UserRoles{
			UserID: user.UserID,
			RoleID: roles[i].RoleID,
		}
		tx := db.Create(&userroles)
		if tx.Error != nil {
			return nil, tx.Error
		}
	}

	return &user, nil
}
