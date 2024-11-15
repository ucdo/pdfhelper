package route

import (
	"Octopus/PdfHelper/api"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	r.POST("/shot_img", api.ScreenShut)
}
