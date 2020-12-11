package semver_test

import (
	"testing"

	"github.com/foxcapades/gVersion/v1/pkg/semver"
)

func TestVersion_String(t *testing.T) {
	tests := []struct {
		expect string
		input  semver.Version
	}{
		{"0.0.1", semver.Version{Patch: 1}},
		{"0.1.0", semver.Version{Minor: 1}},
		{"1.0.0", semver.Version{Major: 1}},
		{"1.0.1", semver.Version{Major: 1, Patch: 1}},
		{"1.1.0", semver.Version{Major: 1, Minor: 1}},
		{"1.1.1", semver.Version{Major: 1, Minor: 1, Patch: 1}},
		{"0.0.0-alpha", semver.Version{Prerelease: []string{"alpha"}}},
		{"0.0.0-alpha.v1", semver.Version{Prerelease: []string{"alpha", "v1"}}},
		{"0.0.0+2020-09-18", semver.Version{Build: []string{"2020-09-18"}}},
		{"0.0.0+2020-09-18.b21", semver.Version{Build: []string{"2020-09-18", "b21"}}},
		{"0.0.0-alpha+2020-09-18", semver.Version{Prerelease: []string{"alpha"}, Build: []string{"2020-09-18"}}},
		{"0.0.0-alpha.v1+2020-09-18.b21", semver.Version{Prerelease: []string{"alpha", "v1"}, Build: []string{"2020-09-18", "b21"}}},
	}

	for _, test := range tests {
		t.Run(test.expect, func(t *testing.T) {
			val := test.input.String()
			if test.expect != test.input.String() {
				t.Errorf("Expected %s got %s", test.expect, val)
			}
		})
	}
}

func TestVersion_VString(t *testing.T) {
	tests := []struct {
		expect string
		input  semver.Version
	}{
		{"v0.0.1", semver.Version{Patch: 1}},
		{"v0.1.0", semver.Version{Minor: 1}},
		{"v1.0.0", semver.Version{Major: 1}},
		{"v1.0.1", semver.Version{Major: 1, Patch: 1}},
		{"v1.1.0", semver.Version{Major: 1, Minor: 1}},
		{"v1.1.1", semver.Version{Major: 1, Minor: 1, Patch: 1}},
		{"v0.0.0-alpha", semver.Version{Prerelease: []string{"alpha"}}},
		{"v0.0.0-alpha.v1", semver.Version{Prerelease: []string{"alpha", "v1"}}},
		{"v0.0.0+2020-09-18", semver.Version{Build: []string{"2020-09-18"}}},
		{"v0.0.0+2020-09-18.b21", semver.Version{Build: []string{"2020-09-18", "b21"}}},
		{"v0.0.0-alpha+2020-09-18", semver.Version{Prerelease: []string{"alpha"}, Build: []string{"2020-09-18"}}},
		{"v0.0.0-alpha.v1+2020-09-18.b21", semver.Version{Prerelease: []string{"alpha", "v1"}, Build: []string{"2020-09-18", "b21"}}},
	}

	for _, test := range tests {
		t.Run(test.expect, func(t *testing.T) {
			val := test.input.String()
			if test.expect != test.input.VString() {
				t.Errorf("Expected %s got %s", test.expect, val)
			}
		})
	}
}

func TestVersion_Equivalent(t *testing.T) {
	tests := []struct {
		name string
		a, b semver.Version
		se bool
	}{
		{
			name: "equal major, minor, and patch",
			a: semver.Version{Major: 1, Minor: 10, Patch: 2},
			b: semver.Version{Major: 1, Minor: 10, Patch: 2},
			se: true,
		},
		{
			name: "one side has a build tag",
			a: semver.Version{Major: 1, Minor: 10, Patch: 2, Build: []string{"24"}},
			b: semver.Version{Major: 1, Minor: 10, Patch: 2},
			se: true,
		},
		{
			name: "one side has a prerelease tag",
			a: semver.Version{Major: 1, Minor: 10, Patch: 2, Prerelease: []string{"24"}},
			b: semver.Version{Major: 1, Minor: 10, Patch: 2},
			se: false,
		},
		{
			name: "different major 1",
			a: semver.Version{Major: 2, Minor: 10, Patch: 2},
			b: semver.Version{Major: 1, Minor: 10, Patch: 2},
			se: false,
		},
		{
			name: "different major 2",
			a: semver.Version{Major: 1, Minor: 10, Patch: 2},
			b: semver.Version{Major: 2, Minor: 10, Patch: 2},
			se: false,
		},
		{
			name: "different minor 1",
			a: semver.Version{Major: 1, Minor: 11, Patch: 2},
			b: semver.Version{Major: 1, Minor: 10, Patch: 2},
			se: false,
		},
		{
			name: "different minor 2",
			a: semver.Version{Major: 1, Minor: 10, Patch: 2},
			b: semver.Version{Major: 1, Minor: 11, Patch: 2},
			se: false,
		},
		{
			name: "different patch 1",
			a: semver.Version{Major: 1, Minor: 10, Patch: 3},
			b: semver.Version{Major: 1, Minor: 10, Patch: 2},
			se: false,
		},
		{
			name: "different patch 2",
			a: semver.Version{Major: 1, Minor: 10, Patch: 2},
			b: semver.Version{Major: 1, Minor: 10, Patch: 3},
			se: false,
		},
		{
			name: "different prerelease 1",
			a: semver.Version{Major: 1, Minor: 10, Patch: 2, Prerelease: []string{"hi"}},
			b: semver.Version{Major: 1, Minor: 10, Patch: 2, Prerelease: []string{"ho"}},
			se: false,
		},
		{
			name: "different prerelease 2",
			a: semver.Version{Major: 1, Minor: 10, Patch: 2, Prerelease: []string{"ho"}},
			b: semver.Version{Major: 1, Minor: 10, Patch: 2, Prerelease: []string{"hi"}},
			se: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.a.Equivalent(&test.b) != test.se {
				t.Fail()
			}
		})
	}
}

func TestVersion_Equal(t *testing.T) {
	tests := []struct {
		name string
		a, b semver.Version
		se bool
	}{
		{
			name: "equal major, minor, and patch",
			a: semver.Version{Major: 1, Minor: 10, Patch: 2},
			b: semver.Version{Major: 1, Minor: 10, Patch: 2},
			se: true,
		},
		{
			name: "one side has a build tag",
			a: semver.Version{Major: 1, Minor: 10, Patch: 2, Build: []string{"24"}},
			b: semver.Version{Major: 1, Minor: 10, Patch: 2},
			se: false,
		},
		{
			name: "one side has a prerelease tag",
			a: semver.Version{Major: 1, Minor: 10, Patch: 2, Prerelease: []string{"24"}},
			b: semver.Version{Major: 1, Minor: 10, Patch: 2},
			se: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.a.Equal(&test.b) != test.se {
				t.Fail()
			}
		})
	}
}

func TestVersion_IsAfter(t *testing.T) {
	tests := []struct {
		name string
		a, b semver.Version
		se bool
	}{
		{
			name: "different major 1",
			a: semver.Version{Major: 2, Minor: 10, Patch: 2},
			b: semver.Version{Major: 1, Minor: 10, Patch: 2},
			se: true,
		},
		{
			name: "different major 2",
			a: semver.Version{Major: 1, Minor: 10, Patch: 2},
			b: semver.Version{Major: 2, Minor: 10, Patch: 2},
			se: false,
		},
		{
			name: "different minor 1",
			a: semver.Version{Major: 1, Minor: 12, Patch: 2},
			b: semver.Version{Major: 1, Minor: 11, Patch: 2},
			se: true,
		},
		{
			name: "different minor 2",
			a: semver.Version{Major: 1, Minor: 11, Patch: 2},
			b: semver.Version{Major: 1, Minor: 12, Patch: 2},
			se: false,
		},
		{
			name: "different patch 1",
			a: semver.Version{Major: 1, Minor: 10, Patch: 3},
			b: semver.Version{Major: 1, Minor: 10, Patch: 2},
			se: true,
		},
		{
			name: "different patch 2",
			a: semver.Version{Major: 1, Minor: 10, Patch: 2},
			b: semver.Version{Major: 1, Minor: 10, Patch: 3},
			se: false,
		},
		{
			name: "prerelease 1",
			a: semver.Version{Major: 1, Minor: 10, Patch: 2},
			b: semver.Version{Major: 1, Minor: 10, Patch: 2, Prerelease: []string{"24"}},
			se: true,
		},
		{
			name: "prerelease 2",
			a: semver.Version{Major: 1, Minor: 10, Patch: 2, Prerelease: []string{"24"}},
			b: semver.Version{Major: 1, Minor: 10, Patch: 2},
			se: false,
		},
		{
			name: "build 1",
			a: semver.Version{Major: 1, Minor: 10, Patch: 2},
			b: semver.Version{Major: 1, Minor: 10, Patch: 2, Build: []string{"24"}},
			se: false,
		},
		{
			name: "build 2",
			a: semver.Version{Major: 1, Minor: 10, Patch: 2, Build: []string{"24"}},
			b: semver.Version{Major: 1, Minor: 10, Patch: 2},
			se: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.a.IsAfter(&test.b) != test.se {
				t.Fail()
			}
		})
	}
}


func TestVersion_IsBefore(t *testing.T) {
	tests := []struct {
		name string
		a, b semver.Version
		se bool
	}{
		{
			name: "different major 1",
			a: semver.Version{Major: 2, Minor: 10, Patch: 2},
			b: semver.Version{Major: 1, Minor: 10, Patch: 2},
			se: false,
		},
		{
			name: "different major 2",
			a: semver.Version{Major: 1, Minor: 10, Patch: 2},
			b: semver.Version{Major: 2, Minor: 10, Patch: 2},
			se: true,
		},
		{
			name: "different minor 1",
			a: semver.Version{Major: 1, Minor: 12, Patch: 2},
			b: semver.Version{Major: 1, Minor: 11, Patch: 2},
			se: false,
		},
		{
			name: "different minor 2",
			a: semver.Version{Major: 1, Minor: 11, Patch: 2},
			b: semver.Version{Major: 1, Minor: 12, Patch: 2},
			se: true,
		},
		{
			name: "different patch 1",
			a: semver.Version{Major: 1, Minor: 10, Patch: 3},
			b: semver.Version{Major: 1, Minor: 10, Patch: 2},
			se: false,
		},
		{
			name: "different patch 2",
			a: semver.Version{Major: 1, Minor: 10, Patch: 2},
			b: semver.Version{Major: 1, Minor: 10, Patch: 3},
			se: true,
		},
		{
			name: "prerelease 1",
			a: semver.Version{Major: 1, Minor: 10, Patch: 2},
			b: semver.Version{Major: 1, Minor: 10, Patch: 2, Prerelease: []string{"24"}},
			se: false,
		},
		{
			name: "prerelease 2",
			a: semver.Version{Major: 1, Minor: 10, Patch: 2, Prerelease: []string{"24"}},
			b: semver.Version{Major: 1, Minor: 10, Patch: 2},
			se: true,
		},
		{
			name: "build 1",
			a: semver.Version{Major: 1, Minor: 10, Patch: 2},
			b: semver.Version{Major: 1, Minor: 10, Patch: 2, Build: []string{"24"}},
			se: false,
		},
		{
			name: "build 2",
			a: semver.Version{Major: 1, Minor: 10, Patch: 2, Build: []string{"24"}},
			b: semver.Version{Major: 1, Minor: 10, Patch: 2},
			se: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.a.IsBefore(&test.b) != test.se {
				t.Fail()
			}
		})
	}
}
