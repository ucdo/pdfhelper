package utils

import (
	Helper "Octopus/PdfHelper/helper"
	"Octopus/PdfHelper/structures"
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/chromedp/chromedp"
	"gopkg.in/yaml.v2"
	"os"
	"runtime"
	"strings"
)

type Utils struct {
}

var (
	config structures.Config // 配置文件
	header string            // 页眉
	footer string            // 页脚
)

func init() {
	fillConfig()          // 读取配置文件
	fillHeaderAndFooter() // 填充页眉页脚
}

// fillConfig 读取并填充配置文件
func fillConfig() {
	file, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.NewDecoder(bytes.NewReader(file)).Decode(&config)
	if err != nil {
		panic(err)
	}
	if !fileExists(config.ImageSavePath) {
		panic("请配置正确的文件保存位置")
	}
}

// fillHeaderAndFooter 判断文件是否存在，存在则读取并保存，否则为空
func fillHeaderAndFooter() {
	if fileExists(config.HeaderTemplate) {
		Header, err := os.ReadFile(config.HeaderTemplate)
		if err != nil {
			panic(err)
		}
		header = string(Header)
	}

	if fileExists(config.FooterTemplate) {
		Footer, err := os.ReadFile(config.FooterTemplate)
		if err != nil {
			panic(err)
		}
		footer = string(Footer)
	}

	config.HeaderTemplate = header
	config.FooterTemplate = footer
}

// extractChromium 根据运行时系统选择指定chrome的位置
func extractChromium() (string, error) {
	switch runtime.GOOS {
	case "windows":
		chromePath := config.WindowsChromePath // 请根据实际情况修改
		return chromePath, nil
	case "linux":
		chromePath := config.LinuxChromePath // 请根据实际情况修改
		return chromePath, nil
	default:
		return "", errors.New("unsupported OS")
	}

}

// fileExists 检查文件是否存在
func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	// 如果错误是 os.ErrNotExist，则文件不存在
	return !os.IsNotExist(err)
}

// 配置浏览器启动
func options(chromiumPath string) []chromedp.ExecAllocatorOption {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Headless,
		chromedp.DisableGPU,
		chromedp.ExecPath(chromiumPath),
		chromedp.NoFirstRun,
		chromedp.NoDefaultBrowserCheck,
	)

	if config.Debug {
		opts = append(opts, chromedp.Flag("headless", false)) // 禁用无头模式，启用带 GUI 的浏览器
	}

	return opts
}

// Do 截屏操作
func Do(urlStr string, imgType string, id int) error {
	chromiumPath, err := extractChromium()
	if err != nil || !fileExists(chromiumPath) {
		return errors.New(fmt.Sprintf("Failed to extract Chromium:%s", err))
	}

	// 获取浏览器配置
	opts := options(chromiumPath)

	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	switch strings.ToUpper(imgType) {
	case "PDF":
		x := Helper.NewPdfHelper(config)
		err = x.Print(ctx, urlStr, id)
		return err
	case "PNG":
		x := Helper.NewPngHelper(config)
		err = x.Print(ctx, urlStr, id)
		return err
	default:
		return errors.New(fmt.Sprintf("Unsupported image type:%s", imgType))
	}
}
