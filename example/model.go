package example

type User struct {
	ID   int
	Name string
}

type Data struct {
	CreatorName string
	OwnerIDs    []int
}

// HasPermissionRequest 查看 User 是否拥有与 Data 相关的权限
type HasPermissionRequest struct {
	User User
	Data Data

	// 第一轮迭代
	UserIsSuper bool // 如果 user 是超管的话也算有权限

	// 第二轮迭代
	UserIsBoss bool // 如果 user 是 Boss 的话也算有权限
}

type HasPermissionResponse struct {
	HasPermission bool
}
