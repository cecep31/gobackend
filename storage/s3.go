package storage

store := s3.New(s3.Config{
	Bucket:   os.Getenv("BUCKET"),
	Endpoint: os.Getenv("S3_ENDPOINT"),
	Region:   os.Getenv("S3_REGION"),
	Reset:    false,
	Credentials: s3.Credentials{
		AccessKey:       os.Getenv("S3_ACCESS_KEY"),
		SecretAccessKey: os.Getenv("S3_SECRET_KEY"),
	},
})