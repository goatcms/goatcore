package orm

import "testing"

func TestEmptyFindAll(t *testing.T) {
	context, err := newTestScope()
	if err != nil {
		t.Error(err)
		return
	}
	rows, err := context.dao.FindAll(context.bdtx)
	if err != nil {
		t.Error(err)
		return
	}
	result, err := loadResult(rows)
	if err != nil {
		t.Error(err)
		return
	}
	compareSlice(t, result, []TestEntity{})
}

func TestFixturesFindAll(t *testing.T) {
	context, err := newTestScope()
	if err != nil {
		t.Error(err)
		return
	}
	if err := lodaFixtures(context); err != nil {
		t.Error(err)
		return
	}
	rows, err := context.dao.FindAll(context.bdtx)
	if err != nil {
		t.Error(err)
		return
	}
	result, err := loadResult(rows)
	if err != nil {
		t.Error(err)
		return
	}
	compareSlice(t, result, []TestEntity{
		TestEntity{10, "title1", "content1", "path1"},
		TestEntity{20, "title2", "content2", "path2"},
	})
}

func TestFindByID(t *testing.T) {
	var entity TestEntity
	context, err := newTestScope()
	if err != nil {
		t.Error(err)
		return
	}
	if err := lodaFixtures(context); err != nil {
		t.Error(err)
		return
	}
	row := context.dao.FindByID(context.bdtx, 10)
	if row.Err() != nil {
		t.Error(row.Err())
		return
	}
	row.StructScan(&entity)
	expected := TestEntity{10, "title1", "content1", "path1"}
	if !compareEntity(entity, expected) {
		t.Errorf("result is %v and expected %v", entity, expected)
	}
}

// Insert store given articles to database
func TestInsert(t *testing.T) {
	var (
		id           int64
		err          error
		insertEntity TestEntity
		resultEntity TestEntity
	)
	insertEntity = TestEntity{0, "title11", "content11", "path11"}
	context, err := newTestScope()
	if err != nil {
		t.Error(err)
		return
	}
	if id, err = context.dao.Insert(context.bdtx, insertEntity); err != nil {
		t.Error(err)
		return
	}
	insertEntity.ID = id
	row := context.dao.FindByID(context.bdtx, id)
	if row.Err() != nil {
		t.Error(row.Err())
		return
	}
	row.StructScan(&resultEntity)
	if !compareEntity(resultEntity, insertEntity) {
		t.Errorf("result is %v and expected %v", resultEntity, insertEntity)
	}
}

// TestInsertWithID store given articles to database (with its ID)
func TestInsertWithID(t *testing.T) {
	var (
		id           int64 = 6666
		err          error
		insertEntity TestEntity
		resultEntity TestEntity
	)
	insertEntity = TestEntity{id, "title11", "content11", "path11"}
	context, err := newTestScope()
	if err != nil {
		t.Error(err)
		return
	}
	if err = context.dao.InsertWithID(context.bdtx, insertEntity); err != nil {
		t.Error(err)
		return
	}
	row := context.dao.FindByID(context.bdtx, id)
	if row.Err() != nil {
		t.Error(row.Err())
		return
	}
	row.StructScan(&resultEntity)
	if !compareEntity(resultEntity, insertEntity) {
		t.Errorf("result is %v and expected %v", resultEntity, insertEntity)
	}
}

// Update data of article
func TestUpdate(t *testing.T) {
	var (
		expectEntity = TestEntity{10, "title1x", "content1x", "path1x"}
		resultEntity TestEntity
	)
	context, err := newTestScope()
	if err != nil {
		t.Error(err)
		return
	}
	if err := lodaFixtures(context); err != nil {
		t.Error(err)
		return
	}
	if err := context.dao.Update(context.bdtx, expectEntity); err != nil {
		t.Error(err)
		return
	}
	row := context.dao.FindByID(context.bdtx, 10)
	if row.Err() != nil {
		t.Error(row.Err())
		return
	}
	row.StructScan(&resultEntity)

	if !compareEntity(resultEntity, expectEntity) {
		t.Errorf("result is %v and expected %v", resultEntity, expectEntity)
	}
}

// Delete remove specyfic record
func TestDelete(t *testing.T) {
	context, err := newTestScope()
	if err != nil {
		t.Error(err)
		return
	}
	if err := lodaFixtures(context); err != nil {
		t.Error(err)
		return
	}
	if err := context.dao.Delete(context.bdtx, 10); err != nil {
		t.Error(err)
		return
	}
	rows, err := context.dao.FindAll(context.bdtx)
	if err != nil {
		t.Error(err)
		return
	}
	result, err := loadResult(rows)
	if err != nil {
		t.Error(err)
		return
	}
	compareSlice(t, result, []TestEntity{
		TestEntity{20, "title2", "content2", "path2"},
	})
}
