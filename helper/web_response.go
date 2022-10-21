package helper

import "efishcommerce/model/web"

func WebResponse(code int, status string, message string, data interface{}) web.WebResponse {
	meta := web.Meta{
		Code:    code,
		Status:  status,
		Message: message,
	}

	webResponse := web.WebResponse{
		Meta: meta,
		Data: data,
	}

	return webResponse
}
