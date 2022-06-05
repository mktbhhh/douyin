// Code generated by Kitex v0.3.1. DO NOT EDIT.

package userservice

import (
	"context"
	"douyin/kitex_gen/user"
	"fmt"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	"github.com/cloudwego/kitex/pkg/streaming"
	"google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

var userServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "UserService"
	handlerType := (*user.UserService)(nil)
	methods := map[string]kitex.MethodInfo{
		"UserRegister": kitex.NewMethodInfo(userRegisterHandler, newUserRegisterArgs, newUserRegisterResult, false),
		"UserLogin":    kitex.NewMethodInfo(userLoginHandler, newUserLoginArgs, newUserLoginResult, false),
		"UserInfo":     kitex.NewMethodInfo(userInfoHandler, newUserInfoArgs, newUserInfoResult, false),
		"CheckUser":    kitex.NewMethodInfo(checkUserHandler, newCheckUserArgs, newCheckUserResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "user",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.3.1",
		Extra:           extra,
	}
	return svcInfo
}

func userRegisterHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.DouyinUserRegisterRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).UserRegister(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *UserRegisterArgs:
		success, err := handler.(user.UserService).UserRegister(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UserRegisterResult)
		realResult.Success = success
	}
	return nil
}
func newUserRegisterArgs() interface{} {
	return &UserRegisterArgs{}
}

func newUserRegisterResult() interface{} {
	return &UserRegisterResult{}
}

type UserRegisterArgs struct {
	Req *user.DouyinUserRegisterRequest
}

func (p *UserRegisterArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UserRegisterArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *UserRegisterArgs) Unmarshal(in []byte) error {
	msg := new(user.DouyinUserRegisterRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UserRegisterArgs_Req_DEFAULT *user.DouyinUserRegisterRequest

func (p *UserRegisterArgs) GetReq() *user.DouyinUserRegisterRequest {
	if !p.IsSetReq() {
		return UserRegisterArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UserRegisterArgs) IsSetReq() bool {
	return p.Req != nil
}

type UserRegisterResult struct {
	Success *user.DouyinUserRegisterResponse
}

var UserRegisterResult_Success_DEFAULT *user.DouyinUserRegisterResponse

func (p *UserRegisterResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UserRegisterResult")
	}
	return proto.Marshal(p.Success)
}

func (p *UserRegisterResult) Unmarshal(in []byte) error {
	msg := new(user.DouyinUserRegisterResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UserRegisterResult) GetSuccess() *user.DouyinUserRegisterResponse {
	if !p.IsSetSuccess() {
		return UserRegisterResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UserRegisterResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.DouyinUserRegisterResponse)
}

func (p *UserRegisterResult) IsSetSuccess() bool {
	return p.Success != nil
}

func userLoginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.DouyinUserLoginRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).UserLogin(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *UserLoginArgs:
		success, err := handler.(user.UserService).UserLogin(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UserLoginResult)
		realResult.Success = success
	}
	return nil
}
func newUserLoginArgs() interface{} {
	return &UserLoginArgs{}
}

func newUserLoginResult() interface{} {
	return &UserLoginResult{}
}

type UserLoginArgs struct {
	Req *user.DouyinUserLoginRequest
}

func (p *UserLoginArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UserLoginArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *UserLoginArgs) Unmarshal(in []byte) error {
	msg := new(user.DouyinUserLoginRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UserLoginArgs_Req_DEFAULT *user.DouyinUserLoginRequest

func (p *UserLoginArgs) GetReq() *user.DouyinUserLoginRequest {
	if !p.IsSetReq() {
		return UserLoginArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UserLoginArgs) IsSetReq() bool {
	return p.Req != nil
}

type UserLoginResult struct {
	Success *user.DouyinUserLoginResponse
}

var UserLoginResult_Success_DEFAULT *user.DouyinUserLoginResponse

func (p *UserLoginResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UserLoginResult")
	}
	return proto.Marshal(p.Success)
}

func (p *UserLoginResult) Unmarshal(in []byte) error {
	msg := new(user.DouyinUserLoginResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UserLoginResult) GetSuccess() *user.DouyinUserLoginResponse {
	if !p.IsSetSuccess() {
		return UserLoginResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UserLoginResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.DouyinUserLoginResponse)
}

func (p *UserLoginResult) IsSetSuccess() bool {
	return p.Success != nil
}

func userInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.DouyinUserRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).UserInfo(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *UserInfoArgs:
		success, err := handler.(user.UserService).UserInfo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UserInfoResult)
		realResult.Success = success
	}
	return nil
}
func newUserInfoArgs() interface{} {
	return &UserInfoArgs{}
}

func newUserInfoResult() interface{} {
	return &UserInfoResult{}
}

type UserInfoArgs struct {
	Req *user.DouyinUserRequest
}

func (p *UserInfoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UserInfoArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *UserInfoArgs) Unmarshal(in []byte) error {
	msg := new(user.DouyinUserRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UserInfoArgs_Req_DEFAULT *user.DouyinUserRequest

func (p *UserInfoArgs) GetReq() *user.DouyinUserRequest {
	if !p.IsSetReq() {
		return UserInfoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UserInfoArgs) IsSetReq() bool {
	return p.Req != nil
}

type UserInfoResult struct {
	Success *user.DouyinUserResponse
}

var UserInfoResult_Success_DEFAULT *user.DouyinUserResponse

func (p *UserInfoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UserInfoResult")
	}
	return proto.Marshal(p.Success)
}

func (p *UserInfoResult) Unmarshal(in []byte) error {
	msg := new(user.DouyinUserResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UserInfoResult) GetSuccess() *user.DouyinUserResponse {
	if !p.IsSetSuccess() {
		return UserInfoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UserInfoResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.DouyinUserResponse)
}

func (p *UserInfoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func checkUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.CheckUserRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).CheckUser(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CheckUserArgs:
		success, err := handler.(user.UserService).CheckUser(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CheckUserResult)
		realResult.Success = success
	}
	return nil
}
func newCheckUserArgs() interface{} {
	return &CheckUserArgs{}
}

func newCheckUserResult() interface{} {
	return &CheckUserResult{}
}

type CheckUserArgs struct {
	Req *user.CheckUserRequest
}

func (p *CheckUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CheckUserArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *CheckUserArgs) Unmarshal(in []byte) error {
	msg := new(user.CheckUserRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CheckUserArgs_Req_DEFAULT *user.CheckUserRequest

func (p *CheckUserArgs) GetReq() *user.CheckUserRequest {
	if !p.IsSetReq() {
		return CheckUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CheckUserArgs) IsSetReq() bool {
	return p.Req != nil
}

type CheckUserResult struct {
	Success *user.CheckUserResponse
}

var CheckUserResult_Success_DEFAULT *user.CheckUserResponse

func (p *CheckUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CheckUserResult")
	}
	return proto.Marshal(p.Success)
}

func (p *CheckUserResult) Unmarshal(in []byte) error {
	msg := new(user.CheckUserResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CheckUserResult) GetSuccess() *user.CheckUserResponse {
	if !p.IsSetSuccess() {
		return CheckUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CheckUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.CheckUserResponse)
}

func (p *CheckUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) UserRegister(ctx context.Context, Req *user.DouyinUserRegisterRequest) (r *user.DouyinUserRegisterResponse, err error) {
	var _args UserRegisterArgs
	_args.Req = Req
	var _result UserRegisterResult
	if err = p.c.Call(ctx, "UserRegister", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UserLogin(ctx context.Context, Req *user.DouyinUserLoginRequest) (r *user.DouyinUserLoginResponse, err error) {
	var _args UserLoginArgs
	_args.Req = Req
	var _result UserLoginResult
	if err = p.c.Call(ctx, "UserLogin", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UserInfo(ctx context.Context, Req *user.DouyinUserRequest) (r *user.DouyinUserResponse, err error) {
	var _args UserInfoArgs
	_args.Req = Req
	var _result UserInfoResult
	if err = p.c.Call(ctx, "UserInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CheckUser(ctx context.Context, Req *user.CheckUserRequest) (r *user.CheckUserResponse, err error) {
	var _args CheckUserArgs
	_args.Req = Req
	var _result CheckUserResult
	if err = p.c.Call(ctx, "CheckUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
