package structures

type ScreenShot struct {
	ID      int    `json:"id" binding:"required"`      // 保存的文件名称
	Uri     string `json:"uri" binding:"required"`     // http://xxx.xx
	ImgType string `json:"imgType" binding:"required"` // 截图的类型： png/pdf
}
