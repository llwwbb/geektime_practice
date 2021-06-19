package data

import (
	"context"
	"github.com/llwwbb/geektime_practice/go_advanced_training/week4/app/user/service/internal/biz"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Id   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
	Age  int                `bson:"age"`
}

var _ biz.UserRepo = &userRepo{}

type userRepo struct {
	coll *mongo.Collection
}

func NewUserRepo(data *Data) biz.UserRepo {
	return &userRepo{coll: data.db.Collection("user")}
}

func (u *userRepo) GetUserById(ctx context.Context, id string) (*biz.User, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var user User
	err = u.coll.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &biz.User{
		Id:   user.Id.Hex(),
		Name: user.Name,
		Age:  user.Age,
	}, nil
}
