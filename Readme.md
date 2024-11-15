# pdfHelper
本项目旨在解决内网环境，尤其是无界面，无网络且 Linux 环境中，生成和保存 PDF 的需求。项目无需安装额外依赖，配置完成后即可作为独立服务运行。用户可以根据需求进行高度定制，使用时只需要具备基础的 CSS、JavaScript、HTML 和 Go 知识。

## 可配置参数：
    header：页眉内容
    footer：页脚内容
    margin：边距设置
    图片保存路径：配置图片保存的位置
    debug：仅在 Windows 上支持 GUI 调试，Linux 上使用 headless-chromium
    等待元素加载完成：设置等待某个元素加载完成后再进行操作截图操作

## go版本
    Go 1.23
### 安装 Go 环境

请确保你已经安装了 Go 1.23 或更高版本。你可以在 [Go 官方网站](https://golang.org/dl/) 下载并安装 Go。

安装完成后，可以通过以下命令检查 Go 是否正确安装：

```bash
go version
```

### 主要使用的第三方依赖
1. chromedp
2. gin
3. yaml

#### 安装依赖
如果你使用了第三方库，通常项目会有一个 `go.mod` 文件。用户只需执行以下命令，就会自动安装所有在 `go.mod` 中列出的依赖：
```bash
go mod tidy
```

## 在windows上使用
1. 安装chrome
2. 编译 ```go build .``` 
3. 配置`config.yaml`

## 在windows7上使用
1. 安装chrome
2. 安装Go `1.20.14` 或者以下版本
3. 编译  ```go build .```
4. 配置`config.yaml`

## 在linux上使用
1. 编译 ```$env:GOOS="linux"; $env:GOARCH="amd64"; go build -o myprogram```
2. chrome配置（二选一）：`cp headless-chromium /usr/local/bin/headless-chromium`
3. chrome配置（二选一）：添加到环境变量并使用`source`命令更新
4. 需要注意的是：如果没法运行，则使用 `chmod 0644 ` 对 `config.yaml`等文件权限进行管理

## 免责声明

本项目是免费的开源软件，供个人或企业使用。使用此工具时，您需要自行承担风险。我们不对因使用本软件而产生的任何直接或间接损失、损害或责任负责。

如果您有任何问题或疑虑，请自行评估软件的适用性和稳定性。