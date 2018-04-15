package apis

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"path/filepath"
// 	"strings"

// 	"github.com/dgrijalva/jwt-go"
// 	"github.com/gorilla/mux"
// 	//	"encoding/json"
// )

// // User model
// type UserRequest struct {
// 	UserID   string `form:"userid" json:"UserID" binding:"required"`
// 	Password string `form:"password" json:"password" binding:"required"`
// 	// Roles    []string `json:"password" `
// }

// // AuthCustomClaims claims for auth including roles
// type AuthCustomClaims struct {
// 	User UserRequest `json:"user"`
// 	jwt.StandardClaims
// }

// func RegisterAuthAPI(router *mux.Router) {
// 	router.HandleFunc("/api/auths", authapi.ServeAPIAuths)
// 	router.HandleFunc("/api/", serveAPI)
// }

// // ServeAPIAuths serve rest auth requests
// func ServeAPIAuths(w http.ResponseWriter, r *http.Request) {
// 	path := filepath.Clean(r.URL.Path)
// 	path = strings.TrimPrefix(path, "/api/")
// 	fmt.Fprintf(w, "Auth request:"+path)
// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&u); err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid request payload")

// 	}
// }
