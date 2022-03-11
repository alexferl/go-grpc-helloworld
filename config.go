package go_grpc_helloworld

import (
	"net"

	xconfig "github.com/alexferl/golib/config"
	xlog "github.com/alexferl/golib/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Config holds all configuration for our program
type Config struct {
	Config            *xconfig.Config
	Logging           *xlog.Config
	BindAddress       net.IP
	BindPort          uint
	GracefulTimeout   uint
	HealthBindAddress net.IP
	HealthBindPort    uint
}

// NewConfig creates a Config instance
func NewConfig() *Config {
	return &Config{
		Config:            xconfig.New("go_grpc_helloworld"),
		Logging:           xlog.DefaultConfig,
		BindAddress:       net.ParseIP("127.0.0.1"),
		BindPort:          50051,
		GracefulTimeout:   30,
		HealthBindAddress: net.ParseIP("127.0.0.1"),
		HealthBindPort:    50052,
	}
}

// addFlags adds all the flags from the command line
func (c *Config) addFlags(fs *pflag.FlagSet) {
	fs.IPVar(&c.BindAddress, "bind-address", c.BindAddress, "The IP address to listen at.")
	fs.UintVar(&c.BindPort, "bind-port", c.BindPort, "The port to listen at.")
	fs.UintVar(&c.GracefulTimeout, "graceful-timeout", c.GracefulTimeout,
		"Timeout for graceful shutdown.")
	fs.IPVar(&c.HealthBindAddress, "health-bind-address", c.HealthBindAddress, "The IP address to listen at.")
	fs.UintVar(&c.HealthBindPort, "health-bind-port", c.HealthBindPort, "The port to listen at.")
}

func (c *Config) BindFlags() {
	c.addFlags(pflag.CommandLine)
	c.Logging.BindFlags(pflag.CommandLine)

	err := c.Config.BindFlags()
	if err != nil {
		panic(err)
	}

	err = xlog.New(&xlog.Config{
		LogLevel:  viper.GetString("log-level"),
		LogOutput: viper.GetString("log-output"),
		LogWriter: viper.GetString("log-writer"),
	})
	if err != nil {
		panic(err)
	}
}
