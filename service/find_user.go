package service

import "errors"

func findUserByUsername(username string) (map[string]interface{}, error) {
	db, err := GetDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT ID,email,passwd FROM users WHERE name=?", username)
	if err != nil {
		return nil, errors.New("db error: " + err.Error())
	}
	defer rows.Close()
	var (
		userID   int
		email    string
		password string
	)
	if rows.Next() {
		err = rows.Scan(&userID, &email, &password)
		if err != nil {
			return nil, errors.New("db rows.Scan error: " + err.Error())
		}
	} else {
		// 查询不到此用户名的记录
		return nil, nil
	}
	return map[string]interface{}{
		"ID":     userID,
		"name":   username,
		"email":  email,
		"passwd": password,
	}, nil
}

func findUserByEmail(email string) (map[string]interface{}, error) {
	db, err := GetDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT ID,name,passwd FROM users WHERE email=?", email)
	if err != nil {
		return nil, errors.New("db error: " + err.Error())
	}
	defer rows.Close()
	var (
		userID   int
		name     string
		password string
	)
	if rows.Next() {
		err = rows.Scan(&userID, &name, &password)
		if err != nil {
			return nil, errors.New("db rows.Scan error: " + err.Error())
		}
	} else {
		// 查询不到此用户名的记录
		return nil, nil
	}
	return map[string]interface{}{
		"ID":     userID,
		"name":   name,
		"email":  email,
		"passwd": password,
	}, nil
}
