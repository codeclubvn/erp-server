package infrastructure

import (
	"erp/config"
	libSendinblue "github.com/sendinblue/APIv3-go-library/lib"
)

type Sendinblue struct {
	*libSendinblue.APIClient
}

func NewSendinblue(config *config.Config) *Sendinblue {
	return &Sendinblue{
		initConfiguration(config),
	}
}

func initConfiguration(config *config.Config) *libSendinblue.APIClient {
	cfg := libSendinblue.NewConfiguration()
	cfg.AddDefaultHeader("api-key", config.Sendinblue.ApiKey)
	cfg.AddDefaultHeader("partner-key", config.Sendinblue.ApiKey)
	sib := libSendinblue.NewAPIClient(cfg)
	return sib
}
