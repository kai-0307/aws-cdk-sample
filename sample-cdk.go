package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type SampleCdkStackProps struct {
	awscdk.StackProps
}

func NewSampleCdkStack(scope constructs.Construct, id string, props *SampleCdkStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// S3バケットの作成(例)
	awss3.NewBucket(stack, jsii.String("TestBucket"), &awss3.BucketProps{
		Versioned:         jsii.Bool(true),
		Encryption:        awss3.BucketEncryption_S3_MANAGED,
		BlockPublicAccess: awss3.BlockPublicAccess_BLOCK_ALL(),
		RemovalPolicy:     awscdk.RemovalPolicy_DESTROY,
	})

	return stack
}

func main() {
	app := awscdk.NewApp(nil)

	NewSampleCdkStack(app, "SampleCdkStack", &SampleCdkStackProps{
		awscdk.StackProps{
			Env: &awscdk.Environment{
				Account: jsii.String("123456789012"), // AWSアカウントID(テキトー)
				Region:  jsii.String("ap-northeast-1"),
			},
		},
	})

	app.Synth(nil)
}
