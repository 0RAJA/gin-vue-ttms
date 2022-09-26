package config

import (
	"time"
)

type All struct {
	Token      Token      `yaml:"Token"`
	AliyunOSS  AliyunOSS  `yaml:"AliyunOSS"`
	Server     Server     `yaml:"Server"`
	Log        Log        `yaml:"Log"`
	Postgresql Postgresql `yaml:"Postgresql"`
	Redis      Redis      `yaml:"Redis"`
	Email      Email      `yaml:"Email"`
	Rule       Rule       `yaml:"Rule"`
	App        App        `yaml:"App"`
	Page       Page       `yaml:"Page"`
	Worker     Worker     `yaml:"Worker"`
	Auto       Auto       `yaml:"Auto"`
	Limit      Limit      `yaml:"Limit"`
	AliPay     AliPay     `yaml:"AliPay"`
}
type Limit struct {
	IPLimit  IPLimit  `yaml:"IPLimit"`
	APILimit APILimit `yaml:"APILimit"`
}

type IPLimit struct {
	Cap     int64 `yaml:"Cap"`
	GenNum  int64 `yaml:"GenNum"`
	GenTime int64 `yaml:"GenTime"`
	Cost    int64 `yaml:"Cost"`
}

type APILimit struct {
	Upload []Bucket `yaml:"Upload"`
	Email  []Bucket `yaml:"Email"`
}

type Bucket struct {
	Count    int           `yaml:"Count"`
	Duration time.Duration `yaml:"Duration"`
	Burst    int           `yaml:"Burst"`
}

type Auto struct {
	MovieVisitCountDuration             time.Duration `yaml:"MovieVisitCountDuration"`
	CommentStarDuration                 time.Duration `yaml:"CommentStarDuration"`
	UserMarkDuration                    time.Duration `yaml:"UserMarkDuration"`
	CodeMarkDuration                    time.Duration `yaml:"CodeMarkDuration"`
	MoviesOrderByBoxOfficeDuration      time.Duration `yaml:"MoviesOrderByBoxOfficeDuration"`
	MoviesOrderByUserMovieCountDuration time.Duration `yaml:"MoviesOrderByUserMovieCountDuration"`
	DeleteOutDatePlansDuration          time.Duration `yaml:"DeleteOutDatePlansDuration"`
	AddVisitCountDuration               time.Duration `yaml:"AddVisitCountDuration"`
}

type App struct {
	Name      string    `yaml:"Name"`
	Version   string    `yaml:"Version"`
	StartTime time.Time `yaml:"StartTime"`
}

type Redis struct {
	Address   string        `yaml:"Address"`
	DB        int           `yaml:"DB"`
	Password  string        `yaml:"Password"`
	PoolSize  int           `yaml:"PoolSize"`
	CacheTime time.Duration `yaml:"CacheTime"`
}

type Log struct {
	Level         string `yaml:"Level"`
	LogSavePath   string `yaml:"LogSavePath"`
	LowLevelFile  string `yaml:"LowLevelFile"`
	LogFileExt    string `yaml:"LogFileExt"`
	HighLevelFile string `yaml:"HighLevelFile"`
	MaxSize       int    `yaml:"MaxSize"`
	MaxAge        int    `yaml:"MaxAge"`
	MaxBackups    int    `yaml:"MaxBackups"`
	Compress      bool   `yaml:"Compress"`
}

type Page struct {
	MaxPageSize     int32  `yaml:"MaxPageSize"`
	PageKey         string `yaml:"PageKey"`
	PageSizeKey     string `yaml:"PageSizeKey"`
	DefaultPageSize int32  `yaml:"DefaultPageSize"`
}

type Token struct {
	Key                  string        `yaml:"key"`
	AssessTokenDuration  time.Duration `yaml:"AssessTokenDuration"`
	RefreshTokenDuration time.Duration `yaml:"RefreshTokenDuration"`
	AuthorizationKey     string        `yaml:"AuthorizationKey"`
	AuthorizationType    string        `yaml:"AuthorizationType"`
}

type Email struct {
	Password string   `yaml:"Password"`
	IsSSL    bool     `yaml:"IsSSL"`
	From     string   `yaml:"From"`
	To       []string `yaml:"To"`
	Host     string   `yaml:"Host"`
	Port     int      `yaml:"Port"`
	UserName string   `yaml:"UserName"`
}

type Server struct {
	RunMode               string        `yaml:"RunMode"`
	Address               string        `yaml:"Address"`
	ReadTimeout           time.Duration `yaml:"ReadTimeout"`
	WriteTimeout          time.Duration `yaml:"WriteTimeout"`
	DefaultContextTimeout time.Duration `yaml:"DefaultContextTimeout"`
}

type Rule struct {
	DefaultCoverURL                 string        `yaml:"DefaultCoverURL"`
	UsernameLenMax                  int           `yaml:"UsernameLenMax"`
	UsernameLenMin                  int           `yaml:"UsernameLenMin"`
	PasswordLenMax                  int           `yaml:"PasswordLenMax"`
	PasswordLenMin                  int           `yaml:"PasswordLenMin"`
	CommentLenMax                   int           `yaml:"CommentLenMax"`
	RowsMax                         int16         `yaml:"RowsMax"`
	ColsMax                         int16         `yaml:"ColsMax"`
	MovieNameLenMax                 int           `yaml:"MovieNameLenMax"`
	ContentLenMax                   int           `yaml:"ContentLenMax"`
	AreaLenMax                      int           `yaml:"AreaLenMax"`
	TagsLenMax                      int           `yaml:"TagsLenMax"`
	TagLenMax                       int           `yaml:"TagLenMax"`
	MaxFileSize                     int64         `yaml:"MaxFileSize"`
	MoviesOrderByBoxOfficePage      int32         `yaml:"MoviesOrderByBoxOfficePage"`
	MoviesOrderByBoxOfficeSize      int32         `yaml:"MoviesOrderByBoxOfficeSize"`
	MoviesOrderByUserMovieCountSize int32         `yaml:"MoviesOrderByUserMovieCountSize"`
	MoviesOrderByUserMovieCountPage int32         `yaml:"MoviesOrderByUserMovieCountPage"`
	LockTicketTime                  time.Duration `yaml:"LockTicketTime"`
	AvatarLenMax                    int           `yaml:"AvatarLenMax"`
	InviteCodeTime                  time.Duration `yaml:"InviteCodeTime"`
}

type Worker struct {
	TaskChanCapacity   int `yaml:"TaskChanCapacity"`
	WorkerChanCapacity int `yaml:"WorkerChanCapacity"`
	WorkerNum          int `yaml:"WorkerNum"`
}

type AliyunOSS struct {
	BucketUrl       string `yaml:"BucketUrl"`
	BasePath        string `yaml:"BasePath"`
	Endpoint        string `yaml:"Endpoint"`
	AccessKeyId     string `yaml:"AccessKeyID"`
	AccessKeySecret string `yaml:"AccessKeySecret"`
	BucketName      string `yaml:"BucketName"`
}

type Postgresql struct {
	DriverName string `yaml:"DriverName"`
	SourceName string `yaml:"SourceName"`
}

type AliPay struct {
	KPrivateKey          string `yaml:"kPrivateKey"`
	AppPublicCertPath    string `yaml:"AppPublicCertPath"`
	AliPayRootCertPath   string `yaml:"AliPayRootCertPath"`
	AliPayPublicCertPath string `yaml:"AliPayPublicCertPath"`
	NotifyURL            string `yaml:"NotifyURL"`
	ReturnURL            string `yaml:"ReturnURL"`
	IsProduction         bool   `yaml:"IsProduction"`
	KAppID               string `yaml:"kAppID"`
}
