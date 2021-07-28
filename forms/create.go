package forms

type ValidationError map[string]string
type ValidationErrors map[string]ValidationError
type Validator func() ValidationError

type UserCreateForm struct {
	Username        string `form:"username" binding:"required"`
	Password        string `form:"password1" binding:"required"`
	PasswordConfirm string `form:"password2" binding:"required"`
	Email           string `form:"email" binding:"required"`
	InviteCode      string `form:"inviteCode"`
}

func (f *UserCreateForm) GetValidators() map[string]Validator {
	return map[string]Validator{
		"username": f.ValidateUsername,
		"password": f.ValidatePassword,
	}
}

func (f *UserCreateForm) Validate() ValidationErrors {
	errs := make(ValidationErrors)

	for field, validator := range f.GetValidators() {
		if err := validator(); err != nil {
			errs[field] = err
		}
	}

	if len(errs) > 0 {
		return errs
	}
	return nil
}

func (f *UserCreateForm) ValidatePassword() ValidationError {
	errs := make(ValidationError)

	if f.Password != f.PasswordConfirm {
		errs["equal"] = "password1 and password2 must be equals"
	}

	if length := len(f.Password); length < 6 {
		errs["length"] = "password's length must be more than 6 characters"
	}

	if len(errs) > 0 {
		return errs
	}
	return nil
}

func (f *UserCreateForm) ValidateUsername() ValidationError {
	errs := make(ValidationError)
	length := len(f.Username)

	switch {
	case length < 3:
		errs["length"] = "Username length must be more than 2 characters"
	case length > 15:
		errs["length"] = "Username length must be less than 15 characters"
	}

	if len(errs) > 0 {
		return errs
	}
	return nil
}
