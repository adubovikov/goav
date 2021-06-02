package main

import (
	"log"

	"github.com/adubovikov/goav/avcodec"
	"github.com/adubovikov/goav/avdevice"
	"github.com/adubovikov/goav/avfilter"
	"github.com/adubovikov/goav/avformat"
	"github.com/adubovikov/goav/avutil"
	"github.com/adubovikov/goav/swresample"
	"github.com/adubovikov/goav/swscale"
)

func main() {

	// Register all formats and codecs
	avformat.AvRegisterAll()
	avcodec.AvcodecRegisterAll()

	log.Printf("AvFilter Version:\t%v", avfilter.AvfilterVersion())
	log.Printf("AvDevice Version:\t%v", avdevice.AvdeviceVersion())
	log.Printf("SWScale Version:\t%v", swscale.SwscaleVersion())
	log.Printf("AvUtil Version:\t%v", avutil.AvutilVersion())
	log.Printf("AvCodec Version:\t%v", avcodec.AvcodecVersion())
	log.Printf("Resample Version:\t%v", swresample.SwresampleLicense())

}
