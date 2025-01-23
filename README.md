# プロジェクト構造

```
aws-cdk-sample/
├── README.md
├── cdk.json
├── go.mod
├── go.sum
├── sample-cdk.go
└── sample-cdk-test.go
```

# Go で AWS CDK 環境構築手順

## 前提条件のインストール

```bash
# Node.jsとnpmのインストール (CDKはNode.js環境が必要)
# macOSの場合
brew install node

# AWS CDK CLIのインストール
npm install -g aws-cdk

# Goのインストール
brew install go

# AWS CLIのインストール
brew install awscli

# AWS認証情報の設定
aws configure
```

## プロジェクトの初期化

```bash
# プロジェクトディレクトリの作成
mkdir aws-cdk-sample
cd aws-cdk-sample

# CDKプロジェクトの初期化 (Go言語を指定)
cdk init app --language go

# Go依存関係の初期化
go mod init aws-cdk-sample
go mod tidy
```

## 必要な AWS CDK Go モジュールのインストール

```bash
# 基本的なAWS CDKモジュール
go get github.com/aws/aws-cdk-go/awscdk/v2
go get github.com/aws/constructs-go/constructs/v10
go get github.com/aws/jsii-runtime-go
```

## テストの実行方法

```
go test -v
```
