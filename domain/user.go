package domain

type User struct {
	ID          int    `json:"id"`
	Email       string `json:"email"`
	Password    string `json:"password" gorm:"<-:create"`
	Name        string `json:"name"`
	IsActive    bool   `json:"is_active"`
	IsSuperUser bool   `json:"is_super_user"`
}

type Salt struct {
	ID    int
	Email string
	Salt  string
}

type UserRequest struct {
	Email string `json:"email" binding:"required"`
}

type UserResponse struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

func (User) TableName() string {
	return "user"
}

func (UserResponse) TableName() string {
	return "user"
}

func (Salt) TableName() string {
	return "salt"
}

type UserRepository interface {
	Create(user *User) error
	Fetch() ([]UserResponse, error)
	GetByEmail(email string) (UserResponse, error)
	GetByID(id string) (UserResponse, error)
}
