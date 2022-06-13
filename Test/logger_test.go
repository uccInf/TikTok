package main

import (
	log "TikTok/logger"
	"testing"
)

func TestLogger(*testing.T) {
	log.Info("aaaaa")
	log.Warn("bbb")
	log.Error("accc")
	log.Warn("warning")
}
