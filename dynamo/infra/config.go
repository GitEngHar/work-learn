package infra

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"log"
	"os"
)

type Config struct {
	TableName string
	Ctx       context.Context
	Cli       *dynamodb.Client
}

func NewConfig() *Config {
	tableName := os.Getenv("DYNAMODB_TABLE_NAME")
	if tableName == "" {
		log.Fatal("DYNAMODB_TABLE_NAME environment variable not set")
	}
	return &Config{TableName: tableName}
}

func (c *Config) SetConfig(ctx context.Context) *Config {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}
	cli := dynamodb.NewFromConfig(cfg)
	c.Cli = cli
	c.Ctx = ctx
	return c
}
