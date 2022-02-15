package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	error2 "pc-ziegert.de/media_service/common/error"
	"pc-ziegert.de/media_service/common/log"
	model2 "pc-ziegert.de/media_service/common/model"
	"pc-ziegert.de/media_service/service/service"
)

// ImageController handles requests for image endpoints.
type ImageController struct {
	iServ *service.ImageService
	tServ *service.TenantService
}

// NewImageController create a new user controller.
func NewImageController(iServ *service.ImageService, tServ *service.TenantService) *ImageController {
	return &ImageController{
		iServ: iServ,
		tServ: tServ,
	}
}

func (imageCtrl *ImageController) GetImageHandler() gin.HandlerFunc {
	// swagger:operation GET /image/{token}/{size} "Get converted resource" getImage
	//
	// Convert DRACOON rsource.
	//
	// For valid response try as size integer IDs with value range(1 , Integer.MAX_VALUE). Other values will generated exceptions
	//
	// ---
	// consumes:
	//   - "*/*"
	// produces:
	//   - application/octet-stream
	// parameters:
	// - name: crop
	//   in: query
	//   description: Allow cropping to provide the possibility of e.g. squared images
	//   required: false,
	//   type: boolean
	//   default: false,
	//   allowEmptyValue: false,
	//   x-example: false
	// - name: size
	//   in: path
	//   description: Size in format width x height (32x32)
	//   required: true
	//   type: string
	// - name: token
	//   in: path
	//   description: Media token from DRACOON Core Service
	//   required: true
	//   type: string
	// responses:
	//   '200':
	//     description: "**Success**"
	//     schema:
	//       type: string
	//       format: byte
	//   '202':
	//     description: "**Converting in process**"
	//     schema:
	//       type: string
	//       format: byte
	//   '401':
	//     description: "**Unauthorized**"
	//     schema:
	//       "$ref": "#/definitions/ErrorResponse"
	//   '403':
	//     description: "**Forbidden**"
	//     schema:
	//       "$ref": "#/definitions/ErrorResponse"
	//   '404':
	//     description: "**Not found**  \n[-1100] Tenant is not whitelisted \n"
	//     schema:
	//       "$ref": "#/definitions/ErrorResponse"
	//   '500':
	//     description: "**Server error**\n
	//       ⦁ [-1500] File is corrupted.\n
	//       ⦁ [-1501] File is encrypted and can not be converted\n
	//       ⦁ [-1501] Unknown error when converting the file\n
	//       ⦁ [-2000] Max file size exceeded to convert"
	//     schema:
	//       "$ref": "#/definitions/ErrorResponse"
	// schemes:
	//   - https
	// deprecated: false
	return func(ctx *gin.Context) {
		ts, err := imageCtrl.tServ.TenantState(ctx)
		if err != nil {
			err := error2.NewError(error2.ValIdInvalid, fmt.Sprintf("Cant get tenant state"))
			log.Debug(err.StackTrace())
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}
		if *ts == model2.TenantStateInvalid {
			err := error2.NewError(error2.ValIdInvalid, fmt.Sprintf("Tenant invalid"))
			log.Debug(err.StackTrace())
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}

		token := getTokenPathVar(ctx)
		size := getSizePathVar(ctx)
		crop := getCropVar(ctx)

		tenantDomain := imageCtrl.iServ.GetTenantDomain(ctx, token)
		encryptedToken := imageCtrl.iServ.GetEncryptedToken(ctx, token)

		width, height := imageCtrl.iServ.SplitSize(ctx, size)

		mediaToken := imageCtrl.iServ.DecryptToken(ctx, encryptedToken)

		image, _ := imageCtrl.iServ.GetImage(ctx, mediaToken, width, height, crop)

		resBody := make(map[string]interface{})
		resBody["token"] = token
		resBody["tenantDomain"] = tenantDomain
		resBody["encryptedToken"] = encryptedToken
		resBody["height"] = height
		resBody["width"] = width
		resBody["crop"] = crop
		resBody["mediaToken"] = mediaToken
		resBody["image"] = image

		if image.State != model2.ImageStateCached {
			ctx.JSON(http.StatusOK, resBody)
			return
		}

		ctx.JSON(http.StatusAccepted, resBody)
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
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "")
	}
}
