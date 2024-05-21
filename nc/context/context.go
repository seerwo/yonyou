package context

import (
	"github.com/seerwo/yonyou/credential"
	"github.com/seerwo/yonyou/nc/config"
)

//Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}
