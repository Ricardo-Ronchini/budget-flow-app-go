package contexts

type APIClient struct {
	// Error func(code int, message string) error
	Response []any
}

type JSONResponse struct {
	Success bool
	Data    Result
}

type Result struct {
	Items  any
	Errors Errors
}

type Errors struct {
	Code    int
	Message string
}

type Session struct {
	UserID    string
	UserLevel string
	Language  string
}

func (c *APIClient) Session() *Session {
	return &Session{
		UserID:    "auth.GetUserID()",
		UserLevel: "",
	}
}

func (a *APIClient) Error(code int, msg string) *JSONResponse {
	return &JSONResponse{
		Success: false,
		Data: Result{
			Errors: Errors{
				Code:    code,
				Message: msg,
			},
		},
	}
}

func (a *APIClient) Ok(data any) *JSONResponse {
	return &JSONResponse{
		Success: true,
		Data: Result{
			Items: data,
		},
	}
}
