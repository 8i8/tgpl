package eval

import "testing"

func TestRingBufferLen(t *testing.T) {
	const fname = "TestRingBufferLen"
	buf := NewBuffer(3)
	buf.Add("one")
	buf.Add("two")
	buf.Add("three")
	buf.MoveUp("four")
	calc := buf.List()
	str := calc.String()
	exp := "two\nthree\nfour\n"
	if str != exp {
		t.Errorf("%s: want %q got %q", fname, exp, str)
	}
}

func TestRingbufferAdd(t *testing.T) {
	const fname = "TestRingBufferAdd"
	buf := NewBuffer()
	buf.Add("one")
	buf.Add("two")
	buf.Add("three")
	buf.Add("four")
	buf.Add("five")
	buf.Add("six")
	calc := buf.List()
	str := calc.String()
	exp := "two\nthree\nfour\nfive\nsix\n"
	if str != exp {
		t.Errorf("%s: want %q got %q", fname, exp, str)
	}
}

func TestRingbufferMoveUpNotThere(t *testing.T) {
	const fname = "TestRingbufferMoveUp"
	buf := NewBuffer()
	buf.Add("one")
	buf.Add("two")
	buf.Add("three")
	buf.Add("four")
	buf.Add("five")
	buf.MoveUp("six")
	calc := buf.List()
	str := calc.String()
	exp := "two\nthree\nfour\nfive\nsix\n"
	if str != exp {
		t.Errorf("%s: want %q got %q", fname, exp, str)
	}
}

func TestRingbufferMoveUpEqual(t *testing.T) {
	const fname = "TestRingbufferMoveUp"
	buf := NewBuffer()
	buf.Add("one")
	buf.Add("two")
	buf.Add("three")
	buf.Add("four")
	buf.Add("five")
	buf.Add("six")
	buf.MoveUp("six")
	calc := buf.List()
	str := calc.String()
	exp := "two\nthree\nfour\nfive\nsix\n"
	if str != exp {
		t.Errorf("%s: want %q got %q", fname, exp, str)
	}
}

func TestRingbufferMoveUpUnder(t *testing.T) {
	const fname = "TestRingbufferMoveUp"
	buf := NewBuffer()
	buf.Add("one")
	buf.Add("two")
	buf.Add("three")
	buf.Add("four")
	buf.Add("five")
	buf.Add("six")
	buf.MoveUp("four")
	calc := buf.List()
	str := calc.String()
	exp := "two\nthree\nfive\nsix\nfour\n"
	if str != exp {
		t.Errorf("%s: want %q got %q", fname, exp, str)
	}
}

func TestRingbufferMoveUpOver(t *testing.T) {
	const fname = "TestRingbufferMoveUp"
	buf := NewBuffer()
	buf.Add("one")
	buf.Add("two")
	buf.Add("three")
	buf.Add("four")
	buf.Add("five")
	buf.Add("six")
	buf.Add("seven")
	buf.Add("eight")
	buf.MoveUp("six")
	calc := buf.List()
	str := calc.String()
	exp := "four\nfive\nseven\neight\nsix\n"
	if str != exp {
		t.Errorf("%s: want %q got %q", fname, exp, str)
	}
}

func TestRingBufferLoadUnload(t *testing.T) {
	const fname = "TestRingBufferLoadUnload"
	buf := NewBuffer()
	str := "Zm91ciYmZml2ZSYmc2V2ZW4mJmVpZ2h0JiZzaXg"
	err := buf.Load(str)
	if err != nil {
		t.Errorf("%s: want nil got %q", fname, err)
	}
	exp := buf.Unload()
	if exp != str {
		t.Errorf("%s: want %q got %q", fname, str, exp)
	}
}

func TestRingBufferLoad(t *testing.T) {
	const fname = "TestRingbufferMoveUp"
	buf := NewBuffer()
	buf.Add("one")
	buf.Add("two")
	buf.Add("three")
	buf.Add("four")
	buf.Add("five")
	buf.Add("six")
	buf.Add("seven")
	buf.Add("eight")
	buf.MoveUp("six")
	str := buf.Unload()
	exp := "four\nfive\nseven\neight\nsix\n"
	exp1 := "Zm91ciYmZml2ZSYmc2V2ZW4mJmVpZ2h0JiZzaXg"
	if str != exp1 {
		t.Errorf("%s: want %q got %q", fname, exp, str)
	}
	buf.Reset()
	err := buf.Load(exp1)
	if err != nil {
		t.Errorf("%s: want nil got %q", fname, err)
	}
	exp2 := buf.Unload()
	if exp2 != exp1 {
		t.Errorf("%s: want %q got %q", fname, exp, exp2)
	}
	calc := buf.List()
	str = calc.String()
	exp = "four\nfive\nseven\neight\nsix\n"
	if str != exp {
		t.Errorf("%s: want %q got %q", fname, exp, str)
	}
}

func TestRingBufferAddAfterLoad(t *testing.T) {
	const fname = "TestRingBufferAddAfterLoad"
	buf := NewBuffer()
	str := "cGxvdChwb3coMixzaW4oeSkpKnBvdygyLHNpbih4KSkvMTIpJiZwbG" +
		"90KHNpbigteCkqcG93KDEuNSwtcikpJiZwbG90KHNpbihyKS9yKSY" +
		"mcGxvdChzaW4oeCp5LzEwKS8xMCk"
	err := buf.Load(str)
	if err != nil {
		t.Errorf("%s: want nil got %q", fname, err)
	}
	buf.MoveUp("one")
	calc := buf.List()
	str = calc.String()
	exp := "plot(pow(2,sin(y))*pow(2,sin(x))/12)\n" +
		"plot(sin(-x)*pow(1.5,-r))\n" +
		"plot(sin(r)/r)\n" +
		"plot(sin(x*y/10)/10)\n" +
		"one\n"
	if str != exp {
		t.Errorf("%s: want %q got %q", fname, exp, str)
	}
}
