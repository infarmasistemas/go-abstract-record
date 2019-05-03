package instance

import (
	"fmt"
)

type InstanceOps struct {
	object      interface{}
	objectArray interface{}
}

func (i *InstanceOps) Prepare(object interface{}, objectArray interface{}) {
	i.object = object
	i.objectArray = objectArray
}

// TODO: Code a way to instantiate objects using new
func (i *InstanceOps) New(values ...interface{}) error {
	fmt.Print(values)
	return nil
}

func defineAttributes() {

}