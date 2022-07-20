# go-lambdaのデプロイ環境テスト用
goで書いたスクリプトをGithubActionsを使ってlambdaへデプロイできることを確認するための検証用リポジトリ。
もともとprivateで作成したのでリポジトリ名にprivate入っているけど、秘匿情報はsecretsに入れているため問題ないはず。

main.goが処理本体。

動作はlambdaのテストにて確認する想定で、以下の内容を設定したイベントで確認可能
```
{
  "firstName": "Takapon",
  "lastName": "Test"
}
```
"Hello Takapon Test!"
