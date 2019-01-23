package model

import (
	"time"
)

type Post struct {
	ID        int `gorm:"primary_key"`
    UserID    int
    User      User
    Body      string     `gorm:"type:varchar(180)"`
    Timestamp *time.Time `sql:"DEFAULT:current_timestamp"`
}

func GetPostsByUserID(id int) (*[]Post, error) {
    var posts []Post
    err := db.Preload("User").Where("user_id=?", id).Find(&posts).Error;
    if err != nil {
        return nil, err
    }
    return &posts, nil
}