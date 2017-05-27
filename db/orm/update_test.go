package orm

import "testing"

func TestUpdate(t *testing.T) {
	scope, err := newTestScope()
	if err != nil {
		t.Error(err)
		return
	}
	createTable, err := NewCreateTable(scope.table, scope.driver)
	if err != nil {
		t.Error(err)
		return
	}
	if err = createTable(scope.tx); err != nil {
		t.Error(err)
		return
	}
	insertWithID, err := NewInsertWithID(scope.table, scope.driver)
	if err != nil {
		t.Error(err)
		return
	}
	inEntity := &TestEntity{10, "title1", "content1", "path1"}
	if err = insertWithID(scope.tx, inEntity); err != nil {
		t.Error(err)
		return
	}
	update, err := NewUpdate(scope.table, scope.driver)
	if err != nil {
		t.Error(err)
		return
	}
	updatedEntity := &TestEntity{10, "title2", "content2", "path2"}
	if err = update(scope.tx, updatedEntity); err != nil {
		t.Error(err)
		return
	}
	findByID, err := NewFindByID(scope.table, scope.driver)
	if err != nil {
		t.Error(err)
		return
	}
	row, err := findByID(scope.tx, 10)
	if err != nil {
		t.Error(err)
		return
	}
	outEntity := &TestEntity{}
	if err = row.StructScan(outEntity); err != nil {
		t.Error(err)
		return
	}
	if outEntity.Content != updatedEntity.Content {
		t.Errorf("Content must be the same %v == %v", outEntity.Content, updatedEntity.Content)
		return
	}
	if outEntity.Title != updatedEntity.Title {
		t.Errorf("Title must be the same %v == %v", outEntity.Title, updatedEntity.Title)
		return
	}
	if outEntity.ID != 10 || updatedEntity.ID != 10 {
		t.Errorf("ID must be the same %v == 10 && %v == 10", outEntity.ID, updatedEntity.ID)
		return
	}
	if outEntity.ID != updatedEntity.ID {
		t.Errorf("Title must be the same %v == %v", outEntity.ID, updatedEntity.ID)
		return
	}
	if outEntity.Image != updatedEntity.Image {
		t.Errorf("Image must be the same %v == %v", outEntity.Image, updatedEntity.Image)
		return
	}

}
