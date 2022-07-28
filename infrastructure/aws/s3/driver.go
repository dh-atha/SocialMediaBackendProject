package s3

import (
	"socialmediabackendproject/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func ConnectAws(cfg *config.AppConfig) *session.Session {
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(cfg.AWS_REGION),
			Credentials: credentials.NewStaticCredentials(
				cfg.AWS_ACCESS_KEY_ID,
				cfg.AWS_SECRET_ACCESS_KEY,
				"", // a token will be created when the session it's used.
			),
		})
	if err != nil {
		panic(err)
	}
	return sess
}
