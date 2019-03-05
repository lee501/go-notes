定义struct,对应数据库中的表结构
type Account struct{
  Id    int64
  Name  string  `xorm:”unique”`
  Balance float64
  Version int `xorm:”version”`  //乐观锁
}
// 乐观锁是 xorm 提供的一个比较实用的功能
// 通过在 tag 中指定 version 来开启它。开启之后，每次对记录进行更新的时候，该字段的值就会自动递增 1


1.创建orm引擎
  a. var x *xorm.Engine定义个engine实例
  b. x,err:=xorm.NewEngine()创建orm链接mysql
    1)通过配置项来处理

      type Config struct {
        User     string
        Password string
        Host     string
        Name     string
        Log      string
      }

      _, err := models.SetEngine(&models.Config{
        User:     setting.Get("db_user"),
        Password: setting.Get("db_password"),
        Host:     setting.Get("db_host"),
        Name:     setting.Get("db_name"),
        Log:      setting.Get("db_log"),
      })
      func SetEngine(config *Config) (*xorm.Engine, error)
    2). 同步表结构
      if err = x.Sync2(new(Account)); err != nil 
2.增删改查操作
  a. 插入数据 _, err := x.Insert(&Account{Name: name, Balance: balance}) 
  b. 删除数据 _, err := x.Delete(&Account{Id: id})
  c. 通过Get方法获取数据
    1).  根据Id来获得单条数据:
          a:=&Account{}
          has, err := x.Id(id).Get(a)

    2).  根据where获取单条数据
          a := new(Account)
          has, err := x.Where("name=?", "adn").Get(a)

    3).  根据Account结构体中非空数据来获取单条数据
          a := &Account{Id:1}
          has, err := x.Get(a)
  d. 更新数据库（Update接受的参数是指针）
      a.Balance += deposit
      _, err = x.Update(a)

3. 批量查询数据
  err = x.Desc("balance").Find(&as)
  // Desc 排序。
  // Find方法的第一个参数为slice的指针或Map指针

4. 事务及回滚
  // 创建 Session 对象
  sess := x.NewSession()
  defer sess.Close()
  // 开启事务
  if err = sess.Begin(); err != nil {    
    return err
  }
  if _, err = sess.Update(a1); err != nil {    // 发生错误时进行回滚
    sess.Rollback()    
    return err
  } 
  // 完成事务
  return sess.Commit()

5. Count方法统计记录条数
  a := new(Account)
  //返回满足id>1的Account的记录条数
  total, err := x.Where("id >?", 1).Count(a)
  // 返回Account所有记录条数
  total,err = x.Count(a)

6. Iterate方法提供逐条执行查询到的记录的方法，他所能使用的条件和Find方法完全相同
  err := x.Where("id > ?=)", 30).Iterate(new(Account), func(i int, bean interface{})error{
    user := bean.(*Account) //自行断言
    // do somthing use i and user
  })

7. Cols 方法可以指定查询特定字段
  x.Cols("name").Iterate(new(Account), printFn)

  var printFn = func(idx int, bean interface{}) error {   
     //dosomething
      return nil
  }

8. Omit排除特定字段
  x.Omit("name").Iterate(new(Account), printFn)

9. 分页查询结果偏移
  x.Limit(3, 2).Iterate(new(Account), printFn)

10. 日志保存到文件
  f, err := os.Create("sql.log")
  if err != nil { 
    log.Fatalf("Fail to create log file: %v\n", err)    
    return
  }
  x.Logger = xorm.NewSimpleLogger(f)

11. 支持 LRU 缓存
  // 获取到 ORM 引擎之后，如下操作
  cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
  x.SetDefaultCacher(cacher)

12. 事件钩子
  // 进行插入记录之前 和 完成插入记录之后 被调用
  func (a *Account) BeforeInsert() {
    log.Printf("before insert: %s", a.Name)
  }

  func (a *Account) AfterInsert() {
    log.Printf("after insert: %s", a.Name)
  }
