https://docs.aws.amazon.com/ja_jp/code-library/latest/ug/go_2_dynamodb_code_examples.html

`AWS SDK for Go V2 ` を使用して、実装する

### テーブル

想定としてはSession IDを日付と共に格納し、SessionIDごとにmessageを格納しておく

テーブル名 : `session_info`
プライマリキー: `string: id`
ソートキー: `string: create_date`
他の値: `string: message`


### 利用するライブラリ

```
"github.com/aws/aws-sdk-go-v2/aws"
"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
"github.com/aws/aws-sdk-go-v2/service/dynamodb"
"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
```


### 返り値がバインドされる構造体の宣言
jsonのように `dynamodbav` で値を宣言する

```
type Item struct {
	ID         string `dynamodbav:"id"`
	CreateDate string `dynamodbav:"create_date"`
	Message    string `dynamodbav:"message"`
}
```



### *dynamodb.Client の操作
実行時にはいずれも TableNameでテーブル名を指定する
- Query
    - キーを指定して該当するすべてのレコードを取得する
        - 値は複合キーでなくてもいい
    - ScanIndexForwardを指定するとSortキーの内容で値をソートできる
- GetItem
    - 複合キー (主キーとソートキーを合わせたキー)を用いて単一のレコードを取得する
- PutItem
    - 代入するプライマリキー、ソートキー、他の値をセットしてPutする
- UpdateItem
    - 複合キーを指定して更新対象を決める
    - UpdateExpression
        - `SET #M = :m` のようにクエリをかく
    - ExpressionAttributeNames
        - クエリのキーを代入する
    - ExpressionAttributeNames
    - クエリの値を代入する
- DeleteItem
    - 複合キーを指定して値を削除する


### 単一キーと複合キーの決め方

- 単一キー
    - 設計とクエリがシンプル
    - 同一PKの複数アイテムを持てないので履歴管理に弱い
- 複合キー
    - 時系列検索や履歴保持が可能
    - キー設計とクエリがやや複雑

### パーティションキーとソートキーの選定
- ソートキー
    - ソートは並び替えられるような値、日付やバージョン
- パーティションキー
    - 一意な値


### Keyを指定するとき

複合キーの指定
```
Key: map[string]types.AttributeValue{
			"id":          &types.AttributeValueMemberS{Value: item.ID},
			"create_date": &types.AttributeValueMemberS{Value: item.CreateDate},
		}
```

### ローカルで検証する場合


```shell
docker run --rm -p 8000:8000 amazon/dynamodb-local
```

