package request

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func ParseIDParam(param string) (int64, error) {
	// 将 id 参数转换为 int64 类型
	intId, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return 0, err
	}

	return intId, nil
}

type Paginate struct {
	Page     int `form:"page"`
	PageSize int `form:"page_size"`
}

func GetPageAndPageSize(c *gin.Context) (int, int) {
	var parameter Paginate
	if err := c.ShouldBind(&parameter); err != nil {
		parameter.Page = 1
		parameter.PageSize = 20
	}
	if parameter.Page <= 0 {
		parameter.Page = 1
	}
	if parameter.PageSize <= 0 {
		parameter.PageSize = 20
	}
	return parameter.Page - 1, parameter.PageSize
}
