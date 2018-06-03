package log

import (
	"testing"

	log "github.com/inconshreveable/log15"
)

func TestLog15(t *testing.T) {
	srvlog := log.New("module", "app/server")
	srvlog.SetHandler(log.MultiHandler(
		log.LvlFilterHandler(log.LvlDebug, log.StdoutHandler),
		log.LvlFilterHandler(log.LvlError, log.Must.FileHandler("errors.log", log.LogfmtFormat())),
	))
	srvlog.Debug("logtest", "level", "debug")
	srvlog.Info("logtest", "level", "info")
	srvlog.Warn("logtest", "level", "warn")
	srvlog.Error("logtest", "level", "error")
	srvlog.Crit("logtest", "level", "critical")
}
