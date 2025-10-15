package responser

type Response struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
	IsSuccess  bool   `json:"isSuccess"`
	Data       any    `json:"data,omitempty"`
}

func NewResponse() Response {
	return Response{}
}

func (r Response) SetMessage(message string) Response {
	r.Message = message
	return r
}

func (r Response) SetIsSuccess(isSuccess bool) Response {
	r.IsSuccess = isSuccess
	return r
}

func (r Response) SetStatusCode(statusCode int) Response {
	r.StatusCode = statusCode
	return r
}

func (r Response) SetData(data any) Response {
	r.Data = data
	return r
}

func (r Response) Build() Response {
	return Response{
		IsSuccess:  r.IsSuccess,
		Message:    r.Message,
		StatusCode: r.StatusCode,
		Data:       r.Data,
	}
}
