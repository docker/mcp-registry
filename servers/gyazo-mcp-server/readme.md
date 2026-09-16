# Gyazo MCP Server (Unofficial)

Gyazo 向けの非公式 MCP (Model Context Protocol) サーバーです。

既存の Gyazo エントリとは別に、より広い機能範囲をカバーする実装として登録しています。
画像の一覧取得・検索・個別取得・最新取得・アップロード・削除・oEmbed メタデータ取得に加えて、
`gyazo-mcp:///image_id` 形式のリソースアクセス、MCP OAuth broker、streamable HTTP / stdio の両 transport を提供します。

## 機能

- **8 つのツール**: 画像の一覧取得・検索・個別取得・最新取得・アップロード・削除・oEmbed メタデータ取得・ユーザー情報
- **リソース**: `gyazo-mcp:///image_id` 形式で画像本体とメタデータにアクセス
- **認証**: MCP 標準 OAuth に準拠した認証 broker を内蔵。Personal Access Token (PAT) によるシンプルな認証も対応
- **トランスポート**: streamable HTTP と stdio の両方に対応

## クイックスタート

### Docker

```bash
docker run -d \
  -p 127.0.0.1:18449:18449 \
  -e GYAZO_MCP_PERSONAL_ACCESS_TOKEN=your-token \
  ishiimasaki646/gyazo-mcp-server
```

MCP endpoint: `http://127.0.0.1:18449/mcp`

### cargo install

```bash
cargo install gyazo-mcp-server
gyazo-mcp-server config init   # 対話形式の初期設定
gyazo-mcp-server env init      # クレデンシャルの設定
gyazo-mcp-server                # HTTP サーバー起動
```

## 認証

以下のいずれかで認証します。

1. **MCP OAuth**: MCP login 対応クライアントから自動的に OAuth フローが実行されます
2. **Personal Access Token**: `GYAZO_MCP_PERSONAL_ACCESS_TOKEN` 環境変数に PAT を設定します

PAT は [Gyazo 開発者ページ](https://gyazo.com/api) で発行できます。

## ツール一覧

| ツール | 説明 |
|--------|------|
| `gyazo_whoami` | 認証ユーザーの情報を取得 |
| `gyazo_search` | キャプチャを全文検索 (Pro プラン) |
| `gyazo_list_images` | 画像一覧を取得 |
| `gyazo_get_image` | 画像 ID または URL で 1 件取得 |
| `gyazo_delete_image` | 画像を削除 |
| `gyazo_get_latest_image` | 最新の画像を取得 |
| `gyazo_upload_image` | base64 画像をアップロード |
| `gyazo_get_oembed_metadata` | oEmbed メタデータを取得 |

## 詳細

- リポジトリ: https://github.com/ishii-masaki-646/gyazo-mcp-server
- crates.io: https://crates.io/crates/gyazo-mcp-server
- ライセンス: MIT OR Apache-2.0
