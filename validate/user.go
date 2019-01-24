// 用户相关表单验证
package validate

import (
	"fmt"
	"regexp"

	"blog/model"
)

// 验证用户名
func checkUsername(username string) string {
	return checkLen("Username", username, 3, 20)
}

// 验证密码
func checkPassword(password string) string {
	return checkLen("Password", password, 6, 50)
}

// 验证邮箱
func checkEmail(email string) string {
	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, email); !m {
		return fmt.Sprintf("Email field not a valid email")
	}
	return ""
}

// 登录验证
func CheckLogin(username, password string) []string {
	var errs []string
	if errCheck := checkUsername(username); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	if errCheck := checkPassword(password); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}

	if !model.CheckLogin(username, password) {
		errs = append(errs, "username or password not correct,please input again")
	}

	return errs
}

// 注册验证
func CheckRegister(username, email, pwd1, pwd2 string) []string {
	var errs []string
	if pwd1 != pwd2 {
		errs = append(errs, "2 password does not match")
	}
	if errCheck := checkUsername(username); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	if errCheck := checkPassword(pwd1); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	if errCheck := checkEmail(email); len(errCheck) > 0 {
		errs = append(errs, errCheck)
	}
	if !model.CheckUserExist(username) {
		errs = append(errs, "username is already be used")
	}
	return errs
}
