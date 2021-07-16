# go-lambdaのデプロイ環境テスト用
main.goが処理本体。

動作はlambdaのテストにて確認する想定で、以下の内容を設定したイベントで確認可能
```
{
  "firstName": "Takapon",
  "lastName": "Test"
}
```
"Hello Takapon Test!"
