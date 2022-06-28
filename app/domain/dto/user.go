package dto

import (
	"github.com/sergeygardner/meal-planner-api/domain/kind"
	"time"
)

type UserCredentialsDTO struct {
	Username string `protobuf:"bytes,1,opt,name=username,proto3" bson:"username" json:"username"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" bson:"password" json:"password"`
}

type UserRegisterDTO struct {
	UserCredentialsDTO
	Name       string    `protobuf:"bytes,3,opt,name=name,proto3" bson:"name" json:"name"`
	Surname    string    `protobuf:"bytes,4,opt,name=surname,proto3" bson:"surname" json:"surname"`
	MiddleName string    `protobuf:"bytes,5,opt,name=middle_name,proto3" bson:"middle_name" json:"middle_name"`
	Birthday   time.Time `protobuf:"bytes,6,opt,name=birthday,proto3" bson:"birthday" json:"birthday"`
}

type UserDTO struct {
	UserRegisterDTO
	DateInsert time.Time       `protobuf:"bytes,7,opt,name=date_insert,proto3" bson:"date_insert" json:"date_insert"`
	DateUpdate time.Time       `protobuf:"bytes,8,opt,name=date_update,proto3" bson:"date_update" json:"date_update"`
	Status     kind.UserStatus `protobuf:"bytes,9,opt,name=status,proto3" bson:"status" json:"status"`
	Active     bool            `protobuf:"bytes,11,opt,name=active,proto3" bson:"active" json:"active"`
	Roles      kind.UserRoles  `protobuf:"bytes,12,opt,name=roles,proto3" bson:"roles" json:"roles"`
}
