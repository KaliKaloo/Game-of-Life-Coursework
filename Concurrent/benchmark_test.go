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

		{ImageWidth: 512, ImageHeight: 512, Threads: 1},
		{ImageWidth: 512, ImageHeight: 512, Threads: 2},
		{ImageWidth: 512, ImageHeight: 512, Threads: 3},
		{ImageWidth: 512, ImageHeight: 512, Threads: 4},
		{ImageWidth: 512, ImageHeight: 512, Threads: 5},
		{ImageWidth: 512, ImageHeight: 512, Threads: 6},
		{ImageWidth: 512, ImageHeight: 512, Threads: 7},
		{ImageWidth: 512, ImageHeight: 512, Threads: 8},
		{ImageWidth: 512, ImageHeight: 512, Threads: 9},
		{ImageWidth: 512, ImageHeight: 512, Threads: 10},
		{ImageWidth: 512, ImageHeight: 512, Threads: 11},
		{ImageWidth: 512, ImageHeight: 512, Threads: 12},
		{ImageWidth: 512, ImageHeight: 512, Threads: 13},
		{ImageWidth: 512, ImageHeight: 512, Threads: 14},
		{ImageWidth: 512, ImageHeight: 512, Threads: 15},
		{ImageWidth: 512, ImageHeight: 512, Threads: 16},
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
