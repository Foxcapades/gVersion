package semver_test

import (
	"fmt"
	"testing"

	"github.com/foxcapades/gVersion/v1/pkg/semver"
)

func TestParse(t *testing.T) {
	tests := []struct {
		version string
		output  semver.Version
	}{
		{"v0.0.1", semver.Version{Patch: 1}},
		{"v0.1.0", semver.Version{Minor: 1}},
		{"v1.0.0", semver.Version{Major: 1}},
		{"v1.0.1", semver.Version{Major: 1, Patch: 1}},
		{"v1.1.0", semver.Version{Major: 1, Minor: 1}},
		{"v1.1.1", semver.Version{Major: 1, Minor: 1, Patch: 1}},

		{"v0.0.10", semver.Version{Patch: 10}},
		{"v0.10.0", semver.Version{Minor: 10}},
		{"v10.0.0", semver.Version{Major: 10}},
		{"v10.0.10", semver.Version{Major: 10, Patch: 10}},
		{"v10.10.0", semver.Version{Major: 10, Minor: 10}},
		{"v10.10.10", semver.Version{Major: 10, Minor: 10, Patch: 10}},

		{"v0.0.0-alpha", semver.Version{Prerelease: []string{"alpha"}}},
		{"v0.0.0-alpha.v1", semver.Version{Prerelease: []string{"alpha", "v1"}}},
		{"v0.0.0+2020-09-18", semver.Version{Build: []string{"2020-09-18"}}},
		{"v0.0.0+2020-09-18.b21", semver.Version{Build: []string{"2020-09-18", "b21"}}},
		{"v0.0.0-alpha+2020-09-18", semver.Version{Prerelease: []string{"alpha"}, Build: []string{"2020-09-18"}}},
		{"v0.0.0-alpha.v1+2020-09-18.b21", semver.Version{Prerelease: []string{"alpha", "v1"}, Build: []string{"2020-09-18", "b21"}}},
	}
	for _, test := range tests {
		t.Run(test.version, func(t *testing.T) {
			vs, err := semver.Parse(test.version)

			if err != nil {
				t.Error("expected no error, got ", err)
			}

			if vs.Major != test.output.Major {
				t.Errorf("Major: Expected %d, got %d", test.output.Major, vs.Major)
			}

			if vs.Minor != test.output.Minor {
				t.Errorf("Minor: Expected %d, got %d", test.output.Minor, vs.Minor)
			}

			if vs.Patch != test.output.Patch {
				t.Errorf("Patch: Expected %d, got %d", test.output.Patch, vs.Patch)
			}

			if len(vs.Prerelease) != len(test.output.Prerelease) {
				t.Errorf(
					"Expected %d prerelease segments, got %d",
					len(test.output.Prerelease),
					len(vs.Prerelease),
				)
			}

			for i := range test.output.Prerelease {
				if vs.Prerelease[i] != test.output.Prerelease[i] {
					t.Errorf(
						"Expected prerelease segment %d to equal %s, got %s",
						i, test.output.Prerelease[i], vs.Prerelease[i],
					)
				}
			}

			if len(vs.Build) != len(test.output.Build) {
				t.Errorf(
					"Expected %d prerelease segments, got %d",
					len(test.output.Build),
					len(vs.Build),
				)
			}

			for i := range test.output.Build {
				if vs.Build[i] != test.output.Build[i] {
					t.Errorf(
						"Expected build segment %d to equal %s, got %s",
						i, test.output.Build[i], vs.Build[i],
					)
				}
			}
		})
	}
}

func TestParse1(t *testing.T) {
	tests := [][2]string {
		{"va.0.1", "invalid format for a semantic version string code 2"},
		{"v0.a.0", "invalid format for a semantic version string code 2"},
		{"v1.0.a", "invalid format for a semantic version string code 2"},
		{"v1.0.0.1", "invalid format for a semantic version string code 1"},
	}

	for _, test := range tests {
		_, err := semver.Parse(test[0])

		if err == nil {
			t.Error("expected no error, got ", err)
		}
		if err.Error() != test[1] {
			t.Error(err.Error(), test[1])
		}
	}
}

var hold semver.Version
func Benchmark(b *testing.B) {
	benchmarks := []string {
		"v0.0.1",
		"v0.1.0",
		"v1.0.0",
		"v1.0.1",
		"v1.1.0",
		"v1.1.1",
		"v0.0.0-alpha",
		"v0.0.0-alpha.v1",
		"v0.0.0+2020-09-18",
		"v0.0.0+2020-09-18.b21",
		"v0.0.0-alpha+2020-09-18",
		"v0.0.0-alpha.v1+2020-09-18.b21",
	}

	for _, bm := range benchmarks {
		b.Run(bm, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				hold, _ = semver.Parse(bm)
			}
		})
	}
}

func ExampleParse() {
	version := "v1.23.0-alpha+b58"

	ver, _ := semver.Parse(version)

	fmt.Println(ver.Major)
	fmt.Println(ver.Minor)
	fmt.Println(ver.Patch)
	fmt.Println(ver.Prerelease)
	fmt.Println(ver.Build)

	// Output:
	// 1
	// 23
	// 0
	// [alpha]
	// [b58]
}
