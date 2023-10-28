package split

import "testing"

func TestSplit(t *testing.T) {
	cases := []struct {
		sql string
		cnt int
	}{
		// ignore empty sql as ';' or '\s;'
		// expect 2
		{`SELECT * FROM a WHERE a=';';   ;;SELECT 2;`, 2},
		// no terminal token
		// expect 1
		{`SELECT 1;SELECT 2`, 1},
	}

	for i := range cases {
		items := SplitWithScanner(cases[i].sql)
		cnt := len(items)
		if cnt != cases[i].cnt {
			t.Fatalf("expected %d, got %d", cases[i].cnt, cnt)
		}
	}
}
