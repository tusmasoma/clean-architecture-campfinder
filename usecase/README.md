## Usecaseレイヤ

usecaseレイヤは、adapter(interface)レイヤから情報を受け取り、entityレイヤのオブジェクトを操作して任意のビジネスロジックを実行する責務を持ちます。レイヤードアーキテクチャのApplication(usercase)層に相当します。

## インプットポートとアウトプットポート(とDB用のポート)
クリーンアーキテクチャのusecaseレイヤはインプットポートとアウトプットポートという仕組みを持っています。インプットポートとアウトプットポートはサーバーの入力と出力に対応するポートです。ポートの実体は、インターフェースとなります。

<img src="https://github.com/tusmasoma/clean-architecture-campfinder/assets/104899572/17cf917a-2887-4337-b069-1adbf3a3ed70" width="400px" height="400px">

- InterfaceAdapterレイヤの入力(controller)に関する実装は、InputPortに依存し、そのインターフェースの実装は、Usecaseレイヤ内のユースケースで行う
- InterfaceAdapterレイヤの出力(presenter)はOutputPortを実装し、Usecaseレイヤ内のユースケースはOutputPortに依存します

### InputPort
InputPortは、サーバの入力に関するポートです。そのため、InterfaceAdapterレイヤの入力(controller)に関する実装はInputPortに依存します。InputPortを実装するのはUsecaseレイヤ内のユースケースです。

InputPortを導入しなくてもレイヤ間の依存関係がおかしくなることはないので、クリーンアーキテクチャにInputPortを実装することは必須ではありません。

しかし、InputPortを実装するメリットがあります。

1. レイヤ間を疎結合にできる
   - レイヤ間を疎結合にするメリットとして、テストのしやすさがあります。controllerをテストする際に依存するユースケースをテスト用のモックにすることでユニットテストが容易になります。
  
2. レイヤ間の依存を最小限にできる
   - インプットポートは具体的な実装と1対1で対応するとは限りません。そのため1つの実装に対して複数のインプットポートを利用することができます。

   例)
   ```go
   package usecase

   type UserInputPort interface {
       GetUserByID(id string) error
   }

   type ItemInputPort interface {
       GetItemByID(id string) error
   }
   ```

   Usecaseレイヤが複数のインターフェースを持つことで、controllerは以下のように必要最低限のインターフェースに依存することが可能になります。

   ```go
   package controller

   // controllerはusecase.UserInputPortにしか依存しない
   type User struct {
       input usecase.UserInputPort
   }

   func NewUser(input usecase.UserInputPort) *User {
       return &User{
           input: input,
       }
   }

   func (u *User) GetUserByID(w http.ResponseWriter, r *http.Request) {
       userID := getUserIDFrom(r)
       u.input.GetUserByID(userID)
   }
   ```

### OutputPort
OutputPortはサーバの出力に関するポートです。OutputPortはInterfaceAdapterレイヤのpresenterが実装します。Usecaseレイヤ内のユースケースはOutputPortに依存します。

```go
package usecase

type UserOutputPort iterface {
    Render(user *entity.User)
}
```

```go
package presenter

type User struct {
    w http.ResponseWriter
}

func NewUser(w http.ResponseWriter) usecase.UserOutputPort {
    return &User{
        w: w,
    }
}

func (u *User) Render(user *entity.User) {
    fmt.Fprint(u.w, user.Name)  // HTTPのレスポンスでentity.User.Nameを出力
}
```




