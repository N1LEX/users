package models

import (
	"butaforia.io/forms"
	"butaforia.io/token"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"github.com/cristalhq/jwt/v3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

type User struct {
	BaseModel
	Username       string `json:"username",gorm:"uniqueIndex"`
	Password       string `gorm:"size:255,not null"`
	Email          string `json:"email",gorm:"size:255,unique,not null"`
	ActivationCode string `gorm:"size:20,unique"`
	InviteCode     string `gorm:"size:20,unique"`
	Warnings       uint8  `json:"warnings"`
	IsBanned       bool   `json:"isBanned"`
	IsAdmin        bool   `json:"isAdmin"`
}

type UserShortData struct {
	BaseModel
	Username string `json:"username"`
	Email    string `json:"email"`
	Warnings uint8  `json:"warnings"`
	IsBanned bool   `json:"isBanned"`
	IsAdmin  bool   `json:"isAdmin"`
}

func (u *User) GetShortData() *UserShortData {
	return &UserShortData{
		BaseModel: u.BaseModel,
		Username:  u.Username,
		Email:     u.Email,
		Warnings:  u.Warnings,
		IsBanned:  u.IsBanned,
		IsAdmin:   u.IsAdmin,
	}
}

func GetByUsername(username string) (*User, error) {
	u := &User{}
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	tx := db.Where("username = ?", username).First(u)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return u, nil
}

func CreateUser(postForm *forms.UserCreateForm) (*User, error) {
	u := ParseCreateForm(postForm)
	u.SetPassword()
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	tx := db.Table("users").Create(u)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return u, nil
}

func ParseCreateForm(f *forms.UserCreateForm) *User {
	var u User
	data, _ := json.Marshal(f)
	if err := json.Unmarshal(data, &u); err != nil {
		panic(err)
	}
	return &u
}

func (u *User) GetStandardClaims() *jwt.StandardClaims {
	return &jwt.StandardClaims{
		ID:        token.GenerateTokenID(),
		Audience:  []string{u.Username},
		ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Hour * 24)},
	}
}

func (u *User) NewAccessToken() string {
	return token.GenerateToken(u.GetStandardClaims())
}

func (u *User) NewRefreshToken() string {
	c := u.GetStandardClaims()
	c.ExpiresAt = &jwt.NumericDate{Time: time.Now().Add(time.Hour * 24 * 30)}
	return token.GenerateToken(c)
}

func (u *User) SetPassword() {
	p := []byte(u.Password)
	h := sha256.New()
	h.Write(p)
	u.Password = base64.URLEncoding.EncodeToString(h.Sum(nil))
}

func (u *User) IsValidCredentials(username, password string) bool {
	return u.Username == username && u.Password == password
}
