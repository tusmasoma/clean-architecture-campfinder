## InterfaceAdapterレイヤ
InterfaceAdapterレイヤは、アダプター(変換器)の集合体です。これは、usecaseとentityに最も便利な形式から、データベースやWebのような外部エージェントに最も便利な形式にデータを変換します。逆も然り。

つまり、このレイヤは、外部エージェントからアプリケーションへの「入力」、アプリケーションから外部エージェントへの「出力」の出入り口になる層です。

```
├── interfaces
    └── controller
    |   └── user.go
    └── gateway
    |   └── user.go
    └── presenter
        └── user.go
```

上記のように今回、InterfaceAdapterレイヤはcontroller、gateway、presenterという3つのパッケージを定義しています。それぞれどうのような責務を持つかは公式には明記されていません。

ここでは、それぞれのパッケージの責務を以下で明記します。

- controller
  - 入力に対するアダプターです。おもにWebを想定しているのでcontrollerという命名にしているのだと考えれます。
- presenter
  - 出力に対するアダプターです。html、json、csvなどの出力形式に対する実装を持ちます
- gateway
  - DBなどの永続化の技術を扱うアダプター

しかし、controller、gateway、presenterではアダプターは足りません。例えば、メール送信、PUSH通知などWebアプリケーションで扱う技術などがあります。おそらくcontroller、gateway、presenterは単なる例でしょう。そのため、実際の開発では自分たちに必要ようなアダプターを用意しましょう。


