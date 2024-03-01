package app

type User struct {
	Username string `valid:"username"`
	Email    string `valid:"email,required"`
	Password string `valid:"required-6"`
}
