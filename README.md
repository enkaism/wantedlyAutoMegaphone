# wantedlyAutoMegaphone
自動で応援するやつ

## 使い方

1. clone
```
$ git clone https://github.com/enkaism/wantedlyAutoMegaphone
$ cd wantedlyAutoMegaphone
```

2. 環境変数を設定(direnvつかうとらく)
```
export FACEBOOK_EMAIL=XXX@XXX.XXX
export FACEBOOK_PASSWORD=XXXXXXXX
export TIMEOUT_SECONDS=X(デフォルトは3)
```

3. run
```
$ go run main.go https://www.wantedly.com/projects/XXXXXX https://www.wantedly.com/projects/XXXXXX https://www.wantedly.com/projects/XXXXXX
```
