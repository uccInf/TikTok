package constdef

const ServerPort = 8080

const StaticServerPath = "/static"

const StaticLocalPath = "./public"

// database config
const (
	UserName          = "root"
	PassWord          = "12345678"
	Ip                = "192.168.1.25"
	Port              = 3306
	DataBaseName      = "TikTok"
	UserTableName     = "users"
	CommentsTableName = "comments"
	VideosTableName   = "videos"
)

const Replace = "*"

const SecretKey = "TikTok"

const (
	MacOS   = 1
	Windows = 2
)

const CurrentOs = MacOS
