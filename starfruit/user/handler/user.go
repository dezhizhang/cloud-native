package handler

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"starfruit.top/user/driver"
	"starfruit.top/user/model"
	"starfruit.top/user/proto"
	"starfruit.top/user/utils"
	"time"
)

type UserService struct {
}

// modelToProtoResponse 转model转换成proto
func modelToProtoResponse(user model.User) proto.ResponseUser {
	return proto.ResponseUser{
		Id:        int32(user.Id),
		Username:  user.Username,
		Gender:    user.Gender,
		Mobile:    user.Mobile,
		Role:      uint32(user.Role),
		BirthDay:  uint64(user.BirthDay.Unix()),
		CreatedAt: uint64(user.CreatedAt.Unix()),
		UpdatedAt: uint64(user.UpdatedAt.Unix()),
		DeletedAt: uint64(user.DeletedAt.Time.Unix()),
		IsDeleted: user.IsDeleted,
	}
}

// GetUserById 通过id 查询用户
func (u *UserService) GetUserById(ctx context.Context, req *proto.IdRequest) (*proto.ResponseUser, error) {
	var user model.User
	err := driver.DB.Where("id = ?", req.Id).First(&user).Error
	if err != nil {
		return nil, err
	}
	// 将用户model转换成proto
	rsp := modelToProtoResponse(user)
	return &rsp, nil
}

// GetUserByMobile 通过mobile获取用户
func (u *UserService) GetUserByMobile(ctx context.Context, req *proto.MobileRequest) (*proto.ResponseUser, error) {
	var user model.User
	err := driver.DB.Where("mobile = ?", req.Mobile).Find(&user).Error
	if err != nil {
		return &proto.ResponseUser{}, err
	}
	// 转换成proto
	rsp := modelToProtoResponse(user)
	return &rsp, nil
}

// CreateUser 创建用户
func (u *UserService) CreateUser(ctx context.Context, req *proto.CreateRequest) (*proto.ResponseUser, error) {

	var user model.User
	tx := driver.DB.Where("mobile = ?", req.Mobile).First(&user)
	if tx.Error != nil {
		return &proto.ResponseUser{}, status.Errorf(codes.AlreadyExists, "用户已存在")
	}

	// 将请求参数赋值给user
	user.Mobile = req.Mobile
	user.Username = req.Username
	user.Password = utils.GenMd5(req.Password)

	// 创建用户
	err := driver.DB.Create(&user).Error
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// 创建成功后返回用户信息
	rsp := modelToProtoResponse(user)
	return &rsp, nil
}

// UpdateUser 更新用户
func (u *UserService) UpdateUser(ctx context.Context, req *proto.UpdateRequest) (*proto.UpdateResponse, error) {
	var user model.User
	err := driver.DB.Where("id = ?", req.Id).First(&user).Error
	if err != nil {
		return nil, status.Error(codes.NotFound, "用户不存在")
	}

	// 转换时间类型
	birthDay := time.Unix(int64(req.BirthDay), 0)
	user = model.User{
		Username: req.Username,
		Mobile:   req.Mobile,
		Gender:   req.Gender,
		Role:     int(req.Role),
		BirthDay: &birthDay,
	}
	// 保存数据
	err = driver.DB.Save(&user).Error
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &proto.UpdateResponse{Updated: true}, nil
}

// GetUserList 获取用户列表
func (u *UserService) GetUserList(ctx context.Context, req *proto.RequestUser) (*proto.ResponseUserList, error) {
	var users []model.User

	rsp := driver.DB.Find(&users)
	if rsp.Error != nil {
		return nil, rsp.Error
	}

	total := rsp.RowsAffected
	result := &proto.ResponseUserList{}
	result.Total = int32(total)

	// 查询分页结果
	driver.DB.Scopes(utils.Pagination(int(req.Page), int(req.Size))).Find(&users)

	// 添加用户方法
	for _, user := range users {
		user := modelToProtoResponse(user)
		result.User = append(result.User, &user)
	}

	return result, nil
}
