package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	m "pc-ziegert.de/media_service/common/model"
	"pc-ziegert.de/media_service/service/service"
	u "pc-ziegert.de/media_service/service/utils"
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

func (iCtrl *ImageController) HeadImageHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ts, err := iCtrl.tServ.TenantState(ctx)
		if err != nil {
			panic(err)
		}
		if *ts == m.TenantStateInvalid {
			err = iCtrl.tServ.BackoffTenant(ctx)
			if err != nil {
				panic(err)
			}
			panic(err)
		}

		t, err := u.GetTokenPathVar(ctx)
		if err != nil {
			panic(err)
		}

		s, err := u.GetSizePathVar(ctx)
		if err != nil {
			panic(err)
		}

		w, h, err := iCtrl.iServ.SplitSize(ctx, s)
		if err != nil {
			panic(err)
		}

		cr := u.GetCropVar(ctx)

		et, err := iCtrl.iServ.GetEncryptedToken(ctx, t)
		if err != nil {
			panic(err)
		}

		mt, err := iCtrl.iServ.DecryptToken(ctx, et)
		if err != nil {
			newErr := iCtrl.tServ.BackoffTenant(ctx)
			if newErr != nil {
				panic(newErr)
			}
			panic(err)
		}

		i, err := iCtrl.iServ.GetImageByToken(ctx, mt, w, h, cr)
		if err != nil {
			newErr := iCtrl.tServ.BackoffTenant(ctx)
			if newErr != nil {
				panic(newErr)
			}
			panic(err)
		}

		co, err := i.GetConversionBySize(w, h, cr)
		if err != nil {
			newErr := iCtrl.tServ.BackoffTenant(ctx)
			if newErr != nil {
				panic(newErr)
			}
			panic(err)
		}

		if co.State == m.ConversionStateCached {
			ctx.Status(http.StatusOK)
			return
		}

		ctx.Status(http.StatusAccepted)
		return
	}
}

func (iCtrl *ImageController) GetImageHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ts, err := iCtrl.tServ.TenantState(ctx)
		if err != nil {
			panic(err)
		}
		if *ts == m.TenantStateInvalid {
			err = iCtrl.tServ.BackoffTenant(ctx)
			if err != nil {
				panic(err)
			}
			panic(err)
		}

		t, err := u.GetTokenPathVar(ctx)
		if err != nil {
			panic(err)
		}

		s, err := u.GetSizePathVar(ctx)
		if err != nil {
			panic(err)
		}

		w, h, err := iCtrl.iServ.SplitSize(ctx, s)
		if err != nil {
			panic(err)
		}

		cr := u.GetCropVar(ctx)

		et, err := iCtrl.iServ.GetEncryptedToken(ctx, t)
		if err != nil {
			panic(err)
		}

		mt, err := iCtrl.iServ.DecryptToken(ctx, et)
		if err != nil {
			newErr := iCtrl.tServ.BackoffTenant(ctx)
			if newErr != nil {
				panic(newErr)
			}
			panic(err)
		}

		i, err := iCtrl.iServ.GetImageByToken(ctx, mt, w, h, cr)
		if err != nil {
			newErr := iCtrl.tServ.BackoffTenant(ctx)
			if newErr != nil {
				panic(newErr)
			}
			panic(err)
		}

		co, err := i.GetConversionBySize(w, h, cr)
		if err != nil {
			newErr := iCtrl.tServ.BackoffTenant(ctx)
			if newErr != nil {
				panic(newErr)
			}
			panic(err)
		}

		if co.State == m.ConversionStateCached {
			d, err := iCtrl.iServ.GenerateS3DownloadUrl(ctx, i, co)
			if err != nil {
				newErr := iCtrl.tServ.BackoffTenant(ctx)
				if newErr != nil {
					panic(newErr)
				}
				panic(err)
			}
			ctx.Redirect(http.StatusTemporaryRedirect, d)
			return
		}

		ctx.Status(http.StatusAccepted)
		return
	}
}

func (iCtrl *ImageController) HeadNodePreviewHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.Status(http.StatusNotImplemented)
		return
	}
}

func (iCtrl *ImageController) GetNodePreviewHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.Status(http.StatusNotImplemented)
		return
	}
}
