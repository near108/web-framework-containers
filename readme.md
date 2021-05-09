# ウェブフレームワークの技術調査
ウェブフレームワークにおける技術調査や、バージョン間の差異を検証するためのリポジトリ

## 対象フレームワーク一覧

| 言語 | FW | コンテナ名 | ポート番号 | 作成日 | 更新日 |
| --- | --- | --- | --- | --- | --- |
| Python | Django | django | 8001 | 2021/5/2 | 2021/5/3 |

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

#### 2. 設定ファイルを編集する

#### 3. コンテナを起動する

下記スクリプトを実行する

Windowsの場合

```
dockerstart.exe
```

linuxの場合

```
dockerstart.sh
```

#### ※起動したコンテナに接続する

docker-compose exec {コンテナ名} bash

例）Django

```
docker-compose exec django bash
```

#### 5. ブラウザからアクセスする
`http://127.0.0.1:{ポート番号}`

例）Django

[http://127.0.0.1:8001]()

### プロジェクトを停止する

#### 1. コンテナを停止する

下記スクリプトを実行する

Windowsの場合

```
dockerstop.exe
```

linuxの場合

```
dockerstop.sh
```

