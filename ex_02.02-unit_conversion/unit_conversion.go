package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"

	"tgpl/ex_02.02-unit_conversion/metric"
)

func main() {
	if len(os.Args) > 1 {
		for _, in := range os.Args[1:] {
			num, _ := strconv.ParseFloat(in, 64)
			convert(num)
		}
	} else {
		in := bufio.NewScanner(os.Stdin)
		for in.Scan() {
			num, _ := strconv.ParseFloat(in.Text(), 64)
			convert(num)
		}
	}
}

func mmTo(w io.Writer, mm metric.Millimeter) {
	fmt.Fprintf(w, "(%v == %v)", mm, metric.MMToCM(mm))
	fmt.Fprintf(w, "(%v == %v)", mm, metric.MMToM(mm))
	fmt.Fprintf(w, "(%v == %v)", mm, metric.MMToKM(mm))
	fmt.Fprintf(w, "(%v == %v)", mm, metric.MMToIN(mm))
	fmt.Fprintf(w, "(%v == %v)", mm, metric.MMToFT(mm))
	fmt.Fprintf(w, "(%v == %v)", mm, metric.MMToYD(mm))
	fmt.Fprintf(w, "(%v == %v)", mm, metric.MMToMI(mm))
	fmt.Fprintf(w, "\n")
}

func cmTo(w io.Writer, cm metric.Centimeter) {
	fmt.Fprintf(w, "(%v == %v)", cm, metric.CMToMM(cm))
	fmt.Fprintf(w, "(%v == %v)", cm, metric.CMToM(cm))
	fmt.Fprintf(w, "(%v == %v)", cm, metric.CMToKM(cm))
	fmt.Fprintf(w, "(%v == %v)", cm, metric.CMToIN(cm))
	fmt.Fprintf(w, "(%v == %v)", cm, metric.CMToFT(cm))
	fmt.Fprintf(w, "(%v == %v)", cm, metric.CMToYD(cm))
	fmt.Fprintf(w, "(%v == %v)", cm, metric.CMToMI(cm))
	fmt.Fprintf(w, "\n")
}

func mTo(w io.Writer, m metric.Meter) {
	fmt.Fprintf(w, "(%v == %v)", m, metric.MToMM(m))
	fmt.Fprintf(w, "(%v == %v)", m, metric.MToCM(m))
	fmt.Fprintf(w, "(%v == %v)", m, metric.MToKM(m))
	fmt.Fprintf(w, "(%v == %v)", m, metric.MToIN(m))
	fmt.Fprintf(w, "(%v == %v)", m, metric.MToFT(m))
	fmt.Fprintf(w, "(%v == %v)", m, metric.MToYD(m))
	fmt.Fprintf(w, "(%v == %v)", m, metric.MToMI(m))
	fmt.Fprintf(w, "\n")
}

func kmTo(w io.Writer, km metric.Kilometer) {
	fmt.Fprintf(w, "(%v == %v)", km, metric.KMToMM(km))
	fmt.Fprintf(w, "(%v == %v)", km, metric.KMToCM(km))
	fmt.Fprintf(w, "(%v == %v)", km, metric.KMToM(km))
	fmt.Fprintf(w, "(%v == %v)", km, metric.KMToIN(km))
	fmt.Fprintf(w, "(%v == %v)", km, metric.KMToFT(km))
	fmt.Fprintf(w, "(%v == %v)", km, metric.KMToYD(km))
	fmt.Fprintf(w, "(%v == %v)", km, metric.KMToMI(km))
	fmt.Fprintf(w, "\n")
}

func inTo(w io.Writer, in metric.Inch) {
	fmt.Fprintf(w, "(%v == %v)", in, metric.INToMM(in))
	fmt.Fprintf(w, "(%v == %v)", in, metric.INToCM(in))
	fmt.Fprintf(w, "(%v == %v)", in, metric.INToM(in))
	fmt.Fprintf(w, "(%v == %v)", in, metric.INToKM(in))
	fmt.Fprintf(w, "(%v == %v)", in, metric.INToFT(in))
	fmt.Fprintf(w, "(%v == %v)", in, metric.INToYD(in))
	fmt.Fprintf(w, "(%v == %v)", in, metric.INToMI(in))
	fmt.Fprintf(w, "\n")
}

func ftTo(w io.Writer, ft metric.Foot) {
	fmt.Fprintf(w, "(%v == %v)", ft, metric.FTToMM(ft))
	fmt.Fprintf(w, "(%v == %v)", ft, metric.FTToCM(ft))
	fmt.Fprintf(w, "(%v == %v)", ft, metric.FTToM(ft))
	fmt.Fprintf(w, "(%v == %v)", ft, metric.FTToKM(ft))
	fmt.Fprintf(w, "(%v == %v)", ft, metric.FTToIN(ft))
	fmt.Fprintf(w, "(%v == %v)", ft, metric.FTToYD(ft))
	fmt.Fprintf(w, "(%v == %v)", ft, metric.FTToMI(ft))
	fmt.Fprintf(w, "\n")
}

func ydTo(w io.Writer, yd metric.Yard) {
	fmt.Fprintf(w, "(%v == %v)", yd, metric.YDToMM(yd))
	fmt.Fprintf(w, "(%v == %v)", yd, metric.YDToCM(yd))
	fmt.Fprintf(w, "(%v == %v)", yd, metric.YDToM(yd))
	fmt.Fprintf(w, "(%v == %v)", yd, metric.YDToKM(yd))
	fmt.Fprintf(w, "(%v == %v)", yd, metric.YDToIN(yd))
	fmt.Fprintf(w, "(%v == %v)", yd, metric.YDToFT(yd))
	fmt.Fprintf(w, "(%v == %v)", yd, metric.YDToMI(yd))
	fmt.Fprintf(w, "\n")
}

func miTo(w io.Writer, mi metric.Mile) {
	fmt.Fprintf(w, "(%v == %v)", mi, metric.MIToMM(mi))
	fmt.Fprintf(w, "(%v == %v)", mi, metric.MIToCM(mi))
	fmt.Fprintf(w, "(%v == %v)", mi, metric.MIToM(mi))
	fmt.Fprintf(w, "(%v == %v)", mi, metric.MIToKM(mi))
	fmt.Fprintf(w, "(%v == %v)", mi, metric.MIToIN(mi))
	fmt.Fprintf(w, "(%v == %v)", mi, metric.MIToFT(mi))
	fmt.Fprintf(w, "(%v == %v)", mi, metric.MIToYD(mi))
	fmt.Fprintf(w, "\n")
}

func convert(v float64) {
	mmTo(os.Stdout, metric.Millimeter(v))
	cmTo(os.Stdout, metric.Centimeter(v))
	mTo(os.Stdout, metric.Meter(v))
	kmTo(os.Stdout, metric.Kilometer(v))
	inTo(os.Stdout, metric.Inch(v))
	ftTo(os.Stdout, metric.Foot(v))
	ydTo(os.Stdout, metric.Yard(v))
	miTo(os.Stdout, metric.Mile(v))
}
