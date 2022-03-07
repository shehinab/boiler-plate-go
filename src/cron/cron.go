package cron

import (
	"gitlab.com/abhishek.k8/crud/src/config"
)

//Cron struct
type Cron struct{}

//Init - init Cron job
func (ct *Cron) Init() {
	if config.AppConfig.Environment == "Development" {
		return
	}

}
