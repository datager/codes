select count(*), 'faces'
from faces
union all
select count(*), 'faces_index'
from faces_index
union all
select count(*), 'vehicle_capture'
from vehicle_capture
union all
select count(*), 'vehicle_capture_index'
from vehicle_capture_index
union all
select count(*), 'nonmotor_capture'
from nonmotor_capture
union all
select count(*), 'nonmotor_capture_index'
from nonmotor_capture_index
union all
select count(*), 'pedestrian_capture'
from pedestrian_capture
union all
select count(*), 'pedestrian_capture_index'
from pedestrian_capture_index;

select 15 + 85 + 15 + 13 + 13 + 117 == 250 w;

-- 删索引
-- faces
alter table faces
    drop constraint pk_faces;
drop index face_face_reid_ts_idx;
drop index face_ts_idx;
drop index face_uts_idx;
-- faces_index
drop index "face_index_face_reid_idx";
drop index "face_index_org_code_idx";
drop index "face_index_ts_idx";
drop index "face_index_uts_idx";
drop index "face_index_vid_idx";
drop index "idx_faces_index_gin";
-- vehicle_capture
alter table vehicle_capture
    drop constraint "pk_vehicle_capture";
drop index "vehicle_capture_vehicle_reid_ts_idx";
-- vehicle_capture_index
drop index "idx_vehicle_capture_index_gin";
drop index "idx_vehicle_idx_plate_text_gin";
drop index "vehicle_capture_index_org_code_idx";
drop index "vehicle_capture_index_reid_ts_idx";
drop index "vehicle_capture_index_ts_idx";
drop index "vehicle_capture_index_uts_idx";
drop index "vehicle_vid_ts";
-- nonmotor_capture
alter table nonmotor_capture
    drop constraint "pk_nonmotor_capture";
drop index "nonmotor_capture_reid_ts_idx";
drop index "nonmotor_capture_ts_idx";
-- nonmotor_capture_index
drop index "idx_nonmotor_capture_index_gin";
drop index "idx_nonmotor_idx_gin";
drop index "nonmotor_capture_index_org_code_idx";
drop index "nonmotor_capture_index_reid_idx";
drop index "nonmotor_index_ts_idx";
-- pedestrian_capture
alter table pedestrian_capture
    drop constraint "pk_pedestrian_capture";
drop index "pedestrian_capture_reid_ts_idx";
drop index "pedestrian_capture_ts_idx";
-- pedestrian_capture_index
drop index "idx_pedestrian_capture_index_gin";
drop index "pedestrian_capture_index_org_code_idx";
drop index "pedestrian_capture_index_reid_idx";
drop index "pedestrian_index_ts_idx";




select * from dbtool_shard_move_explain_backup order by step_number asc;
