package middlewares

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
	"runtime"
)

//RecoverHandle 处理应用的panic
func RecoverHandle(c *fiber.Ctx) error {
	// Catch panics
	defer func() {
		if r := recover(); r != nil {
			//记录错误信息
			defaultStackTraceHandler(r)
			//返回给客户端的响应
			c.Status(500).JSON(fiber.Map{
				"code":    500,
				"message": "服务器内部错误",
				"data":    nil,
			})
		}
	}()
	return c.Next()
}

func defaultStackTraceHandler(e interface{}) {
	buf := make([]byte, 1024)
	buf = buf[:runtime.Stack(buf, false)]
	_, _ = os.Stderr.WriteString(fmt.Sprintf("panic: %v\n%s\n", e, buf))
}
