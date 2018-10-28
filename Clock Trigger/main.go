package main

import (
	"time"
)

var interval = 10
var processingTime time.Time

func main() {
	// Run configuration (Initialize database connection)
	configure()
	trackCount := len(getAllTracks())

	timer := time.Now()

	for {
		time.Sleep(10 * time.Second)

		// Has 10 minutes passed?
		if itIsTime(timer) {
			processingTime = time.Now() // Start our timer
			newTracks := newTracksSinceLastCheck(trackCount)

			// Do we have any new tracks?
			if len(newTracks) != 0 {
				// Send new tracks to slack webhook
				invokeWebhook(newTracks)
			}

			// reset
			timer = time.Now()
			trackCount = len(getAllTracks())
		}
	}
}
