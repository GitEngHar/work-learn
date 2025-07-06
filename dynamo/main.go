package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"log"
	"time"
	"work-learn/dynamo/infra"
)

type Item struct {
	ID         string `dynamodbav:"id"`
	CreateDate string `dynamodbav:"create_date"`
	Message    string `dynamodbav:"message"`
}

func createMockData(id, message string) Item {
	return Item{
		ID:         id,
		CreateDate: time.Now().Format(time.RFC3339),
		Message:    message,
	}
}

func main() {
	ctx := context.Background()
	cfg := infra.NewConfig().SetConfig(ctx)
	ensureTable(cfg)
	item := createMockData("1a", "hello")
	if err := putItem(cfg, item); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("➕ Putted : %s \n", item)
	dynamoValue, err := getItem(cfg, item.ID)
	if err != nil {
		log.Fatal(err)
	}
	if err := updateMessage(cfg, *dynamoValue, "Updated"); err != nil {
		log.Fatalf("Error updating message: %v", err)
	}
	fmt.Printf("✅ Updated \n")

	// 確認のために再度取得する
	newDynamoValue, err := getItem(cfg, item.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("👀 Confirm: %s \n", *newDynamoValue)
	err = deleteItem(cfg, *newDynamoValue)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("🗑️ Deleted")
}

// ensureTable テーブルが存在しない場合に作成
func ensureTable(cfg *infra.Config) {
	_, err := cfg.Cli.DescribeTable(cfg.Ctx, &dynamodb.DescribeTableInput{
		TableName: aws.String(cfg.TableName),
	})
	if err != nil {
		return // すでに存在しているので作成しない
	}
	// TODO: 存在していない場合の処理
}

// putItem アイテムをput
func putItem(cfg *infra.Config, it Item) error {
	av, err := attributevalue.MarshalMap(it)
	if err != nil {
		return err
	}
	_, err = cfg.Cli.PutItem(cfg.Ctx, &dynamodb.PutItemInput{
		TableName: aws.String(cfg.TableName),
		Item:      av,
	})
	return err
}

// 単一アイテムを取得する
func getItem(cfg *infra.Config, id string) (*Item, error) {

	// 単一検索 (PK+SKが両方必須)
	//key, err := attributevalue.MarshalMap(map[string]string{
	//	"id": id,
	//})
	//if err != nil {
	//	return nil, err
	//}
	//out, err := cfg.Cli.GetItem(cfg.Ctx, &dynamodb.GetItemInput{
	//	TableName:      aws.String(cfg.TableName),
	//	Key:            key,
	//	ConsistentRead: aws.Bool(true), // なんだこれ
	//})

	out, err := cfg.Cli.Query(cfg.Ctx, &dynamodb.QueryInput{
		TableName:              aws.String(cfg.TableName),
		KeyConditionExpression: aws.String("id = :v"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":v": &types.AttributeValueMemberS{Value: id},
		},
		ScanIndexForward: aws.Bool(false), //降順で値をソートする
	})
	var items []Item
	if err = attributevalue.UnmarshalListOfMaps(out.Items, &items); err != nil {
		return nil, err
	}
	if items == nil {
		log.Fatal("Itemが見つかりませんでした")
	}

	var item Item
	item = items[0] // 最新の一件を返す
	return &item, nil
}

func updateMessage(cfg *infra.Config, it Item, newMessage string) error {
	_, err := cfg.Cli.UpdateItem(cfg.Ctx, &dynamodb.UpdateItemInput{
		TableName: aws.String(cfg.TableName),
		Key: map[string]types.AttributeValue{
			"id":          &types.AttributeValueMemberS{Value: it.ID},
			"create_date": &types.AttributeValueMemberS{Value: it.CreateDate},
		},
		// 新しいSETする値
		UpdateExpression: aws.String("SET #M = :m"),
		ExpressionAttributeNames: map[string]string{
			"#M": "message",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":m": &types.AttributeValueMemberS{Value: newMessage},
		},
		ReturnValues: types.ReturnValueUpdatedNew,
	})
	return err
}

func deleteItem(cfg *infra.Config, item Item) error {
	_, err := cfg.Cli.DeleteItem(cfg.Ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(cfg.TableName),
		Key: map[string]types.AttributeValue{
			"id":          &types.AttributeValueMemberS{Value: item.ID},
			"create_date": &types.AttributeValueMemberS{Value: item.CreateDate},
		},
	})
	return err
}
