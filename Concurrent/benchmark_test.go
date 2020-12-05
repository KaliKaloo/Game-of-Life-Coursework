package main

import (
	"fmt"
	"os"
	"testing"

	"uk.ac.bris.cs/gameoflife/gol"
)

const turnNum = 1000

var value bool

func Benchmark(b *testing.B) {
	tests := []gol.Params{
		{ImageWidth: 256, ImageHeight: 256, Threads: 1},
		{ImageWidth: 256, ImageHeight: 256, Threads: 2},
		{ImageWidth: 256, ImageHeight: 256, Threads: 3},
		{ImageWidth: 256, ImageHeight: 256, Threads: 4},
		{ImageWidth: 256, ImageHeight: 256, Threads: 5},
		{ImageWidth: 256, ImageHeight: 256, Threads: 6},
		{ImageWidth: 256, ImageHeight: 256, Threads: 7},
		{ImageWidth: 256, ImageHeight: 256, Threads: 8},
		{ImageWidth: 256, ImageHeight: 256, Threads: 9},
		{ImageWidth: 256, ImageHeight: 256, Threads: 10},
		{ImageWidth: 256, ImageHeight: 256, Threads: 11},
		{ImageWidth: 256, ImageHeight: 256, Threads: 12},
		{ImageWidth: 256, ImageHeight: 256, Threads: 13},
		{ImageWidth: 256, ImageHeight: 256, Threads: 14},
		{ImageWidth: 256, ImageHeight: 256, Threads: 15},
		{ImageWidth: 256, ImageHeight: 256, Threads: 16},
	}
	for _, t := range tests {
		value = true
		os.Stdout = nil
		t.Turns = turnNum
		name := fmt.Sprintf("%dx%dx%d-%d", t.ImageWidth, t.ImageHeight, turnNum, t.Threads)

		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				if value == true {
					events := make(chan gol.Event)
					gol.Run(t, events, nil)

					value = false

					for event := range events {
						switch e := event.(type) {
						case gol.FinalTurnComplete:
							fmt.Println("type ", e)
							value = true

						}
					}
				}
			}
		})
	}
}
