package structures

type Config struct {
	LinuxChromePath   string  `yaml:"linux_chrome_path"`
	WindowsChromePath string  `yaml:"windows_chrome_path"`
	ImageSavePath     string  `yaml:"image_save_path"`
	HeaderTemplate    string  `yaml:"header_template"`
	FooterTemplate    string  `yaml:"footer_template"`
	Debug             bool    `yaml:"debug"`
	WaitElement       string  `yaml:"wait_element"`
	MarginTop         float64 `yaml:"margin_top"`
	MarginLeft        float64 `yaml:"margin_left"`
	MarginRight       float64 `yaml:"margin_right"`
	MarginBottom      float64 `yaml:"margin_bottom"`
}
