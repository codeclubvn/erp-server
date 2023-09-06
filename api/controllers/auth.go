package controller

import (
	"encoding/json"
	dto "erp/dto/auth"
	service "erp/service"
	"fmt"
	"github.com/jinzhu/copier"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/copier"
	"go.uber.org/zap"
)

var googleOauthConfig = oauth2.Config{
	RedirectURL:  "http://localhost:8080/api/auth/google/callback", // Replace with your callback URL
	ClientID:     "555161563716-820tf0ghjucnn9mirti5imlvkai8t6kv.apps.googleusercontent.com",
	ClientSecret: "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",
	Scopes: []string{"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile"},
	Endpoint: google.Endpoint,
}

type AuthController struct {
	AuthService service.AuthService
	logger      *zap.Logger
	BaseController
}

func NewAuthController(c *gin.RouterGroup, authService service.AuthService, logger *zap.Logger) *AuthController {
	controller := &AuthController{
		AuthService: authService,
		logger:      logger,
	}
	g := c.Group("/auth")
	g.POST("/register", controller.Register)
	g.POST("/login", controller.Login)
	g.GET("/google/login", controller.GoogleLogin)
	g.GET("/google/callback", controller.GoogleCallback)
	return controller
}

func (b *AuthController) Register(c *gin.Context) {
	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		b.ResponseError(c, http.StatusBadRequest, []error{err})
		return
	}
	b.logger.Info(fmt.Sprintf("Request info %+v", req))
	_, err := b.AuthService.Register(c.Request.Context(), req)

	if err != nil {
		b.ResponseError(c, http.StatusBadRequest, []error{err})
		return
	}
	b.Response(c, http.StatusOK, "success", nil)
}

func (b *AuthController) Login(c *gin.Context) {
	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		b.ResponseError(c, http.StatusBadRequest, []error{err})
		return
	}

	res, err := b.AuthService.Login(c.Request.Context(), req)

	if err != nil {
		b.ResponseError(c, http.StatusBadRequest, []error{err})
		return
	}
	b.Response(c, http.StatusOK, "success", res)
}

func (b *AuthController) GoogleLogin(c *gin.Context) {
	url := googleOauthConfig.AuthCodeURL("", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusFound, url)
}

func (b *AuthController) GoogleCallback(c *gin.Context) {
	code := c.Query("code")
	token, err := googleOauthConfig.Exchange(c, code)
	if err != nil {
		b.ResponseError(c, http.StatusInternalServerError, []error{err})
		return
	}

	client := googleOauthConfig.Client(c, token)

	userInfo, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo?alt=json&access_token" + token.AccessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer userInfo.Body.Close()

	data, err := ioutil.ReadAll(userInfo.Body)
	var response dto.UserGoogleRequest
	err = json.Unmarshal(data, &response)
	b.logger.Info(fmt.Sprintf("google data %+v", response))
	var req dto.UserGoogleRequest

	err = copier.Copy(&req, &response)
	if err != nil {
		b.logger.Error("Cannot register with google")
	}
	_, err = b.AuthService.RegisterByGoogle(c.Request.Context(), req)
	if err != nil {
		res, err := b.AuthService.LoginByGoogle(c.Request.Context(), dto.LoginByGoogleRequest{
			Email:    req.Email,
			GoogleId: req.GoogleID,
		})
		if err != nil {
			b.Response(c, http.StatusInternalServerError, "Cannot login by google account", nil)
		}
		b.Response(c, http.StatusOK, "success", res)
	}
	b.Response(c, http.StatusOK, "success", nil)
}
