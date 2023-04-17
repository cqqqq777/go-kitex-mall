package initialize

import (
	"context"
	"github.com/cqqqq777/go-kitex-mall/cmd/api/pkg/upload/config"
	"github.com/cqqqq777/go-kitex-mall/shared/log"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func newMinio() *minio.Client {
	mc := config.GlobalServiceConfig.MinioInfo
	minioClient, err := minio.New(mc.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(mc.AccessKeyID, mc.SecretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Zlogger.Fatalf("init minio failed err:%s", err.Error())
	}

	exists, err := minioClient.BucketExists(context.Background(), mc.Bucket)
	if err != nil {
		log.Zlogger.Fatal(err)
	}
	if !exists {
		err = minioClient.MakeBucket(context.Background(), mc.Bucket, minio.MakeBucketOptions{Region: "cn-north-1"})
		if err != nil {
			log.Zlogger.Fatal(err)
		}
	}
	policy := `{"Version": "2012-10-17","Statement": [{"Action": ["s3:GetObject"],"Effect": "Allow","Principal": {"AWS": ["*"]},"Resource": ["arn:aws:s3:::` + mc.Bucket + `/*"],"Sid": ""}]}`
	err = minioClient.SetBucketPolicy(context.Background(), mc.Bucket, policy)
	if err != nil {
		log.Zlogger.Fatal(err)
	}

	return minioClient
}
