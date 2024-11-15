package Helper

import (
	"context"
	"github.com/chromedp/chromedp"
)

type ScreenShot interface {
	Task(urlStr string, res *[]byte) chromedp.Tasks
	Print(ctx context.Context, urlStr string, id int) error
}
