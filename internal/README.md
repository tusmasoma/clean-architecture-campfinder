## internalディレクトリとは

まず、注意としてinternalディレクトリはクリーンアーキテクチャ特有のディレクトリではなくgo言語にてよく使われるディレクトリ構成の一つです。

internalディレクトリとは、パッケージのインポートスコープを制限するために使用されます。internal ディレクトリ内に配置されたパッケージは、そのディレクトリをルートとするパッケージツリー内からのみインポート可能です。つまり、internal ディレクトリを使用することで、パッケージの使用を特定のアプリケーションやライブラリ内に限定し、外部からのアクセスを防ぐことができます。

以下を例に説明します。

```
a
├── b
│   ├── d
│   └── internal
│       └── e
├── c
└── go.mod
```
上記の場合、internalディレクトリのルートはbディレクトリなので、bパッケージツリー内からのみinternalパッケージ以下のパッケージを参照できます。
従って、パッケージbとパッケージd以下およびinternalパッケージ以下のパッケージからの参照に限定され、パッケージaやパッケージcからは参照できません。

## internal ディレクトリに配置するコードやファイル
internal ディレクトリに配置するコードやファイルは、主にそのアプリケーションやライブラリ内でのみ使用されるものであり、外部からのアクセスを制限したいものです。

以下に具体的な例を挙げます：

1. `プライベートなユーティリティ関数`: 共通のユーティリティ関数やヘルパー関数であっても、プロジェクト外で使用されることを意図していない場合は、internal ディレクトリに配置します。
2. `内部データモデル`: アプリケーションの内部ロジックで使用されるデータ構造やモデルで、外部のAPIとして公開されないもの。
3. `ビジネスロジック`: アプリケーション固有のビジネスロジックやドメインロジックを含むコード。これらはプロジェクトの内部実装の詳細であり、外部から直接利用されることはありません。
4. `非公開のAPIクライアント`: 外部サービスとの通信を担うクライアントであっても、そのインターフェースがプロジェクト内部でのみ使用される場合。
5. `設定と初期化のコード`: アプリケーションの起動時に必要な設定の読み込みや初期化処理を行うコード。これらはしばしば内部的な詳細を含むため、外部からアクセスする必要はありません。
6. `テスト用の補助コード`: ユニットテストや統合テストで使用されるモックオブジェクトやテストデータなど、テスト専用のコード。
