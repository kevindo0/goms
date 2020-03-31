package slogin

import (
	"fmt"
)

func Login(username, password string) bool {
	fmt.Printf("username: %s, password: %s\n", username, password)
	if (username == "" || password == "") {
		return false
	}
	return true
}