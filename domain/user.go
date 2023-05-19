package domain

type User struct {
	ID          int    `json:"id"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	IsActive    bool   `json:"is_active"`
	IsSuperUser bool   `json:"is_super_user"`
}

type UserRequest struct {
	Email string `json:"email" binding:"required"`
}

type UserResponse struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

type GoogleUser struct {
	ID         string `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	GivenName  string `json:"given_name,omitempty"`
	FamilyName string `json:"family_name,omitempty"`
	Picture    string `json:"picture,omitempty"`
	Locale     string `json:"locale,omitempty"`
}

func (User) TableName() string {
	return "user"
}

func (UserResponse) TableName() string {
	return "user"
}

type UserRepository interface {
	Create(user *User) error
	Fetch() ([]UserResponse, error)
	GetByEmail(email string) (User, error)
	GetByID(id int) (User, error)
}
