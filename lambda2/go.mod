module multi-lambda/lambda2

go 1.17

require (
	github.com/aws/aws-lambda-go v1.32.0
	multi-lambda/shared v0.0.0-00010101000000-000000000000
)

replace multi-lambda/shared => ../shared
