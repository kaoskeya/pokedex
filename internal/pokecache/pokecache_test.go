package pokecache

import (
	"bytes"
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	cases := []struct {
		init       map[string][]byte
		check      []string
		expectVal  []([]byte)
		expectFlag []bool
	}{
		{
			init: map[string][]byte{
				"a": []byte("value1"),
				"b": []byte("value2"),
			},
			check:      []string{"a", "c"},
			expectVal:  [][]byte{[]byte("value1"), nil},
			expectFlag: []bool{true, false},
		},
		{
			init:       map[string][]byte{},
			check:      []string{"a"},
			expectVal:  [][]byte{nil},
			expectFlag: []bool{false},
		},
		{
			init: map[string][]byte{
				"x": []byte("value1"),
				"y": []byte("value1"),
			},
			check:      []string{"x", "y"},
			expectVal:  [][]byte{[]byte("value1"), []byte("value1")},
			expectFlag: []bool{true, true},
		},
	}

	for caseIter, caseVal := range cases {
		c := NewCache(5 * time.Second)

		for k, v := range caseVal.init {
			c.Add(k, v)
		}

		for i := range caseVal.check {
			val, flag := c.Get(caseVal.check[i])
			if !bytes.Equal(val, caseVal.expectVal[i]) {
				t.Fatalf("fail: case %d outputs mismatch, expected: %v, got: %v", caseIter+1, string(caseVal.expectVal[i]), string(val))
			}
			if flag != caseVal.expectFlag[i] {
				t.Fatalf("fail: case %d outputs mismatch, expected: %v, got: %v", caseIter+1, caseVal.expectFlag[i], flag)
			}
		}
	}
}

func TestReap(t *testing.T) {
	for _, tc := range []struct {
		name  string
		age   time.Duration
		found bool
	}{
		{name: "fresh", age: time.Second, found: true},
		{name: "expired", age: 10 * time.Second, found: false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			cache := NewCache(time.Hour)
			cache.Add("page", []byte("locations"))
			cache.reap(time.Now().Add(tc.age), 5*time.Second)
			_, ok := cache.Get("page")
			if ok != tc.found {
				t.Fatalf("Get after reaping = %v; want %v", ok, tc.found)
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	cache := NewCache(20 * time.Millisecond)
	cache.Add("page", []byte("locations"))
	deadline := time.Now().Add(time.Second)
	for time.Now().Before(deadline) {
		if _, ok := cache.Get("page"); !ok {
			return
		}
		time.Sleep(5 * time.Millisecond)
	}
	t.Fatal("expected the background loop to remove the expired entry")
}
