package example

type hasPermissionFunc func(user User, data Data) bool

type hasPermissionFuncBuilder []hasPermissionFunc

func newHasPermissionFuncBuilder() hasPermissionFuncBuilder {
	return make(hasPermissionFuncBuilder, 0)
}

func (b hasPermissionFuncBuilder) Add(f hasPermissionFunc) hasPermissionFuncBuilder {
	return append(b, f)
}

func (b hasPermissionFuncBuilder) ConditionalAdd(f hasPermissionFunc, condition bool) hasPermissionFuncBuilder {
	if !condition {
		return b
	}
	return b.Add(f)
}

func hasAnyPermission(user User, data Data, fs ...hasPermissionFunc) bool {
	for _, f := range fs {
		if f(user, data) {
			return true
		}
	}
	return false
}

func isCreator(user User, data Data) bool {
	return user.Name == data.CreatorName
}

func isOwner(user User, data Data) bool {
	for _, id := range data.OwnerIDs {
		if id == user.ID {
			return true
		}
	}
	return false
}

func (h handlerImpl) isSuper(user User, _ Data) bool {
	profile := h.profileService.GetProfileById(user.ID)
	return profile.IsSuper()
}

func oneOfBosses(bosses []string) hasPermissionFunc {
	return func(user User, _ Data) bool {
		for _, boss := range bosses {
			if boss == user.Name {
				return true
			}
		}
		return false
	}
}
