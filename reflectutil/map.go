package reflectutil
/*
import (
	"reflect"
)

type UpdateCallback func(oldValue reflect.Value) (reflect.Value, error)

type MapReflect struct {
  v reflect.Value
}

func NewMapReflect(v reflect.Value) (*MapReflect, error) {
  if v.Type.Kind() != reflect.Map {
    return nil, fmt.Errorf("Value is not a map (for NewMapReflect)")
  }
  return &MapReflect{
    v: v
  }, nil
}

func (r *MapReflect) updateKey(key reflect.Value, cb UpdateCallback) error {
  oldValue := r.v.MapIndex(key)
  newValue, err := cb(oldValue)
  if err != nil {
    return err
  }
  r.v.MapIndex(key)
  return nil
}
*/
