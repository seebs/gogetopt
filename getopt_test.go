package gogetopt

import (
	// "errors"
	"testing"
)

// a test case maps an option string and set of
// arguments to expected outputs
type testcase struct {
	optstring string
	args      []string
	opts      map[string]Option
	remaining []string
	err       string
}

var testcases = []testcase{
	{optstring: "aa", err: "duplicate option specifiers for 'a'"},
	{optstring: "a:#", err: "option type specifier without option"},
	{optstring: "a?", err: "invalid option specifier '?'"},
	{optstring: "a", args: []string{"-aa"}, err: "duplicate option 'a'"},
	{optstring: "ab", args: []string{"foo"}, remaining: []string{"foo"}},
	{optstring: "a", args: []string{"-a", "foo"}, remaining: []string{"foo"},
		opts: map[string]Option{"a": Option{}}},
	{optstring: "a+", args: []string{"-aa", "foo"}, remaining: []string{"foo"},
		opts: map[string]Option{"a": Option{Int: 2}}},
	{optstring: "n#", args: []string{"-n", "2", "foo"}, remaining: []string{"foo"},
		opts: map[string]Option{"n": Option{Value: "2", Int: 2}}},
	{optstring: "n.", args: []string{"-n", "2", "foo"}, remaining: []string{"foo"},
		opts: map[string]Option{"n": Option{Value: "2", Float: 2.0}}},
	{optstring: "a", args: []string{"--", "-a"}, remaining: []string{"-a"}},
}

func diffRemaining(got []string, expected []string) bool {
	if len(got) != len(expected) {
		return true
	}
	for k, v := range got {
		if expected[k] != v {
			return true
		}
	}
	return false
}

func TestCases(t *testing.T) {
	for idx, tc := range testcases {
		t.Logf("test case %d: optstring '%s'", idx, tc.optstring)
		opts, remaining, err := GetOpt(tc.args, tc.optstring)
		if diffRemaining(remaining, tc.remaining) {
			t.Logf("remaining args mismatch: expecting %v, got %v",
				tc.remaining, remaining)
			t.Fail()
		}
		if tc.err == "" && err != nil {
			t.Logf("unexpected error: '%s'", err.Error())
			t.Fail()
		}
		if tc.err != "" && err == nil {
			t.Logf("missing expected error: '%s'", tc.err)
			t.Fail()
		}
		if err != nil {
			if err.Error() != tc.err {
				t.Logf("error mismatch: expected '%s', got '%s'", tc.err, err.Error())
				t.Fail()
			}
			// if an error occurred, and it was expected, don't check return
			// value, because we don't specify that
			continue
		}
		for k, v := range tc.opts {
			if opts[k] == nil || *opts[k] != v {
				t.Logf("option mismatch for opts[%s]: got %v, expecting %v",
					k, opts[k], v)
				t.Fail()
			}
			delete(opts, k)
		}
		if len(opts) > 0 {
			var first string
			for k := range opts {
				first = k
				break
			}
			t.Logf("unexpected option(s) [%d] in response: '%s'",
				len(opts), first)
			t.Fail()
		}
	}
}
