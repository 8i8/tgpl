package eval

import "math"

type fn struct {
	params  int
	usage   string
	special string
}

func test() {
	math.Atanh(3)
}

var fnData = map[string]fn{
	"cos": {1, `cos returns the cosin of the radian argument x.`,
		`
Special cases are:
	Cos(±Inf) = NaN
	Cos(NaN) = NaN`},

	"asin": {1, `asin returns the arcsine, in radians, of x.`,
		`
Special cases are:
	asin(±0) = ±0
	asin(x) = NaN if x < -1 or x > 1`},

	"acos": {1, `acos returns the arccosine, in radians, of x.`,
		`
Special case is:
	acos(x) = NaN if x < -1 or x > 1`},

	"atan": {1, `atan returns the arctangent, in radians, of x.`,
		`
Special cases are:
        atan(±0) = ±0
        atan(±Inf) = ±Pi/2`},

	"asinh": {1, `asinh returns the inverse hyperbolic sine of x.`,
		`
Special cases are:
  	sinh(±0) = ±0
  	sinh(±Inf) = ±Inf
  	sinh(NaN) = NaN`},

	"acosh": {1, `acosh returns the inverse hyperbolic cosine of x.`,
		`
Special cases are:
  	acosh(+Inf) = +Inf
  	acosh(x) = NaN if x < 1
  	acosh(NaN) = NaN`},

	"atanh": {1, `atanh returns the inverse hyperbolic tangent of x.`,
		`
Special cases are:
	Atanh(1) = +Inf
	Atanh(±0) = ±0
	Atanh(-1) = -Inf
	Atanh(x) = NaN if x < -1 or x > 1
	Atanh(NaN) = NaN`},

	"atan2": {2, `atan2 returns the arc tangent of y/x, using the signs of the two to
determine the quadrant of the return value.`,
		`
 Special cases are (in order):

	Atan2(y, NaN) = NaN
	Atan2(NaN, x) = NaN
	Atan2(+0, x>=0) = +0
	Atan2(-0, x>=0) = -0
	Atan2(+0, x<=-0) = +Pi
	Atan2(-0, x<=-0) = -Pi
	Atan2(y>0, 0) = +Pi/2
	Atan2(y<0, 0) = -Pi/2
	Atan2(+Inf, +Inf) = +Pi/4
	Atan2(-Inf, +Inf) = -Pi/4
	Atan2(+Inf, -Inf) = 3Pi/4
	Atan2(-Inf, -Inf) = -3Pi/4
	Atan2(y, +Inf) = 0
	Atan2(y>0, -Inf) = +Pi
	Atan2(y<0, -Inf) = -Pi
	Atan2(+Inf, x) = +Pi/2
	Atan2(-Inf, x) = -Pi/2`},

	"abs": {1, `Abs returns the absolute value of x.`,
		`
Special cases are:
	Abs(±Inf) = +Inf
	Abs(NaN) = NaN`},

	"ceil": {1, `Ceil returns the least integer value greater than or equal to x.`,
		`
 Special cases are:
	Ceil(±0) = ±0
	Ceil(±Inf) = ±Inf
	Ceil(NaN) = NaN`},

	"cbrt": {1, `Cbrt returns the cube root of x.`,
		`
Special cases are:
	Cbrt(±0) = ±0
	Cbrt(±Inf) = ±Inf
	Cbrt(NaN) = NaN`},

	"copysign": {2, `Copysign returns a value with the magnitude of x and the sign of y.`, ``},

	"dim": {2, `Dim returns the maximum of x-y or 0.`,
		`
Special cases are:
	Dim(+Inf, +Inf) = NaN
	Dim(-Inf, -Inf) = NaN
	Dim(x, NaN) = Dim(NaN, x) = NaN`},

	"exp": {1, `Exp returns e**x, the base-e exponential of x.`,
		`
Special cases are:
	Exp(+Inf) = +Inf
	Exp(NaN) = NaN
Very large values overflow to 0 or +Inf.
Very small values underflow to 1.`},

	"exp2": {1, `Exp2 returns 2**x, the base-2 exponential of x.`,
		`
Special cases are the same as Exp.`},

	"expm1": {1, `Expm1 returns e**x - 1, the base-e exponential of x minus 1.
It is more accurate than Exp(x) - 1 when x is near zero.`,
		`
Special cases are:
	Expm1(+Inf) = +Inf
	Expm1(-Inf) = -1
	Expm1(NaN) = NaN
Very large values overflow to -1 or +Inf.`},

	"FMA": {3, `FMA returns x * y + z, computed with only one rounding.
(That is, FMA returns the fused multiply-add of x, y, and z.)`, ``},

	"floor": {1, `Floor returns the greatest integer value less than or equal to x.`,
		`
Special cases are:
	Floor(±0) = ±0
	Floor(±Inf) = ±Inf
	Floor(NaN) = NaN`},

	"gamma": {1, `Gamma returns the Gamma function of x.`,
		`
 Special cases are:
	Gamma(+Inf) = +Inf
	Gamma(+0) = +Inf
	Gamma(-0) = -Inf
	Gamma(x) = NaN for integer x < 0
	Gamma(-Inf) = NaN
	Gamma(NaN) = NaN`},

	"hypot": {1, `Hypot returns Sqrt(p*p + q*q), taking care to avoid
unnecessary overflow and underflow.`,
		`
Special cases are:
	Hypot(±Inf, q) = +Inf
	Hypot(p, ±Inf) = +Inf
	Hypot(NaN, q) = NaN
	Hypot(p, NaN) = NaN`},

	"inf": {1, `Inf returns positive infinity if sign >= 0, negative infinity if sign < 0.`, ``},

	"J0": {1, `J0 returns the order-zero Bessel function of the first kind.`,
		`
Special cases are:
	J0(±Inf) = 0
	J0(0) = 1
	J0(NaN) = NaN`},

	"J1": {1, `J1 returns the order-one Bessel function of the first kind.i`,
		`
Special cases are:
	J1(±Inf) = 0
	J1(NaN) = NaN`},

	"Jn": {2, `Jn returns the order-n Bessel function of the first kind.`,
		`
Special cases are:
	Jn(n, ±Inf) = 0
	Jn(n, NaN) = NaN`},

	"ldexp": {2, `Ldexp is the inverse of Frexp. It returns frac × 2**exp.`,
		`
Special cases are:
	Ldexp(±0, exp) = ±0
	Ldexp(±Inf, exp) = ±Inf
	Ldexp(NaN, exp) = NaN`},

	"log": {1, `Log returns the natural logarithm of x.`,
		`
Special cases are:
	Log(+Inf) = +Inf
	Log(0) = -Inf
	Log(x < 0) = NaN
	Log(NaN) = NaN`},

	"log10": {1, `Log10 returns the decimal logarithm of x.
The special cases are the same as for Log.`, ``},

	"log1p": {1, `Log1p returns the natural logarithm of 1 plus its argument x.
It is more accurate than Log(1 + x) when x is near zero.`,
		`
Special cases are:
	Log1p(+Inf) = +Inf
	Log1p(±0) = ±0
	Log1p(-1) = -Inf
	Log1p(x < -1) = NaN
	Log1p(NaN) = NaN`},

	"log2": {1, `Log2 returns the binary logarithm of x.
The special cases are the same as for Log.`, ``},

	"logb": {1, `Logb returns the binary exponent of x.`,
		`
Special cases are:
	Logb(±Inf) = +Inf
	Logb(0) = -Inf
	Logb(NaN) = NaN`},

	"max": {2, `Max returns the larger of x or y.`,
		`
Special cases are:
	Max(x, +Inf) = Max(+Inf, x) = +Inf
	Max(x, NaN) = Max(NaN, x) = NaN
	Max(+0, ±0) = Max(±0, +0) = +0
	Max(-0, -0) = -0`},

	"min": {2, `Min returns the smaller of x or y.`,
		`
Special cases are:
	Min(x, -Inf) = Min(-Inf, x) = -Inf
	Min(x, NaN) = Min(NaN, x) = NaN
	Min(-0, ±0) = Min(±0, -0) = -0`},

	"mod": {2, `Mod returns the floating-point remainder of x/y.  The magnitude of the
result is less than y and its sign agrees with that of x.`,
		`
Special cases are:
	Mod(±Inf, y) = NaN
	Mod(NaN, y) = NaN
	Mod(x, 0) = NaN
	Mod(x, ±Inf) = x
	Mod(x, NaN) = NaN`},

	"nextafter": {2, `Nextafter returns the next representable float64 value after x towards y.`,
		`
Special cases are:
	Nextafter(x, x)   = x
	Nextafter(NaN, y) = NaN
	Nextafter(x, NaN) = NaN`},

	"pow": {2, `Pow returns x**y, the base-x exponential of y.`,
		`
Special cases are (in order):
	Pow(x, ±0) = 1 for any x
	Pow(1, y) = 1 for any y
	Pow(x, 1) = x for any x
	Pow(NaN, y) = NaN
	Pow(x, NaN) = NaN
	Pow(±0, y) = ±Inf for y an odd integer < 0
	Pow(±0, -Inf) = +Inf
	Pow(±0, +Inf) = +0
	Pow(±0, y) = +Inf for finite y < 0 and not an odd integer
	Pow(±0, y) = ±0 for y an odd integer > 0
	Pow(±0, y) = +0 for finite y > 0 and not an odd integer
	Pow(-1, ±Inf) = 1
	Pow(x, +Inf) = +Inf for |x| > 1
	Pow(x, -Inf) = +0 for |x| > 1
	Pow(x, +Inf) = +0 for |x| < 1
	Pow(x, -Inf) = +Inf for |x| < 1
	Pow(+Inf, y) = +Inf for y > 0
	Pow(+Inf, y) = +0 for y < 0
	Pow(-Inf, y) = Pow(-0, -y)
	Pow(x, y) = NaN for finite x < 0 and finite non-integer y`},

	"pow10": {1, `Pow10 returns 10**n, the base-10 exponential of n.`,
		`
Special cases are:
	Pow10(n) =    0 for n < -323
	Pow10(n) = +Inf for n > 308`},

	"remainder": {2, `Remainder returns the IEEE 754 floating-point remainder of x/y.`,
		`
Special cases are:
	Remainder(±Inf, y) = NaN
	Remainder(NaN, y) = NaN
	Remainder(x, 0) = NaN
	Remainder(x, ±Inf) = x
	Remainder(x, NaN) = NaN
	`},

	"round": {1, `Round returns the nearest integer, rounding half away from zero.`,
		`
Special cases are:
	Round(±0) = ±0
	Round(±Inf) = ±Inf
	Round(NaN) = NaN`},

	"roundtoeven": {1, `RoundToEven returns the nearest integer, rounding ties to even.`,
		`
Special cases are:
	RoundToEven(±0) = ±0
	RoundToEven(±Inf) = ±Inf
	RoundToEven(NaN) = NaN`},

	"sin": {1, `Sin returns the sine of the radian argument x.`,
		`
Special cases are:
	Sin(±0) = ±0
	Sin(±Inf) = NaN
	Sin(NaN) = NaN`},

	"sinh": {1, `Sinh returns the hyperbolic sine of x.`,
		`
Special cases are:
	Sinh(±0) = ±0
	Sinh(±Inf) = ±Inf
	Sinh(NaN) = NaN`},

	"sqrt": {1, `Sqrt returns the square root of x.`,
		`
Special cases are:
	Sqrt(+Inf) = +Inf
	Sqrt(±0) = ±0
	Sqrt(x < 0) = NaN
	Sqrt(NaN) = NaN
	`},

	"tan": {1, `Tan returns the tangent of the radian argument x.`,
		`
Special cases are:
	Tan(±0) = ±0
	Tan(±Inf) = NaN
	Tan(NaN) = NaN`},

	"tanh": {1, `Tanh returns the hyperbolic tangent of x.`,
		`
Special cases are:
	Tanh(±0) = ±0
	Tanh(±Inf) = ±1
	Tanh(NaN) = NaN`},

	"trunc": {1, `Trunc returns the integer value of x.`,
		`
Special cases are:
	Trunc(±0) = ±0
	Trunc(±Inf) = ±Inf
	Trunc(NaN) = NaN`},

	"Y0": {1, `Y0 returns the order-zero Bessel function of the second kind.`,
		`
Special cases are:
	Y0(+Inf) = 0
	Y0(0) = -Inf
	Y0(x < 0) = NaN
	Y0(NaN) = NaN`},

	"Y1": {1, `Y1 returns the order-one Bessel function of the second kind.`,
		`
Special cases are:
	Y1(+Inf) = 0
	Y1(0) = -Inf
	Y1(x < 0) = NaN
	Y1(NaN) = NaN`},

	"Yn": {2, `Yn returns the order-n Bessel function of the second kind.`,
		`
Special cases are:
	Yn(n, +Inf) = 0
	Yn(n ≥ 0, 0) = -Inf
	Yn(n < 0, 0) = +Inf if n is odd, -Inf if n is even
	Yn(n, x < 0) = NaN
	Yn(n, NaN) = NaN`},
}
