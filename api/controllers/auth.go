package controller

import (
	"encoding/json"
	"erp/config"
	dto "erp/dto/auth"
	service "erp/service"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/google"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/copier"
	"go.uber.org/zap"
)

type AuthController struct {
	AuthService service.AuthService
	logger      *zap.Logger
	config      *config.Config
	BaseController
}

func NewAuthController(c *gin.RouterGroup, authService service.AuthService, logger *zap.Logger, config *config.Config) *AuthController {
	controller := &AuthController{
		AuthService: authService,
		logger:      logger,
		config:      config,
	}
	g := c.Group("/auth")
	g.POST("/register", controller.Register)
	g.POST("/login", controller.Login)
	g.GET("/google/login", controller.GoogleLogin)
	g.GET("/google/callback", controller.GoogleCallback)
	g.GET("/facebook/login", controller.facebookLogin)
	g.GET("/facebook/callback", controller.facebookCallback)
	return controller
}

func (b *AuthController) getGoogleOAuthConfig() oauth2.Config {
	return oauth2.Config{
		RedirectURL:  b.config.GoogleOAuth.RedirectURL, // Replace with your callback URL
		ClientID:     b.config.GoogleOAuth.ClientID,
		ClientSecret: b.config.GoogleOAuth.ClientSecret,
		Scopes:       b.config.GoogleOAuth.Scopes,
		Endpoint:     google.Endpoint,
	}
}
func (b *AuthController) getFacebookConfig() oauth2.Config {
	return oauth2.Config{
		ClientID:     b.config.FacebookOAuth.AppID,
		ClientSecret: b.config.FacebookOAuth.AppSecret,
		Endpoint:     facebook.Endpoint,
		RedirectURL:  b.config.FacebookOAuth.RedirectURL,
		Scopes:       []string{"email"},
	}
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
	authConfig := b.getGoogleOAuthConfig()
	url := authConfig.AuthCodeURL("", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusFound, url)
}

func (b *AuthController) GoogleCallback(c *gin.Context) {
	code := c.Query("code")
	authConfig := b.getGoogleOAuthConfig()
	token, err := authConfig.Exchange(c, code)
	if err != nil {
		b.logger.Error(fmt.Sprintf("Cannot register with google: %+v", err))
		b.ResponseError(c, http.StatusInternalServerError, []error{err})
		return
	}

	client := authConfig.Client(c, token)

	userInfo, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo?alt=json&access_token" + token.AccessToken)
	if err != nil {
		b.logger.Error(fmt.Sprintf("Cannot register with google: %+v", err))
		b.Response(c, http.StatusInternalServerError, "Cannot login by google", nil)
		return
	}

	defer userInfo.Body.Close()

	var data map[string]interface{}
	decoder := json.NewDecoder(userInfo.Body)
	if err := decoder.Decode(&data); err != nil {
		// Handle JSON decoding error
		b.logger.Error(fmt.Sprintf("Cannot register with google: %+v", err))
		b.Response(c, http.StatusInternalServerError, "Cannot login by google", nil)
		return
	}
	fmt.Println(data)
	var response dto.UserGoogleRequest
	if err := mapstructure.Decode(data, &response); err != nil {
		// Handle JSON unmarshaling error
		b.logger.Error(fmt.Sprintf("Cannot unmarshal JSON response: %+v", err))
		b.Response(c, http.StatusInternalServerError, "Cannot login by google", nil)
		return
	}

	var req dto.UserGoogleRequest

	err = copier.Copy(&req, &response)
	if err != nil {
		b.logger.Error("Cannot register with google")
		b.Response(c, http.StatusInternalServerError, "Cannot login by google", nil)
		return
	}
	_, err = b.AuthService.RegisterByGoogle(c.Request.Context(), req)
	if err != nil {
		res, err := b.AuthService.LoginByGoogle(c.Request.Context(), dto.LoginByGoogleRequest{
			Email:    req.Email,
			GoogleId: req.GoogleID,
		})
		if err != nil {
			b.Response(c, http.StatusInternalServerError, "Cannot login by google account", nil)
			return
		}
		b.Response(c, http.StatusOK, "success", res)
		return
	}
	b.Response(c, http.StatusOK, "success", nil)
}

func (b *AuthController) facebookLogin(c *gin.Context) {
	authConfig := b.getFacebookConfig()
	url := authConfig.AuthCodeURL("", oauth2.AccessTypeOffline)
	fmt.Println("url la: ", url)
	c.Redirect(http.StatusFound, url)
}

func (b *AuthController) facebookCallback(c *gin.Context) {
	code := c.DefaultQuery("code", "")
	if code == "" {
		b.Response(c, http.StatusBadRequest, "Missing code parameter", nil)
		return
	}
	authConfig := b.getFacebookConfig()
	token, err := authConfig.Exchange(c, code)
	if err != nil {
		b.Response(c, http.StatusBadRequest, fmt.Sprintf("Error exchanging code: %s", err.Error()), nil)
		return
	}
	client := authConfig.Client(c, token)
	resp, err := client.Get(b.config.FacebookOAuth.GraphAPIURL)
	if err != nil {
		b.Response(c, http.StatusInternalServerError, fmt.Sprintf("Error fetching user data: %s", err.Error()), nil)
		return
	}
	defer resp.Body.Close()

	var userData map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userData); err != nil {
		b.Response(c, http.StatusInternalServerError, fmt.Sprintf("Error decoding user data: %s", err.Error()), nil)
		return
	}
	b.Response(c, http.StatusOK, "success", userData)
}
