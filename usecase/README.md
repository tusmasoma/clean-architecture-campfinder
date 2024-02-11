## Usecaseレイヤ

usecaseレイヤは、adapter(interface)レイヤから情報を受け取り、entityレイヤのオブジェクトを操作して任意のビジネスロジックを実行する責務を持ちます。レイヤードアーキテクチャのApplication(usercase)層に相当します。

## インプットポートとアウトプットポート(とDB用のポート)
クリーンアーキテクチャのusecaseレイヤはインプットポートとアウトプットポートという仕組みを持っています。インプットポートとアウトプットポートはサーバーの入力と出力に対応するポートです。ポートの実体は、インターフェースとなります。

- InterfaceAdapterレイヤの入力(controller)に関する実装は、InputPortに依存し、そのインターフェースの実装は、Usecaseレイヤ内のユースケースで行う
- InterfaceAdapterレイヤの出力(presenter)はOutputPortを実装し、Usecaseレイヤ内のユースケースはOutputPortに依存します
