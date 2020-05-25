package models

type UserLogin struct{
	Email string
	Password []byte
}

func (u UserLogin) IsEmpty() bool {
	return u.Email == "" && len(u.Password) == 0
}