//go:build windows
// +build windows

package main

import (
	"os"
	"os/signal"
	"syscall"
)

func reloadSignal() {

	// Signal handling
	hup := make(chan os.Signal, 1)
	signal.Notify(hup, syscall.SIGHUP)
	go func() {
		for {
			select {
			case <-hup:
				logger.Debug("Signal: HUP")
				logger.Info("ReLoading config")
				if err := sc.ReloadConfig(logger, *configFile); err != nil {
					logger.Error("Reloading config skipped", "err", err)
					continue
				} else {
					monitorPING.DelTargets()
					_ = monitorPING.CheckActiveTargets()
					monitorPING.AddTargets()
					monitorMTR.DelTargets()
					_ = monitorMTR.CheckActiveTargets()
					monitorMTR.AddTargets()
					monitorTCP.DelTargets()
					_ = monitorTCP.CheckActiveTargets()
					monitorTCP.AddTargets()
					monitorHTTPGet.DelTargets()
					monitorHTTPGet.AddTargets()
				}
			}
		}
	}()
}
