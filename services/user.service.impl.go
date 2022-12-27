package service

import (
	"example.com/aorts/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	usercollection *mongo.Collection
	ctx            contxt.Context
}

func (u *UserServiceImpl) CreateUser(user *model.User) error {
	return nil
}

func (u *UserServiceImpl) GetUser(name *string) (*models.User, error) {
	return nil, nil
}

func (u *UserServiceImpl) UpdateUser(user *models.User) error {
	return nil
}

func (u *UserServiceImpl) DeleteUser(name *string) error {

}
