package osenv

import (
	"os"
	"testing"
	"time"
)

func TestEnvString(t *testing.T) {
	if err := os.Setenv("TESTENVSTRING", "go test"); err != nil {
		t.Fatal(err)
	}
	type args struct {
		key    string
		defval string
	}
	tests := []struct {
		name    string
		args    args
		wantVal string
	}{
		{"default", args{"MOCKBA", "Moscow"}, "Moscow"},
		{"existing", args{"TESTENVSTRING", "someval"}, "go test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotVal := String(tt.args.key, tt.args.defval); gotVal != tt.wantVal {
				t.Errorf("EnvString() = %v, want %v", gotVal, tt.wantVal)
			}
		})
	}
}

func TestEnvBool(t *testing.T) {
	var testVals = map[string]string{
		"TESTENVTRUE":    "true",
		"TESTENVFALSE":   "false",
		"TESTENVINVALID": "invalid",
		"TESTENVNOTSET":  "",
	}

	for k, v := range testVals {
		if err := os.Setenv(k, v); err != nil {
			t.Fatal(err)
		}
	}
	type args struct {
		key    string
		defval bool
	}
	tests := []struct {
		name    string
		args    args
		wantVal bool
	}{
		{"true/false->true", args{"TESTENVTRUE", false}, true},
		{"false/true->false", args{"TESTENVFALSE", true}, false},
		{"invalid/false->false", args{"TESTENVINVALID", false}, false},
		{"invalid/true->true", args{"TESTENVINVALID", true}, true},
		{"notset/true->true", args{"TESTENVNOTSET", true}, true},
		{"notexist/true->true", args{"TESTENVNOTEXIST", true}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotVal := Bool(tt.args.key, tt.args.defval); gotVal != tt.wantVal {
				t.Errorf("EnvBool() = %v, want %v", gotVal, tt.wantVal)
			}
		})
	}
}

func TestDuration(t *testing.T) {
	var testVals = map[string]string{
		"TEST1H":         "1h",
		"TEST35S":        "35s",
		"TESTDURINVALID": "some value",
	}

	for k, v := range testVals {
		if err := os.Setenv(k, v); err != nil {
			t.Fatal(err)
		}
	}

	type args struct {
		key     string
		defavlt time.Duration
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		{"default", args{"N$T_HERE", time.Duration(5 * time.Second)}, time.Duration(5 * time.Second)},
		{"1h", args{"TEST1H", 42 * time.Hour}, 1 * time.Hour},
		{"35s", args{"TEST35S", 42 * time.Hour}, 35 * time.Second},
		{"invalid", args{"TESTDURINVALID", 42 * time.Hour}, 42 * time.Hour},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Duration(tt.args.key, tt.args.defavlt); got != tt.want {
				t.Errorf("Duration() = %v, want %v", got, tt.want)
			}
		})
	}
}
