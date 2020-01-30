package conf

import (
	"flag"
	"os"
	"strconv"
	"time"

	xtime "chatroom/xframe/time"

	"github.com/BurntSushi/toml"
)

var (
	confPath string
	debug    bool

	// Conf config
	Conf *Config
)

func init() {
	var (
		defDebug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	)
	flag.StringVar(&confPath, "conf", "conf/router-example.toml", "default config path.")
	flag.BoolVar(&debug, "debug", defDebug, "server debug. or use DEBUG env variable, value: true/false etc.")
}

// Init init config.
func Init() (err error) {
	Conf = Default()
	md, err := toml.DecodeFile(confPath, &Conf)
	println(md.Keys())
	return
}

// Default new a config with specified defualt value.
func Default() *Config {
	return &Config{
		Debug: debug,
		RPCServer: &RPCServer{
			Network:           "tcp",
			Addr:              ":3109",
			Timeout:           xtime.Duration(time.Second),
			IdleTimeout:       xtime.Duration(time.Second * 60),
			MaxLifeTime:       xtime.Duration(time.Hour * 2),
			ForceCloseWait:    xtime.Duration(time.Second * 20),
			KeepAliveInterval: xtime.Duration(time.Second * 60),
			KeepAliveTimeout:  xtime.Duration(time.Second * 20),
		},
		Xlog: &Xlog{
			Path:  "default.log",
			Level: "debug",
		},
	}
}

type Config struct {
	Debug     bool
	RPCServer *RPCServer
	Xlog      *Xlog
}

// RPCServer is RPC server config.
type RPCServer struct {
	Network           string
	Addr              string
	Timeout           xtime.Duration
	IdleTimeout       xtime.Duration
	MaxLifeTime       xtime.Duration
	ForceCloseWait    xtime.Duration
	KeepAliveInterval xtime.Duration
	KeepAliveTimeout  xtime.Duration
}

type Xlog struct {
	Path  string
	Level string
}
