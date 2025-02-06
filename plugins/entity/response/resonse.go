package response

type HandlerResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func HandlerSuccess(Data interface{}) *HandlerResponse {
	return &HandlerResponse{
		Code:    0,
		Message: "success",
		Data:    Data,
	}
}

func HandlerError(message string) *HandlerResponse {
	return &HandlerResponse{
		Code:    -1,
		Message: message,
		Data:    nil,
	}
}

type LogicResponse struct {
	Data interface{}
}

func LogicResult(data interface{}) *LogicResponse {
	return &LogicResponse{
		Data: data,
	}
}
