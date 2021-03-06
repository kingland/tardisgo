// THIS IS NOT PRETTY, IT IS A WORK-IN-PROGRESS

// NOTE : No Output = success

// TODO separate this jumble of tests into a set of smaller ones

// This package should only test the core language functionality, all standard package tests moved elsewhere
package main

import (
	"errors"
	"fmt"
	"runtime"
	"unicode"
	"unicode/utf8"
	"unsafe"

	//"haxegoruntime"

	"github.com/tardisgo/tardisgo/haxe/hx"
)

//const tardisgoLibRuntimePath = "github.com/tardisgo/tardisgo/golibruntime"

const tardisgoHeader = "/* TARDIS Go general header*/"

const tardisgoHaxeHeader = `// Haxe specific header for each file
`

const ShowKnownErrors = false

func loc(l string) string {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return "???"
	}
	return file + ":" + hx.CallString("", "Std.string", 1, line) + " " + l
}

func TEQ(l string, a, b interface{}) bool {
	l = loc(l)
	if a != b {
		fmt.Println("TEQ error " + l + " ")
		fmt.Println(a)
		fmt.Println(b)
		return false
	}
	return true
}

func TEQuint64(l string, a, b uint64) bool {
	l = loc(l)
	if a != b {
		fmt.Println("TEQui64 error " + l + " ")
		fmt.Println("high a", uint(a>>32))
		fmt.Println("low a", uint(a&0xFFFFFFFF))
		fmt.Println("high b", uint(b>>32))
		fmt.Println("low b", uint(b&0xFFFFFFFF))
		return false
	}
	return true
}
func TEQint64(l string, a, b int64) bool {
	l = loc(l)
	if a != b {
		fmt.Println("TEQi64 error " + l + " ")
		fmt.Println("high a", int(a>>32))
		fmt.Println("low a", int(a&0xFFFFFFFF))
		fmt.Println("high b", int(b>>32))
		fmt.Println("low b", int(b&0xFFFFFFFF))
		return false
	}
	return true
}
func TEQuint32(l string, a, b uint32) bool {
	l = loc(l)
	if a != b {
		fmt.Println("TEQui32 error " + l + " ")
		fmt.Println(a)
		fmt.Println(b)
		return false
	}
	return true
}
func TEQint32(l string, a, b int32) bool {
	l = loc(l)
	if a != b {
		fmt.Println("TEQi32 error " + l + " ")
		fmt.Println(a)
		fmt.Println(b)
		return false
	}
	return true
}
func TEQbyteSlice(l string, a, b []byte) bool {
	l = loc(l)
	if len(a) != len(b) {
		fmt.Println("TEQbyteSlice error "+l+" ", a, b)
		return false
	}
	ret := true
	for i := range a {
		if a[i] != b[i] {
			fmt.Println("TEQbyteSlice error "+l+" ", a, b)
			ret = false
		}
	}
	return ret
}
func TEQruneSlice(l string, a, b []rune) bool {
	l = loc(l)
	if len(a) != len(b) {
		fmt.Println("TEQruneSlice error "+l+" ", a, b)
		return false
	}
	ret := true
	for i := range a {
		if a[i] != b[i] {
			fmt.Println("TEQruneSlice error "+l+" ", a, b)
			ret = false
		}
	}
	return ret
}
func TEQintSlice(l string, a, b []int) bool {
	l = loc(l)
	//fmt.Println("TEQintSlice DEBUG: " + l + " ")
	if len(a) != len(b) {
		fmt.Println("TEQintSlice error "+l+" ", a, b)
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			fmt.Println("TEQintSlice error "+l+" ", a, b)
			return false
		}
	}
	return true
}
func TEQfloat(l string, a, b, maxDif float64) bool {
	l = loc(l)
	if a == b {
		return true
	}
	dif := a - b
	if dif < 0 {
		dif = -dif
	}
	if dif > maxDif {
		fmt.Println("TEQfloat error " + l + " ")
		fmt.Println(a)
		fmt.Println(b)
		return false
	}
	return true
}

// CONSTANT TEST DATA
const Name string = "this is my name"
const ests bool = true
const Pi float64 = 3.14159265358979323846
const zero = 0.0 // untyped floating-point constant
const (
	size int = 1024
	eof      = -1 // untyped integer constant
)
const a, b, c = 3, 4, "foo" // a = 3, b = 4, c = "foo", untyped integer and string constants
const u, v float64 = 0, 3   // u = 0.0, v = 3.0
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Partyday
	numberOfDays // this constant is not exported
)
const ( // iota is reset to 0
	c0 = iota // c0 == 0
	c1 = iota // c1 == 1
	c2 = iota // c2 == 2
)
const (
	_a = 1 << iota // a == 1 (iota has been reset)
	_b = 1 << iota // b == 2
	_c = 1 << iota // c == 4
)
const (
	_u         = iota * 42 // u == 0     (untyped integer constant)
	_v float64 = iota * 42 // v == 42.0  (float64 constant)
	_w         = iota * 42 // w == 84    (untyped integer constant)
)
const _x = iota // x == 0 (iota has been reset)
const _y = iota // y == 0 (iota has been reset)
const (
	bit0, mask0 = 1 << iota, 1<<iota - 1 // bit0 == 1, mask0 == 0
	bit1, mask1                          // bit1 == 2, mask1 == 1
	_, _                                 // skips iota == 2
	bit3, mask3                          // bit3 == 8, mask3 == 7
)
const ren = '人'
const Θ float64 = 3 / 2  // Θ == 1.0   (type float64, 3/2 is integer division)
const Π float64 = 3 / 2. // Π == 1.5   (type float64, 3/2. is float division)
const d = 1 << 3.0       // d == 8     (untyped integer constant)
const e = 1.0 << 3       // e == 8     (untyped integer constant)
const h = "foo" > "bar"  // h == true  (untyped boolean constant)
const j = true           // j == true  (untyped boolean constant)
const k = 'w' + 1        // k == 'x'   (untyped rune constant)
const l = "hi"           // l == "hi"  (untyped string constant)
const m = string(k)      // m == "x"   (type string)

func testConst() {
	TEQ("", Name, "this is my name")
	TEQ("", ests, true)
	TEQfloat("", Pi, 3.14159265358979323846, 0.00000000000001)
	TEQ("", zero, 0.0) // untyped floating-point constant
	TEQ("", size, 1024)
	TEQ("", eof, -1) // untyped integer constant
	// a = 3, b = 4, c = "foo", untyped integer and string constants
	TEQ("", a, 3)
	TEQ("", b, 4)
	TEQ("", c, "foo")
	// u = 0.0, v = 3.0
	TEQ("", u, 0.0)
	TEQ("", v, 3.0)
	TEQ("", Sunday, 0)
	TEQ("", Monday, 1)
	TEQ("", Tuesday, 2)
	TEQ("", Wednesday, 3)
	TEQ("", Thursday, 4)
	TEQ("", Friday, 5)
	TEQ("", Partyday, 6)
	TEQ("", numberOfDays, 7) // this constant is not exported
	TEQ("", c0, 0)           // c0 == 0
	TEQ("", c1, 1)           // c1 == 1
	TEQ("", c2, 2)           // c2 == 2
	TEQ("", _a, 1)           // a == 1 (iota has been reset)
	TEQ("", _b, 2)           // b == 2
	TEQ("", _c, 4)           // c == 4
	TEQ("", _u, 0)           // u == 0     (untyped integer constant)
	TEQ("", _v, 42.0)        // v == 42.0  (float64 constant)
	TEQ("", _w, 84)          // w == 84    (untyped integer constant)
	TEQ("", _x, 0)           // x == 0 (iota has been reset)
	TEQ("", _y, 0)           // y == 0 (iota has been reset)
	TEQ("", bit0, 1)
	TEQ("", mask0, 0) // bit0 == 1, mask0 == 0
	TEQ("", bit1, 2)
	TEQ("", mask1, 1) // bit1 == 2, mask1 == 1
	//_, _                                 // skips iota == 2
	TEQ("", bit3, 8)
	TEQ("", mask3, 7) // bit3 == 8, mask3 == 7
	TEQ("", ren, '人')
	TEQ("", Θ, 1.0)  // Θ == 1.0   (type float64, 3/2 is integer division)
	TEQ("", Π, 1.5)  // Π == 1.5   (type float64, 3/2. is float division)
	TEQ("", d, 8)    // d == 8     (untyped integer constant)
	TEQ("", e, 8)    // e == 8     (untyped integer constant)
	TEQ("", h, true) // h == true  (untyped boolean constant)
	TEQ("", j, true) // j == true  (untyped boolean constant)
	TEQ("", k, 'x')  // k == 'x'   (untyped rune constant)
	TEQ("", l, "hi") // l == "hi"  (untyped string constant)
	TEQ("", m, "x")  // m == "x"   (type string)
}

var testUTFlength = "123456789"

func testUTF() {
	var (
		rA, rB, r  []rune
		uS, s1, s2 string
	)
	rA = []rune{0x767d, 0x9d6c, 0x7fd4}
	uS = string(rA) // "\u767d\u9d6c\u7fd4" == "白鵬翔"
	rB = []rune(uS)
	TEQruneSlice("", rA, rB)

	s1 = "香港发生工厂班车砍人案12人受伤"
	r = []rune(s1)
	s2 = string(r)
	TEQ("", s1, s2)

	TEQ("", len(s1), 44)
	TEQ("", len(testUTFlength), 9)

	hellø := "hellø"
	TEQ("", string([]byte{'h', 'e', 'l', 'l', '\xc3', '\xb8'}), hellø)
	TEQbyteSlice("", []byte("hellø"), []byte{'h', 'e', 'l', 'l', '\xc3', '\xb8'})

	TEQ("", "ø", hellø[4:])
}

var TestInit = "init() ran OK"
var primes = [6]int{2, 3, 5, 7, 9, 2147483647}
var iFace interface{} = nil

func testInit() {
	TEQ("", TestInit, "init() ran OK")
	TEQintSlice("", primes[:], []int{2, 3, 5, 7, 9, 2147483647})
	TEQ("", 9, primes[4]) // also testing array access with a constant index
	TEQ("", nil, iFace)
}

var PublicStruct struct {
	a int
	b bool
	c string
	d float64
	e interface{}
	f [12]int
	g [6]string
	h [14]struct {
		x bool
		y [3]float64
		z [6]interface{}
	}
}

func testStruct() {
	var PrivateStruct struct {
		a int
		b bool
		c string
		d float64
		e interface{}
		f [12]int
		g [6]string
		h [14]struct {
			x bool
			y [3]float64
			z [6]interface{}
		}
	}
	// check that everything is equally initialized
	TEQ("", PublicStruct.a, PrivateStruct.a)
	TEQ("", PublicStruct.b, PrivateStruct.b)
	TEQ("", PublicStruct.c, PrivateStruct.c)
	TEQfloat("", PublicStruct.d, PrivateStruct.d, 0.01)
	TEQ("", PublicStruct.e, PrivateStruct.e)
	//fmt.Println("", PublicStruct.f[:], PrivateStruct.f[:])
	TEQintSlice("", PublicStruct.f[:], PrivateStruct.f[:])
	PublicStruct.a = 42
	PrivateStruct.a = 42
	TEQ("", PublicStruct.a, PrivateStruct.a)
	PublicStruct.c = Name
	PrivateStruct.c = Name
	TEQ("", PublicStruct.c, PrivateStruct.c)
	for i := range PrivateStruct.h {
		for j := range PrivateStruct.h[i].y {
			PrivateStruct.h[i].y[j] = 42.0 * float64(i) * float64(j)
			PublicStruct.h[i].y[j] = 42.0 * float64(i) * float64(j)
			TEQfloat("", PrivateStruct.h[i].y[j], PublicStruct.h[i].y[j], 1.0)
		}
	}
}
func Sqrt(x float64) float64 {
	z := x
	for i := 0; i < 1000; i++ {
		z -= (z*z - x) / (2.0 * x)
	}
	return z
}

func testFloat() { // and also slices!
	TEQfloat("", Sqrt(1024), 32.0, 0.1)
	threeD := make([][][]float64, 10)
	for i := range threeD {
		threeD[i] = make([][]float64, 10)
		for j := range threeD[i] {
			threeD[i][j] = make([]float64, 10)
			for k := range threeD[i][j] {
				threeD[i][j][k] = float64(i) * float64(j) * float64(k)
				TEQfloat("", threeD[i][j][k], float64(i)*float64(j)*float64(k), 0.1)
			}
		}
	}
	// TODO add more here
}

func noCaller() float64 { // this should be removed by a good target compiler...
	U_ := Sqrt(float64(64))
	return U_
}

//var aPtr *int // TODO this should generate an error

func twoRets(x int) (a int, b string) {
	return 42 * x, "forty-two"
}

func testMultiRet() {
	r1, r2 := twoRets(1)
	TEQ("", r1, 42)
	TEQ("", r2, "forty-two")
}

func testAppend() {
	s0 := []int{0, 0}
	s1 := append(s0, 2) // append a single element     s1 == []int{0, 0, 2}
	TEQintSlice("", []int{0, 0, 2}, s1)
	s2 := append(s1, 3, 5, 7) // append multiple elements    s2 == []int{0, 0, 2, 3, 5, 7}
	TEQintSlice("", []int{0, 0, 2, 3, 5, 7}, s2)
	s3 := append(s2, s0...) // append a slice              s3 == []int{0, 0, 2, 3, 5, 7, 0, 0}
	TEQintSlice("", []int{0, 0, 2, 3, 5, 7, 0, 0}, s3)
	var t []interface{}
	t = append(t, 42, 3.1415, "foo", nil) //       				 t == []interface{}{42, 3.1415, "foo"}
	TEQ("", t[0], 42)
	TEQ("", t[1], 3.1415)
	TEQ("", t[2], "foo")
	TEQ("", t[3], nil)

	var b []byte
	b = append(b, "bar"...)
	TEQbyteSlice("", b, []byte{'b', 'a', 'r'})
}

func testHeader() {
	// not sure how to test this
}

func testCopy() {
	var a = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	var s = make([]int, 6)
	n1 := copy(s, a[0:]) // n1 == 6, s == []int{0, 1, 2, 3, 4, 5}
	TEQ("", n1, 6)
	TEQintSlice("", s, []int{0, 1, 2, 3, 4, 5})
	n2 := copy(s, s[2:]) // n2 == 4, s == []int{2, 3, 4, 5, 4, 5}
	TEQ("", n2, 4)
	TEQintSlice("", s, []int{2, 3, 4, 5, 4, 5})
	var b = make([]byte, 5)
	n3 := copy(b, "Hello, World!") // n3 == 5, b == []byte("Hello")
	TEQ("", n3, 5)
	TEQbyteSlice("", b, []byte("Hello"))
}

func testInFuncPtr() { // there is no way to stop this use of pointers...
	var ss = 12
	var ssa = &ss
	TEQ("", *ssa, 12)
}

func testCallByValue(a struct{ b int }, x [10]int, y []int, z int) {
	a.b = 42
	x[0] = 43
	y[0] = 44
	z = 45
}

func testCallByReference(a *struct{ b int }, x *[10]int, y []int, z *int) {
	a.b = 46
	x[0] = 47
	y[0] = 48
	*z = 49
}
func testTweakFloatByReference(i *float64) {
	if *i == 0 {
		*i = 0
	} else {
		*i = Sqrt(*i)
	}
}
func testCallBy() {
	var a struct {
		b int
	}
	var x [10]int
	var y []int = make([]int, 1)
	var z int
	testCallByValue(a, x, y, z)
	TEQ("", a.b, 0)
	TEQ("", x[0], 0)
	TEQ("", y[0], 44)
	TEQ("", z, 0)

	testCallByReference(&a, &x, y, &z)
	TEQ("", a.b, 46)
	TEQ("", x[0], 47)
	TEQ("", y[0], 48)
	TEQ("", z, 49)

	var xx [10]float64
	for i := range x {
		xx[i] = float64(i * i)
		testTweakFloatByReference(&xx[i])
		TEQfloat("", xx[i], float64(i), 0.1)
	}
}

func testMap() { // and map-like constucts
	// vowels[ch] is true if ch is a vowel
	vowels := [128]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'y': true}
	for k, v := range vowels {
		switch k {
		case 'a', 'e', 'i', 'o', 'u', 'y':
			TEQ("", true, v)
		default:
			TEQ("", false, v)
		}
	}

	filter := [10]float64{-1, 4: -0.1, -0.1, 9: -1}
	TEQfloat("", filter[5], -0.1, 0.01)

	// frequencies in Hz for equal-tempered scale (A4 = 440Hz)
	noteFrequency := map[string]float64{
		"C0": 16.35, "D0": 18.35, "E0": 20.60, "F0": 21.83,
		"G0": 24.50, "A0": 27.50, "B0": 30.87,
	}
	noteFrequency["Test"] = 42.42
	TEQ("", len(noteFrequency), 8)
	for k, v := range noteFrequency {
		r := 0.0
		switch k {
		case "C0":
			r = 16.35
		case "D0":
			r = 18.35
		case "E0":
			r = 20.60
		case "F0":
			r = 21.83
		case "G0":
			r = 24.50
		case "A0":
			r = 27.50
		case "B0":
			r = 30.87
		case "Test":
			r = 42.42
		default:
			r = -1
		}
		if !TEQfloat(""+" Value itterator in map", v, r, 0.01) {
			break
		}
	}
	x, isok := noteFrequency["Test"]
	TEQfloat("", 42.42, x, 0.01)
	TEQ("", true, isok)
	_, notok := noteFrequency["notHere"]
	TEQ("", false, notok)
	delete(noteFrequency, "Test")
	_, isok = noteFrequency["Test"]
	TEQ("", false, isok)

	if true { // just to get a new scope
		// frequencies in Hz for equal-tempered scale (A4 = 440Hz)
		noteFrequency2 := map[float64]string{
			16.35: "C0", 18.35: "D0", 20.60: "E0", 21.83: "F0",
			24.50: "G0", 27.50: "A0", 30.87: "B0",
		}
		noteFrequency2[42.42] = "Test"
		TEQ("", len(noteFrequency2), 8)
		for k, v := range noteFrequency2 {
			r := ""
			switch k {
			case 16.35:
				r = "C0"
			case 18.35:
				r = "D0"
			case 20.60:
				r = "E0"
			case 21.83:
				r = "F0"
			case 24.50:
				r = "G0"
			case 27.50:
				r = "A0"
			case 30.87:
				r = "B0"
			case 42.42:
				r = "Test"
			default:
				r = "NOT FOUND"
			}
			if !TEQ(""+" Value itterator in map", v, r) {
				break
			}
		}
		x, isok := noteFrequency2[42.42]
		TEQ("", "Test", x)
		TEQ("", true, isok)
		_, notok := noteFrequency2[-42]
		TEQ("", false, notok)
		delete(noteFrequency2, 42.42)
		_, isok = noteFrequency2[43.42]
		TEQ("", false, isok)
	}
}

type MyFloat float64
type MyFloat2 MyFloat

var namedGlobal MyFloat

type IntArray [8]int

type (
	Point struct {
		x, y float64
	}
	Polar Point
)

var myPolar Polar

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (f MyFloat) UncalledMethod() {
	panic("Why are we here?")
}

func (mf *MyFloat) set42() {
	*mf = 42
}

// this is required as Go does not allow the MyFloat2 type to inheret MyFloat.Abs
func (f MyFloat2) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
func (f MyFloat2) Scale(x float64) float64 {
	return float64(f) * x
}

func (ia *IntArray) set42() {
	for i := range ia {
		ia[i] = 42
	}
}
func (p Polar) BearVal() bool {
	return p.x == p.y
}

// from the language spec section Method Values
type T struct {
	a int
}

func (tv T) Mv(a int) int          { return a } // value receiver
func (tp *T) Mp(f float32) float32 { return f } // pointer receiver

var t T
var pt *T

func testNamed() {
	var ia IntArray
	for i := range ia {
		ia[i] = i
	}
	TEQintSlice("", ia[:], []int{0, 1, 2, 3, 4, 5, 6, 7})
	var namedLocal MyFloat = 41.42
	namedGlobal = 42.42
	namedLocal += 1.0
	TEQfloat("", float64(namedGlobal), float64(namedLocal), 0.0002)
	myPolar.x = 11.11
	myPolar.y = 10.11
	myPolar.y++
	TEQfloat("", float64(myPolar.x), float64(myPolar.y), 0.0002)
	// method expression tests...
	TEQ("", myPolar.BearVal(), true)
	f := MyFloat(-555)
	g := MyFloat2(-555)
	TEQfloat("", f.Abs(), g.Abs(), 0.0002)
	fi := Abser(f)
	gi := Abser(g)
	TEQfloat("", fi.Abs(), gi.Abs(), 0.0002)

	ia.set42()
	f.set42()
	TEQfloat("", float64(ia[3]), float64(f), 0.0002)

	// from the language spec section on method values (requires ssa.MakeClosure instruction)
	f1 := t.Mv
	TEQ("", f1(7), t.Mv(7))
	pt = &t
	f2 := pt.Mp
	TEQ("", f2(7), pt.Mp(7))
	f3 := pt.Mv
	TEQ("", f3(7), (*pt).Mv(7))
	f4 := t.Mp
	TEQ("", f4(7), (&t).Mp(7))

	// more from the language spec on Method expressions

	TEQ("", t.Mv(7), T.Mv(t, 7))
	TEQ("", t.Mv(7), (T).Mv(t, 7))

	f1a := T.Mv
	TEQ("", t.Mv(7), f1a(t, 7))
	f2a := (T).Mv
	TEQ("", t.Mv(7), f2a(t, 7))

}

var hypot1 = func(x, y float64) float64 {
	return Sqrt(x*x + y*y)
}

func testFuncPtr() {
	var hypot2 = func(x, y float64) float64 {
		return Sqrt(x*x + y*y)
	}
	TEQfloat("", hypot1(3, 4), hypot2(3, 4), 0.2)
}

var int64_max int64 = 0x7FFFFFFFFFFFFFFF
var int32_max int32 = 0x7FFFFFFF
var int16_max int16 = 0x7FFF
var int8_max int8 = 0x7F
var uint64_max uint64 = 0xFFFFFFFFFFFFFFFF
var uint32_max uint32 = 0xFFFFFFFF // This value too big and too ambiguous for cpp when held as an Int...
var uint16_max uint16 = 0xFFFF
var uint8_max uint8 = 0xFF
var int8_mostNeg int8 = -128
var int16_mostNeg int16 = -32768
var int32_mostNeg int32 = -2147483648
var int64_mostNeg int64 = -9223372036854775808

var five int = 5
var three int = 3

var uint64Global uint64
var uint64GlobalArray [4]uint64

func testIntOverflow() { //TODO add int64
	TEQ(""+" int16 overflow test 1", int16_max+1, int16_mostNeg)
	TEQ(""+" int8 overflow test 1", int8_max+1, int8_mostNeg)
	TEQ(""+" uint16 overflow test 2", uint16(uint16_max+1), uint16(0))
	TEQ(""+" uint8 overflow test 2", uint8(uint8_max+1), uint8(0))
	TEQ(""+" int8 overflow test 3", int8(int8_mostNeg-1), int8_max)
	TEQ(""+" int16 overflow test 3", int16(int16_mostNeg-1), int16_max)

	TEQint64(""+" int64 overflow test 1 ", int64_max+1, int64_mostNeg)
	TEQint32(""+" int32 overflow test 1 ", int32_max+1, int32_mostNeg)
	TEQuint64(""+" uint64 overflow test 2 ", uint64(uint64_max+1), uint64(0))
	TEQuint32(""+" uint32 overflow test 2 ", uint32(uint32_max+1), uint32(0))

	TEQint32(""+" int32 overflow test 3 ", int32(int32_mostNeg-int32(1)), int32_max)
	TEQint64(""+" int64 overflow test 3 ", int64(int64_mostNeg-int64(1)), int64_max)

	//Math.imul test case at https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Math/imul
	TEQuint32("", ((uint32_max << 1) * 5), uint32_max+1-10)
	TEQuint32("", (((uint32_max << 1) + 1) * 5), uint32_max+1-5)

	/* from Go spec:
	   For two integer values x and y, the integer quotient q = x / y and remainder r = x % y satisfy the following relationships:

	   x = q*y + r  and  |r| < |y|
	   with x / y truncated towards zero ("truncated division").

	    x     y     x / y     x % y
	    5     3       1         2
	   -5     3      -1        -2
	    5    -3      -1         2
	   -5    -3       1        -2

	*/

	TEQ("", five/three, 1)
	TEQ("", five%three, 2)
	TEQ("", (-five)/three, -1)
	TEQ("", (-five)%three, -2)
	TEQ("", five/(-three), -1)
	TEQ("", five%(-three), 2)
	TEQ("", (-five)/(-three), 1)
	TEQ("", (-five)%(-three), -2)

	TEQint64("", int64(five)/int64(three), int64(1))
	TEQint64("", int64(five)%int64(three), int64(2))
	TEQint64("", int64(-five)/int64(three), int64(-1))
	TEQint64("", int64(-five)%int64(three), int64(-2))
	TEQint64("", int64(five)/int64(-three), int64(-1))
	TEQint64("", int64(five)%int64(-three), int64(2))
	TEQint64("", int64(-five)/int64(-three), int64(1))
	TEQint64("", int64(-five)%int64(-three), int64(-2))

	/*
		As an exception to this rule, if the dividend x is the most negative value for the int type of x,
		the quotient q = x / -1 is equal to x (and r = 0).
	*/
	TEQint64(""+" int64 div special case", int64_mostNeg/int64(-1), int64_mostNeg)
	TEQint64(""+" int64 mod special case", int64_mostNeg%int64(-1), 0)
	TEQint32(""+" int32 div special case", int32(int32_mostNeg/int32(-1)), int32_mostNeg)
	TEQint32(""+" int32 mod special case", int32_mostNeg%int32(-1), 0)
	if int16(int16_mostNeg/int16(-1)) != int16_mostNeg {
		fmt.Println("" + " int16 div special case")
	}
	if int16(int16_mostNeg%int16(-1)) != int16(0) {
		fmt.Println("" + " int16 mod special case")
	}
	if int8(int8_mostNeg/int8(-1)) != int8_mostNeg {
		fmt.Println("" + " int8 div special case")
	}
	if int8(int8_mostNeg%int8(-1)) != int8(0) {
		fmt.Println("" + " int8 mod special case")
	}

	/*THESE VALUES ARE NOT IN THE SPEC, SO UNTESTED
	if uint64(int64_mostNeg)/0xFFFFFFFFFFFFFFFF == uint64(int64_mostNeg) {
		fmt.Println("" + " uint64 div special case")
	}
	if uint64(int64_mostNeg)%0xFFFFFFFFFFFFFFFF == uint64(0) {
		fmt.Println("" + " uint64 mod special case")
	}
	if uint32(int32_mostNeg)/0xFFFFFFFF == uint32(int32_mostNeg) {
		fmt.Println("" + " uint32 div special case")
	}
	if uint32(int32_mostNeg)%0xFFFFFFFF == uint32(0) {
		fmt.Println("" + " uint32 mod special case")
	}
	if uint16(int16_mostNeg)/0xFFFF == uint16(int16_mostNeg) {
		fmt.Println("" + " uint16 div special case")
	}
	if uint16(int16_mostNeg)%0xFFFF == uint16(0) {
		fmt.Println("" + " uint16 mod special case")
	}
	if uint8(int8_mostNeg)/0xFF == uint8(int8_mostNeg) {
		fmt.Println("" + " uint8 div special case")
	}
	if uint8(int8_mostNeg)%0xFF == uint8(0) {
		fmt.Println("" + " uint8 mod special case")
	}
	*/

	//TODO more tests for unsigned comparisons, need to check all possibilities are covered
	TEQ("", uint8(int8_mostNeg) > uint8(0), true)
	TEQ("", uint8(int8_mostNeg) < uint8(0), false)
	TEQ("", uint16(int16_mostNeg) > uint16(0), true)
	TEQ("", uint16(int16_mostNeg) < uint16(0), false)

	TEQ("", uint32(int32_mostNeg) > uint32(0), true)
	TEQ("", uint32(int32_mostNeg) < uint32(0), false)
	TEQ(""+" uint64(int64_mostNeg) > uint64(0) ", uint64(int64_mostNeg) > uint64(0), true)
	TEQ("", uint64(int64_mostNeg) < uint64(0), false)

	//TEQint64("", int64(int64_mostNeg), int64(uint64(0x8000000000000000)))
	//fmt.Println(float64(int64_mostNeg))
	//fmt.Println(int64_mostNeg)
	uint64Global = uint64(int64_mostNeg)
	TEQuint64("", uint64Global, uint64(0x8000000000000000))

	for i := range uint64GlobalArray {
		uint64GlobalArray[i] = uint64(int64_mostNeg)
		TEQ(""+" uint64(int64_mostNeg) > uint64(0) [Array] ", uint64(uint64GlobalArray[i]) > uint64(0), true)
	}

	// TODO test for equality too & check these constants are not being resolved by the compiler, rather than genereating tests!
	TEQ("", uint8(int8_mostNeg)-uint8(42) > uint8(0), true)
	TEQ("", uint8(int8_mostNeg)-uint8(42) < uint8(three), false)
	TEQ("", uint16(int16_mostNeg)-uint16(42) > uint16(0), true)
	TEQ("", uint16(0xffff)-uint16(five) < uint16(three), false)
	TEQ("", uint32(0xffffffff)-uint32(five) > uint32(0), true)
	TEQ("", uint32(0xffffffff)-uint32(five) < uint32(three), false)
	TEQ("", uint64(0xffffffffffffffff)-uint64(five) > uint64(0), true)
	TEQ("", uint64(0xffffffffffffffff)-uint64(five) < uint64(three), false)
	TEQ("", uint8(0xff) > uint8(0xfe)-uint8(five), true)
	TEQ("", uint8(five) < uint8(three), false)
	TEQ("", uint16(0xffff) > uint16(0xfffe)-uint16(five), true)
	TEQ("", uint16(10000)-uint16(five) < uint16(1000), false)
	TEQ("", uint32(0xffffffff) > uint32(0xfffffffe)-uint32(five), true)
	TEQ("", uint32(12)-uint32(five) < uint32(three), false)
	TEQ("", uint64(0xffffffffffffffff) > uint64(0xfffffffffffffffe)-uint64(five), true)
	TEQ("", uint64(12)-uint64(five) < uint64(three), false)

	// test Float / Int64 conversions
	fiveI64 := int64(five)
	TEQfloat("", float64(fiveI64), 5.0, 0.1)
	TEQfloat("", float64(int32_mostNeg), float64(-2147483648.0), 0.1)

	TEQint64(""+" big -ve int64 division",
		int64_mostNeg/int64(100000), int64(-9223372036854775808)/int64(100000))

	TEQfloat(""+" PHP error",
		float64(int64_mostNeg/int64(100000)), float64(int64(-9223372036854775808)/int64(100000)), float64(1.0))
	TEQfloat(""+" PHP error ",
		float64(int64_max/200), float64(int64(0x7fffffffffffffff)/200), float64(10.0))
	TEQfloat("", float64(int64_mostNeg+1), float64(int64(-9223372036854775808+1)), float64(2000.0))
	TEQfloat("", float64(int64_mostNeg), float64(int64(-9223372036854775808)), float64(2000.0))
	TEQfloat("", float64(uint64Global), float64(int64(0x7fffffffffffffff)), float64(2000.0))
	uint64Global = 0xFFFFFFFFFFFFFFFF
	TEQfloat("", float64(uint64Global), float64(uint64(0xffffffffffffffff)), float64(2000.0))

	// tests below removed to avoid also loading the math package
	//TEQint64(""+" NaN ->int64 conversion", int64(math.NaN()), -9223372036854775808)
	//TEQuint64(""+" NaN ->uint64 conversion (error on php)", uint64(math.NaN()), 9223372036854775808)

	myPi := float64(3)
	myPi64 := int64(myPi)
	myPu64 := uint64(myPi)
	limit := float64(1 << 52)
	loops := 0
	for myPi < limit {
		loops++
		a := TEQint64(""+" +ve float->int64 conversion  ", int64(myPi), myPi64)
		b := TEQint64(""+" -ve float->int64 conversion  ", int64(-myPi), -myPi64)
		c := TEQuint64(""+" float->uint64 conversion  ", uint64(myPi), myPu64)
		d := TEQfloat(""+" int64->Float conversion  ", myPi, float64(myPi64), 0)
		e := TEQfloat(""+" uint64->Float conversion  ", myPi, float64(myPu64), 0)
		if a == false || b == false || c == false || d == false || e == false {
			fmt.Println("i64 loops=", loops, "myPi=", myPi, "myPi64=", myPi64, "myPu64=", myPu64)
			break
		}
		myPi *= myPi
		myPi64 *= myPi64
		myPu64 *= myPu64
	}

	itter := 0
	for u := uint64(1); itter < 53; u = u<<1 + 1 {
		f := float64(u)
		fu := uint64(f)
		if u != fu {
			fmt.Println("uint64/float64 conversion error", itter, u, f, fu, u == fu)
		}
		itter++
	}
}

func testSlices() {
	// from the Go tour...
	p := []int{2, 3, 5, 7, 11, 13}
	TEQintSlice("", p[1:4], []int{3, 5, 7})
	TEQintSlice("", p[:3], []int{2, 3, 5})
	TEQintSlice("", p[4:], []int{11, 13})

	a := make([]int, 5)
	TEQintSlice("", a, []int{0, 0, 0, 0, 0})
	TEQ("", len(a), 5)
	TEQ("", cap(a), 5)
	b := make([]int, 0, 5)
	TEQintSlice("", b, []int{})
	TEQ("", len(b), 0)
	TEQ("", cap(b), 5)
	c := b[:2]
	TEQintSlice("", c, []int{0, 0})
	TEQ("", len(c), 2)
	TEQ("", cap(c), 5)
	d := c[2:5]
	TEQintSlice("", d, []int{0, 0, 0})
	TEQ("", len(d), 3)
	TEQ("", cap(d), 3)

	var z []int
	TEQ("", len(z), 0)
	TEQ("", cap(z), 0)
	TEQ("", z == nil, true)

}

func testUTF8() {
	b := []byte("Hello, 世界")
	r, size := utf8.DecodeLastRune(b)
	TEQ("", '界', r)
	TEQ("", size, 3)
	b = b[:len(b)-size]
	r, size = utf8.DecodeLastRune(b)
	TEQ("", '世', r)
	TEQ("", size, 3)
	b = b[:len(b)-size]
	r, size = utf8.DecodeLastRune(b)
	TEQ("", ' ', r)
	TEQ("", size, 1)

	//fmt.Println("len(Zi)=", len("字"), hx.CodeInt(`'字'.length;`))

	str := "Hello, 世界"
	r, size = utf8.DecodeLastRuneInString(str)
	TEQ("", '界', r)
	TEQ("", size, 3)
	str = str[:len(str)-size]
	r, size = utf8.DecodeLastRuneInString(str)
	TEQ("", '世', r)
	TEQ("", size, 3)
	str = str[:len(str)-size]
	r, size = utf8.DecodeLastRuneInString(str)
	TEQ("", ' ', r)
	TEQ("", size, 1)

	ru := '世'
	buf := make([]byte, 3)
	n := utf8.EncodeRune(buf, ru)
	TEQ("", n, 3)
	TEQbyteSlice("", buf, []byte{228, 184, 150})

	buf = []byte{228, 184, 150} // 世
	TEQ("", true, utf8.FullRune(buf))
	TEQ("", false, utf8.FullRune(buf[:2]))

	str = "世"
	TEQ("", true, utf8.FullRuneInString(str))
	//if ShowKnownErrors || hx.GetInt("", "'字'.length") == 3 {
	TEQ(""+" NOTE: known error handling incorrect strings on UTF16 platforms", false, utf8.FullRuneInString(str[:2]))
	//}
	buf = []byte("Hello, 世界")
	TEQ("", 13, len(buf))
	TEQ("", 9, utf8.RuneCount(buf))

	str = "Hello, 世界"
	TEQ("", 13, len(str))
	TEQ("", 9, utf8.RuneCountInString(str))

	TEQ("", 1, utf8.RuneLen('a'))
	TEQ("", 3, utf8.RuneLen('界'))

	buf = []byte("a界")
	TEQ("", true, utf8.RuneStart(buf[0]))
	TEQ("", true, utf8.RuneStart(buf[1]))
	TEQ("", false, utf8.RuneStart(buf[2]))

	valid := []byte("Hello, 世界")
	invalid := []byte{0xff, 0xfe, 0xfd}
	TEQ("", true, utf8.Valid(valid))
	TEQ("", false, utf8.Valid(invalid))

	valid_rune := 'a'
	invalid_rune := rune(0xfffffff)
	TEQ("", true, utf8.ValidRune(valid_rune))
	TEQ("", false, utf8.ValidRune(invalid_rune))

	valid_string := "Hello, 世界"
	invalid_string := string([]byte{0xff, 0xfe, 0xfd})
	TEQ("", true, utf8.ValidString(valid_string))
	//if ShowKnownErrors || hx.GetInt("", "'字'.length") == 3 {
	TEQ(""+" NOTE: known error handling incorrect strings on UTF16 platforms", false, utf8.ValidString(invalid_string))
	//}
}

func testChan() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	close(c)
	TEQ("", <-c, 1)
	TEQ("", <-c, 2)
	v, ok := <-c
	TEQ("", v, 0)
	TEQ("", ok, false)

	ch := make(chan bool, 2)
	ch <- true
	ch <- true
	close(ch)
	rangeCount := 0
	for v := range ch {
		TEQ("", v, true)
		rangeCount++
	}
	TEQ("", rangeCount, 2)

	//TODO much more to come here...
}

func testComplex() {

	var x, y, z complex64
	var ss complex128

	x = 1 + 2i
	TEQfloat("", float64(real(x)), 1, 0.1)
	TEQfloat("", float64(imag(x)), 2, 0.1)

	y = complex(3, 4)
	TEQfloat("", float64(real(y)), 3, 0.1)
	TEQfloat("", float64(imag(y)), 4, 0.1)

	//this previously failed in the SSA interpreter
	z = -x
	TEQfloat("", float64(real(z)), -1, 0.1)
	TEQfloat("", float64(imag(z)), -2, 0.1)

	z = x + y
	TEQfloat("", float64(real(z)), 4, 0.1)
	TEQfloat("", float64(imag(z)), 6, 0.1)

	z = x - y
	TEQfloat("", float64(real(z)), -2, 0.1)
	TEQfloat("", float64(imag(z)), -2, 0.1)

	z = x + y - y
	TEQfloat("", float64(real(z)), float64(real(x)), 0.1)
	TEQfloat("", float64(imag(z)), float64(imag(x)), 0.1)
	/*
		z = x * y
		printf64("real(x*y)", float64(real(z)))
		printf64("imag(x*y)", float64(imag(z)))

		z = x / y
		printf64("real(x/y)", float64(real(z)))
		printf64("imag(x/y)", float64(imag(z)))
	*/
	z = x * y / y
	TEQfloat("", float64(real(z)), float64(real(x)), 0.1)
	TEQfloat("", float64(imag(z)), float64(imag(x)), 0.1)

	TEQ("", x == y, false)

	TEQ("", x != y, true)

	ss = complex128(x)
	tt := complex128(y)
	TEQ("", ss != tt, true)
}

var aString = "A"
var aaString = "AA"
var bbString = "BB"

func testString() {
	TEQ("", aString <= "A", true)
	TEQ("", aString <= aaString, true)
	TEQ("", aString > aaString, false)
	TEQ("", aString == aaString, false)
	TEQ("", aString+aString == aaString, true)
	TEQ("", bbString < aaString, false)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// fib returns a function that returns
// successive Fibonacci numbers.
func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func testClosure() {
	// example from the go tour
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		pos(i)
		neg(-2 * i)
	}
	TEQ("", pos(0), 45)
	TEQ("", neg(0), -90)

	// example from http://jordanorelli.tumblr.com/post/42369331748/function-types-in-go-golang
	x := 5
	fn := func(y int) {
		TEQ("", x, y)
	}
	fn(5)
	x++
	fn(6)

	f := fib()
	TEQ("", f(), 1)
	TEQ("", f(), 1)
	TEQ("", f(), 2)
	TEQ("", f(), 3)
	TEQ("", f(), 5)
}

func testVariadic(values ...int) {
	total := 0
	for i := range values {
		total += values[i]
	}
	TEQ("", total, 42)
}

func testMath() {
	// comment out for quicker testing
	/*
		if int(math.Sqrt(16.0)) != 4 {
			fmt.Println("" + ": Incorrect square root of 16")
		}
	*/

	// test defer close
	x := make(chan interface{})
	defer close(x) // to make sure it is not removed by Dead Code Elimination
}

func testInterface() {
	var i interface{}

	i = "test"
	if i.(string) != "test" {
		fmt.Println("testInterface string not equal 'test':")
		fmt.Println(i)
	}

	i = int(42)
	if i.(int) != 42 {
		fmt.Println("testInterface int not equal 42:")
		fmt.Println(i)
	}

	j, ok := i.(rune)
	if ok {
		fmt.Println("error rune!=int")
	}
	TEQ("", j, rune(0))
}

// from the go tour
type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) float64 {
	v.X = v.X * f
	v.Y = v.Y * f
	return v.Abs()
}

func (v MyFloat) Scale(f float64) float64 {
	return float64(v) * f
}

type Abser interface {
	Abs() float64
	Scale(x float64) float64
}

//from the go.tools/go/types documentation
type T0 struct {
	X float64
}

var p0 *T0

type Value struct {
	typ *rtype
	flag
}

type flag uintptr

func (f flag) mv(x flag) flag { f = x; return f }

type rtype struct {
	hash      uint32
	_         uint8  // test unused/padding
	ptrToThis *rtype // test self-ref
	flag
}

var rt *rtype

func (v Value) testFieldFn() {
	f0 := flag.mv(0, 42)
	f1 := flag.mv
	TEQ("", f1(f0, 42), f0)
	rt = new(rtype)
	v.typ = rt
	v.typ.hash = 42
	x := rt.hash
	TEQ("", v.typ.hash, x)
	y := &rt.hash
	z := (*y)
	TEQ("", z, x)
	var ff interface{} = &v.flag
	*(ff.(*flag)) = 42
	var ffh interface{} = v.typ
	ffh.(*rtype).flag = f1(0, 42)
	TEQ("", *(ff.(*flag)), ffh.(*rtype).flag)
}

func testInterfaceMethods() {
	var v0 Value
	v0.testFieldFn()

	v := &Vertex{3, 4}
	TEQfloat("", v.Abs(), 5, 0.00001)
	TEQfloat("", v.Scale(5), 25, 0.001)
	TEQfloat("", v.X, 15, 0.0000001)
	TEQfloat("", v.Y, 20, 0.0000001)

	var a Abser
	f := MyFloat(-42)
	vt := Vertex{3, 4}

	a = f // a MyFloat implements Abser
	x, ok := a.(Abser)
	//fmt.Println(reflect.TypeOf(x).String()) => main.MyFloat
	if !ok {
		fmt.Println("Error in testInterfaceMethods(): MyFloat should be in Abser interface")
	}
	TEQ(""+"testInterfaceMethods():MyFloat in Abser", a, f)
	TEQ(""+"testInterfaceMethods():MyFloat.Abs()", a.Abs(), float64(42))
	TEQ(""+"testInterfaceMethods():x.Abs()", x.Abs(), float64(42))
	TEQ(""+"testInterfaceMethods():MyFloat.Scale(10)", a.Scale(10), float64(-420))
	TEQ(""+"testInterfaceMethods():x.Scale(10)", x.Scale(10), float64(-420))

	a = &vt // a *Vertex implements Abser
	y, ok := a.(Abser)
	//fmt.Println(reflect.TypeOf(y).String()) //=> *main.Vertex
	//fmt.Println(y)
	if !ok {
		fmt.Println("Error in testInterfaceMethods(): Vertex should be in Abser interface")
	}
	TEQ(""+"testInterfaceMethods():*Vertex in Abser", a, &vt)
	TEQfloat(""+"testInterfaceMethods():*Vertex.Abs()", a.Abs(), float64(5), 0.000001)
	TEQfloat(""+"testInterfaceMethods():y.Abs()", y.Abs(), float64(5), 0.000001)
	TEQfloat(""+"testInterfaceMethods():*Vertex.Scale(10)", a.Scale(10), float64(50), 0.000001)
	//fmt.Println(y)
	TEQfloat(""+"testInterfaceMethods():y.Scale(10)", y.Scale(10),
		float64(653.351098686295472), 0.01) // crazey number because Sqrt fn is an approximisation
	//fmt.Println(y)

	// a=vt // a Vertex, does NOT

	//from the go.tools/go/types documentation
	p0 = new(T0) // TODO should fail with this line missing, but does not (globals pre-initialised when they should not be)
	p0.X = 42
	TEQfloat("", p0.X, 42.0, 0.01)

}

func testStrconv() {
	/*
		TEQ(""+"testStrconv():Itoa", "424242", strconv.Itoa(424242))

		TEQ("", strings.HasPrefix("say what", "say"), true)
		TEQ(""+" string.Contains (error on js)", strings.Contains("say what", "ay"), true)
		TEQ(""+" string.Contains (error on js)", strings.Contains("seafood", "foo"), true)
		TEQ("", strings.Contains("seafood", "bar"), false)
		TEQ("", strings.Contains("seafood", ""), true)
		TEQ("", strings.Contains("", ""), true)
		TEQ("", strings.Contains("equal?", "equal?"), true)

		TEQ("", bytes.HasPrefix([]byte("say what"), []byte("say")), true)
		TEQ("", bytes.Contains([]byte("say what"), []byte("ay")), true)
	*/
}

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // send sum to c
}

func testTour64() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // receive from c

	TEQ("", x+y, 12) // x & y could arrive in any order...
}

func testDefer_a() {
	i := 0
	defer TEQ("", i, 0)
	i++
	return
}
func testDefer_b(ch chan int) {
	for i := 0; i < 4; i++ {
		defer func(j int) { ch <- j }(i)
	}
}
func testDefer_c() (i int) {
	defer func() { i++ }()
	return 1
}
func protect(g func(int)) {
	defer func() {
		TEQ("", recover(), "test panic")
	}()
	g(0)
}

func g(i int) {
	if i > 3 {
		err := errors.New("test panic")
		panic(err.Error())
	}
	for j := 0; j < i; j++ {
		defer testDefer_d()
	}
	g(i + 1)
}

var tddCount = 0

func testDefer_d() {
	tddCount++ // just to give the routine something to do
}

func testDefer() {
	// examples from http://blog.golang.org/defer-panic-and-recover
	testDefer_a()
	b := make(chan int, 4)
	testDefer_b(b)
	TEQ("", <-b, 3)
	TEQ("", <-b, 2)
	TEQ("", <-b, 1)
	TEQ("", <-b, 0)
	TEQ("", testDefer_c(), 2)
	protect(g)
	TEQ("", tddCount, 6)
}

// these two names were failing in java as being duplicates, now failing in PHP...
func Ilogb(x float64) int {
	return int(Sqrt(x))
}
func ilogb(x float64) int {
	return int(Sqrt(x))
}
func testCaseSensitivity() {
	//moved to a separate test file
	//TEQ("", ilogb(64), Ilogb(64))
}

var (
	aGrCtr int32
	//aGrCtrMux sync.Mutex
	//aGrWG     sync.WaitGroup
)

func aGoroutine(a int) {
	if a == 4 {
		//panic("test panic in goroutine 4")
	}
	for i := 0; i < a; i++ {
		runtime.Gosched()
	}
	//(&aGrCtrMux).Lock()
	//atomic.AddInt32(&aGrCtr, -1)
	aGrCtr--
	//(&aGrCtrMux).Unlock()

	//aGrWG.Done()
}

const numGR = 5

func testManyGoroutines() {
	var n = numGR
	aGrCtr = numGR * 2 // set up the goroutine counter
	for i := 0; i < n; i++ {
		//aGrWG.Add(1)
		go aGoroutine(i)
	}
	for i := n; i > 0; i-- {
		//aGrWG.Add(1)
		go aGoroutine(i)
	}
}

//
// Code from http://golangtutorials.blogspot.co.uk/2011/06/channels-in-go-range-and-select.html
//
func makeCakeAndSend(cs chan string, flavor string, count int) {
	for i := 1; i <= count; i++ {
		TEQ("Delay", i, i)
		cakeName := flavor + " Cake " + string('0'+i)
		cs <- cakeName //send a strawberry cake
	}
	close(cs)
}

func receiveCakeAndPack(strbry_cs chan string, choco_cs chan string) {
	strbry_closed, choco_closed := false, false

	for {
		//if both channels are closed then we can stop
		if strbry_closed && choco_closed {
			return
		}
		//fmt.Println("Waiting for a new cake ...")
		select {
		case cakeName, strbry_ok := <-strbry_cs:
			if !strbry_ok {
				strbry_closed = true
				//fmt.Println(" ... Strawberry channel closed!")
			} else {
				//fmt.Println("Received from Strawberry channel.  Now packing", cakeName)
				_ = cakeName
			}
		case cakeName, choco_ok := <-choco_cs:
			if !choco_ok {
				choco_closed = true
				//fmt.Println(" ... Chocolate channel closed!")
			} else {
				//fmt.Println("Received from Chocolate channel.  Now packing", cakeName)
				_ = cakeName
			}
		default:
			//fmt.Println("no cake!")
		}
	}
}

func testChanSelect() {
	strbry_cs := make(chan string)
	choco_cs := make(chan string)

	//two cake makers
	go makeCakeAndSend(choco_cs, "Chocolate", 3)   //make 3 chocolate cakes and send
	go makeCakeAndSend(strbry_cs, "Strawberry", 3) //make 3 strawberry cakes and send

	//one cake receiver and packer
	receiveCakeAndPack(strbry_cs, choco_cs) //pack all cakes received on these cake channels

	//sleep for a while so that the program doesn’t exit immediately
	//time.Sleep(2 * 1e9)
}

//end code from http://golangtutorials.blogspot.co.uk/2011/06/channels-in-go-range-and-select.html

//From the go tour http://tour.golang.org/#69
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			//fmt.Println("quit")
			return
		}
	}
}

func tourfib() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			/*fmt.Println(*/ <-c /*)*/
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

// end tour

func testUintDiv32() {
	for seed := int32(-2); seed <= 2; seed++ {
		var uifs, pwr2 uint32
		uifs = uint32(seed)
		if seed != 0 {
			one := uint32(1)
			TEQuint32("testUintDiv32() uint x/x==unity ", 1, uifs/uifs)
			TEQint32("testUintDiv32() int x/x==unity ", 1, seed/seed)
			TEQuint32("testUintDiv32() uint x/1==x ", uifs/one, uifs)
			TEQint32("testUintDiv32() int x/1==x ", seed/int32(one), seed)
			if seed > 0 {
				TEQuint32("testUintDiv32() uint +ve roundtrip ", uifs, (uifs*uifs)/uifs)
			}
			TEQint32("testUintDiv32() int +ve roundtrip ", seed, (seed*seed)/seed)
		}
		pwr2 = uint32(1)
		for i := uint32(0); i < 32; i++ {
			if !TEQuint32("testUintDiv32() T1 uint ", (uifs)>>i, (uifs)/pwr2) {
				fmt.Println("ProblemT1 uint i=", int(i))
			}
			if !TEQuint32("testUintDiv32() uint shift equivalence roundtrip  ", (uifs*pwr2)>>i, (uifs<<i)/pwr2) {
				fmt.Println("Problem uint seed,i=", seed, i)
			}
			if i < 31 {
				if seed >= 0 {
					if !TEQint32("testUintDiv32() T1 int ", (seed)>>i, (seed)/int32(pwr2)) {
						fmt.Println("ProblemT1 int i=", int(i))
					}
				}
				if !TEQint32("testUintDiv32() int shift equivalence roundtrip  ",
					(seed*int32(pwr2))>>i, (seed<<i)/int32(pwr2)) {
					fmt.Println("Problem int seed,i=", seed, i)
				}
			}
			pwr2 <<= 1
		}
	}
}
func testUintDiv64() {
	for seed := int64(-2); seed <= 2; seed++ {
		var uifs, pwr2 uint64
		uifs = uint64(seed)
		if seed != 0 {
			one := uint64(1)
			TEQuint64("testUintDiv64() uint x/x==unity ", 1, uifs/uifs)
			TEQint64("testUintDiv64() int x/x==unity ", 1, seed/seed)
			TEQuint64("testUintDiv64() uint x/1==x ", uifs/one, uifs)
			TEQint64("testUintDiv64() int x/1==x ", seed/int64(one), seed)
			if seed > 0 {
				TEQuint64("testUintDiv64() +ve roundtrip ", uifs, (uifs*uifs)/uifs)
			}
			TEQint64("testUintDiv64() int +ve roundtrip ", seed, (seed*seed)/seed)
		}
		pwr2 = uint64(1)
		for i := uint64(0); i < 64; i++ {
			if !TEQuint64("testUintDiv64() uint ", uifs>>i, uifs/pwr2) {
				fmt.Println("Problem seed,i=", seed, i)
			}
			if !TEQuint64("testUintDiv64() uint shift equivalence roundtrip  ", (uifs*pwr2)>>i, (uifs<<i)/pwr2) {
				fmt.Println("Problem seed,i=", seed, i)
			}
			if i < 63 {
				if seed >= 0 {
					if !TEQint64("testUintDiv64() T1 int ", (seed)>>i, (seed)/int64(pwr2)) {
						fmt.Println("ProblemT1 int i=", int(i))
					}
				}
				if !TEQint64("testUintDiv64() int shift equivalence roundtrip  ",
					(seed*int64(pwr2))>>i, (seed<<i)/int64(pwr2)) {
					fmt.Println("Problem int seed,i=", seed, i)
				}
			}
			pwr2 <<= 1
		}
	}
}

type foo struct {
	a int
}

func bar() *foo {
	//fmt.Println("bar() called")
	return &foo{42}
}

type foo2 struct {
	a struct {
		b int
	}
}

func testPtr() {
	q := &bar().a
	//fmt.Println("pointer created")
	*q = 40
	TEQ("", *q, 40) // should be 40

	// type foo as above
	f := foo{6}
	r := &f // this isn't creating a pointer
	f = foo{4}
	TEQ("", r.a, 4) // should be 4

	f2 := foo2{}
	p2 := &f2.a
	q2 := &p2.b // referring to the `p` variable instead of its pointer value at the time; &(*p).b (should be equivalent) is handled correctly though
	p2 = nil
	TEQ("", *q2, 0) // should be 0

	var f3 struct{ a [3]int }
	f3.a = [3]int{6, 6, 6}
	s3 := f3.a[:]
	f3.a = [3]int{4, 4, 4}
	TEQ("", s3[1], 4) // should be 4

}

type tbe struct {
	a int
	b string
}

func (x tbe) zip() string {
	ret := ""
	for z := 0; z < x.a; z++ {
		ret += x.b
	}
	return ret
}

type testtbe struct {
	tbe // embedded struct
	c   bool
}

func testEmbed() {
	var t testtbe
	t.a = 3
	t.b = "Grunt"
	TEQ("", "GruntGruntGrunt", t.zip())
}

func testUnsafe() { // adapted from http://stackoverflow.com/questions/19721008/golang-unsafe-dynamic-byte-array

	// Arbitrary size
	n := 4

	// Create a slice of the correct size
	m := make([]int32, n)

	// Use convoluted indirection to cast the first few bytes of the slice
	// to an unsafe pointer
	mPtr := unsafe.Pointer(&m[0])

	// Check it worked
	m[0] = 987
	// (we have to recast the uintptr to a *int to examine it)
	TEQint32("", m[0], *(*int32)(mPtr))

	if runtime.GOARCH != "neko" { // to avoid errors on the automated test
		TEQuint32("Only works in fullunsafe mode", 219, (uint32)(*(*uint8)(mPtr)))
	}

	// error on pointer arithmetic
	//uip := uintptr(mPtr)
	//uip++
}

func tc64(f float64) float64 {
	if runtime.GOOS == "nacl" {
		return hx.CallFloat("", "Go_haxegoruntime_FFloat64frombits.hx", 1,
			hx.CallDynamic("", "Go_haxegoruntime_FFloat64bits.hx", 1, f))
	}
	return f
}

const (
	SmallestNormalFloat64   = 2.2250738585072014e-308 // 2**-1022
	LargestSubnormalFloat64 = SmallestNormalFloat64 - SmallestNonzeroFloat64

	MaxFloat32             = 3.40282346638528859811704183484516925440e+38  // 2**127 * (2**24 - 1) / 2**23
	SmallestNonzeroFloat32 = 1.401298464324817070923729583289916131280e-45 // 1 / 2**(127 - 1 + 23)

	MaxFloat64             = 1.797693134862315708145274237317043567981e+308 // 2**1023 * (2**53 - 1) / 2**52
	SmallestNonzeroFloat64 = 4.940656458412465441765687928682213723651e-324 // 1 / 2**(1023 - 1 + 52)
)

func testFloatConv() {
	if runtime.GOARCH != "neko" {
		TEQ("SmallestNormalFloat64", SmallestNormalFloat64, tc64(SmallestNormalFloat64))
		TEQ("LargestSubnormalFloat64", LargestSubnormalFloat64, tc64(LargestSubnormalFloat64))
		TEQ("MaxFloat32", MaxFloat32, tc64(MaxFloat32))
		TEQ("SmallestNonzeroFloat32", SmallestNonzeroFloat32, tc64(SmallestNonzeroFloat32))
		TEQ("MaxFloat64", MaxFloat64, tc64(MaxFloat64))
		TEQ("SmallestNonzeroFloat64", SmallestNonzeroFloat64, tc64(SmallestNonzeroFloat64))
		TEQ("42.42", 42.42, tc64(42.42))
		if runtime.GOOS == "nacl" {
			pi := tc64(hx.GetFloat("", "Math.POSITIVE_INFINITY"))
			if pi <= MaxFloat64 {
				fmt.Println("testFloatConv() POSITIVE_INFINITY invalid")
			}
			ni := tc64(hx.GetFloat("", "Math.NEGATIVE_INFINITY"))
			if ni >= SmallestNonzeroFloat64 {
				fmt.Println("testFloatConv() NEGATIVE_INFINITY invalid")
			}
			if hx.GetFloat("", "Math.NaN") == tc64(hx.GetFloat("", "Math.NaN")) {
				fmt.Println("testFloatConv() NaN == NaN")
			}
		}
	}
}

type ObjKey [2]int

func testObjMap() {
	m := map[ObjKey]int32{
		ObjKey{1, 2}: 3,
	}
	TEQint32("", 3, m[ObjKey{1, 2}])

	cm := map[complex128]int32{
		1 + 2i: 3,
	}

	TEQint32("", 3, cm[1+2i])

}

type unaligned7 struct {
	h [17]testtbe
	i int32
	j uint16
	k uint8
}

func testUnaligned() {
	var x [3]unaligned7
	//y := interface{}(x)
	//z := reflect.ValueOf(y)
	for d := 0; d < len(x); d++ {
		x[d].k = uint8(d)
		TEQuint32("x[d]k!=d", uint32(x[d].k), uint32(d))
		//TEQuint32("reflect(x[d]k)!=d", uint32(z.Index(d).FieldByName("k").Int()), uint32(d))
		for hh := 0; hh < 17; hh++ {
			x[d].h[hh].a = hh
			TEQuint32("x[d].h[hh].a!=hh", uint32(x[d].h[hh].a), uint32(hh))
			//TEQuint32("reflect(x[d].h[hh].a)!=hh", uint32(z.Index(d).FieldByName("h").Index(hh).Int()), uint32(hh))
		}
	}
	x1 := make([]unaligned7, 7)
	//y1 := interface{}(x1)
	//z1 := reflect.ValueOf(y1)
	for d := 0; d < len(x1); d++ {
		x1[d].k = uint8(d)
		TEQuint32("x1[d]k!=d", uint32(x1[d].k), uint32(d))
		//TEQuint32("reflect(x1[d]k)!=d", uint32(z1.Index(d).FieldByName("k").Int()), uint32(d))
		for hh := 0; hh < 17; hh++ {
			x1[d].h[hh].a = hh
			TEQuint32("x1[d].h[hh].a!=hh", uint32(x1[d].h[hh].a), uint32(hh))
			//TEQuint32("reflect(x1[d].h[hh].a)!=hh", uint32(z1.Index(d).FieldByName("h").Index(hh).Int()), uint32(hh))
		}
	}
}

func main() {
	var array [4][5]int
	array[3][2] = 12
	if array[3][2] != 12 {
		fmt.Println("Array handling error:", array[3][2])
	}
	//fmt.Println("Start test running in: " + runtime.GOARCH)
	testManyGoroutines() // here to run alongside the other code execution
	tourfib()
	testCaseSensitivity()
	testInit()
	testConst()
	testUTF()
	testFloat()
	testMultiRet()
	testAppend()
	testStruct()
	testHeader()
	testCopy()
	testInFuncPtr()
	testCallBy()
	testMap()
	testNamed()
	testFuncPtr()
	testIntOverflow()
	testSlices()
	testChan()
	testComplex()
	testUTF8()
	testString()
	testClosure()
	testVariadic(42)
	testVariadic(40, 2)
	testVariadic(42, -5, 3, 2)
	testInterface()
	testInterfaceMethods()
	testStrconv()
	testTour64()
	testUintDiv32()
	testUintDiv64()
	testDefer()
	testPtr()
	testChanSelect()
	testEmbed()
	testUnsafe()
	testObjMap()
	testFloatConv()
	testUnaligned()
	//aGrWG.Wait()
	TEQint32(""+" testManyGoroutines() (NOT sync/atomic) counter:", aGrCtr, 0)
	if runtime.GOOS == "nacl" { // really a haxe emulation of nacl
		TEQ("", hx.CodeInt("", "42;"), int(42))
		TEQ("", hx.CodeString("", "'test';"), "test")
		TEQ(""+"Num Haxe GR post-wait", runtime.NumGoroutine(), 1)
		//panic("show GRs active")
	} else {
		//fmt.Println(runtime.NumGoroutine())
		TEQ(""+"Num Haxe GR post-wait", runtime.NumGoroutine(), 3)
	}
	TEQ("", unicode.IsSpace(' '), true) // makes the test longer but more complete
	//fmt.Println("End test running in: " + runtime.GOARCH)
	//fmt.Println("Σ sigma, " + "再见！Previous two chinese characters should say goodbye! (testing unicode output)")
	//fmt.Println()
	//hx.Call("", "Tgotypes.setup", 0)
	//hx.Call("", "Go_haxegoruntime_typetest.hx", 0)
	//testTypes()
	//fmt.Println(hx.GetFloat("", "Object.MinFloat64"))
}

/*
func testTypes() {
	for id := 0; id < len(haxegoruntime.TypeTable); id++ {
		r := unsafe.Pointer(haxegoruntime.TypeTable[id])
		fmt.Println("DEBUG type", id, r)
	}
}
*/
