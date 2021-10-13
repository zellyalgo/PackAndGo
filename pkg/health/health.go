package health

import(

	"github.com/gin-gonic/gin"
)

//Health method, does nothing to return 200 if the server can recieve the request
func Check(c *gin.Context) {
}