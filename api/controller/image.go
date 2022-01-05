package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"pc-ziegert.de/media_service/service"
)

// ImageController handles requests for image endpoints.
type ImageController struct {
	iServ *service.ImageService
}

// NewImageController create a new user controller.
func NewImageController(is *service.ImageService) *ImageController {
	return &ImageController{is}
}

func (imageCtrl *ImageController) GetImageHandler() gin.HandlerFunc {
	// swagger:operation GET /mediaserver/image/{token}/{size} images getImage
	//
	// Convert DRACOON rsource.
	//
	// ---
	//
	// security:
	// - basic: []
	//
	// produces:
	// - application/json
	//
	// responses:
	//   '200':
	//     "$ref": "#/responses/UpdateUserResponse"
	//   '400':
	//     description: "__Bad Request__\n\n
	//       ⦁ [-303]: Invalid ID\n
	//       ⦁ [-309]: Negative number\n
	//       ⦁ [-310]: Negative or zero number\n
	//       ⦁ [-311]: Empty string\n
	//       ⦁ [-312]: Too long string\n
	//       ⦁ [-313]: Invalid date\n
	//       ⦁ [-317]: Invalid username\n
	//       ⦁ [-318]: Invalid password"
	//     schema:
	//       "$ref": "#/definitions/Error"
	//   '401':
	//     description: "__Unauthorized__\n\n
	//       ⦁ [-101]: Invalid authentication data\n
	//       ⦁ [-102]: Invalid credentials"
	//     schema:
	//       "$ref": "#/definitions/Error"
	//   '403':
	//     description: "__Forbidden__\n\n
	//       ⦁ [-202]: No right to update user"
	//     schema:
	//       "$ref": "#/definitions/Error"
	//   '404':
	//     description: "__Not Found__\n\n
	//       ⦁ [-409]: User not found"
	//     schema:
	//       "$ref": "#/definitions/Error"
	//   '409':
	//     description: "__Conflict__\n\n
	//       ⦁ [-410]: User already exists"
	//     schema:
	//       "$ref": "#/definitions/Error"
	//   '412':
	//     description: "__Precondition Failed__\n\n
	//       ⦁ [-103]: User not activated"
	//     schema:
	//       "$ref": "#/definitions/Error"
	//   default:
	//     "$ref": "#/responses/ErrorResponse"
	return func(ctx *gin.Context) {
		token := getTokenPathVar(ctx)
		size := getSizePathVar(ctx)
		crop := getCropVar(ctx)

		tenantDomain := imageCtrl.iServ.GetTenantDomain(ctx, token)
		encryptedToken := imageCtrl.iServ.GetEncryptedToken(ctx, token)

		width, height := imageCtrl.iServ.SplitSize(ctx, size)

		mediaToken := imageCtrl.iServ.DecryptToken(ctx, encryptedToken)

		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Writer.WriteHeader(http.StatusOK)
		resBody := make(map[string]interface{})
		resBody["token"] = token
		resBody["tenantDomain"] = tenantDomain
		resBody["encryptedToken"] = encryptedToken
		resBody["height"] = height
		resBody["width"] = width
		resBody["crop"] = crop
		resBody["mediaToken"] = mediaToken
		rBody, _ := json.Marshal(resBody)
		ctx.Writer.Write(rBody)
	}
}

func (imageCtrl *ImageController) HeadImageHandler() gin.HandlerFunc {
	// swagger:operation HEAD /mediaserver/image/{token}/{size} entries listEntries
	//
	// Lists all entries.
	//
	// Only entries a user can see are returned.
	//
	// # Filtering
	//
	// The result can be filtered via following fields:
	// | field name  | operators             | data type / allowed values |
	// | ----------- | --------------------- | -------------------------- |
	// | userId      | eq (equal)            | int                        |
	// | typeId      | eq (equal)            | int                        |
	// | startTime   | bt (between)          | datetime strings           |
	// | activityId  | i (is), eq (equal)    | null, int                  |
	// | description | i (is), cn (contains) | null, string               |
	// &#9432; Filters are connected via logical conjunction (AND).
	//
	// __Filter Syntax:__
	// [field name];[operator];[value-1];...;[value-n]
	//
	// __Example:__
	// Get entries for a specific time interval: startTime;bt;2019-01-01T00:00:00;2019-01-05T00:00:00
	//
	// # Sorting
	//
	// The result can be sorted via following fields:
	// | field name | operators  |
	// | ---------- | ---------- |
	// | startTime  | asc / desc |
	// &#9432; Sorting by multiple fields is not supported.
	//
	// __Sort Syntax:__
	// [field name];[operator]
	//
	// __Example:__
	// Sort entries descending by their date: date;desc
	//
	// ---
	//
	// security:
	// - basic: []
	//
	// produces:
	// - application/json
	//
	// parameters:
	// - name: filter
	//   in: query
	//   description: Filtering applied to the entries result.
	//   required: false
	//   type: string
	// - name: sort
	//   in: query
	//   description: Sorting applied to the entries result.
	//   required: false
	//   type: string
	// - name: offset
	//   in: query
	//   description: Start of the entries result page.
	//   required: false
	//   type: integer
	//   format: int32
	// - name: limit
	//   in: query
	//   description: Size of the entries result page. (default=50)
	//   required: false
	//   type: integer
	//   format: int32
	//
	// responses:
	//   '200':
	//     "$ref": "#/responses/GetEntriesResponse"
	//   '400':
	//     description: "__Bad Request__\n\n
	//       ⦁ [-302]: Invalid page number\n
	//       ⦁ [-304]: Invalid filter\n
	//       ⦁ [-305]: Invalid sort\n
	//       ⦁ [-306]: Invalid offset\n
	//       ⦁ [-307]: Invalid limit"
	//     schema:
	//       "$ref": "#/definitions/Error"
	//   '401':
	//     description: "__Unauthorized__\n\n
	//       ⦁ [-101]: Invalid authentication data\n
	//       ⦁ [-102]: Invalid credentials"
	//     schema:
	//       "$ref": "#/definitions/Error"
	//   '403':
	//     description: "__Forbidden__\n\n
	//       ⦁ [-207]: No right to get entries of other users"
	//     schema:
	//       "$ref": "#/definitions/Error"
	//   '412':
	//     description: "__Precondition Failed__\n\n
	//       ⦁ [-103]: User not activated"
	//     schema:
	//       "$ref": "#/definitions/Error"
	//   default:
	//     "$ref": "#/responses/ErrorResponse"
	return func(context *gin.Context) {
		context.Writer.Header().Set("Content-Type", "application/json")
		context.Writer.WriteHeader(http.StatusOK)
	}
}
