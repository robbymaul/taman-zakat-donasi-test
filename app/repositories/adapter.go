package repositories

import (
	restyclient "donasitamanzakattest/app/client/resty"
	"donasitamanzakattest/config"
	"fmt"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Adapter struct {
	JakartaLoc *time.Location
	HttpClient *restyclient.RestyClient
	Minio      *minio.Client
}

func NewRepositoryAdapter(cfg *config.Config) (*Adapter, error) {
	adapter := new(Adapter)

	location, err := time.LoadLocation(cfg.ServerTimeZone)
	if err != nil {
		return nil, err
	}

	adapter.JakartaLoc = location

	adapter.HttpClient = restyclient.NewRestyClient()

	//adapter.Senangpay = senangpay.NewSenangpay(cfg.SenangpayUrl, cfg.SenangpaySecretKey, cfg.SenangpayMerchantId, adapter.HttpClient)
	//adapter.Midtrans = midtrans.NewMidtrans(adapter.HttpClient, cfg)
	//adapter.Strategy = strategy.NewStrategy(adapter.Senangpay, adapter.Midtrans, nil)

	minioClient, err := minio.New(
		cfg.S3Endpoint,
		&minio.Options{
			Creds:  credentials.NewStaticV4(cfg.S3AccessKeyId, cfg.S3SecretAccessKey, ""),
			Secure: cfg.S3UseSSL,
		},
	)
	if err != nil {
		return nil, newError(fmt.Sprintf("error set up minio client:"), err.Error())
	}

	adapter.Minio = minioClient

	return adapter, nil
}
