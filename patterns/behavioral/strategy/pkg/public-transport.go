package pkg

import "fmt"

type PublicTransportStrategy struct{}

func (r *PublicTransportStrategy) Route(startPoint, endPoint int) {
	var (
		avrSpeed  = 40
		total     = endPoint - startPoint
		totalTime = total * 40
	)

	fmt.Printf("Public transport: [%d] to B: [%d] Avr speed: [%d] Total: [%d] Total time: [%d] min\n",
		startPoint, endPoint, avrSpeed, total, totalTime)
}
