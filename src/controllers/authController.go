package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/dto"
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/services/interfaces"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "https://garage-best-team-ever.tk/google-callback",
		ClientID:     os.Getenv("GOOGLE_AUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_AUTH_CLIENT_SECRET"),
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/plus.me"},
		Endpoint: google.Endpoint,
	}
	randomState = "random"
)

type AuthController struct {
	googleAuthService interfaces.GoogleAuthServiceProvider
}

func NewAuthController(authService interfaces.GoogleAuthServiceProvider) *AuthController {
	return &AuthController{
		authService,
	}
}

func (controller *AuthController) GoogleLogin(writer http.ResponseWriter, request *http.Request) {
	url := googleOauthConfig.AuthCodeURL(randomState)
	fmt.Println(url)
	http.Redirect(writer, request, url, http.StatusTemporaryRedirect)
}

func (controller *AuthController) GoogleCallback(writer http.ResponseWriter, request *http.Request) {
	if request.FormValue("state") != randomState {
		fmt.Println("state is not valid")
		http.Redirect(writer, request, "/", http.StatusTemporaryRedirect)
		return
	}
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, request.FormValue("code"))
	fmt.Println(token.AccessToken)
	if err != nil {
		fmt.Printf("could not get token : %s \n", err.Error())
		http.Redirect(writer, request, "/", http.StatusTemporaryRedirect)
		return
	}
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		fmt.Printf("could not create get request : %s \n", err.Error())
		http.Redirect(writer, request, "/", http.StatusTemporaryRedirect)
		return
	}
	defer resp.Body.Close()
	user := &dto.User{}
	if err := json.NewDecoder(resp.Body).Decode(user); err != nil {
		errorJsonRespond(writer, http.StatusBadRequest, errJsonDecode)
		return
	}
	// create user if not exist
	if err := controller.googleAuthService.ResolveUser(user); err != nil {
		errorJsonRespond(writer, http.StatusBadRequest, err)
		return
	}

	accessToken, err := controller.googleAuthService.Authenticate(user)
	if err != nil {
		errorJsonRespond(writer, http.StatusInternalServerError, err)
		return
	}

	http.SetCookie(writer, &http.Cookie{
		Name:       "accessToken",
		Value:      accessToken,
		Path:       "/",
		RawExpires: time.Now().Add(controller.googleAuthService.GetAccessTokenTTL()).String(),
	})

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("could not parse response : %s \n", err.Error())
		http.Redirect(writer, request, "/", http.StatusTemporaryRedirect)
		return
	}
	fmt.Fprintf(writer, "Response: %s", content)
}

func (controller *AuthController) AuthorizationMW(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenCookie, err := r.Cookie("accessToken")
		if err != nil {
			errorJsonRespond(w, http.StatusUnauthorized, err)
			return
		}
		userId, err := controller.googleAuthService.Authorize(tokenCookie.Value)
		if err != nil {
			errorJsonRespond(w, http.StatusUnauthorized, err)
			return
		}
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), contextKeyId, userId)))
	})
}
