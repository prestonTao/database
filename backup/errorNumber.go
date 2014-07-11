package database

/*
Error codes.
Some of these codes may indicate data loss or issues with file system.
These error codes will immediately halt the function when they happen.
*/

//info
const (
	OK = 0 // do not change :)
)

//warn
const (
	ColumnNameNotFound = 200 + iota
	FailedToCopyCertainRows
	FailedToReadCertainRows
)

//error
const (
	CannotOpenDatabaseDirectory = 100 + iota
	CannotReadDatabaseDirectory
	CannotOpenTableFiles
	CannotOpenTableDefFile
	CannotOpenTableDataFile
	InvalidColumnDefinition
	CannotStatTableDefFile
	CannotStatTableDataFile
	CannotSeekTableDataFile
	CannotSeekTableDefFile
	CannotReadTableDataFile
	CannotWriteTableDataFile
	CannotWriteTableDefFile
	CannotFlushTableDefFile
	CannotFlushTableDataFile
	TableDoesNotHaveDelColumn
	TableNameTooLong
	TableAlreadyExists
	CannotCreateTableFile
	CannotCreateTableDir
	CannotRenameTableFile
	CannotRenameTableDir
	CannotRemoveTableFile
	CannotRemoveTableDir
	ColumnAlreadyExists
	ColumnNameTooLong
	TableNotFound
	InvalidColumnLength
	AliasNotFound
	AliasAlreadyExists
	CannotCreateInitFile
	CannotReadSharedLocksDir
	CannotReadExclusiveLocksFile
	CannotUnlockSharedLock
	CannotUnlockExclusiveLock
	CannotCreateFile
	CannotReadFile
	CannotWriteFile
	CannotRemoveSpecialColumn
)

/*
Database logic errors.
Data is ensured to be safe and consistent when these codes are raised.
*/
const (
	DuplicatedPKValue = 301 + iota
	InvalidFKValue
	DeleteRestricted
	UpdateRestricted
	CannotLockInExclusive
	CannotLockInShared
	DuplicatedAlias
)
