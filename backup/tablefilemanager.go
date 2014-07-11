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

/* Manage table files, handles creation/renaming/removing of table files. */

package database

import (
	"os"
)

// Creates table files.
func Create(path string, name string) int {
	if len(name) > MaxTableNameLength {
		return TableNameTooLong
	}
	// Create table files with extension names.
	for _, ext := range TableFiles() {
		_, err := os.Create(path + name + ext)
		if err != nil {
			Err("tablefilemanager", "Create", err)
			return CannotCreateTableFile
		}
	}
	// Create table directories with name suffixes.
	for _, dir := range TableDirs() {
		err := os.Mkdir(path+name+dir, TableDirPerm)
		if err != nil {
			Err("tablefilemanager", "Create", err)
			return CannotCreateTableDir
		}
	}
	return OK
}

// Renames table files.
func Rename(path string, oldName string, newName string) int {
	for _, ext := range TableFiles() {
		err := os.Rename(path+oldName+ext, path+newName+ext)
		if err != nil {
			Err("tablefilemanager", "Rename", err)
			return CannotRenameTableFile
		}
	}
	for _, dir := range TableDirs() {
		err := os.Rename(path+oldName+dir, path+newName+dir)
		if err != nil {
			Err("tablefilemanager", "Rename", err)
			return CannotRenameTableDir
		}
	}
	return OK
}

// Deletes table files
func Delete(path string, name string) int {
	for _, ext := range TableFiles() {
		err := os.Remove(path + name + ext)
		if err != nil {
			Err("tablefilemanager", "Delete", err)
			return CannotRemoveTableFile
		}
	}
	for _, dir := range TableDirs() {
		err := os.RemoveAll(path + name + dir)
		if err != nil {
			Err("tablefilemanager", "Delete", err)
			return CannotRemoveTableDir
		}
	}
	return OK
}
