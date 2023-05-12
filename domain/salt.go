package domain

type Salt struct {
	ID    int
	Email string
	Salt  string
}

func (Salt) TableName() string {
	return "salt"
}

type SaltRepository interface {
	GetSaltByEmail(email string) (salt string, err error)
}
