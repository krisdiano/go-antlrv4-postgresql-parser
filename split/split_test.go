package split

import "testing"

func TestSplit(t *testing.T) {
	cases := []struct {
		sql    string
		cnt    int
		hasErr bool
	}{
		// ignore empty sql as ';' or '\s;'
		// expect 2
		{`;;;SELECT * FROM a WHERE a=';';   ;;SELECT 2;`, 2, false},
		// no terminal token
		// expect 2 and err
		{`SELECT 1;SELECT 2`, 2, true},
		{`SELECT 1;SELECT 2;`, 2, false},
	}

	for i := range cases {
		items, err := SplitWithScanner(cases[i].sql)
		cnt := len(items)
		if cnt != cases[i].cnt {
			t.Fatalf("sql `%s` expected %d, got %d", cases[i].sql, cases[i].cnt, cnt)
		}
		hasErr := err != nil
		if hasErr != cases[i].hasErr {
			t.Fatalf("sql `%s`, expected %v, got %v", cases[i].sql, cases[i].hasErr, hasErr)
		}
	}
}
