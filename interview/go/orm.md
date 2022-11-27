项目中ORM的使用

```
const OpenIdColumn = "open_id"

type User struct{
  Id   string `orm:"pk"`
  Name string
}

// 在init中注册定义的model
func init() {
	orm.RegisterModel(new(User))
}


// TableName 得到表名, beego.orm约定
func (u *User) TableName() string {
	return "user"
}

// sql连接配置
func (u *User) ConnName() string {
	return "user-db"
}

func GetUserById(ctx context.Context, openId string) (*User, error) {
  user := new(User)
  o, err := utils.GetOrmByModel(user)
	if err != nil {
		log.Err("get configs db err").Error(err).Ctx(ctx).Line()
		return nil, err
	}
	err = o.QueryTable(user).Filter(OpenIdColumn, openId).One(user)
	if err != nil {
		if err == orm.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}
```

参考：https://studygolang.com/articles/4401

```
o := orm.NewOrm()
o.Using("default") // 默认使用 default，你可以指定为其他数据库
user := new(User)
user.Name = "guguda"
fmt.Println(o.Insert(user))
orm.RegisterDriver("mysql", orm.DRMySQL)
orm.RegisterDataBase("default", "mysql", "root:root@/orm_test?charset=utf8")
maxIdle := 30
maxConn := 30
orm.RegisterDataBase("default", "mysql", "root:root@/orm_test?charset=utf8", maxIdle, maxConn)

// 数据库的最大空闲连接
orm.SetMaxIdleConns("default", 30)
// 数据库的最大数据库连接
orm.SetMaxOpenConns("default", 30)
// 设置为 UTC 时间
orm.DefaultTimeLoc = time.UTC
var r orm.RawSeter
r = o.Raw("UPDATE user SET name = ? WHERE name = ?", "testing", "slene")

调试模式
orm.Debug = true
var w io.Writer
orm.DebugLog = orm.NewLog(w)
```

