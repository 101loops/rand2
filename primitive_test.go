package randgen

import (
	. "github.com/101loops/bdd"
)

var _ = Describe("Rand Util", func() {

	It("return random string", func() {

		rand := String(5)
		Check(rand, HasLen, 5)

		rand = String(50)
		Check(rand, HasLen, 50)
	})
})
