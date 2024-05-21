package credential

import (
	"encoding/json"
	"fmt"
	"github.com/seerwo/yonyou/cache"
	"github.com/seerwo/yonyou/util"
	"sync"
	"time"
)

const (
	//AccessTokenURL get access_token interface
	accessTokenURL = "https://api.weixin.qq.com/cgi-bin/token"
	//CacheKeyNCPrefix yonyou cache key prefix
	CacheKeyNCPrefix = "goyonyou_nc_"
)

//DefaultAccessToken 默认AccessToken 获取
type DefaultAccessToken struct {
	appID           string
	appSecret       string
	cacheKeyPrefix  string
	cache           cache.Cache
	accessTokenLock *sync.Mutex
}

//NewDefaultAccessToken new DefaultAccessToken
func NewDefaultAccessToken(appID, appSecret, cacheKeyPrefix string, cache cache.Cache) AccessTokenHandle {
	if cache == nil {
		panic("cache is ineed")
	}
	return &DefaultAccessToken{
		appID:           appID,
		appSecret:       appSecret,
		cache:           cache,
		cacheKeyPrefix:  cacheKeyPrefix,
		accessTokenLock: new(sync.Mutex),
	}
}

//ResAccessToken struct
type ResAccessToken struct {
	util.CommonError

	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

func (ak *DefaultAccessToken) GetUrlParam() (accessParam string, err error) {
	accessParam = ak.appID
	return
}

//GetAccessToken get access_token,from cache，or from server.
func (ak *DefaultAccessToken) GetAccessToken() (accessToken string, err error) {
	//add lock，to protect token，cache invalid from server
	ak.accessTokenLock.Lock()
	defer ak.accessTokenLock.Unlock()

	accessTokenCacheKey := fmt.Sprintf("%s_access_token_%s", ak.cacheKeyPrefix, ak.appID)
	val := ak.cache.Get(accessTokenCacheKey)
	if val != nil {
		accessToken = val.(string)
		return
	}

	//cache invalid from server
	var resAccessToken ResAccessToken
	resAccessToken, err = GetTokenFromServer(ak.appID, ak.appSecret)
	if err != nil {
		return
	}

	expires := resAccessToken.ExpiresIn - 1500
	err = ak.cache.Set(accessTokenCacheKey, resAccessToken.AccessToken, time.Duration(expires)*time.Second)
	if err != nil {
		return
	}
	accessToken = resAccessToken.AccessToken
	return
}

//GetTokenFromServer force from server
func GetTokenFromServer(appID, appSecret string) (resAccessToken ResAccessToken, err error) {
	url := fmt.Sprintf("%s?grant_type=client_credential&appid=%s&secret=%s", accessTokenURL, appID, appSecret)
	var body []byte
	body, err = util.HTTPGet(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &resAccessToken)
	if err != nil {
		return
	}
	if resAccessToken.ErrMsg != "" {
		err = fmt.Errorf("get access_token error : errcode=%v , errormsg=%v", resAccessToken.ErrCode, resAccessToken.ErrMsg)
		return
	}
	return
}
