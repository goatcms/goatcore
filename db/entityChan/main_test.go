package entitychan

import (
	"testing"

	"github.com/goatcms/goat-core/db"
	"github.com/goatcms/goat-core/testbase/mocks/mockentity"
	"github.com/goatcms/goat-core/testbase/mocks/mockrows"
)

const (
	testTitle1    = "title1"
	testTitle2    = "title2"
	testCcontent1 = "content1"
	testCcontent2 = "content2"
)

// TestChanCorverterOneThread test convert data for single thread
func TestChanCorverterOneThread(t *testing.T) {
	var rows db.Rows = mockrows.NewRows(mockrows.List{
		mockrows.Row{
			"title":   testTitle1,
			"content": testCcontent1,
		},
		mockrows.Row{
			"title":   testTitle2,
			"content": testCcontent2,
		},
	})
	c := NewChanCorverter(nil, rows, mockentity.NewEntityI)
	c.Go()

	r1 := (<-c.Chan).(*mockentity.Entity)
	if r1.Title != testTitle1 {
		t.Errorf("r1.Title(%v) != %v", r1.Title, testTitle1)
	}
	if r1.Content != testCcontent1 {
		t.Errorf("r1.Content(%v) != %v", r1.Content, testCcontent1)
	}

	r2 := (<-c.Chan).(*mockentity.Entity)
	if r2.Title != testTitle2 {
		t.Errorf("r1.Title(%v) != %v", r2.Title, testTitle2)
	}
	if r2.Content != testCcontent2 {
		t.Errorf("r1.Content(%v) != %v", r2.Content, testCcontent2)
	}

	/*_, ok := (<-c.Chan).(*mockentity.Entity)
	if ok == true {
		t.Errorf("channel should be closed after read all elements")
	}*/
}

// TestChanCorverterMultiThread test convert data for two thread
func TestChanCorverterMultiThread(t *testing.T) {
	var rows db.Rows = mockrows.NewRows(mockrows.List{
		mockrows.Row{
			"title":   testTitle1,
			"content": testCcontent1,
		},
		mockrows.Row{
			"title":   testTitle2,
			"content": testCcontent2,
		},
	})
	c := NewChanCorverter(nil, rows, mockentity.NewEntityI)
	go c.Go()

	r1 := (<-c.Chan).(*mockentity.Entity)
	if r1.Title != testTitle1 {
		t.Errorf("r1.Title(%v) != %v", r1.Title, testTitle1)
	}
	if r1.Content != testCcontent1 {
		t.Errorf("r1.Content(%v) != %v", r1.Content, testCcontent1)
	}

	r2 := (<-c.Chan).(*mockentity.Entity)
	if r2.Title != testTitle2 {
		t.Errorf("r1.Title(%v) != %v", r2.Title, testTitle2)
	}
	if r2.Content != testCcontent2 {
		t.Errorf("r1.Content(%v) != %v", r2.Content, testCcontent2)
	}

	/*_, ok := (<-c.Chan).(*mockentity.Entity)
	if ok == true {
		t.Errorf("channel should be closed after read all elements")
	}*/
}
