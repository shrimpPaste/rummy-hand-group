package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rummy-logic-v3/pkg/request"
)

type BaseResponse struct {
	Success bool `json:"success"`
	Data    any  `json:"data"`
}

type SuccessResponse struct {
	Success bool `json:"success"`
	Data    any  `json:"data"`
}

type FailResponse struct {
	Success  bool   `json:"success"`
	Data     any    `json:"data,omitempty"`
	ErrorTip string `json:"error_tip"`
}

type Paginate struct {
	Total    int64 `json:"total"`
	Page     int   `json:"page"`
	PageSize int   `json:"page_size"`
}

func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, SuccessResponse{
		Success: true,
		Data:    data,
	})
}

func Fail(c *gin.Context, err error) {
	c.JSON(http.StatusOK, FailResponse{
		Success:  false,
		ErrorTip: err.Error(),
	})
}

func SuccessPaginate(c *gin.Context, data any, total int64) {
	page, pageSize := request.GetPageAndPageSize(c)
	c.JSON(http.StatusOK, SuccessResponse{
		Success: true,
		Data: gin.H{
			"data": data,
			"paginate": Paginate{
				Total:    total,
				Page:     (page + 1) / (pageSize + 1),
				PageSize: pageSize,
			},
		},
	})
}

func FailHttpStatus(c *gin.Context, err error, httpStatus int) {
	c.JSON(httpStatus, FailResponse{
		Success:  false,
		Data:     nil,
		ErrorTip: err.Error(),
	})
}
