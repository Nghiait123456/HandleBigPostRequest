package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"gopkg.in/go-playground/validator.v9"
	"log"
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

	rs.ctx.SetStatusCode(rs.code)
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

func ResponseJson(ctx *fasthttp.RequestCtx, bodyJson interface{}, httpCode int) {
	jsonRs, err := json.Marshal(bodyJson)
	if err != nil {
		log.Fatalln("Data input not json string")
	}

	ctx.SetStatusCode(httpCode)
	fmt.Fprintf(ctx, string(jsonRs))
}
