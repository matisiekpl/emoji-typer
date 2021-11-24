// +build !skippackr
// Code generated by github.com/gobuffalo/packr/v2. DO NOT EDIT.

// You can use the "packr2 clean" command to clean up this,
// and any other packr generated files.
package packrd

import (
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/packr/v2/file/resolver"
)

var _ = func() error {
	const gk = "e8ad54d574961e8fa5f52e380e654a7b"
	g := packr.New(gk, "")
	hgr, err := resolver.NewHexGzip(map[string]string{
		"d6469c13c8c86b6885d49e438465d0ef": "1f8b08000000000000ff84914b4e23311086f77d0a8f673b8927db916d6924c409b880d32e70811f8d5d4e082daf3904d7817b21a73b4282054b3feafb1f251d05af07e9c0583d30c6980c4086451340f103c2714a99381b532488a4f8112d3965e180236cce873f0c23121abf29a3f1a076dbbf5c0f522c48b94ff6d40576fa3ae51aa4703b3dc8db94033323618a8a0b8c05ba4a0072c92a3ea5427cb58371aac4e83481e2044fc4576fa6924b99b3c99b115cf216b2e2ff2bf5bb0c8f1533d81f106ba82f8c9b0cefaf6f2fec21058864f2b3f906dc57a2145762a9fb80c4f555b2e6fe73ea9714cbb7de458fab0759fd0298e7df6046c7029462eea0b4b6703d6a36cfe4b06c9778adb163462288ff2e0fabe9d6a4f078c1898e6b6d90a24b48b1962ecedbfd080000ffffcc0e9c94e4010000",
	})
	if err != nil {
		panic(err)
	}
	g.DefaultResolver = hgr

	func() {
		b := packr.New("Templates", "./templates")
		b.SetResolver("index.html", packr.Pointer{ForwardBox: gk, ForwardPath: "d6469c13c8c86b6885d49e438465d0ef"})
	}()
	return nil
}()