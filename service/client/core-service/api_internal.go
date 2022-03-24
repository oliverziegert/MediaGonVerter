package core_service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/url"
	e "pc-ziegert.de/media_service/common/error"
	"pc-ziegert.de/media_service/common/log"
	m "pc-ziegert.de/media_service/common/model"
	"strings"
)

var (
	_ context.Context
)

type InternalApiService service

/*
ConfigApiService Request algorithms
&lt;h3 style&#x3D;&#x27;padding: 5px; background-color: #F6F7F8; border: 1px solid #AAA; border-radius: 5px; display: table-cell;&#x27;&gt;&amp;#128640; Since v4.24.0&lt;/h3&gt;  ### Description: Retrieve a list of available algorithms used for encryption.  ### Precondition: Authenticated user.  ### Postcondition: List of available algorithms is returned.  ### Further Information: None.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *ConfigApiRequestAlgorithmsOpts - Optional Parameters:
     * @param "XSdsAuthToken" (optional.String) -  Authentication token
@return AlgorithmVersionInfoList
*/

type MediaDecryptTokenOpts struct {
	XSdsServiceToken string
}

func (i *InternalApiService) MediaDecryptToken(ctx *gin.Context, token string, localVarOptionals *MediaDecryptTokenOpts) (*m.MediaDecryptedToken, *e.Error) {
	var (
		localVarHttpMethod  = http.MethodGet
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue m.MediaDecryptedToken
	)
	// create path and map variables
	localVarPath := i.client.cfg.BasePath + "/v4/internal/media/decrypt/{token}"
	localVarPath = strings.Replace(localVarPath, "{token}", fmt.Sprintf("%v", token), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json; charset=utf-8"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"text/plain"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals == nil || localVarOptionals.XSdsServiceToken == "" {
		return &localVarReturnValue, e.NewError(e.AuthCredentialsInvalid, "XSdsServiceToken must be provided")
	}

	localVarHeaderParams["X-Sds-Service-Token"] = parameterToString(localVarOptionals.XSdsServiceToken, "")

	r, err := i.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		err := e.WrapError(e.ValIdInvalid, "Failed to register a consumer.", err)
		log.Debug(err.StackTrace())
		return &localVarReturnValue, err
	}

	localVarHttpResponse, err := i.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		err := e.WrapError(e.ValIdInvalid, "Failed to register a consumer.", err)
		log.Debug(err.StackTrace())
		return &localVarReturnValue, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		err := e.WrapError(e.ValIdInvalid, "Failed to register a consumer.", err)
		log.Debug(err.StackTrace())
		return &localVarReturnValue, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = i.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			err := e.WrapError(e.ValIdInvalid, "Failed to register a consumer.", err)
			log.Debug(err.StackTrace())
			return &localVarReturnValue, err
		}
		return &localVarReturnValue, nil
	}

	if localVarHttpResponse.StatusCode >= 300 {
		err := e.NewError(e.ValIdInvalid, "Failed to register a consumer.")
		log.Debug(err.StackTrace())
		return &localVarReturnValue, err
	}

	return &localVarReturnValue, nil
}

type FilesFileIdDownloadsOpts struct {
	XDcInternalServiceToken string
	UserId                  uint64
}

func (i *InternalApiService) FilesFileIDDownloads(ctx *gin.Context, fileId uint64, localVarOptionals *FilesFileIdDownloadsOpts) (*m.FilesFileIDDownloads, *e.Error) {
	var (
		localVarHttpMethod  = http.MethodPost
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue m.FilesFileIDDownloads
	)
	// create path and map variables
	localVarPath := i.client.cfg.BasePath + "/v4/internal/files/{file_id}/downloads"
	localVarPath = strings.Replace(localVarPath, "{file_id}", fmt.Sprintf("%v", fileId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json; charset=utf-8"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json; charset=utf-8"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals == nil {
		return &localVarReturnValue, e.NewError(e.AuthCredentialsInvalid, "FilesFileIdDownloadsOpts must be provided")
	}
	if localVarOptionals.XDcInternalServiceToken == "" {
		return &localVarReturnValue, e.NewError(e.AuthCredentialsInvalid, "XDcInternalServiceToken must be provided")
	}
	localVarHeaderParams["X-DC-Internal-Service-Token"] = parameterToString(localVarOptionals.XDcInternalServiceToken, "")

	if localVarOptionals.UserId > 0 {
		uis := fmt.Sprintf("%d", localVarOptionals.UserId)
		localVarQueryParams.Add("user_id", uis)
	}

	r, err := i.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		err := e.WrapError(e.ValIdInvalid, fmt.Sprintf("Tenant invalid"), err)
		log.Debug(err.StackTrace())
		return &localVarReturnValue, err
	}

	localVarHttpResponse, err := i.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		err := e.WrapError(e.ValIdInvalid, fmt.Sprintf("Tenant invalid"), err)
		log.Debug(err.StackTrace())
		return &localVarReturnValue, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		err := e.WrapError(e.ValIdInvalid, fmt.Sprintf("Tenant invalid"), err)
		log.Debug(err.StackTrace())
		return &localVarReturnValue, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = i.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			err := e.WrapError(e.ValIdInvalid, fmt.Sprintf("Tenant invalid"), err)
			log.Debug(err.StackTrace())
			return &localVarReturnValue, err
		}
		return &localVarReturnValue, nil
	}

	newErr := e.NewError(e.ValIdInvalid, fmt.Sprintf("Unexpected status code"))
	log.Debug(newErr.StackTrace())
	return &localVarReturnValue, newErr

}
