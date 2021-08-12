create or replace function insert_faces(x int) returns void as
$$
begin
    -- function开始
    insert into faces
    select (1584633600000 + random() * 15897600000), --ts: 从2020年3月20号, 到 2020年9月20号
           '467813bc-67f4-4037-8509-3e86eaa73026',   --sensor_id
           uuid_generate_v4(),                       --face_id
           '8f4f57fe-4c30-406e-a2d1-c7d4f3a76f7d',   --face_reid
           '{}',                                     --relation_types
           '{}',                                     --relation_ids
           '',                                       --feature
           0,                                        --confidence
           1,                                        --gender_id
           1,                                        --gender_confidence
           1,                                        --age_id
           1,                                        --age_confidence
           1,                                        --nation_id
           1,                                        --nation_confidence
           1,                                        --glass_id
           1,                                        --glass_confidence
           1,                                        --mask_id
           1,                                        --mask_confidence
           1,                                        --hat_id
           1,                                        --hat_confidence
           1,                                        --halmet_id
           1,                                        --halmet_confidence
           1,                                        --image_type
           '',                                       --image_uri
           '',                                       --thumbnail_image_uri varchar(256) default '',
           '',                                       --cutboard_image_uri varchar(256) default '',
           100,                                      --cutboard_x int default 0,
           100,                                      --cutboard_y int default 0,
           100,                                      --cutboard_width int default 0,
           100,                                      --cutboard_height int default 0,
           100,                                      --cutboard_res_width int default 0,
           100,                                      --cutboard_res_height int default 0,
           1,                                        --is_warned int default 0,
           1; --status int not null default 1,
end; -- function结束
$$ language plpgsql; -- 说明语言类型
select insert_faces_index(1000000);

create or replace function insert_pedestrian_index(x bigint) returns void as
$$
begin
    -- function开始
    insert into pedestrian_capture_index
    select '1111111',
           cast(generate_series(1, x) as varchar),
           (1577808000001 + random() * 15724700000),
           CONCAT('1000000000000000000000000000000', cast((10000 + random() * 10000) as int)),
           random() * 10000,
           random() * 100,
           0,
           random() * 4,
           random() * 4,
           'http://192.168.2.142:8501/api/v2/file/31/93b698558f69cb?_=1584703984111',
           random() * 2,
           random() * 100,
           random() * 57,
           random_int_array(3, 3),
           random_int_array(4, 2),
           random_int_array(4, 2),
           random() * 20,
           random() * 10,
           random() * 20,
           random() * 10,
           random() * 20,
           random() * 10,
           random() * 3,
           random() * 15,
           random() * 10;
end; -- function结束
$$ language plpgsql; -- 说明语言类型

