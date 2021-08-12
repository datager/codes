select * from dbcheck_citus_move_data_check_table;

select sum(first_count_of_day) from dbcheck_citus_move_data_check_table; --2623185

select * from pg_dist_placement order by groupid desc;
select count(*) from pg_dist_placement where groupid = 3;

