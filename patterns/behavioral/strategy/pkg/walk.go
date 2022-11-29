package pkg

import "fmt"

type WalkStrategy struct{}

func (r *WalkStrategy) Route(startPoint, endPoint int) {
	var (
		avrSpeed  = 4
		total     = endPoint - startPoint
		totalTime = total * 60
	)

	fmt.Printf("Walk: [%d] to B: [%d] Avr speed: [%d] Total: [%d] Total time: [%d] min\n",
		startPoint, endPoint, avrSpeed, total, totalTime)
}
