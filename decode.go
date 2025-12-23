// Copyright (c) 2025 Ted Dunning
// Licensed under the MIT License. See LICENSE file in the project root for details.

package main

import (
	"bufio"
	"ds1004-decoder/internal"
	"encoding/binary"
	"flag"
	"fmt"
	"os"
)

func main() {
	outputOption := flag.String("o", "", "output file")
	flag.Parse()
	fname := "usr_wf_data-2.bin"
	if flag.NArg() > 0 {
		fname = flag.Arg(0)
	}
	f, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var header internal.Header
	if err := binary.Read(f, binary.LittleEndian, &header); err != nil {
		panic(err)
	}

	outName := "data.csv"
	if *outputOption != "" {
		outName = *outputOption
	}
	fout, err := os.Create(outName)
	if err != nil {
		panic(err)
	}
	defer fout.Close()

	out := bufio.NewWriter(fout)

	channels := 0
	channelID := []int{}
	for i, on := range header.ChOn {
		if on != 0 {
			channelID = append(channelID, i)
			channels++
		}
	}
	channelData := make([][]uint8, channels)
	for i := range channelData {
		channelData[i] = make([]uint8, header.AnalogSamples)
		read, err := f.Read(channelData[i])
		if err != nil {
			panic(err)
		}
		if read != int(header.AnalogSamples) {
			panic(fmt.Sprintf("Read %d bytes instead of %d", read, header.AnalogSamples))
		}
	}
	_, _ = fmt.Fprintf(out, "time")
	i := 0
	for _, on := range header.ChOn {
		i++
		if on != 0 {
			_, _ = fmt.Fprintf(out, ",channel%d", i)
		}
	}
	_, _ = fmt.Fprintf(out, "\n")
	t := 0.0
	for i := range header.AnalogSamples {
		_, _ = fmt.Fprintf(out, "%g", t)
		for j := range channelData {
			offset := header.VerticalOffset[channelID[j]].Reduce()
			// there are 25 code values per division on the SDS-1204-EX
			scale := header.VoltsPerDivision[channelID[j]].Reduce() / 25
			v := float64(int(channelData[j][i])-128)*scale + offset
			_, _ = fmt.Fprintf(out, ",%g", v)
		}
		_, _ = fmt.Fprintf(out, "\n")
		t = float64(i) / header.AnalogSampleRate.Reduce()
	}
	out.Flush()
}
