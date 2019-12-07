package conf

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"mall/library/database/sql"
	"mall/library/log"
)

type Config struct {
	Log *log.Config
	//BM         *bm.ServerConfig
	//Tracer     *trace.Config
	//Memcache   *Memcache
	MySQL *sql.Config
	//RPCClient2 *RPC
	// grpc server
	//WardenServer *warden.ServerConfig
}

var (
	confPath string
	Conf     Config
)

func init() {
	flag.StringVar(&confPath, "conf", "", "default config path")
}

// Init init conf
func Load() error {
	if confPath != "" {
		return local()
	}
	return fmt.Errorf("invalid conf path")
}

func local() (err error) {
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}
