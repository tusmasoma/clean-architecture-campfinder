## Usecaseレイヤ

usecaseレイヤは、adapter(interface)レイヤから情報を受け取り、entityレイヤのオブジェクトを操作して任意のビジネスロジックを実行する責務を持ちます。レイヤードアーキテクチャのApplication(usercase)層に相当します。

## インプットポートとアウトプットポート(とDB用のポート)
クリーンアーキテクチャのusecaseレイヤはインプットポートとアウトプットポートという仕組みを持っています。インプットポートとアウトプットポートはサーバーの入力と出力に対応するポートです。ポートの実体は、インターフェースとなります。

<img src="https://github.com/tusmasoma/clean-architecture-campfinder/assets/104899572/17cf917a-2887-4337-b069-1adbf3a3ed70" width="400px" height="400px">

- InterfaceAdapterレイヤの入力(controller)に関する実装は、InputPortに依存し、そのインターフェースの実装は、Usecaseレイヤ内のユースケースで行う
- InterfaceAdapterレイヤの出力(presenter)はOutputPortを実装し、Usecaseレイヤ内のユースケースはOutputPortに依存します
