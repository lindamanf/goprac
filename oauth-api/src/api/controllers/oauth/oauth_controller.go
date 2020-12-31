package oauth

import (
	"net/http"

	"work/oauth-api/src/api/services"

	"github.com/gin-gonic/gin"
)

func CreateAccessToken(c *gin.Context) {
	var request AccessTokenRequest
	if err := c.ShouldBindJson(&request); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	token, err := services.OauthService.CreateAccessToken(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, token)
}

func GetAccessToken(c *gin.Context) {
	token, err := services.GetAccessTokenByToken(c.Param("token_id"))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, token)
}
