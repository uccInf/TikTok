package constdef

// Server config
const (
	ServerPort = 8080
	ServerIp   = "192.168.1.25"
)

// DataBase config
const (
	DBUserName        = "root"
	DBPassWord        = "12345678"
	DBIp              = "localhost"
	DBPort            = 3306
	DataBaseName      = "TikTok"
	UserTableName     = "users"
	CommentsTableName = "comments"
	VideosTableName   = "videos"
)

// redis config
const (
	RDBIp       = "127.0.0.1"
	RDBPort     = 6379
	RDBIndex    = 0
	RDBPassWord = ""
)

// Os config
const Os = MacOS
