package main

import "fmt"

// TableManager is to manage Table struct
type TableManager interface {
	countPlayer() int
	addPlayer(int)
	getTableID() int
}

// Table meets TableManager's method
type Table struct {
	ID, Players int
}

// embe sample
type Desk struct {
	*Table
}

func (table *Table) countPlayer() int {
	return table.Players
}

func (table *Table) getTableID() int {
	return table.ID
}

func (table *Table) addPlayer(num int) {
	table.Players += num
}

func DoManage(manager TableManager) {
	manager.addPlayer(-2)
	fmt.Printf("Table %v has %v players left.\n", manager.getTableID(), manager.countPlayer())
}

func NewTable(id, players int) *Table {
	return &Table{id, players}
}

func NewDesk(id, players int) *Desk {
	return &Desk{Table: NewTable(id, players)}
}

// Tableを埋め込んだDesk特有の関数
func (desk Desk) GetDeskInfo() {
	fmt.Printf("Desck %v has %v players left.\n", desk.getTableID(), desk.countPlayer())
}

func main() {
	tables := make([]Table, 0)
	// 失敗する
	// var table Table = &Table{1, 7}
	// DoManage(table)

	// 成功する。何故？
	table1 := NewTable(1, 7)
	table2 := NewTable(2, 8)
	table3 := NewTable(3, 9)
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
		DoManage(&tables[i])
	}

	// DoManagerが呼び出された後に値が変更されていることを確認
	for i, v := range tables {
		fmt.Printf("After DoManage, tables[%v].Players: %v\n", i, v.Players)
	}

	// embedができているかの確認
	desk := NewDesk(4, 10)
	DoManage(desk)
	desk.GetDeskInfo()
}
