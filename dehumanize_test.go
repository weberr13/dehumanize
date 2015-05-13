package dehumanize

import (
	log "github.com/cihub/seelog"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestYearDate(t *testing.T) {

	defer log.Flush()

	Convey("regex compiled", t, func() {
		So(initErr, ShouldBeNil)
	})
	Convey("bytes conversion", t, func() {
		So(SizeConvert("1b", 0), ShouldEqual, 1)
		So(SizeConvert("1b", 1000), ShouldEqual, 1)
		So(SizeConvert("1B", 0), ShouldEqual, 1)
	})
	Convey("kbytes conversion", t, func() {
		So(SizeConvert("1kb", 0), ShouldEqual, 1024)
		So(SizeConvert("1Kb", 0), ShouldEqual, 1024)
		So(SizeConvert("1kB", 0), ShouldEqual, 1024)
		So(SizeConvert("1KB", 0), ShouldEqual, 1024)
		So(SizeConvert("1.2kb", 0), ShouldEqual, 1228)
		So(SizeConvert("1kb", 1000), ShouldEqual, 1000)
		So(SizeConvert("1.2kb", 1000), ShouldEqual, 1200)
	})
	Convey("mbytes conversion", t, func() {
		So(SizeConvert("1mb", 0), ShouldEqual, 1024*1024)
		So(SizeConvert("1Mb", 0), ShouldEqual, 1024*1024)
		So(SizeConvert("1mB", 0), ShouldEqual, 1024*1024)
		So(SizeConvert("1MB", 0), ShouldEqual, 1024*1024)
		So(SizeConvert("1.2mb", 0), ShouldEqual, 1258291)
		So(SizeConvert("1mb", 1000), ShouldEqual, 1000000)
		So(SizeConvert("1.2mb", 1000), ShouldEqual, 1200000)
	})
	Convey("gbytes conversion", t, func() {
		So(SizeConvert("1gb", 0), ShouldEqual, 1024*1024*1024)
		So(SizeConvert("1Gb", 0), ShouldEqual, 1024*1024*1024)
		So(SizeConvert("1gB", 0), ShouldEqual, 1024*1024*1024)
		So(SizeConvert("1GB", 0), ShouldEqual, 1024*1024*1024)
		So(SizeConvert("1.2gb", 0), ShouldEqual, 1288490188)
		So(SizeConvert("1gb", 1000), ShouldEqual, 1000000000)
		So(SizeConvert("1.2gb", 1000), ShouldEqual, 1200000000)
	})
	Convey("tbytes conversion", t, func() {
		So(SizeConvert("1tb", 0), ShouldEqual, 1024*1024*1024*1024)
		So(SizeConvert("1Tb", 0), ShouldEqual, 1024*1024*1024*1024)
		So(SizeConvert("1tB", 0), ShouldEqual, 1024*1024*1024*1024)
		So(SizeConvert("1TB", 0), ShouldEqual, 1024*1024*1024*1024)
		So(SizeConvert("1.2tb", 0), ShouldEqual, 1319413953331)
		So(SizeConvert("1tb", 1000), ShouldEqual, 1000000000000)
		So(SizeConvert("1.2tb", 1000), ShouldEqual, 1200000000000)
	})
	Convey("tbytes conversion", t, func() {
		So(SizeConvert("1pb", 0), ShouldEqual, 1024*1024*1024*1024*1024)
		So(SizeConvert("1Pb", 0), ShouldEqual, 1024*1024*1024*1024*1024)
		So(SizeConvert("1pB", 0), ShouldEqual, 1024*1024*1024*1024*1024)
		So(SizeConvert("1PB", 1000), ShouldEqual, 1000000000000000)
		So(SizeConvert("1.2pb", 1000), ShouldEqual, 1200000000000000)
	})
	Convey("garbage conversion", t, func() {
		So(SizeConvert("1nb", 0), ShouldEqual, 0)
		So(SizeConvert("1.2nb", 0), ShouldEqual, 0)
		So(SizeConvert("", 0), ShouldEqual, 0)
		So(SizeConvert("kb", 0), ShouldEqual, 0)
		So(SizeConvert("1.b", 0), ShouldEqual, 1)
	})
}
