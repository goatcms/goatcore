package dependency

type Builder struct {
	Static   bool
	Factory  Factory
	Instance Instance
}

func (d *Builder) Get() (Instance, error) {
	var err error
	if d.Static == false {
		return d.factory()
	}
	if d.Instance != nil {
		return d.Instance, nil
	}
  d.Instance, err = d.factory()
  if err != nil {
    return nil, err
  }
	return d.Instance, nil
}

func (d *Builder) factory() (Instance, error) {
	var err error
	d.Instance, err = d.Factory()
	if err != nil {
		return nil, err
	}
	return d.Instance, nil
}
