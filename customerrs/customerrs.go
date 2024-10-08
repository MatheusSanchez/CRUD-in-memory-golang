package customerrs

import "errors"

var ErrUserNotFoundById = errors.New("the user with the specified ID does not exist")
var ErrSomethingWentWrong = errors.New("something went wrong")
var ErrInvalidUUID = errors.New("invalid id")
