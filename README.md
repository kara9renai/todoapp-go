# Go言語でTodoアプリのバックエンドAPIを作成

主に自分が練習するために作成しました。

## Setup
ローカルでそのまま開発することもできますが、Dockerコンテナを利用して開発することを勧めます。
下記のコマンドを実行することで、コンテナが立ち上がります。

```
make serve
```
尚、実装していくなかでGo言語の実行環境のみを再ビルドする必要性に駆られると思われます。その場合は、下記のコマンドを実行してください。
```
make buildup
```

## part1 EntityとDTOを定義しよう
TODO
- [ ] ToddEntity定義

todo_entity.goのTodoEntityを定義してください。
構造体には
- int型のID
- string型のTitle
- string型のContent
を定義してください。

- [ ] dto構造体の定義

Todoデータを転送するにあたって利用するDTOを定義しよう

### todo_dto.goの仕様
- TodoResponseには、ID(int型)、Title(string型)、Content(string型)を定義してください。また、各々のfieldにjsonタグをつけてください。
- TodoRequestには、Title(string型)、Content(string型)で定義してください。また、各々のfieldにjsonタグをつけてください。
- TodosResponseには、Todos(TodoResponse)を定義してください。またfieldにjsonタグをつけてください。

全て実装したら、`make test-part1`を実行してください。

## part2 GetTodosを実装しよう
- [ ] Repository層でGetTodosを実装しよう（直接DBと接続する箇所）

GetTodosの仕様
- 指定されたクエリを実装すること
- データベースのカラムのNOT NULL制約に準拠すること

- [ ] Controller層でGetTodosを実装しよう

dtoの構造体を利用して、Repository層でGetしたクエリを取り出します。
取り出したクエリをMarshallIndentで返しましょう。（headerの設定も忘れずに）

- [ ] router.goでRouterの設定

全て実装したら。`make test-part2`を実行してください。

## part3 InsertTodoを実装しよう　

- [ ] Repository層でInsertTodoを実装
- [ ] Controller層にて、dto構造体に準拠しながら、PostTodoを実装
- [ ] router.goでMethodに応じて、PostTodoを実行できるようにする (MethodがPostのときに実行)

全て実装して終えたら、 `make test-part3` を実行してください。

## part4 UpdateTodoを実装しよう

- [ ] Repository層でUpdateTodoを実装
- [ ] Controller層でPostTodoを実装
- [ ] 適切なMethodでPostTodoを実行するようにRoutingを設定

全て実装し終えたら、`make test-part4` を実行してください。

## part5 DeleteTodoを実装しよう

- [ ] Repository層でDeleteTodoを実装
- [ ] Controller層でDeleteTodoを実装
- [ ] 適切なMethodでDeleteTodoを実行するようにRoutingを設定

全て実装し終えたら、`make test-part5` を実行してください。
