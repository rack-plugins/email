package email

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func help(ctx *gin.Context) {
	ctx.String(http.StatusOK, `POST /mailto
-H "content-type: application/json" \
-d '{
	"mailto": "lmr@epurs.com",
	"subject": "test subject",
	"body": "测试发送成功！"
}'
`)
}
