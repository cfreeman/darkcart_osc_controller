/*
 * Copyright (c) Clinton Freeman 2014
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and
 * associated documentation files (the "Software"), to deal in the Software without restriction,
 * including without limitation the rights to use, copy, modify, merge, publish, distribute,
 * sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all copies or
 * substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT
 * NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
 * NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
 * DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package main

import (
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"time"
)

func maestroLink(sequence chan int32) {
	err := rpio.Open()
	if err != nil {
		fmt.Printf("Unable to open IO ports on the Raspberry PI.\n")
	}

	for {
		// Initially, we just reset the maestro for any sequence.
		<-sequence

		// Pulse the rst pin on the maestro for 5 milliseconds to reset.
		pin := rpio.Pin(17)
		pin.Output()
		pin.High()
		time.Sleep(5 * time.Millisecond)
		pin.Low()

		// TODO: Connect to maestro over serial and trigger a specific animation
		// sequence stored on the device.
	}

	// Never reached. Here for completness. Clean up the memory used by the GPIO ports.
	rpio.Close()
}
