package nc

import (
	"github.com/seerwo/yonyou/credential"
	"github.com/seerwo/yonyou/nc/config"
	"github.com/seerwo/yonyou/nc/context"
	"github.com/seerwo/yonyou/nc/oauth"
	"github.com/seerwo/yonyou/nc/sale"
)

//NC nc relate api
type NC struct {
	ctx *context.Context
}

//NewNC instance api related
func NewNC(cfg *config.Config) *NC {
	defaultAkHandle := credential.NewDefaultAccessToken(cfg.AppID, cfg.AppSecret, credential.CacheKeyNCPrefix, cfg.Cache)
	ctx := &context.Context{
		Config: cfg,
		AccessTokenHandle: defaultAkHandle,
	}
	return &NC{ctx: ctx}
}

//GetAccessToken get access_token
func(nc *NC) GetAccessToken()(string, error){
	return nc.ctx.GetAccessToken()
}

//GetOauth oauth2 web oauth
func (nc *NC) GetOauth() *oauth.Oauth {
	return oauth.NewOauth(nc.ctx)
}

//GetSale
func (nc *NC) GetSale() *sale.Sale {
	return sale.NewSale(nc.ctx)
}