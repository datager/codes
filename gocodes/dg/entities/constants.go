package entities

const (
	TableNamePetrolVehicles = "vehicles"
)

const (
	TableNameAccount                      = "account"
	TableNameCivilAttrs                   = "civil_attrs"
	TableNameCivilImages                  = "civil_images"
	TableNameDeviceRole                   = "device_role"
	TableNameDeviceRoleList               = "device_role_list"
	TableNameErrInfo                      = "err_info"
	TableNameFaces                        = "faces"
	TableNameFacesIndex                   = "faces_index"
	TableNameFaceEvents                   = "face_events"
	TableNameFaceEventsIndex              = "face_events_index"
	TableNameFaceRepos                    = "face_repos"
	TableNameFaceRuleExtend               = "face_rule_extend"
	TableNameFaceRuleSensors              = "face_rule_sensors"
	TableNameFuncRole                     = "func_role"
	TableNameIssuedRepo                   = "issued_repo"
	TableNameOperRecord                   = "oper_record"
	TableNameOrgStructure                 = "org_structure"
	TableNamePersonBang                   = "person_bang"
	TableNamePersonFile                   = "person_file"
	TableNameSensors                      = "sensors"
	TableNameWhiteWarn                    = "white_warn"
	TableNameWhiteWarnIndex               = "white_warn_index"
	TableNameRepoUpdate                   = "repo_updates"
	TableNameRepoDailyPicCount            = "repo_daily_pic_count"
	TableNameFootnote                     = "face_footnote"
	TableNameRepoCollision                = "repo_collision"
	TableNameTask                         = "tasks"
	TableNameMonitorRule                  = "monitor_rule"
	TableNameMonitorRuleSensorRelation    = "monitor_rule_sensor_relation"
	TableNameMonitorRuleFaceRepoRelation  = "monitor_rule_face_repo_relation"
	TableNamePedestrianCapture            = "pedestrian_capture"
	TableNamePedestrianIndexCapture       = "pedestrian_capture_index"
	TableNameDeleteTask                   = "delete_task"
	TableNameNonmotorCapture              = "nonmotor_capture"
	TableNameNonmotorIndexCapture         = "nonmotor_capture_index"
	TableNameVehicleCaptureIndexEntity    = "vehicle_capture_index"
	TableNameVehicleCaptureEntity         = "vehicle_capture"
	TableNamePlatformSyncGBEntity         = "platform_sync_gb"
	TableNamePlatformSyncPVGEntity        = "platform_sync_pvg"
	TableNameAthenaTaskEntity             = "athena_task"
	TableNameVehicleFile                  = "vehicle_file"
	TableNameVehicleFrequencyTask         = "vehicle_frequency_task"
	TableNameVehicleFrequencySummary      = "vehicle_frequency_summary"
	TableNameVehicleFrequencyDetail       = "vehicle_frequency_detail"
	TableNamePersonFrequencyTask          = "person_frequency_task"
	TableNamePersonFrequencySummary       = "person_frequency_summary"
	TableNamePersonFrequencyDetail        = "person_frequency_detail"
	TableNameVehicleBrandDict             = "vehicle_brand_dict"
	TableNameCaptureCountStatistics       = "capture_count_statistics"
	TableNameVehicleCollisionSummary      = "vehicle_areacol_summary"
	TableNamePersonCollisionSummary       = "person_areacol_summary"
	TableNameVehicleCollisionDetail       = "vehicle_areacol_detail"
	TableNamePersonCollisionDetail        = "person_areacol_detail"
	TableNameVehicleCollisionTask         = "vehicle_collision_task"
	TableNamePersonCollisionTask          = "person_collision_task"
	TableNameVehicleFirstAppearanceTask   = "vehicle_first_appearance_task"
	TableNameVehicleFirstAppearanceDetail = "vehicle_first_appearance_detail"
	TableNameVehicleAttrs                 = "vehicle_attrs"
	TableNameVehicleDayInNightOutTask     = "vehicle_dayinnightout_task"
	TableNameVehicleDayInNightOutSummary  = "vehicle_dayinnightout_summary"
	TableNameVehicleDayInNightOutDetail   = "vehicle_dayinnightout_detail"
	TableNameVehicleFakePlateTask         = "vehicle_fake_plate_task"
	TableNameVehicleFakePlateSummary      = "vehicle_fake_plate_summary"
	TableNameVehicleFakePlateDetail       = "vehicle_fake_plate_detail"
	TableNameFileCountStatistics          = "file_count_statistics"
	TableNameUserCaptureFavorite          = "user_capture_favorite"
	TableNameVehicleEvents                = "vehicle_events"
	TableNameVehicleEventsIndex           = "vehicle_events_index"
	TableNameSensorPreview                = "sensor_preview"
	TableNameVehicleHiddenTask            = "vehicle_hidden_task"
	TableNameVehicleHiddenSummary         = "vehicle_hidden_summary"
	TableNameVehicleHiddenDetail          = "vehicle_hidden_detail"
	TableNamePathAnalyzeTask              = "path_analyze_task"
	TableNameVehiclePathAnalyzeSummary    = "vehicle_path_analyze_summary"
	TableNameVehiclePathAnalyzeDetail     = "vehicle_path_analyze_detail"
	TableNamePersonPathAnalyzeSummary     = "person_path_analyze_summary"
	TableNamePersonPathAnalyzeDetail      = "person_path_analyze_detail"
	TableNameUserConfig                   = "user_config"
	TableNameUnCommitFaceView             = "faces_not_commit_view"
	TableNameUnCommitFace                 = "faces_not_commit"
	TableNamePlatformSyncOrgEntity        = "platform_sync_org"

	TableNameFaceCaptureHourCount       = "face_capture_hour_count"
	TableNameFaceCaptureDayCount        = "face_capture_day_count"
	TableNameVehicleCaptureHourCount    = "vehicle_capture_hour_count"
	TableNameVehicleCaptureDayCount     = "vehicle_capture_day_count"
	TableNameNonmotorCaptureHourCount   = "pedestrian_capture_hour_count"
	TableNameNonmotorCaptureDayCount    = "pedestrian_capture_day_count"
	TableNamePedestrianCaptureHourCount = "nonmotor_capture_hour_count"
	TableNamePedestrianCaptureDayCount  = "nonmotor_capture_day_count"
	TableNameFaceWarnHourCount          = "face_warn_hour_count"
	TableNameFaceWarnDayCount           = "face_warn_day_count"
	TableNameVehicleWarnHourCount       = "vehicle_warn_hour_count"
	TableNameVehicleWarnDayCount        = "vehicle_warn_day_count"
)

const (
	GlobalLockUser                               = "global_lock_user"
	GlobalLockFuncRole                           = "global_lock_func_role"
	GlobalLockDeviceRoleService                  = "global_lock_device_role_service"
	GlobalLockMonitorRuleService                 = "global_lock_monitor_rule_service"
	GlobalLockPlatformSyncGB                     = "global_lock_platform_sync_gb"
	GlobalLockPlatformSyncPVG                    = "global_lock_platform_sync_pvg"
	GlobalLockExportFaceCaptureByCondition       = "global_lock_export_face_capture_by_condition"
	GlobalLockExportPedestrianCaptureByCondition = "global_lock_export_pedestrian_capture_by_condition"
	GlobalLockExportNonmotorCaptureByCondition   = "global_lock_export_nonmotor_capture_by_condition"
	GlobalLockExportVehicleCaptureByCondition    = "global_lock_export_vehicle_capture_by_condition"
	GlobalLockExportFaceEventByCondition         = "global_lock_export_face_event_by_condition"
	GlobalLockExportFaceCaptureByFeature         = "global_lock_export_face_capture_by_feature"
	GlobalLockExportPedestrianCaptureByFeature   = "global_lock_export_pedestrian_capture_by_feature"
	GlobalLockExportNonmotorCaptureByFeature     = "global_lock_export_nonmotor_capture_by_feature"
	GlobalLockExportVehicleCaptureByFeature      = "global_lock_export_vehicle_capture_by_feature"
	GlobalLockExportCivilsByCondition            = "global_lock_export_civils_by_condition"
	GlobalLockVehicleFirstAppearance             = "global_lock_vehicle_first_appearance"
	GlobalLockVehicleFakePlate                   = "global_lock_vehicle_fake_plate"
	GlobalFileCount                              = "global_lock_file_count"
	GlobalUserFavorite                           = "global_lock_user_favorite"
)

const (
	SuperAdminUserID = "0000"
	SuperAdminOrgID  = "0000"
)