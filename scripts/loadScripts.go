package scripts

import (
	"github.com/0xBl4nk/FuzzSwarm2/scripts/ssti"
	"github.com/0xBl4nk/FuzzSwarm2/src"
)

// InitScripts initializes scripts based on the provided name.
func InitScripts(script string, cfg *src.Config) {
	switch script {
	case "ssti":
		src.LogInfo("Initializing SSTI script")
		ssti.LoadSSTIPayloads(cfg)
	default:
		src.LogError("Unknown script: %s", script)
	}
}
