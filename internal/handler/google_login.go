package handler

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2/google"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	goauth2 "golang.org/x/oauth2"
	"google.golang.org/api/oauth2/v2"

	"nitflex/util"
)

var oauthConfig *goauth2.Config

func init() {
	_ = godotenv.Load(".env")
	oauthConfig = &goauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_CALLBACK_URL"),
		Scopes:       []string{oauth2.UserinfoEmailScope, oauth2.UserinfoProfileScope},
		Endpoint:     google.Endpoint,
	}
}

func (h *Handler) GoogleLogin(c *gin.Context) {
	url := oauthConfig.AuthCodeURL(os.Getenv("OAUTH_STATE"))
	fmt.Println("->")
	fmt.Println("->", os.Getenv("GOOGLE_CLIENT_SECRET"))
	fmt.Println("->", os.Getenv("GOOGLE_CALLBACK_URL"))
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (h *Handler) GoogleCallback(c *gin.Context) {
	state := c.Query("state")
	if state != os.Getenv("OAUTH_STATE") {
		c.JSON(http.StatusBadRequest, util.FailResponse("Fail to login via google"))
		return
	}

	code := c.Query("code")
	token, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.FailResponse("Fail to login via google"))
		return
	}

	client := oauthConfig.Client(context.Background(), token)
	service, err := oauth2.New(client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.FailResponse("Fail to login via google"))
		return
	}

	userInfo, err := service.Userinfo.Get().Do()
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.FailResponse("Fail to login via google"))
		return
	}

	// run biz
	response, err := h.biz.GoogleLogin(c.Request.Context(), userInfo)
	if err != nil {
		c.JSON(http.StatusOK, util.FailResponse(err.Error()))
		return
	}

	// response client
	c.JSON(http.StatusOK, util.SuccessResponse(response))
}
