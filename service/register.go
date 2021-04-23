package service

import (
	"errors"
	"fmt"
	"regexp"
)

//RegisterAccount 注册账号
func RegisterAccount(username, password, email string) error {
	if username == "" {
		return errors.New("用户名不能为空")
	}
	if password == "" {
		return errors.New("密码不能为空")
	}
	if email == "" {
		return errors.New("邮箱不能为空")
	}
	matched, err := regexp.MatchString("^[a-zA-Z0-9_]+$", username)
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("用户名只能包含字母、数字、下划线")
	}
	usernameLength := len(username)
	if usernameLength < 4 || usernameLength > 10 {
		return errors.New("登录用户名必须大于4位小于10位")
	}
	//判断用户名是否已存在
	tmpUserInfo, err := findUserByUsername(username)
	if err != nil {
		return err
	}
	if tmpUserInfo != nil {
		return errors.New("此用户名已存在")
	}
	passwordLength := len(password)
	if passwordLength < 6 || passwordLength > 30 {
		return errors.New("登录密码必须大于6位小于30位")
	}
	//判断邮箱格式
	matched, err = regexp.MatchString("^[A-Za-z0-9]+([_\\.][A-Za-z0-9]+)*@([A-Za-z0-9\\-]+\\.)+[A-Za-z]{2,6}$", email)
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("邮箱格式错误")
	}
	emailLength := len(email)
	if emailLength < 4 || emailLength > 30 {
		return errors.New("邮箱地址必须大于4位小于30位")
	}
	//判断邮箱是否已存在
	tmpUserInfo, err = findUserByEmail(email)
	if err != nil {
		return err
	}
	if tmpUserInfo != nil {
		return errors.New("此邮箱已存在")
	}
	//add user
	salt := encodePassword(username, password)
	sqlText := fmt.Sprintf("call adduser('%s', %s, '0', '0', '0', '0', '%s', '0', '0', '0', '0', '0', '0', '0', '', '', %s)",
		username, salt, email, salt)
	db, err := GetDB()
	if err != nil {
		return err
	}
	if _, err := db.Exec(sqlText); err != nil {
		return err
	}
	return nil
}
