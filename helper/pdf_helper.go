package Helper

import (
	"Octopus/PdfHelper/structures"
	"context"
	"fmt"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"log"
	"os"
	"path"
)

type PdfHelper struct {
	structures.Config
}

func NewPdfHelper(config structures.Config) *PdfHelper {
	return &PdfHelper{
		config,
	}
}

func (p PdfHelper) Task(urlStr string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlStr),
		// 等待元素加载完成
		chromedp.WaitReady(p.WaitElement),
	}
}

func (p PdfHelper) Print(ctx context.Context, urlStr string, id int) error {
	log.Println(p.WaitElement)
	var buf []byte
	task := p.Task(urlStr, &buf)
	err := chromedp.Run(ctx, task)
	chromedp.Reload()
	chromedp.Reload()
	err = chromedp.Run(ctx, p.downloadTask(&buf))
	if err != nil {
		log.Println("Failed to capture screenshot:", err)
		return err
	}

	// Save the screenshot
	savePath := path.Join(p.ImageSavePath, fmt.Sprintf("screenshot_%d.pdf", id))
	if err := os.WriteFile(savePath, buf, 0644); err != nil {
		log.Println("Failed to save pdf:", err)
	}

	return nil
}

func (p PdfHelper) downloadTask(res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().
				WithHeaderTemplate(p.HeaderTemplate).
				WithFooterTemplate(p.FooterTemplate).
				WithMarginTop(p.MarginTop).
				WithMarginBottom(p.MarginBottom).
				WithMarginLeft(p.MarginLeft).
				WithMarginRight(p.MarginRight).
				WithDisplayHeaderFooter(true).
				WithPrintBackground(false).
				Do(ctx)
			if err != nil {
				return err
			}
			*res = buf
			return nil
		}),
	}
}
