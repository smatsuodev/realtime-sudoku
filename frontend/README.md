# ディレクトリ構成

クリーンアーキテクチャに基づいています（[参考](https://zenn.dev/koudai/articles/8ee15d10008f55)）

```
lib/
  ├── core/
  │   ├── exception/
  │   ├── extension/
  │   └── util/
  ├── domain/
  │   ├── entity/
  │   ├── value/
  │   ├── repository/
  │   ├── usecase/
  │   └── factory/
  ├── application/
  │   └── usecase/
  ├── infrastructure/
  │   ├── datasource/
  │   ├── repository/
  │   ├── model/
  │   └── factory/
  ├── presentation/
  │   ├── component/
  │   ├── page/
  │   ├── notifier/
  │   ├── state/
  │   └── navigation/
  └── main.dart
```
