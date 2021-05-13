package main

import "fmt"

/*
 * 学習内容
 * ・インターフェイスはメソッドを複数記述し、リスト化する。
 * ・メソッドはレシーバ引数を取る関数。typeで型定義されたものをレシーバ引数として取る。
 */

// byte型で4個の変数が格納できる配列IPAddrを作成
type IPAddr [4]byte

// IPAddrにインターフェイスfmt.Stringを実装する。
// fmt.Sprintfと%dを使用することで配列aをstring型に変更して返り値とする
func (a IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", a[0], a[1], a[2], a[3])
}

func main() {
	// mapの作成 キーはstring型、値はIPAddrで処理を行ったstring型とする
	hosts := map[string]IPAddr{
		"loopback":  {12, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	// 繰り返し処理を実施。キーをname、値をipとして出力する
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
