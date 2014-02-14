package testlib

import (
	"bytes"
	"fmt"
	"image/png"

	"github.com/remogatto/imagetest"
	"github.com/remogatto/mandala"
)

func (t *TestSuite) TestDraw() {
	request := mandala.LoadResourceRequest{
		Filename: "drawable/expected.png",
		Response: make(chan mandala.LoadResourceResponse),
	}

	mandala.ResourceManager() <- request
	response := <-request.Response
	buffer := response.Buffer

	t.True(response.Error == nil, "An error occured during resource opening")

	if buffer != nil {
		exp, err := png.Decode(bytes.NewBuffer(buffer))
		t.True(err == nil, "An error occured during png decoding")

		distance := imagetest.CompareDistance(exp, <-t.testDraw, nil)
		t.True(distance < 0.1, fmt.Sprintf("Distance is %f", distance))
	}
}
