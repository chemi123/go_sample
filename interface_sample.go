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
	fmt.Println("Table:", manager.getTableID())
	manager.addPlayer(-2)
	fmt.Println(manager.countPlayer(), "players left")
}

func main() {
	tables := make([]Table, 0)
	fmt.Println(tables)
	// 失敗する
	// var table Table = &Table{1, 7}
	// doManage(table)

	// 成功する。何故？
	table1 := &Table{1, 7}
	table2 := &Table{2, 8}
	table3 := &Table{3, 9}
	tables = append(tables, *table1)
	tables = append(tables, *table2)
	tables = append(tables, *table3)
	for i, _ := range tables {
		doManage(&tables[i])
	}

	// doManagerが呼び出された後に値が変更されていることを確認
	for i, v := range tables {
		fmt.Printf("After doManage, tables[%v].Players: %v\n", i, v.Players)
	}
}
