package example

import (
	"option_and_builder/boss"
	"option_and_builder/profile"
)

type Handler interface {
	HasPermission(req HasPermissionRequest) HasPermissionResponse
}

func NewHandler() Handler {
	return &handlerImpl{
		profileService: profile.NewService(),
	}
}

type handlerImpl struct {
	profileService profile.Service
}

func (h handlerImpl) HasPermission(req HasPermissionRequest) HasPermissionResponse {
	user := req.User
	data := req.Data

	fs := newHasPermissionFuncBuilder().
		Add(isCreator).
		Add(isOwner).
		// 第一次迭代时，添加这一行 👇
		ConditionalAdd(h.isSuper, req.UserIsSuper).
		// 第二次迭代时，添加这一行 👇
		ConditionalAdd(oneOfBosses(boss.GetBosses()), req.UserIsBoss)
	hasPermission := hasAnyPermission(user, data, fs...)

	return HasPermissionResponse{
		HasPermission: hasPermission,
	}
}
