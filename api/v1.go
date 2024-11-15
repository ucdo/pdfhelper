package api

import (
	"Octopus/PdfHelper/response"
	"Octopus/PdfHelper/structures"
	"Octopus/PdfHelper/utils"
	"errors"
	"github.com/gin-gonic/gin"
)

func ScreenShut(c *gin.Context) {
	imgTypeMap := map[string]struct{}{
		"png": {},
		"pdf": {},
	}
	// 1. 验证以及解析参数
	var params structures.ScreenShot
	if err := c.ShouldBind(&params); err != nil {
		response.Error(c, err)
		return
	}

	if _, ok := imgTypeMap[params.ImgType]; !ok {
		response.Error(c, errors.New("unsupported image type"))
		return
	}

	err := utils.Do(params.Uri, params.ImgType, params.ID)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c, nil)
}
