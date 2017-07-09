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

func (table *Table) countPlayer() int {
	return table.Players
}

func (table *Table) getTableID() int {
	return table.ID
}

func (table *Table) addPlayer(num int) {
	table.Players += num
}

func doManage(manager TableManager) {
	fmt.Println("Talbe:", manager.getTableID())
	manager.addPlayer(-2)
	fmt.Println(manager.countPlayer(), "players left")
}

func main() {
	// 失敗する
	// var table Table = &Table{1, 7}
	// doManage(table)

	// 成功する。何故？
	table := &Table{1, 7}
	doManage(table)

	// doManagerが呼び出された後に値が変更されている。
	// doManagerはtableのポインタが渡されている？
	fmt.Println("After doManage, table.Players:", table.Players)
}
