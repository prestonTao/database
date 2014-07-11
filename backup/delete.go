/*
<DBGo - A flat-file relational database engine implementation in Go programming language>
Copyright (C) <2011>  <Houzuo (Howard) Guo>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

/*
Delete a row, trigger appropriate triggers and log information for rollback.
*/

package database

import ()

type UndoDelete struct {
	Table     *Table
	RowNumber int
}

// An insert operation is undone by marking the inserted row deleted.
func (u *UndoDelete) Undo() int {
	return u.Table.Update(u.RowNumber, map[string]string{"~del": ""})
}

func (tr *Transaction) Delete(t *Table, rowNumber int) int {
	// Execute "before delete" triggers.
	beforeTable, status := tr.DB.Get("~before")
	if status != OK {
		return status
	}
	row, status := t.Read(rowNumber)
	if status != OK {
		return status
	}
	triggerRA := New()
	_, status = triggerRA.Load(beforeTable)
	if status != OK {
		return status
	}
	_, status = triggerRA.Select("TABLE", Eq{}, t.Name)
	if status != OK {
		return status
	}
	status = ExecuteTrigger(tr.DB, t, triggerRA, "DE", row, nil)
	if status != OK {
		return status
	}
	// Update the row.
	status = t.Delete(rowNumber)
	if status != OK {
		return status
	}
	// Execute "after delete" triggers.
	afterTable, status := tr.DB.Get("~after")
	if status != OK {
		return status
	}
	triggerRA = New()
	_, status = triggerRA.Load(afterTable)
	if status != OK {
		return status
	}
	_, status = triggerRA.Select("TABLE", Eq{}, t.Name)
	if status != OK {
		return status
	}
	status = ExecuteTrigger(tr.DB, t, triggerRA, "DE", row, nil)
	if status != OK {
		return status
	}
	// Log the deleted row.
	tr.Log(&UndoDelete{t, rowNumber})
	return OK
}
