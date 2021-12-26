package handler

import (
	"context"
	"user/domain/model"
	"user/domain/service"
	pb "user/proto/user"
)

type User struct {
	UserDataService service.IUserDataService
}

func (u *User) Register(ctx context.Context, request *pb.UserRegisterRequest, response *pb.UserRegisterResponse) error {
	UserRegister := &model.User{
		UserName: request.UserName,
		FirstName: request.FirstName,
		HashPassword: request.Pwd,
	}
	_,err := u.UserDataService.AddUser(UserRegister)
	if err != nil {
		return err
	}
	response.Message =  "添加成功"
	return nil
}

func (u *User) Login(ctx context.Context, request *pb.UserLoginRequest, response *pb.UserLoginResponse) error {
	isOk,err := u.UserDataService.CheckPwd(request.UserName,request.Pwd)
	if err != nil {
		return err
	}
	response.IsSuccess = isOk
	return nil
}

func (u *User) GetUserInfo(ctx context.Context, request *pb.UserInfoRequest, response *pb.UserInfoResponse) error {
	userInfo,err := u.UserDataService.FindUserByName(request.UserName)
	if err != nil {
		return err
	}
	response = UserForResponse(userInfo)
	return nil
}
func UserForResponse(userInfo *model.User) *pb.UserInfoResponse  {
	response := &pb.UserInfoResponse{}
	response.UserName = userInfo.UserName
	response.UserId = userInfo.ID
	response.FirstName= userInfo.FirstName
	return response
}
