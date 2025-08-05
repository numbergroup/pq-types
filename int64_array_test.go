package pq_types

import (
	"database/sql"
	"fmt"

	. "github.com/numbergroup/check"
)

func (s *TypesSuite) TestInt64Array(c *C) {
	type testData struct {
		a Int64Array
		b []byte
	}
	for _, d := range []testData{
		{Int64Array(nil), []byte(nil)},
		{Int64Array{}, []byte(`{}`)},
		{Int64Array{1}, []byte(`{1}`)},
		{Int64Array{1, 0, -3}, []byte(`{1,0,-3}`)},
		{Int64Array{-3, 0, 1}, []byte(`{-3,0,1}`)},
	} {
		s.SetUpTest(c)

		_, err := s.db.Exec("INSERT INTO pq_types (int64_array) VALUES($1)", d.a)
		c.Assert(err, IsNil)

		b1 := []byte("42")
		a1 := Int64Array{42}
		err = s.db.QueryRow("SELECT int64_array, int64_array FROM pq_types").Scan(&b1, &a1)
		c.Check(err, IsNil)
		c.Check(b1, DeepEquals, d.b, Commentf("\nb1  = %#q\nd.b = %#q", b1, d.b))
		c.Check(a1, DeepEquals, d.a)

		// check db array length
		var length sql.NullInt64
		err = s.db.QueryRow("SELECT array_length(int64_array, 1) FROM pq_types").Scan(&length)
		c.Check(err, IsNil)
		c.Check(length.Valid, Equals, len(d.a) > 0)
		c.Check(length.Int64, Equals, int64(len(d.a)))

		// check db array elements
		for i := range len(d.a) {
			q := fmt.Sprintf("SELECT int64_array[%d] FROM pq_types", i+1)
			var el sql.NullInt64
			err = s.db.QueryRow(q).Scan(&el)
			c.Check(err, IsNil)
			c.Check(el.Valid, Equals, true)
			c.Check(el.Int64, Equals, int64(d.a[i]))
		}
	}
}

func (s *TypesSuite) TestInt64ArrayEqualWithoutOrder(c *C) {
	c.Check(Int64Array{1, 0, -3}.EqualWithoutOrder(Int64Array{-3, 0, 1}), Equals, true)
	c.Check(Int64Array{1, 0, -3}.EqualWithoutOrder(Int64Array{1}), Equals, false)
	c.Check(Int64Array{1, 0, -3}.EqualWithoutOrder(Int64Array{1, 0, 42}), Equals, false)
	c.Check(Int64Array{}.EqualWithoutOrder(Int64Array{}), Equals, true)
	c.Check(Int64Array{}.EqualWithoutOrder(Int64Array{1}), Equals, false)
}
