package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type FTConfig struct {
	ListenAddress string `mapstructure:"listen_address"`
	UploadDir     string `mapstructure:"upload_dir"`
	URLPrefix     string `mapstructure:"url_prefix"`
	BaseAuth      struct {
		Enabled  bool   `mapstructure:"enabled"`
		UserName string `mapstructure:"username"`
		Password string `mapstructure:"password"`
	} `mapstructure:"basic_auth"`
}

var (
	h       bool
	CfgFile string
	Config  *FTConfig
)

func init() {
	flag.BoolVar(&h, "h", false, "help")
	flag.StringVar(&CfgFile, "c", "", "-c {configFilePath}")
	// 改变默认的 Usage，flag包中的Usage 其实是一个函数类型。这里是覆盖默认函数实现，具体见后面Usage部分的分析
	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stderr,
		`file_transfer version: 0.0.2
Usage: file_transfer [-h] [-c configFilePath]

Options:
`)
	flag.PrintDefaults()
}

func parseConfigFile() {
	log.Println("CfgFile", CfgFile)

	if CfgFile != "" {
		viper.SetConfigFile(CfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.AddConfigPath("/etc")
		viper.AddConfigPath("/etc/file_transfer")
		viper.SetConfigName("file_transfer")
		//viper.SetConfigFile()
		viper.SetConfigType("yaml")
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		log.Println("can't read config file")
		panic(err)
	}

	log.Println("read config from", viper.ConfigFileUsed())

	ft := FTConfig{}
	err := viper.Unmarshal(&ft)
	if err != nil {
		log.Println("can't parse config file")
		panic(err)
	}

	Config = &ft

	log.Println("UploadDir", ft.UploadDir)
	log.Println("URLPrefix", ft.URLPrefix)
	log.Println("EnableBasicAuth", ft.BaseAuth.Enabled)
	log.Println("starting...")
}

func DealUpload(c *gin.Context) {
	// single file
	file, _ := c.FormFile("file")
	log.Println(file.Filename)
	// Upload the file to specific dst.
	dst := filepath.Join(Config.UploadDir, file.Filename)
	c.SaveUploadedFile(file, dst)
	downloadPath := path.Join(Config.URLPrefix, file.Filename)
	downloadPath = strings.Replace(downloadPath, "/", "//", 1)
	c.String(http.StatusOK, "DownloadPath: %s\n",
		downloadPath)
}

func main() {
	flag.Parse()
	if h {
		flag.Usage()
		return
	}

	parseConfigFile()

	router := gin.Default()
	if Config.BaseAuth.Enabled {
		router.Use(gin.BasicAuth(gin.Accounts{
			Config.BaseAuth.UserName: Config.BaseAuth.Password, //用户名：密码
		}))
	}
	router.StaticFS("/download", http.Dir(Config.UploadDir))
	router.POST("/upload", DealUpload)
	router.Run(Config.ListenAddress)
}
