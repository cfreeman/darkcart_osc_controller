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
	"bitbucket.org/liamstask/gosc"
	"errors"
	"fmt"
)

func oscServer() error {
	osc.HandleFunc("/position", func(msg *osc.Message) {
		p, err := parseFloatArg(msg.Args[0])
		if err != nil {
			fmt.Printf("Unable to parse position argument.")
			return
			// TODO: BZZT. Invalid position argument.
		}

		fmt.Printf("Position: %f\n", p)
		// TODO: Push the position argument into the stepper controller.
	})

	osc.HandleFunc("/height", func(msg *osc.Message) {
		h, err := parseFloatArg(msg.Args[0])
		if err != nil {
			fmt.Printf("Unable to parse height argument.")
			return
			// TODO: BZZT. Invalid height argument.
		}

		fmt.Printf("Height: %f\n", h)
		// TODO: Push the height argument into the linear actuator.
	})

	osc.HandleFunc("/sequence", func(msg *osc.Message) {
		s, err := parseIntArg(msg.Args[0])
		if err != nil {
			fmt.Printf("Unable to parse sequence to trigger.")
			return
			// TODO: BZZT. Invalid height argument.
		}

		fmt.Printf("Sequence: %d\n", s)
		// TODO: Push the sequence into the mini maestro.
	})

	return osc.ListenAndServeUDP(":8000", nil)
}

func parseIntArg(arg interface{}) (int32, error) {
	switch v := arg.(type) {
	default:
		return 0, errors.New("osc message does not contain integer argument.")
	case int32:
		return v, nil
	case int64:
		return int32(v), nil
	}
}

func parseFloatArg(arg interface{}) (float32, error) {
	switch v := arg.(type) {
	default:
		return 0.0, errors.New("OSC message does not contain float argument.")
	case float32:
		return v, nil
	case float64:
		return float32(v), nil
	}
}