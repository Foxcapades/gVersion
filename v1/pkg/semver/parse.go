package semver

import (
	"unsafe"

	"github.com/foxcapades/gVersion/v1/internal/util"
)

var (
	errInvalidSemVerString = "invalid format for a semantic version string"
)

const (
	parseBufferSize = uint8(64)

	segDivider   uint8 = '.'
	preDivider   uint8 = '-'
	buildDivider uint8 = '+'

	digit0 uint8 = '0'
	digit9 uint8 = '9'

	leader uint8 = 'v'
)

func Parse(versionString string) (version Version, err error) {
	input := roBytes(&versionString)
	pos := uint8(0)

	// Skip leading character if it's present.
	if input[0] == leader {
		pos++
	}

	buf := [parseBufferSize]byte{}
	ln := uint8(len(input))

	err = parseVersions(input, &pos, &buf, &version)
	if err != nil {
		return version, err
	}

	if pos < ln && input[pos] == preDivider {
		pos++
		parsePrerelease(input, &pos, &buf, &version)
	}

	if pos < ln && input[pos] == buildDivider {
		pos++
		parseBuild(input, &pos, &buf, &version)
	}

	return
}

const vSegs = 3
func parseVersions(vn []byte, pos *uint8, buf *[parseBufferSize]byte, ver *Version) error {
	parts := [vSegs]*uint8{&ver.Major, &ver.Minor, &ver.Patch}

	pp := uint8(0)
	ln := uint8(len(vn))
	bp := 0

	for ; *pos < ln; *pos++ {
		if pp >= vSegs {
			return parseError{1, errInvalidSemVerString}
		}

		switch true {

		case vn[*pos] >= digit0 && vn[*pos] <= digit9:
			buf[bp] = vn[*pos] - '0'
			bp++

		case vn[*pos] == segDivider:
			*parts[pp] = bufToU8(buf[:bp])
			pp++
			bp = 0

		default:
			switch vn[*pos] {
			case preDivider, buildDivider:
				return nil
			default:
				return parseError{2, errInvalidSemVerString}
			}
		}
	}

	*parts[pp] = bufToU8(buf[:bp])

	return nil
}

func parsePrerelease(vn []byte, pos *uint8, buf *[parseBufferSize]byte, ver *Version) {
	segments := uint8(1)
	preLen := uint8(len(vn))

	for i := *pos; i < preLen; i++ {
		if vn[i] == segDivider {
			segments++
		} else if vn[i] == buildDivider {
			break
		}
	}

	bufPos := 0
	segIndex := 0
	ver.Prerelease = make([]string, segments)

	for ; *pos < preLen; *pos++ {
		if vn[*pos] == segDivider {
			ver.Prerelease[segIndex] = string(buf[:bufPos])
			segIndex++
			bufPos = 0
		} else if vn[*pos] == buildDivider {
			break
		} else {
			buf[bufPos] = vn[*pos]
			bufPos++
		}
	}

	ver.Prerelease[segIndex] = string(buf[:bufPos])
}

func parseBuild(vn []byte, pos *uint8, buf *[parseBufferSize]byte, ver *Version) {
	segments := uint8(1)
	buildLen := uint8(len(vn))

	for i := *pos; i < buildLen; i++ {
		if vn[i] == segDivider {
			segments++
		}
	}

	ver.Build = make([]string, segments)

	bufPos := 0
	segIndex := 0
	ver.Build = make([]string, segments)

	for ; *pos < buildLen; *pos++ {
		if vn[*pos] == segDivider {
			ver.Build[segIndex] = string(buf[:bufPos])
			segIndex++
			bufPos = 0
		} else {
			buf[bufPos] = vn[*pos]
			bufPos++
		}
	}

	ver.Build[segIndex] = string(buf[:bufPos])
}

func bufToU8(buf []byte) (val uint8) {
	up := uint8(0)


	for dn := int8(len(buf)) - 1; dn >= 0; dn-- {
		val += buf[dn] * util.PowU8(10, up)
		up++
	}

	return
}

func roBytes(str *string) []byte {
	return *(*[]byte)(unsafe.Pointer(str))
}

type parseError struct {
	code uint8
	err string
}

func (p parseError) Error() string {
	return p.err + " code " + string(p.code + '0')
}
