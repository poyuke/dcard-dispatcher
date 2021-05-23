package controller

import (
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 上傳檔案
func (e Env) UploadFileHandler(c *gin.Context) {
	code := http.StatusOK
	var rep interface{}
	fileContent := c.PostForm("fileContent")
	hashId, err := e.DAO.SetRedis(fileContent);
	if err != nil {
		code = http.StatusInternalServerError
		rep = gin.H{
			"message": err,
		}
	} else {
		rep = gin.H{
			"jobId": hashId,
		}
	}
	c.JSON(code, rep)
	return
}

// 處理檔案
func (e Env) ProcessFileHandler(c *gin.Context) {
	code := http.StatusOK
	var rep interface{}
	id := c.Param("id")
	status, err := e.DAO.GetRedisByKey(id, "status")
	if err == nil {
		if status == "" || status == "<nil>" {
			code = http.StatusInternalServerError
			rep = gin.H{
				"message": "Job id does not exist",
			}
		} else if status == "waiting" || status == "processing" {
			rep = gin.H{
				"status": status,
			}
		} else if status == "success" {
			sha, err := e.DAO.GetRedisByKey(id, "sha")
			if err != nil {
				code = http.StatusInternalServerError
				rep = gin.H{
					"message": "get sha-1 file content failed",
				}
			} else {
				if sha != "" || status != "<nil>" {
					rep = gin.H{
						"sha": sha,
					}
				} else {
					code = http.StatusInternalServerError
					rep = gin.H{
						"message": "sha-1 file content does not exist",
					}
				}
			}
		}
	} else {
		code = http.StatusInternalServerError
		rep = gin.H{
			"message": err,
		}
	}
	c.JSON(code, rep)
	return
}

// 取得所有檔案狀態
func (e Env) GetJobsListHandler(c *gin.Context) {
	code := http.StatusOK
	var rep interface{}
	jobs, err := e.DAO.GetRedisAllStatus();
	if err != nil {
		code = http.StatusInternalServerError
		rep = gin.H{
			"message": err,
		}
	} else {
		rep = gin.H{
			"jobs": jobs,
		}
	}
	c.JSON(code, rep)
	return
}
