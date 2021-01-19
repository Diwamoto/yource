## テストの実行
`go test ./...` 
もし、ここで`./`を忘れてしまうと`$GOPATH`等の全てのモジュールをテストしてしまうので忘れないように!!!


### Table Driven Test
テスト設計は["Table Driven Test"](https://github.com/golang/go/wiki/TableDrivenTests)に従って
基本的に以下の形で行う。
```
func 〇〇Test(t *testing.T) {
	tests := []struct {
		in,want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		rs := func(tt.in)
		if rs != tt.in {
			t.Error("")//エラー文
		}

	}
}
```
