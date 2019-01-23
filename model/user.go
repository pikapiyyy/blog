package model

import (
	"crypto/md5"
	"encoding/hex"
)

type User struct {
	ID           int    `gorm:"primary_key"`
	Username     string `gorm:"type:varchar(64)"`
	Email        string `gorm:"type:varchar(120)"`
	PasswordHash string `gorm:"type:varchar(128)"`
	Posts        []Post
	Followers    []*User `gorm:"many2many:follower;association_jointable_foreignkey:follower_id"`
}

func GeneratePasswordHash(pwd string) string {
	hasher := md5.New()
	hasher.Write([]byte(pwd))
	pwdHash := hex.EncodeToString(hasher.Sum(nil))

	return pwdHash
}

func (this *User) SetPassword(password string) {
	this.PasswordHash = GeneratePasswordHash(password)
}

// CheckPassword func
func (u *User) CheckPassword(password string) bool {
	return GeneratePasswordHash(password) == u.PasswordHash
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	err := db.Where("username=?", username).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func AddUser(username, password, email string) error {
	user := User{Username: username, Email: email}
	user.SetPassword(password)

	return db.Create(&user).Error
}

func CheckLogin(username, password string) bool {
	user, err := GetUserByUsername(username)
	if err != nil {
		return false
	}

	return user.CheckPassword(password)
}

func CheckUserExist(username string) bool {
	_, err := GetUserByUsername(username)
	if err != nil {
		return true
	}
	return false
}
