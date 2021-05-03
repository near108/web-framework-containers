# ウェブフレームワークの技術調査
ウェブフレームワークにおける技術調査や、バージョン間の差異を検証するためのリポジトリ

## 対象フレームワーク一覧

| 言語 | FW | Ver. | コンテナ名 | ポート番号 | 作成日 | 更新日 |
| --- | --- | --- | --- | --- | --- | --- |
| Python | Django | 3.2 | django3.2 | 8081 | 2021/5/2 | 2021/5/3 |

## 使い方

### 環境構築

#### 1. Dockerをインストール
[Docker公式サイト](https://www.docker.com/)

#### 2. リポジトリをクローン

```
git clone https://github.com/jusplat/web-frameworks.git
```

### プロジェクトの起動

#### 1. プロジェクトルートへ移動する

```
cd web-frameworks
```

#### 2. 起動したいフレームワークのコンテナを起動する

docker-compose up -d {コンテナ名}

例) Djangoのバージョン3.2

```
docker-compose up -d django3.2
```

#### 3. 起動したコンテナに接続する

docker exec -it {コンテナ名} bash

例）Djangoのバージョン3.2

```
docker exec -it django3.2 bash
```

#### 4. フレームワークのプロジェクトルートへ移動する

```
cd container
```

#### 5. サーバ起動スクリプトを実行する

```
sh runserver.sh
```

#### 6. ブラウザからアクセスする
`http://127.0.0.1:{ポート番号}`

例）Djangoのバージョン3.2

[http://127.0.0.1:8081]()

### プロジェクトを停止する

#### 1. コンテナの接続を終了する
```
exit
```

#### 2. コンテナを停止する
```
docker-compose down
```
