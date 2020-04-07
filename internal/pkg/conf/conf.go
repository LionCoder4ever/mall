package conf

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"mall/internal/pkg/database/sql"
	"mall/internal/pkg/log"
)

type Config struct {
	Log   *log.Config
	MySQL *sql.Config
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
