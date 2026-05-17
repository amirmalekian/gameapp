package main

import (
	"encoding/json"
	"fmt"
	"gameapp/entity"
	"gameapp/repository/mysql"
	"gameapp/service/userservice"
	"io"
	"log"
	"net/http"
)

func main() {

	// first method
	http.HandleFunc("/health-check", healthCheckHandler)
	http.HandleFunc("/users/register", userRegisterHandler)
	http.HandleFunc("/users/login", userLoginHandler)

	log.Println("server is listening on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

	// second method => multiplexer
	//mux := http.NewServeMux()
	//mux.HandleFunc("/health-check", healthCheckHandler)
	//mux.HandleFunc("/users/register", userRegisterHandler)
	//mux.HandleFunc("/users/login", userLoginHandler)
	//
	//http.ListenAndServe(":8080", mux)
}

func userRegisterHandler(rw http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		fmt.Fprintf(rw, `{"error": "invalid method"}`)

		return
	}

	data, err := io.ReadAll(req.Body)
	if err != nil {
		rw.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		//	 or use => fmt.Fprintf(rw, `{"error": "invalid method"}`)

		return
	}

	var uReq userservice.RegisterRequest
	err = json.Unmarshal(data, &uReq)
	if err != nil {
		rw.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		//	 or use => fmt.Fprintf(rw, `{"error": "invalid method"}`)

		return
	}

	mysqlRepo := mysql.New()
	userSvc := userservice.New(mysqlRepo)

	resp, err := userSvc.Register(uReq)
	if err != nil {
		rw.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))

		return
	}

	rw.Write([]byte(fmt.Sprintf(`{"message": "user registered, %v"}`, resp)))

}

func healthCheckHandler(rw http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(rw, `{"message": "everything is ok!"}`)
}

func userLoginHandler(rw http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		fmt.Fprintf(rw, `{"error": "invalid method"}`)
	}

	data, err := io.ReadAll(req.Body)
	if err != nil {
		rw.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
	}

	var lReq userservice.LoginRequest
	err = json.Unmarshal(data, &lReq)
	if err != nil {
		rw.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))

		return
	}

	mysqlRepo := mysql.New()
	userSvc := userservice.New(mysqlRepo)

	resp, err := userSvc.Login(lReq)
	if err != nil {
		rw.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))

		return
	}

	rw.Write([]byte(fmt.Sprintf(`{"message": "user registered, %v"}`, resp)))

	rw.Write([]byte(fmt.Sprintf(`{"message": "user credentials is OK, %v"}`, lReq)))
}

func testUserMysqlRepo() {
	mysqlRepo := mysql.New()

	createdUser, err := mysqlRepo.Register(entity.User{
		ID:          0,
		PhoneNumber: "0912345678",
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
