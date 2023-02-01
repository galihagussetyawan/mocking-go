package mockinggo

type UserRepository interface {
	FindAll() *[]User
}
