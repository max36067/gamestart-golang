package models

type User struct {
	ID          int    `json:"id"`
	Email       string `json:"email"`
	Password    string `json:"password" gorm:"<-:create"`
	Name        string `json:"name"`
	IsActive    bool   `json:"is_active"`
	IsSuperUser bool   `json:"is_super_user"`
}

type SignUpInput struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
}

type SignInInput struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
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
