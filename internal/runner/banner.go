// Package runner executes the enumeration process.
package runner

import (
	"fmt"

	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/nuclei/v3/pkg/catalog/config"
	pdcpauth "github.com/projectdiscovery/utils/auth/pdcp"
	updateutils "github.com/projectdiscovery/utils/update"
)

var banner = fmt.Sprintf(`
             _           _       _     _
 __   ___   | |_ __  ___(_) __ _| |__ | |_
 \ \ / / | | | '_ \/ __| |/ _` + "`" + ` | '_ \| __|
  \ V /| |_| | | | \__ \ | (_| | | | | |_
   \_/  \__,_|_| |_|___/_|\__, |_| |_|\__|
                           |___/            %s
`, config.Version)

// showBanner is used to show the banner to the user
func showBanner() {
	gologger.Print().Msgf("%s\n", banner)
	gologger.Print().Msgf("\t\tprojectdiscovery.io\n\n")
}

// NucleiToolUpdateCallback updates nuclei binary/tool to latest version
func NucleiToolUpdateCallback() {
	showBanner()
	updateutils.GetUpdateToolCallback(config.BinaryName, config.Version)()
}

// AuthWithPDCP is used to authenticate with PDCP
func AuthWithPDCP() {
	showBanner()
	pdcpauth.CheckNValidateCredentials(config.BinaryName)
}
