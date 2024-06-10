package auth

import (
	"github.com/anthdm/gothkit/kit"
	v "github.com/anthdm/gothkit/validate"
)

var authSchema = v.Schema{
	"email":    v.Rules(v.Email),
	"password": v.Rules(v.Required),
}

var signupSchema = v.Schema{
	"email": v.Rules(v.Email),
	"password": v.Rules(
		//v.ContainsSpecial,
		//v.ContainsUpper,
		v.Min(7),
		v.Max(50),
	),
	"firstName": v.Rules(v.Min(2), v.Max(50)),
	"lastName":  v.Rules(v.Min(2), v.Max(50)),
}

func HandleAuthIndex(kit *kit.Kit) error {
	return kit.Render(AuthIndex(AuthIndexPageData{}))
}

func HandleAuthCreate(kit *kit.Kit) error {
	var values LoginFormValues
	errors, ok := v.Request(kit.Request, &values, authSchema)
	if !ok {
		return kit.Render(LoginForm(values, errors))
	}
	return nil
}

func HandleSignupIndex(kit *kit.Kit) error {
	return kit.Render(SignupIndex(SignupIndexPageData{}))
}

func HandleSignupCreate(kit *kit.Kit) error {
	var values SignupFormValues
	errors, ok := v.Request(kit.Request, &values, signupSchema)
	if !ok {
		return kit.Render(SignupForm(values, errors))
	}
	if values.Password != values.PasswordConfirm {
		errors.Add("passwordConfirm", "passwords do not match")
		return kit.Render(SignupForm(values, errors))
	}
	user, err := createUserFromFormValues(values)
	if err != nil {
		return err
	}
	return kit.Render(ConfirmEmail(user.Email))
}
