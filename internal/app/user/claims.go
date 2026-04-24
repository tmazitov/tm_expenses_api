package user

import (
	"errors"

	"github.com/tmazitov/ayda-order-service.git/internal/domain/user"
)

type UserClaims map[string]any

func newUserClaims(u *user.User) UserClaims {
	return UserClaims{
		"user_id": u.Id(),
	}
}

func claimsToParams(claims map[string]interface{}) (user.UserParams, error) {

	params := user.UserParams{}

	email, ok := claims["email"].(string)
	if !ok {
		return user.UserParams{}, errors.New("invalid payload claim 'email'")
	}
	params.Email = email

	firstName, ok := claims["given_name"].(string)
	if !ok {
		return user.UserParams{}, errors.New("invalid payload claim 'given_name'")
	}
	params.FirstName = firstName

	lastName, _ := claims["family_name"].(string)
	params.LastName = lastName

	sub, ok := claims["sub"].(string)
	if !ok {
		return user.UserParams{}, errors.New("invalid payload claim 'sub'")
	}
	params.Sub = sub

	return params, nil
}
