package middleware

import(
	
	"github.com/gin-gonic/gin"
	"fmt"
	"io/ioutil"
	"bytes"
	
	
	_"net/url"

)
func PrintRequestBody() gin.HandlerFunc {
	return func(c *gin.Context) {
		buf, _ := ioutil.ReadAll(c.Request.Body)
		body := string(buf)
		fmt.Println("Request Body:", body)

		// Write body back for further processing
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(buf))
		c.Next()
	}
}