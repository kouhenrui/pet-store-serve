package common

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"pet-store-serve/src/dto/reqDto"
	"pet-store-serve/src/global"
	"pet-store-serve/src/msg"
	"pet-store-serve/src/service/common"
	util "pet-store-serve/src/utils"
	"strings"
)

var (
	commonService common.CommonServiceInter = &common.CommonServiceImpl{}
)

// @Summary	获取验证码
// @Produce	json
// @Tags		public
// @Success	200 {object} comDto.ResponseData
// @Router		/captcha [get]
func GetCaptcha(c *gin.Context) {
	err, captera := commonService.GetCaptcha()
	//fmt.Println("接获获取", captera)
	if err != nil {
		c.Error(err)
		return
	}

	c.Set("res", captera)

}

/**
 * @Author Khr
 * @Description //TODO 校验验证码
 * @Date 9:21 2023/9/27
 * @Param
 * @return
 **/
func VfCaptcha(c *gin.Context) {
	var vf = reqDto.Captcha{}
	if err := c.ShouldBindJSON(&vf); err != nil {
		c.AbortWithError(http.StatusBadRequest, util.GetValidate(err, &vf))
		return
	}
	c.Set("res", commonService.VfCaptcha(vf))
	//return common.VfCaptcha(vf)
}

// @Summary	上传单个文件
// @Produce	mpfd
// @Tags		public
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File to be uploaded"
// @Success	200 {object} comDto.ResponseData
// @Router		/upload/file [post]
func UploadFile(c *gin.Context) {
	//err := c.Request.ParseMultipartForm(10 << 20)
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
		c.Error(err)
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		c.Error(errors.New(msg.FILE_TYPE_ERROR))

		return
		//c.Error(errors.New(util.FILE_TYPE_ERROR))
	}
	filename := file.Filename
	newFileName := util.GenerateUniqueFileName(filename)
	filePath := filepath.Join(global.FilePath, newFileName)

	c.SaveUploadedFile(file, filePath)
	c.Set("res", filePath)

}

// @Summary	上传图片
// @Produce	mpfd
// @Tags		public
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File to be uploaded"
// @Success	200 {object} comDto.ResponseData
// @Router		/upload/pic [post]
func UploadPic(c *gin.Context) {
	//err := c.Request.ParseMultipartForm(10 << 20)
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
		c.Error(err)
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		c.Error(errors.New(msg.FILE_TYPE_ERROR))

		return
		//c.Error(errors.New(util.FILE_TYPE_ERROR))
	}
	filename := file.Filename
	newFileName := util.GenerateUniqueFileName(filename)
	filePath := filepath.Join(global.FilePath, newFileName)

	c.SaveUploadedFile(file, filePath)
	c.Set("res", filePath)

}

// @Summary	上传多个文件
// @Produce	mpfd
// @Tags		public
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File to be uploaded"
// @Success	200 {object} comDto.ResponseData
// @Router		/uploads/files [post]
func UploadFiles(c *gin.Context) {
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
		c.Error(errors.New(msg.FILE_TOO_LARGE))
	}
	form, err := c.MultipartForm()
	if err != nil {
		c.Error(errors.New(msg.FILE_TYPE_ERROR))
	}
	var filePaths []string
	// 获取所有图片
	files := form.File["files"]
	// 遍历所有图片
	for _, file := range files {
		filename := file.Filename
		newFileName := util.GenerateUniqueFileName(filename)
		filePath := filepath.Join(global.FilePath, newFileName)

		filePaths = append(filePaths, filePath)
		// 逐个存
		if err = c.SaveUploadedFile(file, filePath); err != nil {
			c.Error(errors.New(msg.FILE_SAVE_ERROR))
		}
	}
	c.Set("res", filePaths)

}

// @Summary	上传视频
// @Produce	mpfd
// @Tags		public
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File to be uploaded"
// @Success	200 {object} comDto.ResponseData
// @Router		/upload/files [post]
func UploadVideo(c *gin.Context) {
	//if err := c.Request.ParseMultipartForm(10 << 100); err != nil {
	//	c.Error(err)
	//}
	file, _ := c.FormFile("video")
	//获取上传文件的类型
	filetype := file.Header.Get("Content-Type")
	types := strings.Split(filetype, "/")
	fmt.Println(types, "文件类型")
	if types[0] != "video" {
		c.Error(errors.New(msg.FILE_TYPE_ERROR))
	}
	filename := file.Filename
	newFileName := util.GenerateUniqueFileName(filename)
	filePath := filepath.Join(global.FilePath, newFileName)
	//name := time.Now().Unix()
	//filename := file.Filename
	//suffix := strings.Split(filename, ".")
	//nameSuffix := suffix[1]
	//t := util.ExistIn(nameSuffix, global.VideoType)
	//if !t {
	//	res.Error(http.StatusBadRequest, util.FILE_SUFFIX_ERROR)
	//	return
	//}
	//file.Filename = strconv.FormatInt(name, 10) + "." + nameSuffix
	//filePath := path.Join(global.VideoPath)
	//_, e := os.Stat(filePath)
	//if e != nil {
	//	os.Mkdir(global.VideoPath, os.ModePerm)
	//}
	//pa := path.Join("./"+global.VideoPath+"/", file.Filename)
	err := c.SaveUploadedFile(file, filePath)
	if err != nil {
		return
	}
	c.Set("res", filePath)
}

/**
 * @Author Khr
 * @Description //TODO 文件流下载
 * @Date 16:50 2023/10/16
 * @Param
 * @return
 **/

// @Summary	文件流下载
// @Produce	mpfd
// @Tags		public
// @Success	200 {object} comDto.ResponseData
// @Router		/download [get]
func DownloadFile(c *gin.Context) {
	fileName := c.Query("name")
	// 文件路径
	filePath := global.FilePath + "./" + fileName

	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()
	// 获取文件信息
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 设置响应头，告诉浏览器这是一个文件下载
	c.Header("Content-Disposition", "attachment; filename="+filePath)
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Length", string(fileSize(filePath)))

	// 将文件内容写入响应主体
	http.ServeContent(c.Writer, c.Request, filePath, fileInfo.ModTime(), file)

}
func fileSize(filename string) int64 {
	file, err := os.Stat(filename)
	if err != nil {
		return 0
	}
	return file.Size()
}
