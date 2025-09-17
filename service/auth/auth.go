package auth

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/devdiagon/gomerce/config"
	"github.com/devdiagon/gomerce/types"
	"github.com/devdiagon/gomerce/utils"
	"github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// Define a unique type for the key
type contexKey string

const UserKey contexKey = "userId"

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ComparePasswords(hashed string, plain []byte) bool {
	// hashed: comes from the database
	// plain: comes from the http payload
	err := bcrypt.CompareHashAndPassword([]byte(hashed), plain)
	return err == nil
}

func WithJWTAuth(handlerFunc http.HandlerFunc, store types.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Get the tocken from the user request
		tokenStr := getTokenFromRequest(r)

		//Validate JWT
		token, err := validateToken(tokenStr)
		if err != nil {
			log.Printf("failed to validate token: %v", err)
			permissionDenied(w)
			return
		}

		if !token.Valid {
			log.Printf("failed to validate token: %v", err)
			permissionDenied(w)
			return
		}

		//Fetch the user ID
		claims := token.Claims.(jwt.MapClaims)
		str := claims["userId"].(string)

		userId, _ := strconv.Atoi(str)
		user, err := store.GetUserById(userId)

		if err != nil {
			log.Printf("failed to get user by id: %v", err)
			permissionDenied(w)
			return
		}

		//Set context "userId" to the user ID
		ctx := r.Context()
		ctx = context.WithValue(ctx, UserKey, user.ID)
		r = r.WithContext(ctx)

		handlerFunc(w, r)
	}
}

func getTokenFromRequest(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")

	if tokenAuth != "" {
		return tokenAuth
	}

	return ""
}

func validateToken(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected string method: %v", t.Header["alg"])
		}

		return []byte(config.Envs.JWTSecret), nil
	})
}

func permissionDenied(w http.ResponseWriter) {
	utils.WriteError(w, http.StatusForbidden, fmt.Errorf("permission denied"))
}

func GetUserIdFromContext(ctx context.Context) int {
	userId, ok := ctx.Value(UserKey).(int)
	if !ok {
		return -1
	}

	return userId
}
