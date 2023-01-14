package controller

import (
	"github.com/gin-gonic/gin"
	"go-hello/common"
	"go-hello/model"
	"net/http"
)

func ListCV(context *gin.Context) {
	var filter model.CVFilter
	if err := context.ShouldBind(&filter); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var paging common.Paging
	if err := context.ShouldBind(&paging); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	paging.Fulfill()
	result, err := model.ListCVByCondition(nil, &filter, &paging)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	context.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
}

func CreateCV(context *gin.Context) {
	var createCV model.CV

	if err := context.ShouldBind(&createCV); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	cv := model.CV{
		Name:     createCV.Name,
		Mobile:   createCV.Mobile,
		Email:    createCV.Email,
		Github:   createCV.Github,
		Linkedin: createCV.Linkedin,
		Summary:  createCV.Summary,
		Skills:   createCV.Skills,
	}

	createdUser, err := cv.Create()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	context.JSON(http.StatusOK, gin.H{"user": createdUser})
}
