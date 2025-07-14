package main

type Response struct {
	StatusCode int
	Message    string
}

func CreateUserHandler(u CreateUserUsecase) *Response {
	return createUserHandler(u)
}

func createUserHandler(u CreateUserUsecase) *Response {
	res, err := u.Do()
	if err != nil {
		return &Response{
			StatusCode: 400,
			Message:    err.Error(),
		}
	}
	return &Response{
		StatusCode: 200,
		Message:    res,
	}
}
