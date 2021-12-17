module github.com/teralytics/prometheus-ecs-discovery

go 1.17

require (
	github.com/aws/aws-sdk-go-v2/config v1.1.4
	github.com/aws/aws-sdk-go-v2/credentials v1.1.4
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.3.0
	github.com/aws/aws-sdk-go-v2/service/ecs v1.2.1
	github.com/aws/aws-sdk-go-v2/service/sts v1.2.1
	github.com/aws/smithy-go v1.3.0
	github.com/go-yaml/yaml v2.1.0+incompatible
)

require (
	github.com/aws/aws-sdk-go-v2 v1.3.1 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.0.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.0.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.1.4 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
)
