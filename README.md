# マインクラフトのログファイルからユーザごと稼働時間を調べる
[![Minecraft PlayTime](https://github.com/piyoraik/minecraft-log/actions/workflows/go.yml/badge.svg)](https://github.com/piyoraik/minecraft-log/actions/workflows/go.yml)  
`logs/` ディレクトリにlogファイルを入れることで、ユーザごとのマインクラフトをしている時間を調べることができる。

## 使い方

1. 下記を実行
```console
cp .env.sample .env
mkdir logs
```
2. 作成された`.env` の、`USER` にマインクラフトのユーザ名、 `USER_NAME` に表示用の名前を入れる。
3. 作成された `logs` ディレクトリにログファイルを入れる
4. 下記を実行で表示される
```console
go run main.go
```

## TODO
現状５人の固定なので汎用性のあるような形にしたい。

envで一人一人指定するのは面倒なのでいい感じにしたい。

logsフォルダに移動するのが面倒＆元がzipファイルなので展開までできると便利そう
