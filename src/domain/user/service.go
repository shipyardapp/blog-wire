package user

type Service interface {
	GetRepo

	Add(email string) (*User, error)
}
