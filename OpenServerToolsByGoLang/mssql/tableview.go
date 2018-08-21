package main

import "github.com/lxn/walk"
import "sort"



//--------------------------------------------------------------------
// 模板继承
//--------------------------------------------------------------------

// Called by the TableView from SetModel and every time the model publishes a
// RowsReset event.
func (m *TableViewModel) RowCount() int {
	return len(m.items)
}

// Called by the TableView when it needs the text to display for a given cell.
func (m *TableViewModel) Value(row, col int) interface{} {
	item := m.items[row]

	switch col {
	case 0:
		return item.ServerId

	case 1:
		return item.ServerName

	case 2:
		return item.ServerMachine

	case 3:
		return item.AndroidCount
	case 4:
		return item.GameRoomListIndex
	case 5:
		return item.SqlFileRobotListIndex

	}

	panic("unexpected col")
}

// Called by the TableView to retrieve if a given row is checked.
func (m *TableViewModel) Checked(row int) bool {
	return m.items[row].checked
}

// Called by the TableView when the user toggled the check box of a given row.
func (m *TableViewModel) SetChecked(row int, checked bool) error {
	m.items[row].checked = checked

	return nil
}

// Called by the TableView to sort the model.
func (m *TableViewModel) Sort(col int, order walk.SortOrder) error {
	m.sortColumn, m.sortOrder = col, order

	sort.SliceStable(m.items, func(i, j int) bool {
		a, b := m.items[i], m.items[j]

		c := func(ls bool) bool {
			if m.sortOrder == walk.SortAscending {
				return ls
			}

			return !ls
		}

		switch m.sortColumn {
		case 0:
			return c(a.ServerId < b.ServerId)

		case 1:
			return c(a.ServerName < b.ServerName)

		case 2:
			return c(a.ServerMachine < b.ServerMachine)

		case 3:
			return c(a.AndroidCount < b.AndroidCount)

		case 4:
			return c(a.GameRoomListIndex < b.GameRoomListIndex)
		case 5:
			return c(a.SqlFileRobotListIndex < b.SqlFileRobotListIndex)

		}

		panic("unreachable")
	})

	return m.SorterBase.Sort(col, order)
}

