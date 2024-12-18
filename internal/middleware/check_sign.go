package middleware

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"rummy-logic-v3/global"
	"rummy-logic-v3/pkg/response"
	"rummy-logic-v3/pkg/tool"
	"rummy-logic-v3/pkg/xerr"
	"runtime/debug"
)

func CheckSign() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyBytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			global.Logger.Error(c, "参数校验签名器读取请求体错误: "+err.Error())
			response.Success(c, xerr.InvalidParams.ErrorMsg())
			c.Abort()
			return
		}

		// 将请求体重新写回，以便后续处理函数可以再次读取
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		var req Request
		if err := c.ShouldBind(&req); err != nil {
			global.Logger.Error(c, "校验签名中间件请求参数校验错误: ", err, debug.Stack())
			response.Success(c, xerr.InvalidParams.ErrorMsg())
			c.Abort()
			return
		}

		// 将请求体重新写回，以便后续处理函数可以再次读取
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		//req.AppSecret = configs.NGConf.GameConf.AppSecret

		// 自定义排序函数
		//sort.Slice(keys, func(i, j int) bool {
		//	return customSort(keys[i], keys[j])
		//})

		// 拼接排序后的键值对
		//var builder strings.Builder
		//for _, key := range keys {
		//	value := fmt.Sprintf("%v", req[key])
		//	builder.WriteString(key + "=" + value + "&")
		//}

		signatureString := fmt.Sprintf("%v%v%v", req.AppKey, req.AppSecret, req.Timestamp)

		// 移除最后一个多余的 "&"
		//signatureString := strings.TrimSuffix(builder.String()+global.RequestParamSecret, "&")

		// 生成 MD5 签名
		md5Hash := md5.New()
		md5Hash.Write([]byte(signatureString))
		signature := hex.EncodeToString(md5Hash.Sum(nil))

		if len(req.Sign) < 16 {
			global.Logger.Errorf(c, "签名不足16位，签名错误, %s", req.Sign)
			response.Fail(c, xerr.InvalidParams)
			c.Abort()
			return
		}

		if signature != req.Sign[8:len(req.Sign)-8] {
			global.Logger.Errorf(c, "签名错误应该得到: %s, 实际得到: %s", tool.GenerateRandomString(8)+signature+tool.GenerateRandomString(8), req.Sign)
			response.Success(c, xerr.InvalidParams.ErrorMsg())
			c.Abort()
			return
		}
	}
}

// 自定义排序规则：1-8, a-z, A-Z, 其他符号
func customSort(a, b string) bool {
	// 定义排序优先级
	priority := func(r rune) int {
		switch {
		case r >= '1' && r <= '8':
			return int(r) - '1'
		case r >= 'a' && r <= 'z':
			return int(r) - 'a' + 8
		case r >= 'A' && r <= 'Z':
			return int(r) - 'A' + 34
		default:
			// 符号优先级排在最后
			return int(r) + 100
		}
	}

	// 比较两个字符串
	for i := 0; i < len(a) && i < len(b); i++ {
		if priority(rune(a[i])) != priority(rune(b[i])) {
			return priority(rune(a[i])) < priority(rune(b[i]))
		}
	}

	// 如果当前字符相同，则比较长度
	return len(a) < len(b)
}
