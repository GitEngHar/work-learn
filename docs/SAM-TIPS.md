## SAMって何??

IaCをしようしてサーバレスアプリケーション構築するためのOSSフレームワークです!

### SAMの仕様

- AWS CloudFormation上に構築される
- AWS CloudFormationの拡張で、CLoudFormationの構文の中にAWS SAMを書く
- 抽象的で簡潔な文になっている
- SAMをCloudFormationの構文に変換して、デプロイしている

### 何が嬉しいの??

https://docs.aws.amazon.com/ja_jp/serverless-application-model/latest/developerguide/what-is-sam.html#what-is-sam-feature

### YAMLの構造

```yaml
Transform: AWS::Serverless-2016-10-31

Globals:
# API に共通するプロパティを定義
# API に共通するプロパティを定義
Description:
# 説明テンプレート CloudFormationのDescription

Metadata:
# テンプレート追加情報 CloudFormationのMeatadata

Parameters:
# 実行時にテンプレートに渡す内容
# ResourcesとOutputsのセクションからパラメータを参照できる

Mappings:
# キーと値のマッピング
# ResourcesとOutputsのセクションで組み込み関数を使用するとキーを対応する値と一致できる

Conditions:
# 実稼働ようかテスト稼働八日で依存するリソースを条件付きで作成できたりする

Resources:
# SAMとCloudFormationのリソースを含めることができる
# ResourcesとOutputsのセクションを参照できる
  HelloWorldFunction:
    Type: AWS::Serverless::Function
Outputs:
# スタックのプロパティを表示するたびに返される値
# AWS CloudFormationのテンプレートセクションOutputsに対応

```