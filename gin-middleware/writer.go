package gin_middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
)

type writer struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (writer writer) Write(b []byte) (int, error) {
	writer.body.Write(b)
	return writer.ResponseWriter.Write(b)
}
