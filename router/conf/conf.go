package conf

import (
	"flag"
	"os"
	"strconv"
	"time"

	xtime "router/xframe/time"

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
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}

// Default new a config with specified defualt value.
func Default() *Config {
	return &Config{
		Debug: debug,
		//Discovery: &naming.Config{Region: region, Zone: zone, Env: deployEnv, Host: host},
		RPCClient: &RPCClient{
			Dial:    xtime.Duration(time.Second),
			Timeout: xtime.Duration(time.Second),
		},
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
	}
}

// Config is comet config.
type Config struct {
	Debug     bool
	RPCClient *RPCClient
	RPCServer *RPCServer
	Xlog      *Xlog
}

// RPCClient is RPC client config.
type RPCClient struct {
	Dial    xtime.Duration
	Timeout xtime.Duration
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
