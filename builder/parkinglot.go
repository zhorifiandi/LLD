package builder

import (
	"github.com/zhorifiandi/parking-lot-lld/usecase/mvp"
)

func NewApplication(slots []int) *mvp.Application {
	baseApp := mvp.NewApplication(mvp.ApplicationInputs{})
	for _, slot := range slots {
		baseApp.AddFloor(slot)
	}

	return baseApp
}
