package mockrows

import "github.com/goatcms/goatcore/varutil"

// List represent plain data for mock interface
type List []Row

// Row represent a single object data
type Row map[string]interface{}

// Rows represent a mocked database response
type Rows struct {
	index int
	data  List
}

// NewRows create new mocked rows
func NewRows(data List) *Rows {
	return &Rows{
		index: 0,
		data:  data,
	}
}

// Close close mock database connection
func (r *Rows) Close() error {
	return nil
}

// Next check if next element exist
func (r *Rows) Next() bool {
	return r.index < len(r.data)
}

// StructScan scan current row
func (r *Rows) StructScan(dest interface{}) error {
	if err := varutil.LoadStruct(dest, r.data[r.index], "schema", true); err != nil {
		return err
	}
	r.index++
	return nil
}
