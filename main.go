package main

import (
	"fmt"
	"gameapp/config"
	"gameapp/delivery/httpserver"
	"gameapp/repository/mysql"
	"gameapp/service/authservice"
	"gameapp/service/userservice"
	"time"
)

const (
	JwtSignKey                 = "jwt_secret"
	AccessTokenSubject         = "at"
	RefreshTokenSubject        = "rt"
	AccessTokenExpireDuration  = time.Hour * 24
	RefreshTokenExpireDuration = time.Hour * 24 * 7
)

func main() {
	fmt.Println("start echo server")

	cfg := config.Config{
		HttpServer: config.HTTPServer{Port: 8088},
		Auth: authservice.Config{
			SignKey:               JwtSignKey,
			AccessExpirationTime:  AccessTokenExpireDuration,
			RefreshExpirationTime: RefreshTokenExpireDuration,
			AccessSubject:         AccessTokenSubject,
			RefreshSubject:        RefreshTokenSubject,
		},
		MySQL: mysql.Config{
			Username: "gameapp",
			Password: "gameappt0lk2o20",
			Port:     3308,
			Host:     "localhost",
			DBName:   "gameapp_db",
		},
	}

	authSvc, userSvc := setupServices(cfg)

	server := httpserver.New(cfg, authSvc, userSvc)

	server.Serve()
}

//func userLoginHandler(rw http.ResponseWriter, req *http.Request) {
//	if req.Method != http.MethodPost {
//		fmt.Fprintf(rw, `{"error": "invalid method"}`)
//
//		return
//	}
//
//	data, err := io.ReadAll(req.Body)
//	if err != nil {
//		rw.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
//
//		return
//	}
//
//	var lReq userservice.LoginRequest
//	err = json.Unmarshal(data, &lReq)
//	if err != nil {
//		rw.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
//
//		return
//	}
//
//	authSvc := authservice.New(JwtSignKey, AccessTokenSubject, RefreshTokenSubject,
//		AccessTokenExpireDuration, RefreshTokenExpireDuration)
//
//	mysqlRepo := mysql.New()
//	userSvc := userservice.New(authSvc, mysqlRepo)
//
//	resp, err := userSvc.Login(lReq)
//	if err != nil {
//		rw.Header().Add("Content-Type", "application/json")
//		rw.WriteHeader(http.StatusBadRequest)
//		rw.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
//
//		return
//	}
//
//	data, err = json.Marshal(resp)
//	if err != nil {
//		rw.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
//
//		return
//	}
//
//	rw.Header().Add("Content-Type", "application/json")
//	rw.Write(data)
//}

//func userProfileHandler(rw http.ResponseWriter, req *http.Request) {
//	if req.Method != http.MethodGet {
//		fmt.Fprintf(rw, `{"error": "invalid method"}`)
//	}
//
//	authSvc := authservice.New(JwtSignKey, AccessTokenSubject, RefreshTokenSubject,
//		AccessTokenExpireDuration, RefreshTokenExpireDuration)
//
//	authToken := req.Header.Get("Authorization")
//	claims, err := authSvc.ParseToken(authToken)
//	if err != nil {
//		fmt.Fprintf(rw, `{"error": "token is not valid"}`)
//	}
//
//	mysqlRepo := mysql.New()
//	userSvc := userservice.New(authSvc, mysqlRepo)
//
//	resp, err := userSvc.Profile(userservice.ProfileRequest{UserID: claims.UserID})
//	if err != nil {
//		rw.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
//
//		return
//	}
//
//	data, err := json.Marshal(resp)
//	if err != nil {
//		rw.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
//
//		return
//	}
//
//	rw.Write(data)
//
//}

func setupServices(cfg config.Config) (authservice.Service, userservice.Service) {
	authSvc := authservice.New(cfg.Auth)

	mysqlRepo := mysql.New(cfg.MySQL)
	userSvc := userservice.New(authSvc, mysqlRepo)

	return authSvc, userSvc
}
