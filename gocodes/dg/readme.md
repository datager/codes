# 抓拍
## 数据准备:

* 组织: 3级, 每级10个兄弟, 总计111个

* 设备: 每个组织下属100个设备. 总计11000个

* 抓拍: 车辆索引表

  * 分区跨度1个月
  * 灌入1亿数据, 均匀分散在各sensor下(每sensor 1亿/11000=9090个)
  * 各属性值random分散

## 执行过程
1. 刷库
2. 建分区, 人车非脸各5周, done.
```sql
-- 时间选择[20200203, 20200302) 共4周
select AutoCreatePartitionTable('faces_index', 'findex', thisWeekName(), timeZone(weekBegin(0)), timeZone(weekBegin(1)));

-- 车
-- 非
-- 人
-- 脸
select AutoCreatePartitionTable('faces_index', 'findex', '20200210', timeZone(weekBegin(1)), timeZone(weekBegin(2)));
select AutoCreatePartitionTable('faces_index', 'findex', '20200217', timeZone(weekBegin(2)), timeZone(weekBegin(3)));
select AutoCreatePartitionTable('faces_index', 'findex', '20200224', timeZone(weekBegin(3)), timeZone(weekBegin(4)));
select AutoCreatePartitionTable('faces_index', 'findex', '20200302', timeZone(weekBegin(4)), timeZone(weekBegin(5)));
```

3. 写程序 建org, 建sensor, 灌抓拍
因为需要写入org_code, 所以必须分sensor来灌数据, 所以必须用程序(对11000个设备分11000次写入, 写入时包括sensor/org信息), 就比较慢
```go
controller.CreateOrgSensor() // 建org, 建sensor
controller.InsertCaptureIndexTable(models.DetTypeVehicle) // 灌抓拍

// 思路如下
for sIntID := 1; sIntID < 11000; sIntID++ {
    s := dbEngine.GetBySIntID(sIntID);
    sID, orgCode := s.sID, s.orgCode
 
    dbEngine.Exec("sql") // 此sql为跨度1个月, 填入指定的sID, sIntID, orgCode, 每个sensor写9091个抓拍
}
```

3.1 写入性能: 
    * "NumInsideOneInsertBatch": 2000, "Concurrency": 300, 插入9091条需1.5s, (即一个sensor) 总共11000个sensor, 则需5h (使用中, ./release.sh放入服务器的supervisor里)

4. 备忘
不需要org_code的话, 可以直接用以下SQL写入性能更快
```sql
insert into vehicle_capture_index select
(1580659200000 + random()*15724700000), -- 时间选择[20200203, 20200302) 共4周, st=localtime(20200203)的时间戳ms值, duration=28天=100800000ms
CONCAT('1000000000000000000000000000000',cast((10000+ random()*10000) as int)),
generate_series(1,100000000), -- 1亿数据
random()*10,
random()*2,
random()*20,
random()*10,
random()*20,
random()*10,
random()*5,
random()*5,
random()*10,
random()*0,
random()*10,
random()*20,
random()*10,
random()*5,
random()*10,
random()*15;

--灌完后, update, 以便填入sensor/org信息
select min(sensor_int_id), max(sensor_int_id) from sensors; 

-- 找出自己插入的sensor_int_id范围
select floor(1 * 10999 + 1) as num; -- demo: 1~11000 
select org_code, sensor_id, sensor_int_id from sensors where sensor_int_id = (select floor(random() * 10999 + 1) as num); -- 从sensor表找id

-- 对表里每行执行以下sql, 从而更新(可以用go程序来弄)
update faces_index set (sensor_int_id, sensor_id, ts) = (select sensor_int_id, sensor_id, ts from sensors where sensor_int_id = 19) where face_reid = '5e29b480-acd6-4946-b693-634288b0f32a';
```
    
* 方案2(弃用): 先insert抓拍(无sensor/org信息), 再update sensor/org信息到抓拍里

  * 灌数据时如何填入org_code(111个)与sensor_id, sensor_int_id(2~11001), 且要匹配
  * 为了快, 先在ts跨度1个月内灌数据, 不写入org_code/sensor_id/sensor_int_id, 只需要random属性
  * 然后再对灌好的update, 来分配

  ```sql
  update vehicle_capture_index set (sensor_int_id, sensor_id, ts) = (select sensor_int_id, sensor_id, ts from sensors where sensor_int_id = 19) where face_reid = xxx;
  ```

## 性能验证:

搜索条件: ts跨度1个月, 组织/设备条件如下

第一次查询耗时x, 10次平均查询耗时y, 单位ms

| a表示attr             | a1   | a1&&a2 | a1&&a2...&&a10 | a1\|\|a2 | a1\|\|a2...\|\|a10 |
| --------------------- | ---- | ------ | -------------- | -------- | ------------------ |
| 不选org/sensor        |      |        |                |          |                    |
| 选1个org              |      |        |                |          |                    |
| 选10个org             |      |        |                |          |                    |
| 选1个sensor           |      |        |                |          |                    |
| 选10个sensor          |      |        |                |          |                    |
| 选100个sensor         |      |        |                |          |                    |
| 选1000个sensor        |      |        |                |          |                    |
| 选10000个sensor       |      |        |                |          |                    |
| 选1个org+1个sensor    |      |        |                |          |                    |
| 选1个org+10个sensor   |      |        |                |          |                    |
| 选1个org+100个sensor  |      |        |                |          |                    |
| 选10个org+1个sensor   |      |        |                |          |                    |
| 选10个org+10个sensor  |      |        |                |          |                    |
| 选10个org+100个sensor |      |        |                |          |                    |

在deepface_orgcode库灌1亿数据, 用orgcode版本的loki, 前端是8849接口可以复制验证
