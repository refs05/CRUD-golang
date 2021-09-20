package database

import (
	"crudgoeg/config"
	"crudgoeg/models"
)

func GetUsers() (*[]models.User, error) {
	var users []models.User

	query := config.DB

	if err := query.Unscoped().Find(&users).Error; err != nil {
		return &[]models.User{}, err
	}

	return &users, nil
}

func GetbyIDUser(id int) (*models.User, error) {
	var user models.User

	if err := config.DB.Unscoped().Where("id = ?", id).First(&user).Error; err != nil {
		return &models.User{}, err
	}

	return &user, nil
}

func StoreUser(user models.User) (*models.User, error) {
	if err := config.DB.Save(&user).Error; err != nil {
		return &models.User{}, err
	}

	return &user, nil
}

func Update(id int) (*models.User, error) {
	var user models.User

	if err := config.DB.Model(&user).Where("id = ?", id).Updates(&user).Error; err != nil {
		return &models.User{}, err
	}

	return &user, nil
}

func Delete(id int) (*models.User, error) {
	var user models.User

	if err := config.DB.Where("id = ?", id).Delete(&user).Error; err != nil {
		return &models.User{}, err
	}

	return &user, nil
}