package services

import (
	dao "rpay/pkg/user/dao"
	userRepo "rpay/pkg/user/repository"
	"strconv"
	s "strings"
)

func GetUserByUserId(user_id string) dao.Login_Out {
	return userRepo.GetUserById(user_id)
}

func GetAllUsersByQuery(query string) dao.UserQuery {
	if s.Contains(query, ".com") && s.Contains(query, "@") {
		return userRepo.GetUserByEmail(query)
	} else if isNumeric(query) {
		return userRepo.GetUserByPhone(query)
	} else {
		return userRepo.GetUserByName(query)
	}
}

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
