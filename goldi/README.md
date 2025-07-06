## 何のためのOSS
構造体やサービス同士の“つながり”をコードの外（コンテナ）で定義し、必要なときにまとめて組み立てる

## 主な機能
1. サービス登録
2. パラメータ注入
3. 依存解決 & 検証
4. 取得
5. グルーピング


## 基本実装

- サービス登録
  - `registry := goldi.NewTypeRegistry()`でオブジェクトを生成
- パラメータを注入
  - `registry.RegisterType("key", FacterOrInstance, args)`でサービスやパラメータを登録 
- コンテナ生成
  - `container := goldi.NewContainer(registry, paramsMap)` 
- 依存グラフをチェック
  - `validation.NewContainerValidatior().Validate(container)`
- サービスを取得して使う
  - `svc := container.MustGet("key").(*MyType)` 


### コンテナの登録

- パターン1
```go
// loggerというkeyで&SimpleLoggerをレジストリに設定

container.RegisterType("logger", &SimpleLogger{})
```

- パターン2
```go
// greeterというkeyでapp_nameとloggerを注入してサービスを生成する

container.RegisterType("greeter", NewGreeter, "%app_name%", "@logger")
```

### コンテナの検証

```go
import "github.com/fgrosse/goldi/validation"
validation.NewContainerValidator()
if err := validator.Validate(container); err != nil {
panic(err)
}
```

### コンテナの利用

```go
greeter := container.MustGet("greeter").(*Greeter)
greeter.Greet("World")
```


### まとめ

- コンテナを登録する際に値を代入する。値はkeyとvalueのmap形式
- `%key名%`で値を利用する。 `@service_name`でServiceを注入する
- MustGet("key名")でサービスを取得して、返り値のかたは .(*Type)で指定する 