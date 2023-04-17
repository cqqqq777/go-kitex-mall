package config

type NsqConfig struct {
	Host    string `mapstructure:"host" json:"host"`
	Port    int    `mapstructure:"port" json:"port"`
	Topic   string `mapstructure:"topic" json:"topic"`
	Channel string `mapstructure:"channel" json:"channel"`
}

type MinioConfig struct {
	Endpoint        string `mapstructure:"endpoint" json:"endpoint"`
	AccessKeyID     string `mapstructure:"access_key_id" json:"access_key_id"`
	SecretAccessKey string `mapstructure:"secret_access_key" json:"secret_access_key"`
	Bucket          string `mapstructure:"bucket" json:"bucket"`
	UrlPrefix       string `mapstructure:"url_prefix" json:"url_prefix"`
}

type UploadServiceConfig struct {
	MinioInfo MinioConfig `mapstructure:"minio" json:"minio"`
	NsqInfo   NsqConfig   `mapstructure:"nsq" json:"nsq"`
}
