package config

type Zap struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`                           // 級別
	Format        string `mapstructure:"format" json:"format" yaml:"format"`                        // 輸出
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                        // 前綴
	Director      string `mapstructure:"director" json:"director"  yaml:"director"`                 // 資料夾
	ShowLine      bool   `mapstructure:"show-line" json:"showLine" yaml:"showLine"`                 // 顯示行數
	EncodeLevel   string `mapstructure:"encode-level" json:"encodeLevel" yaml:"encode-level"`       // 編輯馬
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktraceKey" yaml:"stacktrace-key"` // stacktrace key
	LogInConsole  bool   `mapstructure:"log-in-console" json:"logInConsole" yaml:"log-in-console"`  // 是否顯示到console
}
