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
		{"default", args{"BRISBANE", "Brisbane"}, "Brisbane"},
		{"existing", args{"TESTENVSTRING", "someval"}, "go test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotVal := Value(tt.args.key, tt.args.defval); gotVal != tt.wantVal {
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
			if gotVal := Value(tt.args.key, tt.args.defval); gotVal != tt.wantVal {
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
			if got := Value(tt.args.key, tt.args.defavlt); got != tt.want {
				t.Errorf("Duration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTime(t *testing.T) {
	var testVals = map[string]string{
		"TESTTIME":        "2021-03-26T13:47:34Z",
		"INVALIDTIME":     "xxxx-xx-xx",
		"UNSUPPORTEDTIME": "2021-03-26 13:47:34Z",
	}

	var defDate = time.Date(2019, 9, 16, 5, 6, 7, 0, time.UTC)

	for k, v := range testVals {
		if err := os.Setenv(k, v); err != nil {
			t.Fatal(err)
		}
	}

	type args struct {
		key     string
		defavlt time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{"default", args{"N$T_HERE", defDate}, defDate},
		{"empty default", args{"N$T_HERE", time.Time{}}, time.Time{}},
		{"test time is set", args{"TESTTIME", defDate}, time.Date(2021, 03, 26, 13, 47, 34, 0, time.UTC)},
		{"invalid format", args{"INVALIDTIME", defDate}, defDate},
		{"invalid format", args{"UNSUPPORTEDTIME", defDate}, defDate},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Value(tt.args.key, tt.args.defavlt); got != tt.want {
				t.Errorf("Duration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSecret(t *testing.T) {
	const varName = "TEST_SECRET"
	const sTest = "blah"
	os.Setenv(varName, sTest)
	v := Secret(varName, "fail")
	if v != sTest {
		t.Errorf("Secret() failed want=%s got=%s", sTest, v)
	}
	clearedV := os.Getenv(varName)
	if clearedV != "" {
		t.Errorf("value not cleared: %s", clearedV)
	}
}
