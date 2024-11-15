package Helper

import (
	"Octopus/PdfHelper/structures"
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
	"os"
	"path"
	"time"
)

type PngHelper struct {
	structures.Config
}

func NewPngHelper(config structures.Config) *PngHelper {
	return &PngHelper{
		config,
	}
}

func (p PngHelper) Task(urlStr string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlStr),
		chromedp.Sleep(2 * time.Second),
		chromedp.FullScreenshot(res, 90),
	}
}

func (p PngHelper) Print(ctx context.Context, urlStr string, id int) error {
	var buf []byte
	task := p.Task(urlStr, &buf)
	err := chromedp.Run(ctx, task, chromedp.WaitVisible(p.WaitElement, chromedp.ByID), chromedp.ActionFunc(func(ctx context.Context) error {
		p.pngAction()
		return nil
	}))

	if err != nil {
		fmt.Println("Failed to capture screenshot:", err)
		return err
	}

	// Save the screenshot
	savePath := path.Join(p.ImageSavePath, fmt.Sprintf("screenshot_%d.png", id))
	if err := os.WriteFile(savePath, buf, 0644); err != nil {
		log.Println("Failed to save screenshot:", err)
	}

	return nil
}

func (p PngHelper) pngAction() chromedp.Action {
	return chromedp.Evaluate(fmt.Sprintf(`%s %s`, p.HeaderTemplate, p.FooterTemplate), nil)
}
