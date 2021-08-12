-- 用于统计dbtool移动后机器是否还是正常的表
drop table if exists dbcheck_citus_move_data_check_table;
create table dbcheck_citus_move_data_check_table
(
    table_to_statistics                   text      not null default '', -- 哪个表(8大抓拍表)
    day_start_ts_human                    timestamp not null,            -- 哪一天的ts(当天00:00)
    day_start_ts                          bigint    not null default -1, -- 那一天的开始00:00

    first_statistics_time                 timestamp not null,            -- 上一次: 统计的时间
    first_count_of_day                    bigint    not null default -1, -- 上一次: 那一天有几个抓拍
    first_pk_of_min_ts_of_day_start_ts    text      not null default '', -- 上一次: 那一天最小的ts对应的主键是什么(比如faces表就是07:30的那条face_id, faces_index表就是07:35的face_reid)

    second_count_of_day                   bigint    not null default -1, -- 这一次: 那一天有几个抓拍
    second_pk_of_min_ts_of_day_start_ts   text      not null default '', -- 这一次: 那一天最小的ts对应的主键是什么(比如faces表就是07:30的那条face_id, faces_index表就是07:35的face_reid)

    is_equal_between_previous_and_current integer   not null default -1  -- 上一次统计和这一次统计的指标是不是都相等(1相等, 0不相等, 默认-1)
);
create unique index uidx_table_dayts on dbcheck_citus_move_data_check_table (table_to_statistics, day_start_ts);

-- 看下现场的机器配置
select *
from pg_dist_placement;
select *
from pg_dist_shard;
select *
from pg_dist_node;
select *
from pg_dist_shard_placement;
select *
from pg_dist_placement p
         join pg_dist_node n on p.groupid = n.groupid;

--估算一下数据量
select count(*)
from faces;
select count(*)
from faces_index;
select count(*)
from vehicle_capture;
select count(*)
from vehicle_capture_index;
select count(*)
from pedestrian_capture;
select count(*)
from pedestrian_capture_index;
select count(*)
from nonmotor_capture;
select count(*)
from nonmotor_capture_index;