package contexts

type APIClient struct {
	Response []any
}

type JSONResponse struct {
	Success bool   `json:"success"`
	Data    Result `json:"data"`
}

type Result struct {
	Items  any     `json:"items,omitempty"`
	Errors *Errors `json:"errors,omitempty"`
}

type Errors struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Session struct {
	UserID    string `json:"user_id"`
	UserLevel string `json:"user_level"`
	Language  string `json:"language"`
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
			Errors: &Errors{
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
			Items:  data,
			Errors: nil,
		},
	}
}
