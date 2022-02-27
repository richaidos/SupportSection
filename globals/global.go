package globals

import "github.com/richaidos/SupportSection/configuration"

var ISettings *configuration.InitSettings

func Init() {
	ISettings = configuration.GetInitSettings()
}
