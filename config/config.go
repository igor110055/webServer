package config

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"net/url"
	"os"
	"path"
)

var (
	APPVIPER           *viper.Viper
	DB                 *gorm.DB
	CLIENT             *ethclient.Client
	ControllerContract controller
	MORALIS            *Moralis
)

type controller struct {
	Address common.Address
	Abi     abi.ABI
}

type Moralis struct {
	Url     string
	Key     string
	ChainId string
}

func init() {
	APPVIPER = initAppConfig()
}

func initAppConfig() *viper.Viper {
	workDir, _ := os.Getwd()
	appViper := viper.New()
	appViper.SetConfigName("application")
	appViper.SetConfigType("yml")
	appViper.AddConfigPath(path.Join(workDir, "config"))
	err := appViper.ReadInConfig()
	if err != nil {
		panic("Failed to get config！")
	}
	//开启配置热更新
	appViper.WatchConfig()
	appViper.OnConfigChange(func(in fsnotify.Event) {
		//重新初始化链接
		log.Infoln("config change")
		initDB(appViper)
		initClient(appViper)
		initMoralis(appViper)
	})
	//初始化
	initDB(appViper)
	initClient(appViper)
	initMoralis(appViper)
	return appViper
}

func initDB(viper *viper.Viper) {
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	database := viper.GetString("database.databaseName")
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	charset := viper.GetString("database.charset")
	loc := viper.GetString("database.loc")

	sqlStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		username,
		password,
		host,
		port,
		database,
		charset,
		url.QueryEscape(loc),
	)

	db, err := gorm.Open(mysql.Open(sqlStr), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		panic("connected error" + err.Error())
	} else {
		log.Infoln("connect db success")
	}
	//监听增删改的sql日志
	db.Callback().Update().Register("updateSql", callback)
	db.Callback().Delete().Register("deleteSql", callback)
	db.Callback().Create().Register("createSql", callback)
	DB = db
}

func initClient(viper *viper.Viper) {
	url := viper.GetString("node")
	client, err := ethclient.Dial(url)
	if err != nil {
		panic("connected bsc node error" + err.Error())
	}

	log.Infoln("connect client success")
	CLIENT = client
}

func initMoralis(viper *viper.Viper) {
	url := viper.GetString("moralis.url")
	chainId := viper.GetString("moralis.chainId")
	key := viper.GetString("moralis.key")
	MORALIS = &Moralis{Url: url, ChainId: chainId, Key: key}
}

//gorm 打印sql回调方法
func callback(option *gorm.DB) {
	log.Infoln(option.Statement.SQL.String())
}
