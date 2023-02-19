package sold

import (
	"github.com/zhorifiandi/parking-lot-lld/usecase/mvp"
)

type Application struct {
	BaseApplication mvp.Application
}

func NewApplication(slots []int) *Application {
	baseApp := mvp.NewApplication(mvp.ApplicationInputs{})
	for _, slot := range slots {
		baseApp.AddFloor(slot)
	}

	return &Application{
		BaseApplication: *baseApp,
	}
}
