## Entityレイヤ

entityレイヤは、ドメインロジックを実装する責務を持ちます。レイヤードアーキテクチャのドメインレイヤに相当します。そのため、技術的な実装を持つことはありません。

ここでいう技術的な実装とは、「DBにMySQLを使って〜」や「ORMを使って〜」などです。これをentityレイヤに実装しないことで、entityレイヤが特定の技術に依存しないようになります。

```go
package entity

type User struct {
	ID       uuid.UUID
	Name     string
	Email    string
	Password string
}
```

以下のように、バリデーションも実装できる。

```go
package entity

type User struct {
	ID       uuid.UUID
	Name     string
	Email    string
}

func NewUser(name string, email string) (*User, error) {
    return newUser(
        uuid.New(),
        name,
        email,
    )
}

func newUser(id uuid.UUID, name string, email string) (*User, error) {
    // 名前のバリデーション
    if utf8.RuneCountInString(name) < nameLengthMin || utf8.RuneCountInString(name) > nameLengthMax {
        return nil, errors.New("名前の値が不正です。")
    }

    // メールアドレスのバリデーション
    if _, err := mail.ParseAddress(email); err != nil {
        return nil, errors.New("メールアドレスの値が不正です。")
    }

    return &User{
        ID: id,
        Name: name,
        Email: email,
    }, nil
}

const (
	nameLengthMax = 255
	nameLengthMin = 1
)
```

また、EntityレイヤにRepositoryを置くこともある。

```go
package entity

type User struct {
	ID       uuid.UUID
	Name     string
	Email    string
	Password string
}

type UserRepository interface {
	CheckIfUserExists(ctx context.Context, email string) (bool, error)
	GetUserByID(ctx context.Context, id string) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
}
```
