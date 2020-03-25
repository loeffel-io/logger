package gin_middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/loeffel-io/logger/v2"
	"log"
)

func Logger(logger *logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		gWriter := &writer{
			body:           new(bytes.Buffer),
			ResponseWriter: c.Writer,
		}

		c.Writer = gWriter
		c.Next()

		if c.IsAborted() {
			log.Printf("%+v %+v %+v", c.Request.RequestURI, gWriter.body, gWriter.Status())
			logger.Log(fmt.Errorf("asdf"))
		}
	}
}
