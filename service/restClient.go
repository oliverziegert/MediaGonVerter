package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"pc-ziegert.de/media_service/constant"
	"pc-ziegert.de/media_service/log"
	"time"
)

type Client struct {
	BaseURL    string
	apiKey     string
	HTTPClient *http.Client
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type successResponse struct {
	Code int    `json:"code"`
	Data string `json:"data"`
}

func NewClient(baseURL string, apiKey string) *Client {
	return &Client{
		BaseURL: baseURL,
		apiKey:  apiKey,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *Client) sendRequest(ctx *gin.Context, req *http.Request) (successResponse, error) {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "text/plain")
	req.Header.Set("User-Agent", fmt.Sprintf("PC-Ziegert Media Service %s", constant.AppVersion))
	req.Header.Set(constant.CoreServiceForwardHostHeader, ctx.GetString("tenantDomain"))
	//req.Header.Set(constant.CoreServiceForwardPortHeader, ctx.Request.Proto)
	//req.Header.Set(constant.CoreServiceForwardProtoHeader, ctx.Request.Proto)
	req.Header.Set(constant.CoreServiceTokenHeader, c.apiKey)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return successResponse{}, nil
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes errorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return successResponse{}, errors.New(errRes.Message)
		}

		return successResponse{}, fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	bodyString := string(bodyBytes)

	fullResponse := successResponse{
		Data: bodyString,
		Code: res.StatusCode,
	}
	return fullResponse, nil
}
