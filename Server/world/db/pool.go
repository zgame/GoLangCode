// 游戏数据库
package db

//
import (
	"github.com/gomodule/redigo/redis"
	_ "github.com/go-sql-driver/mysql"

	//"util/dbs/database"
	//"util/dbs/redis"
	"../../core/logs"
	"github.com/astaxie/beego/orm"
)

////////////////////////////////////////////////////////////////////////////
//
func Init(path string) {
	// redis
	InitRedis(path)

	// mysql
	InitMysql(path)
}

//
func HealthCheck() error {
	// redis
	if e := HealthCheckRedis(); e != nil {
		return e
	}

	// mysql
	return HealthCheckMysql()
}

////////////////////////////////////////////////////////////////////////////
//
var g_cachePools *redis.Pool

func InitRedis(path string) {
	//fileName := path + "redis.json"
	//redis.InitByFile(fileName)

	//g_cachePools = redis.GetRedisPools("cache")
	//g_cachePools = newPool("127.0.0.1:6379", 0)

	logs.Info("init redis ok!")
}
func newPool(host string, db int) *redis.Pool {
	return &redis.Pool {
		MaxIdle: 50,
		MaxActive: 100,
		Dial: func() (redis.Conn, error) {
			options := redis.DialDatabase(db)
			//redis.DialPassword("zsw123")
			c, err := redis.Dial("tcp", host, options)
			if err != nil {
				panic(err.Error())
			}
			// 密码验证
			//if _, err := c.Do("AUTH", "zsw123"); err != nil {
			//	c.Close()
			//	return nil, err
			//}
			return c, err
		},
	}
}

//
func HealthCheckRedis() error {
	//return redis.HealthCheck()
	return nil
}

//
func getCacheConn(uid string) *redis.Conn {
	//return g_cachePools.GetConn()
	return nil
}

////////////////////////////////////////////////////////////////////////////
//
func InitMysql(path string) {
	//fileName := path + "mysql.json"
	//database.InitByFile(fileName)

	logs.Info("init mysql ok!")
}

//
func HealthCheckMysql() error {
	//return database.HealthCheck()
	return nil
}

//
func getPlayerOrm() orm.Ormer {
	//return database.GetDefOrm()
	return nil
}
