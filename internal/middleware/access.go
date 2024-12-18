package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"rummy-logic-v3/global"
	"rummy-logic-v3/pkg/logger"
	"rummy-logic-v3/pkg/response"
	"rummy-logic-v3/pkg/xerr"
	"time"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

type Request struct {
	Uid       int64  `form:"uid,string" binding:"required"`
	Currency  string `form:"currency" binding:"required"`
	AppKey    string `form:"app_key" binding:"required"`
	Timestamp int64  `form:"timestamp,string" binding:"required"`
	Sign      string `form:"sign" binding:"required"`
	AppSecret string `form:"app_secret"`
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 替换 ResponseWriter，用于捕获响应数据
		writer := &AccessLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = writer

		beginTime := time.Now()

		// 读取并缓存请求体
		bodyBytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			global.Logger.WithCaller(1).Error(c, "访问控制器读取请求体错误: "+err.Error())
			response.Success(c, xerr.InvalidParams.ErrorMsg())
			c.Abort()
			return
		}

		// 将请求体重新写回，以便后续处理函数可以再次读取
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		var req Request
		if err := c.ShouldBind(&req); err != nil {
			global.Logger.WithCaller(1).Errorf(c, "访问控制器解析请求数据错误: %v", err)
			response.Success(c, xerr.InvalidParams.ErrorMsg())
			c.Abort()
			return
		}

		// 将请求体重新写回，以便后续处理函数可以再次读取
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		c.Next()
		endTime := time.Now()

		// 解析响应数据
		var resp map[string]any
		if len(writer.body.Bytes()) > 0 {
			err := json.Unmarshal(writer.body.Bytes(), &resp)
			if err != nil {
				global.Logger.WithCaller(1).Error(c, "访问控制器解析返回数据错误: "+err.Error())
				response.Success(c, xerr.InvalidParams.ErrorMsg())
				return
			}
		}

		// 记录请求和响应日志
		fields := logger.Fields{
			"request":  req,
			"response": resp,
		}

		s := "access log: method:%s, status_code:%d, begin time: %s, end time: %s"
		global.Logger.WithCaller(1).WithFields(fields).Infof(c, s, c.Request.Method, writer.Status(), beginTime.Format(time.DateTime), endTime.Format(time.DateTime))
	}
}

func decodeSkey(ctx *gin.Context, skey string) (int64, error) {
	return 0, nil
	//resp := &service.UserMoneyResp{}
	//redisSrv := service.NewRedisService(ctx)
	//err := redisSrv.GetUserMoney(skey, resp)
	//if err != nil {
	//	return 0, err
	//}
	//
	//id, err := strconv.ParseInt(resp.Id, 10, 64)
	//if err != nil {
	//	global.Logger.Error(ctx, "解析用户id错误: "+err.Error())
	//	return 0, xerr.ServerError
	//}
	//return id, nil
}
