package common

import (
	"log"

	"github.com/wenlng/go-captcha-assets/resources/imagesv2"
	"github.com/wenlng/go-captcha-assets/resources/shapes"
	"github.com/wenlng/go-captcha-assets/resources/tiles"
	"github.com/wenlng/go-captcha/v2/base/option"
	"github.com/wenlng/go-captcha/v2/click"
	"github.com/wenlng/go-captcha/v2/slide"
)

var textCapt click.Captcha
var slideCapt slide.Captcha

func CaptchaSlide() slide.Captcha {
	return slideCapt
}

func InitCaptcha() {
	initTextCaptcha()
	initSlideCapt()

}

func initTextCaptcha() {
	// 初始化验证码
	builder := click.NewBuilder(
		click.WithRangeLen(option.RangeVal{Min: 4, Max: 5}),
		click.WithRangeVerifyLen(option.RangeVal{Min: 3, Max: 3}),
	)

	// fonts
	// fonts, err := fzshengsksjw.GetFont()
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// background images
	imgs, err := imagesv2.GetImages()
	if err != nil {
		log.Fatalln(err)
	}
	shapes, err := shapes.GetShapes()
	if err != nil {
		log.Fatalln(err)
	}
	builder.SetResources(
		// click.WithChars(chars.GetChineseChars()),
		// click.WithFonts([]*truetype.Font{fonts}),
		click.WithBackgrounds(imgs),
		click.WithShapes(shapes),
	)

	textCapt = builder.MakeShape()
}

func initSlideCapt() {
	builder := slide.NewBuilder(
		slide.WithGenGraphNumber(2),
		slide.WithEnableGraphVerticalRandom(true),
	)

	// background images
	imgs, err := imagesv2.GetImages()
	if err != nil {
		log.Fatalln(err)
	}

	graphs, err := tiles.GetTiles()
	if err != nil {
		log.Fatalln(err)
	}

	var newGraphs = make([]*slide.GraphImage, 0, len(graphs))
	for i := 0; i < len(graphs); i++ {
		graph := graphs[i]
		newGraphs = append(newGraphs, &slide.GraphImage{
			OverlayImage: graph.OverlayImage,
			MaskImage:    graph.MaskImage,
			ShadowImage:  graph.ShadowImage,
		})
	}

	// set resources
	builder.SetResources(
		slide.WithGraphImages(newGraphs),
		slide.WithBackgrounds(imgs),
	)

	slideCapt = builder.Make()
}
