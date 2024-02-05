package lib

import (
	"context"
	"erp/cmd/infrastructure"
	"erp/config"
	"fmt"
	libSendinblue "github.com/sendinblue/APIv3-go-library/lib"
)

type sendinblue struct {
	config    *config.Config
	apiClient *infrastructure.Sendinblue
}

type Sendinblue interface {
	SendMail(ctx context.Context, toEmail, toName string, TemplateId int64, params map[string]interface{})
}

func NewSendinblue(config *config.Config, apiClient *infrastructure.Sendinblue) Sendinblue {
	return &sendinblue{
		config:    config,
		apiClient: apiClient,
	}
}

func (s *sendinblue) SendMail(ctx context.Context, toEmail, toName string, TemplateId int64, mapParams map[string]interface{}) {
	var params interface{} = mapParams
	body := libSendinblue.SendSmtpEmail{
		Sender: &libSendinblue.SendSmtpEmailSender{
			Email: s.config.Sendinblue.DefaultFromEmail,
			Name:  s.config.Sendinblue.DefaultFromName,
		},
		To: []libSendinblue.SendSmtpEmailTo{
			{
				Email: toEmail,
				Name:  toName,
			},
		},
		TemplateId: TemplateId,
		Params:     &params,
	}
	obj, rest, err := s.apiClient.TransactionalEmailsApi.SendTransacEmail(ctx, body)
	if err != nil {
		fmt.Println("error send email", err)
	}
	fmt.Println("response send email", obj, rest)
}
