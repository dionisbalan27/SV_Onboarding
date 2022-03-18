package user_repository

import (
	helpers "backend-api/helpers/helpers_user"
	"backend-api/models/user/dto"
	"backend-api/models/user/entity"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers() ([]entity.UserList, error)
	GetUserById(string) (*entity.UserID, error)
	CreateNewUser(entity.User) (*entity.User, error)
	UpdateUserData(entity.User, string) (*entity.User, error)
	DeleteUserById(string) error
	CheckLogin(dto.Login) (string, error)
}

type userRepository struct {
	mysqlConnection *gorm.DB
}

func GetUserRepository(mysqlConn *gorm.DB) UserRepository {
	return &userRepository{
		mysqlConnection: mysqlConn,
	}
}

func (repo *userRepository) GetAllUsers() ([]entity.UserList, error) {

	users := []entity.UserList{}
	err := repo.mysqlConnection.Model(&entity.User{}).Select("users.name, users.active, users.id, roles.title, users.role_id").Joins("left join roles on roles.id = users.role_id").Scan(&users).Error
	if err != nil {
		return nil, err
	}

	if len(users) <= 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return users, nil
}

func (repo *userRepository) GetUserById(id string) (*entity.UserID, error) {
	user := entity.UserID{}

	err := repo.mysqlConnection.Model(&entity.User{}).Where("users.id = ?", id).Select("users.name, users.active, users.id, roles.title, users.role_id, users.email, users.personal_number").Joins("left join roles on roles.id = users.role_id").Scan(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *userRepository) CreateNewUser(user entity.User) (*entity.User, error) {
	user.ID = uuid.New().String()
	user.RoleID = uuid.New().String()

	if err := repo.mysqlConnection.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *userRepository) UpdateUserData(user entity.User, id string) (*entity.User, error) {

	if err := repo.mysqlConnection.Model(&user).Where("id = ?", id).Updates(map[string]interface{}{
		"name":            user.Name,
		"role_id":         user.RoleID,
		"active":          user.Active,
		"email":           user.Email,
		"personal_number": user.Personal_number,
	}).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *userRepository) DeleteUserById(id string) error {
	sql := "DELETE FROM users"

	sql = fmt.Sprintf("%s WHERE id = '%s'", sql, id)

	if err := repo.mysqlConnection.Raw(sql).Scan(entity.User{}).Error; err != nil {

		return err
	}
	// if err := repo.mysqlConnection.Delete(&entity.User{}, id).Error; err != nil  {
	// 	return err
	// }
	return nil
}

func (repo *userRepository) CheckLogin(user dto.Login) (string, error) {
	data := entity.CheckLogin{}

	err := repo.mysqlConnection.Model(&entity.User{}).Where("users.personal_number = ?", user.Personal_number).Select("users.password, users.id").Scan(&data).Error

	if err != nil {
		fmt.Println("Query error")
		return "", err
	}

	match, err := helpers.CheckPasswordHash(user.Password, data.Password)
	if !match {
		fmt.Println("Hash and password doesn't match.")
		return "", err
	}

	return data.ID, nil
}
