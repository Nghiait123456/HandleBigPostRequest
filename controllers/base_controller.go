package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"gopkg.in/go-playground/validator.v9"
)

type ResponseError struct {
	ctx          *fasthttp.RequestCtx
	code         int
	jsonResponse JsonResponseError
}

type JsonResponseError struct {
	ErrorCode int
	Message   string
}

func (rs *ResponseError) ResponseError() {
	rs.ctx.SetStatusCode(rs.code)
	bodyJson, err := json.Marshal(rs.jsonResponse)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(rs.ctx, string(bodyJson))
}

func MessageFrErrorValidate(err error) string {
	if _, ok := err.(*validator.InvalidValidationError); ok {
		fmt.Println(err)
		return "Validate error"
	}

	errorDetail := "Validate error detail"
	for _, err := range err.(validator.ValidationErrors) {
		errorDetail += "field error: " + err.Field() + ", type error: " + err.ActualTag() + ", value error: " + fmt.Sprintf("%v", err.Value())
	}

	return errorDetail
}
