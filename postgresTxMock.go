package postgres

import (
	"database/sql"
)

// PgDbTxMock PgDbTxMock
type PgDbTxMock struct {
	Tx       *sql.Tx
	PgDBMock *PgDBMock
}

// Insert Insert
func (t *PgDbTxMock) Insert(query string, args ...interface{}) (bool, int64) {
	var rtn = false
	var id int64
	if !t.PgDBMock.mockInsertSuccess1Used {
		t.PgDBMock.mockInsertSuccess1Used = true
		rtn = t.PgDBMock.MockInsertSuccess1
		id = t.PgDBMock.MockInsertID1
	} else if !t.PgDBMock.mockInsertSuccess2Used {
		t.PgDBMock.mockInsertSuccess2Used = true
		rtn = t.PgDBMock.MockInsertSuccess2
		id = t.PgDBMock.MockInsertID2
	} else if !t.PgDBMock.mockInsertSuccess3Used {
		t.PgDBMock.mockInsertSuccess3Used = true
		rtn = t.PgDBMock.MockInsertSuccess3
		id = t.PgDBMock.MockInsertID3
	} else if !t.PgDBMock.mockInsertSuccess4Used {
		t.PgDBMock.mockInsertSuccess4Used = true
		rtn = t.PgDBMock.MockInsertSuccess4
		id = t.PgDBMock.MockInsertID4
	} else if !t.PgDBMock.mockInsertSuccess5Used {
		t.PgDBMock.mockInsertSuccess5Used = true
		rtn = t.PgDBMock.MockInsertSuccess5
		id = t.PgDBMock.MockInsertID5
	} else if !t.PgDBMock.mockInsertSuccess6Used {
		t.PgDBMock.mockInsertSuccess6Used = true
		rtn = t.PgDBMock.MockInsertSuccess6
		id = t.PgDBMock.MockInsertID6
	} else if !t.PgDBMock.mockInsertSuccess7Used {
		t.PgDBMock.mockInsertSuccess7Used = true
		rtn = t.PgDBMock.MockInsertSuccess7
		id = t.PgDBMock.MockInsertID7
	} else if !t.PgDBMock.mockInsertSuccess8Used {
		t.PgDBMock.mockInsertSuccess8Used = true
		rtn = t.PgDBMock.MockInsertSuccess8
		id = t.PgDBMock.MockInsertID8
	}
	return rtn, id
}

// Update Update
func (t *PgDbTxMock) Update(query string, args ...interface{}) bool {
	var rtn = false
	if !t.PgDBMock.mockUpdateSuccess1Used {
		t.PgDBMock.mockUpdateSuccess1Used = true
		rtn = t.PgDBMock.MockUpdateSuccess1
	} else if !t.PgDBMock.mockUpdateSuccess2Used {
		t.PgDBMock.mockUpdateSuccess2Used = true
		rtn = t.PgDBMock.MockUpdateSuccess2
	} else if !t.PgDBMock.mockUpdateSuccess3Used {
		t.PgDBMock.mockUpdateSuccess3Used = true
		rtn = t.PgDBMock.MockUpdateSuccess3
	} else if !t.PgDBMock.mockUpdateSuccess4Used {
		t.PgDBMock.mockUpdateSuccess4Used = true
		rtn = t.PgDBMock.MockUpdateSuccess4
	}
	return rtn
}

// Delete Delete
func (t *PgDbTxMock) Delete(query string, args ...interface{}) bool {
	var rtn = false
	if !t.PgDBMock.mockDeleteSuccess1Used {
		t.PgDBMock.mockDeleteSuccess1Used = true
		rtn = t.PgDBMock.MockDeleteSuccess1
	} else if !t.PgDBMock.mockDeleteSuccess2Used {
		t.PgDBMock.mockDeleteSuccess2Used = true
		rtn = t.PgDBMock.MockDeleteSuccess2
	} else if !t.PgDBMock.mockDeleteSuccess3Used {
		t.PgDBMock.mockDeleteSuccess3Used = true
		rtn = t.PgDBMock.MockDeleteSuccess3
	} else if !t.PgDBMock.mockDeleteSuccess4Used {
		t.PgDBMock.mockDeleteSuccess4Used = true
		rtn = t.PgDBMock.MockDeleteSuccess4
	} else if !t.PgDBMock.mockDeleteSuccess5Used {
		t.PgDBMock.mockDeleteSuccess5Used = true
		rtn = t.PgDBMock.MockDeleteSuccess5
	} else if !t.PgDBMock.mockDeleteSuccess6Used {
		t.PgDBMock.mockDeleteSuccess6Used = true
		rtn = t.PgDBMock.MockDeleteSuccess6
	} else if !t.PgDBMock.mockDeleteSuccess7Used {
		t.PgDBMock.mockDeleteSuccess7Used = true
		rtn = t.PgDBMock.MockDeleteSuccess7
	} else if !t.PgDBMock.mockDeleteSuccess8Used {
		t.PgDBMock.mockDeleteSuccess8Used = true
		rtn = t.PgDBMock.MockDeleteSuccess8
	}
	return rtn
}

// Commit Commit
func (t *PgDbTxMock) Commit() bool {
	return t.PgDBMock.MockCommitSuccess
}

// Rollback Rollback
func (t *PgDbTxMock) Rollback() bool {
	return t.PgDBMock.MockRollbackSuccess
}
