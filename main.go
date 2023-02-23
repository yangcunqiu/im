package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh2 "github.com/go-playground/validator/v10/translations/zh"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"im/global"
	"im/model"
	"im/router"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func main() {
	initConfig()
	initDB()
	initOtherConfig()
	initTran()
	initRouter()
}

func initOtherConfig() {
	global.HttpClient = resty.New()
}

// 初始化validate的中文翻译器
func initTran() {
	// 创建翻译器
	uni := ut.New(zh.New())
	// 获取中文简体翻译器
	trans, _ := uni.GetTranslator("zh")
	// 判断gin默认的校验引擎是不是validate
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册中文简体翻译器
		_ = zh2.RegisterDefaultTranslations(v, trans)
		// 注册func, 获取struct中自定义的tag (label), 在输出时会将label的值作为字段名
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			jsonName := field.Tag.Get("json")
			if jsonName == "" {
				jsonName = field.Name
			} else {
				jsonName = strings.Split(jsonName, ",")[0]
			}

			name := field.Tag.Get("label")
			if name == "" {
				return jsonName
			}
			return name + fmt.Sprintf("[%v]", jsonName)
		})
	}
	global.Trans = trans
}

// 初始化配置
func initConfig() {
	// 设置配置文件名称 类型
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	// 设置viper的查找路径(viper会去这个路径查找上面配置的文件名的文件)
	viper.AddConfigPath(".")
	// 读取配置
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	// 把配置信息反序列化到结构体, 方便使用
	err = viper.Unmarshal(&global.Config)
	if err != nil {
		log.Println(err)
	}
	// 运行时监控配置文件的更新
	viper.WatchConfig()
	// 配置文件更新时的回调函数
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("config changed:", e.Name)
		// 重新反序列化
		err = viper.Unmarshal(&global.Config)
		if err != nil {
			log.Println(err)
		}
	})
}

func initRouter() {
	r := gin.Default()
	// 用作测试IP, c.ClientIP()方法会将header中key是test-ip的值作为ip返回
	r.TrustedPlatform = "test-ip"
	// r.Use(middlewares.GlobalExceptionCapture(), gin.Logger())
	// 注册路由
	router.RegisterRouter(r)
	var addr string
	if global.Config.Server.Port == 0 {
		addr = ":8080"
	} else {
		addr = ":" + strconv.Itoa(int(global.Config.Server.Port))
	}
	err := r.Run(addr)
	if err != nil {
		panic(err)
	}
}

func initDB() {
	// gorm日志
	gormLogger := logger.New(
		log.New(os.Stdout, "/r/n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // 慢SQL阈值
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	db, err := gorm.Open(mysql.Open(getDsn()), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用外键
		Logger:                                   gormLogger,
	})
	if err != nil {
		panic(err)
	}
	global.DB = db
	// 同步表结构
	syncTable()
}

func getDsn() string {
	// dsn := "root:pass1234@tcp(106.14.18.18:3306)/im?charset=utf8mb4&parseTime=True&loc=Local"
	m := global.Config.Mysql
	if m.Username == "" {
		panic("数据库用户名不能为空")
	}
	if m.Password == "" {
		panic("数据库密码不能为空")
	}
	if m.Ip == "" {
		panic("数据库IP不能为空")
	}
	if m.Database == "" {
		panic("数据库库名不能为空")
	}
	if m.Port == 0 {
		m.Port = 3306
	}
	var params string
	if m.Params != "" {
		params = "?" + m.Params
	}
	return m.Username + ":" + m.Password + "@tcp(" + m.Ip + ":" + strconv.Itoa(int(m.Port)) + ")" + "/" + m.Database + params
}

func syncTable() {
	err := global.DB.AutoMigrate(
		&model.User{},
		&model.UserLoginInfo{},
		&model.District{},
		&model.CallLog{},
	)
	if err != nil {
		panic(err)
	}
}
