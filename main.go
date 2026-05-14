package main

import (
	"fmt"
	"gameapp/entity"
	"gameapp/repository/mysql"
)

func main() {

}

func testUserMysqlRepo() {
	mysqlRepo := mysql.New()

	createdUser, err := mysqlRepo.Register(entity.User{
		ID:          0,
		PhoneNumber: "2020-12",
		Name:        "hossein",
	})

	if err != nil {
		fmt.Println("register user error:", err)
	} else {
		fmt.Println("registered user:", createdUser)
	}

	isUnique, err := mysqlRepo.IsPhoneNumberUnique(createdUser.PhoneNumber)
	if err != nil {
		fmt.Println("unique error:", err)
	}

	fmt.Println("isUnique:", isUnique)
}
