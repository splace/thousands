// thousands.Int is an int implementing fmt.Formatter with the int's digits separated into groups of 3 by a separator.
// also..
// it can replace trailing groups with the appropriate suffix letter.
// the scaling can be automatically varied to retain precision.
package thousands

import "fmt"
import "strings"
import "strconv"

var suffixes = []byte(string("kMGTPEZY"))

type Int int

const (
	decimals = 3
	binaries = 10
	defWidth = 5
)

// '%s' returns using a char separator (',' or '%+s' for alt sep '.')
// '%v' uses narrow space separator.
// Width specifies number of thousand blocks to remove.
// Precision specifies displayed digits (overrides width) default:4 (this is only approx for binary)
// Examples:
// "X" truncates by X groups of three decimal digits
// "X.Y" truncates by as many groups of three decimal digits that still leaves Y  digits.
// '#' switches to binary (1000 -> 1024)
// Note: scaling is integer, with value just truncated.
func (v Int) Format(f fmt.State, r rune) {
	var s uint
	if v < 0 {
		f.Write([]byte("-"))
		s = uint(-v)
	} else {
		if f.Flag('+') {
			f.Write([]byte("+"))
		}
		s = uint(v)
	}
	p, pset := f.Precision()
	w, wset := f.Width()
	if !wset {
		w = defWidth
	}
	if f.Flag('#') {
		s >>= binaries * p
		if !pset {
			m := (w+binaries)%binaries - w
			p += m
			s >>= m * binaries
		}
	} else {
		s /= power10(uint8(p * decimals))
		if !pset {
			m := (w+decimals)%decimals - w
			p += m
			s /= power10(uint8(m))
		}
	}
	sr := strconv.FormatUint(uint64(s), 10)
	switch r {
	case 's':
		if f.Flag('-') {
			f.Write(CharGroupRTL(sr, '.', decimals))
		} else {
			f.Write(CharGroupRTL(sr, ',', decimals))
		}
	default:
		f.Write(CharGroupRTL(sr, ' ', decimals))
	}
	if p > 0 && p < len(suffixes) {
		f.Write(suffixes[p-1 : p])
		if f.Flag('#') {
			f.Write([]byte("i"))
		}
	}
}

func CharGroupRTL(s string, d rune, c int) []byte {
	var bs strings.Builder
	lsmo := len(s) - 1
	for i, r := range s {
		bs.WriteRune(r)
		if i < lsmo && (lsmo-i)%c == 0 {
			bs.WriteRune(d)
		}
	}
	return []byte(bs.String())
}

func power10(n uint8) uint {
	switch n {
	case 0:
		return 1
	case 1:
		return 1e1
	case 2:
		return 1e2
	case 3:
		return 1e3
	case 4:
		return 1e4
	case 5:
		return 1e5
	case 6:
		return 1e6
	case 7:
		return 1e7
	default:
		return 1e8 * power10(nonOverflowSubtract(n, 8))
	}
}

func nonOverflowSubtract(a, b uint8) uint8 {
	if b > a {
		return 0
	}
	return a - b
}
