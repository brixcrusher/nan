package nan

import "testing"

func TestSQLNullInt64(t *testing.T) {
	v1, v2 := int64(7676), int64(7676)
	doSQLTest(t, v1, v2, &NullInt64{Int64: v1, Valid: true}, &NullInt64{})

	doSQLTest(t, nil, nil, &NullInt64{Valid: false}, &NullInt64{})
	doSQLTest(t, nil, nil, &NullInt64{Valid: false}, &NullInt64{Valid: true})
}
