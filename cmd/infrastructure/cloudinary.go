package infrastructure

import (
	"context"
	"github.com/cloudinary/cloudinary-go/v2"
)

type Cloudinary struct {
	*cloudinary.Cloudinary
}

func NewCloudinary() *Cloudinary {
	cld, _ := credentials()
	return &Cloudinary{
		cld,
	}
}

func credentials() (*cloudinary.Cloudinary, context.Context) {
	// Add your Cloudinary credentials, set configuration parameter
	// Secure=true to return "https" URLs, and create a context
	//===================
	cld, err := cloudinary.New()
	if err != nil {
		panic(err)
	}
	// CLOUDINARY_URL=cloudinary://API-Key:API-Secret@Cloud-name
	// CLOUDINARY_URL=cloudinary://oZ47iHrgrFQq4fe7ksKKlo7tg4A:991793784142871@dsr2xnaj7
	cld.Config.URL.Secure = true
	ctx := context.Background()
	return cld, ctx
}
