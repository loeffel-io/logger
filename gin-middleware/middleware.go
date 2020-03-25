package gin_middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/loeffel-io/logger/v2"
)

func Logger(logger *logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		ginWriter := &writer{
			body:           new(bytes.Buffer),
			ResponseWriter: c.Writer,
		}

		c.Writer = ginWriter
		c.Next()

		if c.IsAborted() {
			logger.Log(fmt.Errorf(
				"aborted (%d) with %s @ %s",
				ginWriter.Status(),
				ginWriter.body.String(),
				c.Request.RequestURI,
			))
		}
	}
}
