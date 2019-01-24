// session操作相关
package util

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	sessionName string
	store       *sessions.CookieStore
)

// 设置session name和id
func init() {
	store = sessions.NewCookieStore([]byte("x5YJc5BXoBoOOXAB"))
	sessionName = "go-blog"
}

// 获取session用户名
func GetSessionUser(r *http.Request) (string, error) {
	var username string
	session, err := store.Get(r, sessionName)
	if err != nil {
		return "", err
	}

	val := session.Values["user"]
	fmt.Println("val:", val)
	username, ok := val.(string)
	if !ok {
		return "", errors.New("can not get session user")
	}
	fmt.Println("username:", username)
	return username, nil
}

// 设置session用户名
func SetSessionUser(w http.ResponseWriter, r *http.Request, username string) error {
	session, err := store.Get(r, sessionName)
	if err != nil {
		return err
	}
	session.Values["user"] = username
	err = session.Save(r, w)
	if err != nil {
		return err
	}
	return nil
}

// 删除某个会话session
func ClearSession(w http.ResponseWriter, r *http.Request) error {
	session, err := store.Get(r, sessionName)
	if err != nil {
		return err
	}

	session.Options.MaxAge = -1

	err = session.Save(r, w)
	if err != nil {
		return err
	}

	return nil
}
