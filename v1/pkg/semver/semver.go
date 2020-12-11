package semver

import (
	"github.com/foxcapades/go-bytify/v0/bytify"
)

// Version holds the components of a SemVer version number.
type Version struct {
	Major      uint8
	Minor      uint8
	Patch      uint8
	Build      []string
	Prerelease []string
}

// Equivalent returns whether this Version is the same as the given version
// when comparing the Major, Minor, Patch, and Prerelease values.
func (v *Version) Equivalent(other *Version) bool {
	switch false {
	case v.Major == other.Major:
		return false
	case v.Minor == other.Minor:
		return false
	case v.Patch == other.Patch:
		return false
	}

	tln := len(v.Prerelease)

	if tln != len(other.Prerelease) {
		return false
	}

	for i := 0; i < tln; i++ {
		if v.Prerelease[i] != other.Prerelease[i] {
			return false
		}
	}

	return true
}

// Equal returns whether this Version is the same as the given version comparing
// all fields, including prerelease or build tags.
func (v *Version) Equal(other *Version) bool {
	if !v.Equivalent(other) {
		return false
	}

	if len(v.Build) != len(other.Build) {
		return false
	}

	return true
}

// IsAfter returns whether the current Version is a later version than the given
// value.
func (v *Version) IsAfter(other *Version) bool {
	switch true {
	case v.Major > other.Major:
		return true
	case v.Minor > other.Minor:
		return true
	case v.Patch > other.Patch:
		return true
	}

	return len(v.Prerelease) == 0 && len(other.Prerelease) > 0
}

// IsAfter returns whether the current Version is an earlier version number than
// the given value.
func (v *Version) IsBefore(other *Version) bool {
	switch true {
	case v.Major < other.Major:
		return true
	case v.Minor < other.Minor:
		return true
	case v.Patch < other.Patch:
		return true
	}

	return len(other.Prerelease) == 0 && len(v.Prerelease) > 0
}

// VString prints the string form of this Version with a leading 'v' character.
func (v *Version) VString() string {
	out := make([]byte, v.outSize()+1)
	out[0] = leader
	v.stringFill(out, 1)

	return string(out)
}

// String prints the string form of this Version.
func (v *Version) String() string {
	out := make([]byte, v.outSize())
	v.stringFill(out, 0)

	return string(out)
}

func (v *Version) outSize() (size uint8) {
	size += bytify.Uint8StringSize(v.Major) + bytify.Uint8StringSize(v.Minor)
	size += bytify.Uint8StringSize(v.Patch) + 2

	if ln := uint8(len(v.Prerelease)); ln > 0 {
		// Add one for the hyphen separator
		size++

		for i := uint8(0); i < ln; {
			size += uint8(len(v.Prerelease[i]))

			if i++; i < ln {
				// Add one for the period separator.
				size++
			}
		}
	}

	if ln := uint8(len(v.Build)); ln > 0 {
		// Add one for the plus separator
		size++

		for i := uint8(0); i < ln; {
			size += uint8(len(v.Build[i]))

			if i++; i < ln {
				// Add one for the period separator.
				size++
			}
		}
	}

	return
}

func (v *Version) stringFill(out []byte, pos uint8) {
	pos += bytify.Uint8ToBytes(v.Major, out[pos:])
	out[pos] = segDivider
	pos++

	pos += bytify.Uint8ToBytes(v.Minor, out[pos:])
	out[pos] = segDivider
	pos++

	pos += bytify.Uint8ToBytes(v.Patch, out[pos:])

	if ln := uint8(len(v.Prerelease)); ln > 0 {
		out[pos] = preDivider
		pos++
		for i := uint8(0); i < ln; {
			copy(out[pos:], v.Prerelease[i])
			pos += uint8(len(v.Prerelease[i]))

			if i++; i < ln {
				out[pos] = segDivider
				pos++
			}
		}
	}

	if ln := uint8(len(v.Build)); ln > 0 {
		out[pos] = buildDivider
		pos++
		for i := uint8(0); i < ln; {
			copy(out[pos:], v.Build[i])
			pos += uint8(len(v.Build[i]))

			if i++; i < ln {
				out[pos] = segDivider
				pos++
			}
		}
	}
}