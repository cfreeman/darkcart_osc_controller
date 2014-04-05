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
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/huin/goserial"
	"io"
	"io/ioutil"
	"strings"
	"time"
)

// arduinoLink makes a link to the arduino so that commands can be transmitted down the serial line to
// an arduino where it can be used to drive various motors and actuators.
func arduinoLink(position chan float32, height chan float32) {
	// Find the device that represents the arduino serial connection.
	c := &goserial.Config{Name: findArduino(), Baud: 9600}
	s, err := goserial.OpenPort(c)

	// When connecting to an older revision arduino, you need to wait a little while it resets.
	time.Sleep(1 * time.Second)

	if err != nil {
		fmt.Printf("Unable to find arduino.\n")
		return
	}

	for {
		select {
		case p := <-position:
			sendArduinoCommand('p', p, s)
		case h := <-height:
			sendArduinoCommand('h', h, s)
		}
	}
}

// sendArduinoCommand transmits a new command over the numonated serial port to the arduino. Returns an
// error on failure. Each command is identified by a single byte and may take one argument (a float).
func sendArduinoCommand(command byte, argument float32, serialPort io.ReadWriteCloser) error {
	if serialPort == nil {
		return nil
	}

	// Package argument for transmission
	bufOut := new(bytes.Buffer)
	err := binary.Write(bufOut, binary.LittleEndian, argument)
	if err != nil {
		return err
	}

	// Transmit command and argument down the pipe.
	for _, v := range [][]byte{[]byte{command}, bufOut.Bytes()} {
		_, err = serialPort.Write(v)
		if err != nil {
			return err
		}
	}

	return nil
}

// findArduino looks for the file that represents the arduino serial connection. Returns the fully qualified path
// to the device if we are able to find a likely candidate for an arduino, otherwise an empty string if unable to
// find an arduino device.
func findArduino() string {
	contents, _ := ioutil.ReadDir("/dev")

	// Look for the arduino device
	for _, f := range contents {
		if strings.Contains(f.Name(), "tty.usbserial") ||
			strings.Contains(f.Name(), "tty.usbmodem") ||
			strings.Contains(f.Name(), "ttyUSB") ||
			strings.Contains(f.Name(), "ttyACM") {
			return "/dev/" + f.Name()
		}
	}

	// Have not been able to find the device.
	return ""
}
