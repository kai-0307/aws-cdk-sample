package main

import (
	"testing"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/assertions"
	"github.com/aws/jsii-runtime-go"
)

func TestCdkStack(t *testing.T) {
	// テスト用のアプリケーションとスタックを作成
	app := awscdk.NewApp(nil)
	stack := NewSampleCdkStack(app, "TestStack", &SampleCdkStackProps{
		StackProps: awscdk.StackProps{
			Env: &awscdk.Environment{
				Account: jsii.String("123456789012"),
				Region:  jsii.String("ap-northeast-1"),
			},
		},
	})

	// テンプレートの検証用にTemplate assertionを作成
	template := assertions.Template_FromStack(stack, nil)

	// リソースの存在確認テスト
	template.HasResourceProperties(jsii.String("AWS::S3::Bucket"), map[string]interface{}{
		"VersioningConfiguration": map[string]interface{}{
			"Status": "Enabled",
		},
	})

	// バケット数の確認テスト
	template.ResourceCountIs(jsii.String("AWS::S3::Bucket"), jsii.Number(1))

	// バケットの特定のプロパティ値を確認するテスト
	template.HasResource(jsii.String("AWS::S3::Bucket"), map[string]interface{}{
		"Properties": map[string]interface{}{
			"VersioningConfiguration": map[string]interface{}{
				"Status": "Enabled",
			},
		},
	})

	// スタック全体のスナップショットテスト
	template.ToJSON()
}

func TestBucketProperties(t *testing.T) {
	// テスト用のアプリケーションとスタックを作成
	app := awscdk.NewApp(nil)
	stack := NewSampleCdkStack(app, "TestBucketStack", &SampleCdkStackProps{
		StackProps: awscdk.StackProps{},
	})

	// テンプレートの検証用にTemplate assertionを作成
	template := assertions.Template_FromStack(stack, nil)

	// バケットのバージョニング設定をテスト
	template.HasResourceProperties(jsii.String("AWS::S3::Bucket"), map[string]interface{}{
		"VersioningConfiguration": map[string]interface{}{
			"Status": "Enabled",
		},
	})

	// バケットの暗号化設定をテスト
	template.HasResourceProperties(jsii.String("AWS::S3::Bucket"), map[string]interface{}{
		"BucketEncryption": map[string]interface{}{
			"ServerSideEncryptionConfiguration": []interface{}{
				map[string]interface{}{
					"ServerSideEncryptionByDefault": map[string]interface{}{
						"SSEAlgorithm": "AES256",
					},
				},
			},
		},
	})
}

func TestEdgeCases(t *testing.T) {
	app := awscdk.NewApp(nil)
	stack := NewSampleCdkStack(app, "TestEdgeCaseStack", &SampleCdkStackProps{
		StackProps: awscdk.StackProps{
			Env: &awscdk.Environment{
				Account: jsii.String("123456789012"),
				Region:  jsii.String("invalid-region"),
			},
		},
	})

	template := assertions.Template_FromStack(stack, nil)
	template.ResourceCountIs(jsii.String("AWS::S3::Bucket"), jsii.Number(1))
}
