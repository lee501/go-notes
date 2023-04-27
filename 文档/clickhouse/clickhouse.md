####Clickhouse 查询优化

#####单表查询
* 使用prewhere替代where；prewhere会自动优化执行过滤阶段的数据读取方式，降低io
    - 当查询列明显多余筛选列时使用prewhere可提供查询性能
  
```sql
select * from work_basic_model  where product='tracker_view' and ( id='eDf8fZky' or code='eDf8fZky' ) 
#替换where关键字
select * from work_basic_model  prewhere product='tracker_view' and ( id='eDf8fZky' or code='eDf8fZky' )
```

* 数据采样 SAMPLE
    - 必须是mergetree engine表才有效，建表时要指定采样策略sample by inHash32(列)
```sql
SELECT
    Title,
    count() * 10 AS PageViews
FROM hits_distributed
SAMPLE 0.1   #代表采样10%的数据, 也可以是具体的条数
WHERE
    CounterID = 34
GROUP BY Title
ORDER BY PageViews DESC LIMIT 1000
```

* 千万以上数据集进行order by查询时需要搭配where条件和limit语句一起使用

* 如非必须不要在结果集上构建虚拟列，虚拟列非常消耗资源浪费性能，可以考虑在前端进行处理，或者在表中构造实际字段进行额外存储
```sql
select id ,pv, uv , pv/uv rate 
```

* 使用 uniqCombined 替代 distinct 性能可提升10倍以上，uniqCombined 底层采用类似HyperLogLog算法实现，如能接收2%左右的数据误差，可直接使用这种去重方式提升查询性能
```sql
SELECT DISTINCT ON (a,b) * FROM t1;
#替换
SELECT uniqCombined ON (a,b) * FROM t1;
```

* 对于一些确定的数据模型，可将统计指标通过物化视图的方式进行构建，这样可避免数据查询时重复计算的过程；物化视图会在有新数据插入时进行更新
```sql
# 通过物化视图提前预计算用户下载量
CREATE MATERIALIZED VIEW download_hour_mv
ENGINE = SummingMergeTree
PARTITION BY toYYYYMM(hour) ORDER BY (userid, hour)
AS SELECT
  toStartOfHour(when) AS hour,
  userid,
  count() as downloads,
  sum(bytes) AS bytes
FROM download WHERE when >= toDateTime('2020-10-01 00:00:00')  #设置更新点, 该时间点之前的数据可以通过insert into select的方式进行插入
GROUP BY userid, hour

## 或者
CREATE MATERIALIZED VIEW db.table_MV TO db.table_new  ## table_new 可以是一张mergetree表
AS SELECT * FROM db.table_old; 

```

#####多表关联
* 当多表联查时，查询的数据仅从其中一张表出时，可考虑使用IN操作而不是JOIN
```sql
select a.* from a where a.uid in (select uid from b)
# 不要写成
select a.* from a left join b on a.uid=b.uid
```

* 多表Join时要满足小表在右的原则，右表关联时被加载到内存中与左表进行比较

* clickhouse在join查询时不会主动发起谓词下推的操作，需要每个子查询提前完成过滤操作；需要注意的是，是否主动执行谓词下推，对性能影响差别很大【新版本中已不再存在此问题，但是需要注意的是谓词位置的不同依然有性能的差异

* 将一些需要关联分析的业务创建成字典表进行join操作，前提是字典表不易太大，因为字典表会常驻内存。
```sql
CREATE TABLE products
(
    product_id UInt64,
    title String,
)
    ENGINE = Dictionary(products)
```
