package gapi

import (
	db "github.com/fsobh/auth/db/sqlc"
	"github.com/fsobh/auth/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.User) *pb.User {
	return &pb.User{

		Username:         user.Username,
		FullName:         user.FullName,
		Email:            user.Email,
		PasswordChangeAt: timestamppb.New(user.PasswordChangeAt), // need to convert to pb timestamp cuz it's not the same as go timestamp
		CreatedAt:        timestamppb.New(user.CreatedAt),
	}
}
