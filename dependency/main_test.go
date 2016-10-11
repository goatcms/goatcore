package dependency

const (
	MyDepName       = "MyDep"
	MyCircleDepName = "MyCircleDep"
)

type MyDepInterface interface {
	IsItOk() bool
	Set(int)
	Get() int
}

type MyDep struct {
	value int
}

type MyCircleDepInterface interface {
}

type MyCircleDep struct {
	instance MyCircleDepInterface
}

func (d *MyDep) IsItOk() bool {
	return true
}

func (d *MyDep) Get() int {
	return d.value
}

func (d *MyDep) Set(v int) {
	d.value = v
}

func MyDepFactory(dp Provider) (Instance, error) {
	return &MyDep{}, nil
}

func MyCircleDepFactory(dp Provider) (Instance, error) {
	instance, err := dp.Get(MyCircleDepName)
	if err != nil {
		return nil, err
	}
	return &MyCircleDep{
		instance: instance,
	}, nil
}

type SimpleObject struct {
	Instance *MyDep `inject:"MyDep"`
}

func (o *SimpleObject) Load(dp *Provider) error {
	ins, err := (*dp).Get(MyDepName)
	if err != nil {
		return err
	}
	o.Instance = ins.(*MyDep)
	return nil
}
