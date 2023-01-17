package profile

type Profile struct{}

func (p Profile) IsSuper() bool {
	return true
}

type Service interface {
	GetProfileById(id int) Profile
}

func NewService() Service {
	return &serviceImpl{}
}

type serviceImpl struct{}

func (s serviceImpl) GetProfileById(_ int) Profile {
	return Profile{}
}
