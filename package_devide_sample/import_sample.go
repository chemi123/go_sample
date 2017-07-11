package main

import "fmt"
import "./sample"

func main() {
	tables := make([]sample.Table, 0)
	// 失敗する
	// var table Table = &Table{1, 7}
	// DoManage(table)

	// 成功する。何故？
	table1 := sample.NewTable(1, 7)
	table2 := sample.NewTable(2, 8)
	table3 := sample.NewTable(3, 9)
	tables = append(tables, *table1)
	tables = append(tables, *table2)
	tables = append(tables, *table3)

	// 以下のやり方だとうまくいかなかった(要素に変更がなかったということ)
	// おそらく理由はvに対してはtable要素のコピーが走っているから
	// ポインタを渡せる方法があれば知りたい
	// for i, v := range tables {
	//         DoManage(&v)
	// }
	for i, _ := range tables {
		sample.DoManage(&tables[i])
	}

	// DoManagerが呼び出された後に値が変更されていることを確認
	for i, v := range tables {
		fmt.Printf("After DoManage, tables[%v].Players: %v\n", i, v.Players)
	}

	// embedができているかの確認
	desk := sample.NewDesk(4, 10)
	sample.DoManage(desk)
	desk.GetDeskInfo()
}
