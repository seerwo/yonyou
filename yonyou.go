package yonyou

import (
	"github.com/seerwo/yonyou/cache"
	"github.com/seerwo/yonyou/nc"
	ncconfig "github.com/seerwo/yonyou/nc/config"
	log "github.com/sirupsen/logrus"
	"os"
)

func init(){
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

//yonyou struct
type Yonyou struct {
	cache  cache.Cache
}

//NewYonyou init
func NewYonyou() *Yonyou {
	return &Yonyou{}
}

//SetCache set cache
func (yy *Yonyou) SetCache(cache cache.Cache) {
	yy.cache = cache
}

//GetNC get nc instance
func (yy *Yonyou) GetNC(cfg *ncconfig.Config) *nc.NC{
	if cfg.Cache == nil {
		cfg.Cache = yy.cache
	}
	return nc.NewNC(cfg)
}