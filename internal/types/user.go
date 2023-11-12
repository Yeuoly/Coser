package types

import "gorm.io/gorm"

type User struct {
	Model
	Username string      `gorm:"unique;not null;size:64" json:"username"`
	Email    string      `gorm:"unique;not null;size:128;uniqueIndex" json:"email"`
	Role     int         `gorm:"not null" json:"role"`
	Profile  UserProfile `gorm:"foreignKey:UserID" json:"profile"`
}

const (
	USER_ROLE_DEFAULT = iota
	USER_ROLE_USER
	USER_ROLE_ADMIN
)

func (u *User) IsAdmin() bool {
	return u.Role == USER_ROLE_ADMIN
}

func (u *User) ToResponse() struct {
	Username string      `json:"username"`
	Role     string      `json:"role"`
	Profile  UserProfile `json:"profile"`
} {
	return struct {
		Username string      `json:"username"`
		Role     string      `json:"role"`
		Profile  UserProfile `json:"profile"`
	}{
		Username: u.Username,
		Role:     map[int]string{USER_ROLE_ADMIN: "admin", USER_ROLE_USER: "user"}[u.Role],
		Profile:  u.Profile,
	}
}

type UserProfile struct {
	gorm.Model
	UserID    uint   `gorm:"unique;not null;uniqueIndex" json:"user_id"`
	Avatar    string `gorm:"not null;size:128" json:"avatar"`
	Sign      string `gorm:"not null;size:128" json:"sign"`
	Memory    int64  `gorm:"not null;default:0" json:"memory"`
	MaxMemory int64  `gorm:"not null;default:20971520" json:"max_memory"` // max memory is 20MB
}

type Password struct {
	gorm.Model
	UserID   uint   `gorm:"unique;not null;uniqueIndex"`
	Password string `gorm:"not null;size:128"`
}

func (p *Password) Check(password string) bool {
	return p.Password == password
}

type UserGithubBind struct {
	gorm.Model
	UserID     uint   `gorm:"unique;not null" json:"user_id"`
	GithubId   int    `gorm:"not null;uniqueIndex" json:"github_id"`
	GithubName string `gorm:"not null;size:128" json:"github_name"`
}
