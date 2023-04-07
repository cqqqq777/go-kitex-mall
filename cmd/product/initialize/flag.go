package initialize

import (
	"flag"

	"github.com/cqqqq777/go-kitex-mall/shared/consts"
	"github.com/cqqqq777/go-kitex-mall/shared/log"
	"github.com/cqqqq777/go-kitex-mall/shared/tools"
)

// InitFlag to init flag
func InitFlag() (string, int) {
	IP := flag.String(consts.IPFlagName, consts.IPFlagValue, consts.IPFlagUsage)
	Port := flag.Int(consts.PortFlagName, 0, consts.PortFlagUsage)
	// Parsing flags and if Port is 0 , then will automatically get an empty Port.
	flag.Parse()
	if *Port == 0 {
		*Port, _ = tools.GetFreePort()
	}
	log.Zlogger.Info("ip: ", *IP)
	log.Zlogger.Info("port: ", *Port)
	return *IP, *Port
}
