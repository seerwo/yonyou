package credential


//AccessTokenHandle AccessToken interface
type AccessTokenHandle interface {
	GetAccessToken() (accessToken string, err error)
	GetUrlParam()(accessParam string, err error)
}
