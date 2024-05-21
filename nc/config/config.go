package config

import "github.com/seerwo/yonyou/cache"

// Config config for nc
type Config struct {
	AppID          string `json:"app_id"`           //appid
	AppSecret      string `json:"app_secret"`       //appsecret
	Token          string `json:"token"`            //token
	EncodingAESKey string `json:"encoding_aes_key"` //EncodingAESKey
	Cache          cache.Cache
}
