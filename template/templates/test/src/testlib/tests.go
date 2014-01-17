package testlib

func (t *TestSuite) TestDraw() {
	t.True(<-t.testDraw)
}
