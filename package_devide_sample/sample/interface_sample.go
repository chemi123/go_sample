package sample

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

// Desk embed Table struct
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

// DoManage takes TableManager
func DoManage(manager TableManager) {
	manager.addPlayer(-2)
	fmt.Printf("Table %v has %v players left.\n", manager.getTableID(), manager.countPlayer())
}

// NewTable returns Table instance
func NewTable(id, players int) *Table {
	return &Table{id, players}
}

// NewDesk returns Desk instance
func NewDesk(id, players int) *Desk {
	return &Desk{Table: NewTable(id, players)}
}

// GetDeskInfo takes Desk receiver
func (desk Desk) GetDeskInfo() {
	fmt.Printf("Desck %v has %v players left.\n", desk.getTableID(), desk.countPlayer())
}
