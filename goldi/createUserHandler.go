package main

type Response struct {
	statusCode int
	message    string
}

func CreateUserHandler(u CreateUserUsecase) *Response {
	return createUserHandler(u)
}

func createUserHandler(u CreateUserUsecase) *Response {
	res, err := u.Do()
	if err != nil {
		return &Response{
			statusCode: 400,
			message:    err.Error(),
		}
	}
	return &Response{
		statusCode: 200,
		message:    res,
	}
}
