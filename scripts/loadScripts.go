package scripts

import (
	"github.com/0xBl4nk/FuzzSwarm2/scripts/ssti"
	"github.com/0xBl4nk/FuzzSwarm2/scripts/sqli"
	"github.com/0xBl4nk/FuzzSwarm2/scripts/xss"
	"github.com/0xBl4nk/FuzzSwarm2/src"
)

// InitScripts initializes scripts based on the provided name.
func InitScripts(script string, cfg *src.Config) {
	switch script {
	case "ssti":
		src.LogInfo("Initializing SSTI script")
		ssti.LoadSSTIPayloads(cfg)
	case "sqli", "sql":
		src.LogInfo("Initializing SQL injection script")
		sqli.LoadSQLiPayloads(cfg)
	case "xss":
		src.LogInfo("Initializing XSS script")
		xss.LoadXSSPayloads(cfg)
	default:
		src.LogError("Unknown script: %s. Available scripts: ssti, sqli, xss", script)
	}
}
