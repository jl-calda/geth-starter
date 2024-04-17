package services

type User struct {
	Id            string
	Name          string
	Email         string
	EmailVerified bool
	Phone         string
	PhoneVerified bool
}

func (u *User) NewUser(id, name, email, phone string) *User {
	return &User{
		Id:    id,
		Name:  name,
		Email: email,
		Phone: phone,
	}
}
