package keyManager

import (
	"context"
	"log"
	"os"

	"github.com/akeylesslabs/akeyless-go/v3"
)

type impl struct {
	client *akeyless.V2ApiService
	auth   *akeyless.Auth
}

func New() Service {
	client := akeyless.NewAPIClient(&akeyless.Configuration{
		Servers: []akeyless.ServerConfiguration{
			{
				URL: "https://api.akeyless.io",
			},
		},
	}).V2Api

	auth := akeyless.NewAuth()
	auth.SetAccessType("api_key")
	auth.SetAccessId(os.Getenv("ACCESS_ID"))
	auth.SetAccessKey(os.Getenv("ACCESS_KEY"))

	return &impl{
		client: client,
		auth:   auth,
	}
}

func (im *impl) GetKeys(ctx context.Context, keys ...string) (map[string]string, error) {
	token, err := im.getToken(ctx)
	if err != nil {
		log.Fatalln("can't get token:", err)
		return nil, err
	}
	body := akeyless.GetSecretValue{
		Names: keys,
		Token: &token,
	}

	out, _, err := im.client.GetSecretValue(ctx).Body(body).Execute()
	if err != nil {
		log.Fatalln("can't get secret value:", err)
		return nil, err
	}

	return out, nil
}

func (im *impl) getToken(ctx context.Context) (string, error) {
	out, _, err := im.client.Auth(ctx).Body(*im.auth).Execute()
	if err != nil {
		log.Fatalln("can't get token:", err)
		return "", err
	}

	return out.GetToken(), nil
}
