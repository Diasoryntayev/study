package pkg

import "fmt"

type RoadStrategy struct{}

func (r *RoadStrategy) Route(startPoint, endPoint int) {
	var (
		avrSpeed   = 30
		trafficJam = 2
		total      = endPoint - startPoint
		totalTime  = total * 40 * trafficJam
	)

	fmt.Printf("Road A: [%d] to B: [%d] Avr speed: [%d] Trrafic jam: [%d] Total: [%d] Total time: [%d] min\n",
		startPoint, endPoint, avrSpeed, trafficJam, total, totalTime)
}
