package oauth

import (
	"encoding/json"
	"fmt"
	"github.com/seerwo/yonyou/nc/context"
	"github.com/seerwo/yonyou/util"
	"net/http"
	"net/url"
)

const (
	redirectOauthURL       = "%s%s%s%s"
	webAppRedirectOauthURL = "%s%s%s%s"
	accessTokenURL         = "%s%s%s"
	refreshAccessTokenURL  = "%s%s"
	userInfoURL            = "%s%s%s%s"
	checkAccessTokenURL    = "%s%s"
)

//Oauth save user oauth message
type Oauth struct {
	*context.Context
}

//NewOauth instance
func NewOauth(context *context.Context) *Oauth {
	auth := new(Oauth)
	auth.Context = context
	return auth
}

//GetRedirectURL get a url
func (oauth *Oauth) GetRedirectURL(redirectURI, scope, state string) (string, error) {
	//url encode
	urlStr := url.QueryEscape(redirectURI)
	return fmt.Sprintf(redirectOauthURL, oauth.AppID, urlStr, scope, state), nil
}

//GetWebAppRedirectURL to jump url
func (oauth *Oauth) GetWebAppRedirectURL(redirectURI, scope, state string) (string, error) {
	urlStr := url.QueryEscape(redirectURI)
	return fmt.Sprintf(webAppRedirectOauthURL, oauth.AppID, urlStr, scope, state), nil
}

//Redirect to jump web oauth
func (oauth *Oauth) Redirect(writer http.ResponseWriter, req *http.Request, redirectURI, scope, state string) error {
	location, err := oauth.GetRedirectURL(redirectURI, scope, state)
	if err != nil {
		return err
	}
	http.Redirect(writer, req, location, http.StatusFound)
	return nil
}

// ResAccessToken get access_token return
type ResAccessToken struct {
	util.CommonError

	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenID       string `json:"openid"`
	Scope        string `json:"scope"`

}

// GetUserAccessToken from web get code and take access_token
func (oauth *Oauth) GetUserAccessToken(code string) (result ResAccessToken, err error) {
	urlStr := fmt.Sprintf(accessTokenURL, oauth.AppID, oauth.AppSecret, code)
	var response []byte
	response, err = util.HTTPGet(urlStr)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetUserAccessToken error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}
	return
}

//RefreshAccessToken refresh access_token
func (oauth *Oauth) RefreshAccessToken(refreshToken string) (result ResAccessToken, err error) {
	urlStr := fmt.Sprintf(refreshAccessTokenURL, oauth.AppID, refreshToken)
	var response []byte
	response, err = util.HTTPGet(urlStr)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetUserAccessToken error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}
	return
}

//CheckAccessToken check access_token
func (oauth *Oauth) CheckAccessToken(accessToken, openID string) (b bool, err error) {
	urlStr := fmt.Sprintf(checkAccessTokenURL, accessToken, openID)
	var response []byte
	response, err = util.HTTPGet(urlStr)
	if err != nil {
		return
	}
	var result util.CommonError
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		b = false
		return
	}
	b = true
	return
}
