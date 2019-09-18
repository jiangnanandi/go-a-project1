package middleware

import (
	"net/http"
	"net/http/httputil"
	"runtime/debug"

	log "github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				httprequest, _ := httputil.DumpRequest(c.Request, false)
				log.Error("panic recovered:httprequest[%s] err[%s] stack[%s]", string(httprequest), err, string(debug.Stack()))
				c.String(http.StatusOK, `{"errNo":99999,"errStr":"unknow_error","data":{}}`)
				c.Abort()
			}
		}()
		c.Next()
	}
}
