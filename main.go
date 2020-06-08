// Example draw-text shows how to draw text to an xgraphics.Image type.
package main

import (
	"flag"

	log "github.com/sirupsen/logrus"

	"net/http"
	_ "net/http/pprof"

	_ "net/http/pprof"

	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/xevent"
	"github.com/shimmerglass/bar3x/bar"
)

func main() {
	cfgPath := flag.String("cfg", "config.yaml", "YAML Config file path")
	flag.Parse()

	cfg, err := getConfig(*cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	X, err := xgbutil.NewConn()
	if err != nil {
		log.Fatal(err)
	}

	xevent.ErrorHandlerSet(X, func(err xgb.Error) {
		log.Errorf("X error: %s", err)
	})

	_, err = bar.CreateBars(cfg, X)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	xevent.Main(X)
}