package logic_01

import (
	"fmt"
	"testing"
)

func TestSoal01ulangan8(t *testing.T) {
	n := 10
	angka := 1
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(angka, "\t")
			angka++
		} else {
			fmt.Print(3, "\t")
		}
	}
}

func TestSoal02ulangan8(t *testing.T) {
	n := 10
	angka := 1
	for i := 0; i < n; i++ {
		if i%2 == 1 {
			fmt.Print(angka, "\t")
			angka++
		} else {
			fmt.Print(3, "\t")

		}
	}
}

func TestSoal03ulangan8(t *testing.T) {
	n := 10
	angka := 1
	x := 99
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(angka, "\t")
			angka++
		} else {
			fmt.Print(x, "\t")
		}
	}
}

func TestSoal04ulangan8(t *testing.T) {
	n := 10
	angka := 1
	x := 777
	for i := 0; i < n; i++ {
		if i%2 == 1 {
			fmt.Print(angka, "\t")
			angka++
		} else {
			fmt.Print(x, "\t")
		}
	}
}

func TestSoal05ulangan8(t *testing.T) {
	n := 15
	angka := 1
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Print("fizz,buzz", "\t")
		} else if i%3 == 0 {
			fmt.Print("fizz", "\t")
			angka++
		} else if i%5 == 0 {
			fmt.Print("buzz", "\t")
			angka++
		} else {
			fmt.Print(angka, "\t")
			angka++
		}
	}
}

func TestSoal6ulangan8(t *testing.T) {
	n := 15
	angkaKiri := 1
	angkaKuadrat := 1
	angkaKanan := 0
	for i := 1; i <= n; i++ {
		if i%4 == 0 {
			angkaKuadrat = i * i
			fmt.Print(angkaKuadrat, "\t")
		} else if i <= 7 {
			angkaKiri = i * 3
			fmt.Print(angkaKiri, "\t")
		} else {
			angkaKanan = angkaKiri
			fmt.Print(angkaKanan, "\t")
			angkaKanan = angkaKiri - 3
			angkaKiri = angkaKiri - 3
		}
	}
}

func TestSoal09ulangan8(t *testing.T) {
	n := 12
	a := 1
	b := 2
	c := 3

	for i := 1; i < +n; i++ {
		if i%3 == 2 {
			fmt.Print(b, "\t")
			b *= 2
		} else if i%3 == 0 {
			fmt.Print(c, "\t")
			c *= 3
		} else {
			fmt.Print(a, "\t")
		}
	}
}

func TestSoal10ulangan8(t *testing.T) {
	n := 12
	a := 1
	b := 2
	c := 3

	for i := 1; i <= n; i++ {
		if i%4 == 1 {
			fmt.Print(c, "\t")
			c += 3
		} else if i%4 == 2 {
			fmt.Print(b, "\t")
			b += 2
		} else if i%4 == 3 {
			fmt.Print(a, "\t")
			a += 1
		} else {
			fmt.Print(999, "\t")
		}
	}
}
