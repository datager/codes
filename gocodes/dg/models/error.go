package models

import "fmt"

const (
	// 用户
	ErrorCodeUserModelNotFound                       = 100000
	ErrorCodeUserModelProhibitOperation              = 100001
	ErrorCodeUserModelUserNameDuplicated             = 100002
	ErrorCodeUserModelInvalidUser                    = 100003
	ErrorCodeUserModelCouldNotDeleteOrUpdateYourSelf = 100004
	ErrorCodeUserModelCouldNotUpdateOrdID            = 100005
	ErrorCodeUserModelUserIDShouldNotEmpty           = 100006
	ErrorCodeUserModelUserNameShouldNotEmpty         = 100007
	ErrorCodeUserModelUserPasswordShouldNotEmpty     = 100008
	ErrorCodeUserModelOrgIDShouldNotEmpty            = 100009
	ErrorCodeUserModelUserFuncRoleShouldNotEmpty     = 100010
	ErrorCodeUserModelFuncRoleIDShouldNotEmpty       = 100011
	ErrorCodeUserModelUserDeviceRoleShouldNotEmpty   = 100012
	ErrorCodeUserModelDeviceRoleIDShouldNotEmpty     = 100013
	ErrorCodeUserModelSuperAdminShouldNotBeOperation = 100014
	// 功能角色
	ErrorCodeFuncRoleModelNotFound                   = 200000
	ErrorCodeFuncRoleModelProhibitOperation          = 200001
	ErrorCodeFuncRoleModelFuncRoleNameDuplicated     = 200002
	ErrorCodeFuncRoleModelCouldNotUpdateOrdID        = 200003
	ErrorCodeFuncRoleModelFuncRoleTypeNotNormal      = 200004
	ErrorCodeFuncRoleModelFuncRoleIDShouldNotEmpty   = 200005
	ErrorCodeFuncRoleModelOrgIDShouldNotEmpty        = 200006
	ErrorCodeFuncRoleModelFuncRoleNameShouldNotEmpty = 200007
	ErrorCodeFuncRoleModelUserCountNotZero           = 200008
	// 设备角色
	ErrorCodeDeviceRoleModelNotFound                               = 300000
	ErrorCodeDeviceRoleModelProhibitOperation                      = 300001
	ErrorCodeDeviceRoleModelDeviceRoleNameDuplicatedInSameOrgLevel = 300002
	ErrorCodeDeviceRoleModelCouldNotUpdateOrdID                    = 300003
	ErrorCodeDeviceRoleModelDeviceRoleTypeNotNormal                = 300004
	ErrorCodeDeviceRoleModelDeviceRoleIDShouldNotEmpty             = 300005
	ErrorCodeDeviceRoleModelOrgIDShouldNotEmpty                    = 300006
	ErrorCodeDeviceRoleModelDeviceRoleNameShouldNotEmpty           = 300007
	ErrorCodeDeviceRoleModelUserCountNotZero                       = 300008
	ErrorCodeDeviceRoleModelSensorIsNotVisibleForCurrentUser       = 300009
	ErrorCodeDeviceRoleModelDeviceRoleTypeInvalid                  = 300010
	ErrorCodeDeviceRoleModelDeviceRoleIDCouldNotBeEmptyWhenEdit    = 300011
	// 组织
	ErrorCodeOrgModelNotFound          = 400000
	ErrorCodeOrgModelProhibitOperation = 400001
	// 设备
	ErrorCodeSensorModelSingleSensorUpdateFail                    = 500000
	ErrorCodeSensorNormalDeviceTypeHasNoJurisdictionOfAddSensor   = 500001
	ErrorCodeSensorModelCountExceedLimit                          = 500002
	ErrorCodeSensorCouldNotDeleteOrUpdateBecauseBindToTask        = 500003
	ErrorCodeSensorCouldNotDeleteOrUpdateBecauseBindToMonitorRule = 500004
	// 登陆者鉴权
	ErrorCodeEmptyLoginUserFound             = 600000
	ErrorCodeEmptyDeviceRoleOfLoginUserFound = 600001
	// 布控
	ErrorCodeMonitorRuleCouldNotDeleteOrEditBecauseOfOrg            = 700000
	ErrorCodeMonitorRuleCouldNotDeleteOrEditBecauseOfSensor         = 700001
	ErrorCodeMonitorRuleCouldNotDeleteOrEditBecauseOfRepo           = 700002
	ErrorCodeMonitorRuleNoVisibleSensorOrNoVisibleRepoWhenAddOrEdit = 700003
	ErrorCodeMonitorRuleNoExistedMonitorRuleByIDFoundWhenEdit       = 700004
	ErrorCodeMonitorRuleAllDstStatusAreNotSame                      = 700005
	ErrorCodeMonitorRuleOneSensorCouldNotBeMonitorRuledTwice        = 700006
	ErrorCodeMonitorRuleGotErrorWhenQuerySensorOfMonitorRuleByID    = 700007
	ErrorCodeMonitorRuleGotErrorWhenQueryRepoOfMonitorRuleByID      = 700008
	ErrorCodeMonitorRuleNameTooLongError                            = 700009
	ErrorCodeMonitorRuleNameAlreadyExistError                       = 700010

	// 库
	ErrorCodeRepoCouldNotGetIssuedReposByEmptyOrgID        = 800000
	ErrorCodeRepoCouldNotGetIssuedToOrgOfRepoByEmptyRepoID = 800001
	// 下发库

	// 比对库

	// 平台同步
	ErrorCodePlatformSyncModelNotFound           = 110000
	ErrorCodePlatformSyncModelWaitForSync        = 110001
	ErrorCodePlatformSyncModelAlreadyDistributed = 110002
	ErrorCodePlatformSyncModelConfigNotFound     = 110003
	// 导出
	ErrorCodeExportModelFieldNoFound              = 120000
	ErrorCodeExportModelLimitInvalid              = 120001
	ErrorCodeExportModelTargetTaskNotFound        = 120003
	ErrorCodeExportModelTargetInnerTypeNotFound   = 120004
	ErrorCodeExportModelFeatureIDMapNOExceedLimit = 120005
	// 频次分析通用(人+车)
	ErrorCodeAthenaFrequencyModelReachMaxFrequencyAnalyzeTaskNumInAthena  = 130000
	ErrorCodeAthenaFrequencyModelReachMaxFrequencyAnalyzeTaskNumOfOneUser = 130001
	ErrorCodeAthenaFrequencyTaskCreateFailWithNoTaskInAthenaAndNoTaskInPg = 130100
	// 车辆频次分析(车)
	// athena类
	ErrorCodeAthenaFrequencyTaskModelGotErrorFromAthena                 = 140000
	ErrorCodeAthenaFrequencyTaskModelGotTaskNotExistErrorFromAthena     = 140001
	ErrorCodeAthenaFrequencyTaskModelGotTaskStateChangedErrorFromAthena = 140002
	ErrorCodeAthenaFrequencyTaskModelGotMaxTaskNumErrorFromAthena       = 140003
	// loki类
	ErrorCodeAthenaVehicleFrequencyModelInsertIntoDBError                                              = 140100
	ErrorCodeAthenaVehicleFrequencyModelCreateTaskCouldNotHaveEmptySensorsError                        = 140101
	ErrorCodeAthenaVehicleFrequencyModelCreateTaskSensorCountReachMaxLimitError                        = 140102
	ErrorCodeAthenaVehicleFrequencyModelCreateTaskStartTsGreaterEqualThanEndTsError                    = 140103
	ErrorCodeAthenaVehicleFrequencyModelCreateTaskDurationDaysBetweenStartTsAndEndTsReachMaxLimitError = 140104
	ErrorCodeAthenaVehicleFrequencyModelCreateTaskIllegalTsError                                       = 140105
	ErrorCodeAthenaVehicleFrequencyModelCreateTaskEmptyTaskNameError                                   = 140106
	ErrorCodeAthenaVehicleFrequencyModelCreateTaskTaskNameTooLongError                                 = 140107
	ErrorCodeAthenaVehicleFrequencyModelCreateTaskIllegalTimesThresholdError                           = 140108
	ErrorCodeAthenaFrequencyModelSomeTaskIDIsNotUserVisibleError                                       = 140109
	ErrorCodeAthenaFrequencyModelTaskDeleteTaskInAthenaFailWithRetry                                   = 140110
	// 车辆频次分析(人)
	ErrorCodeAthenaPersonFrequencyModelInsertIntoDBError = 150100

	// 区域碰撞athena
	ErrorCodeAthenaCollisionTaskModelGotErrorFromAthena                 = 160000
	ErrorCodeAthenaCollisionTaskModelGotTaskNotExistErrorFromAthena     = 160001
	ErrorCodeAthenaCollisionTaskModelGotTaskStateChangedErrorFromAthena = 160002
	ErrorCodeAthenaCollisionTaskModelGotMaxTaskNumErrorFromAthena       = 160003

	//区域碰撞（人+车）
	ErrorCodeAthenaCollisionModelInsertIntoDBError                                              = 160100
	ErrorCodeAthenaCollisionModelCreateTaskCouldNotHaveEmptyTimeSpacesError                     = 160101
	ErrorCodeAthenaCollisionModelCreateTaskSensorCountReachMaxLimitError                        = 160102
	ErrorCodeAthenaCollisionModelCreateTaskTaskNameTooLongError                                 = 160103
	ErrorCodeAthenaCollisionModelCreateTaskEmptyTaskNameError                                   = 160104
	ErrorCodeAthenaCollisionTaskCreateTaskTsAndSensorSuperposition                              = 160105
	ErrorCodeAthenaCollisionModelCreateTaskIllegalTsError                                       = 160106
	ErrorCodeAthenaCollisionModelCreateTaskStartTsGreaterEqualThanEndTsError                    = 160107
	ErrorCodeAthenaCollisionModelCreateTaskDurationDaysBetweenStartTsAndEndTsReachMaxLimitError = 160108
	ErrorCodeAthenaCollisionModelReachMaxCollisionAnalyzeTaskNumInAthena                        = 160109
	ErrorCodeAthenaCollisionModelReachMaxCollisionAnalyzeTaskNumOfOneUser                       = 160110
	ErrorCodeAthenaCollisionTaskCreateFailWithNoTaskInAthenaAndNoTaskInPg                       = 160111
	ErrorCodeAthenaCollisionModelSomeTaskIDIsNotUserVisibleError                                = 160112
	ErrorCodeAthenaCollisionModelTaskDeleteTaskInAthenaFailWithRetry                            = 160113
	ErrorCodeAthenaCollisionModelCreateTaskTaskNameError                                        = 160114
	ErrorCodeGetCollisionTaskResultErr                                                          = 160115
	ErrorCodeAthenaCollisionTaskNotUser                                                         = 160116

	// 首次入城 athena
	ErrorCodeFirstAppearanceModelGotErrorFromAthena                 = 170000
	ErrorCodeFirstAppearanceModelGotTaskNotExistErrorFromAthena     = 170001
	ErrorCodeFirstAppearanceModelGotTaskStateChangedErrorFromAthena = 170002
	ErrorCodeFirstAppearanceModelGotMaxTaskNumErrorFromAthena       = 170003
	// 首次入城 loki
	ErrorCodeFirstAppearanceModelCreateTaskEmptyTaskNameError     = 170100
	ErrorCodeFirstAppearanceModelCreateTaskNameTooLongError       = 170101
	ErrorCodeFirstAppearanceModelCreateTaskSensorIDsEmptyError    = 170102
	ErrorCodeFirstAppearanceModelCreateTaskSensorIDsTooManyError  = 170103
	ErrorCodeFirstAppearanceModelCreateTaskSelectedTimeError      = 170104
	ErrorCodeFirstAppearanceModelCreateTaskBacktrackingTimeError  = 170105
	ErrorCodeFirstAppearanceModelReachSystemMaxTaskLimitError     = 170106
	ErrorCodeFirstAppearanceModelReachUserMaxTaskLimitError       = 170107
	ErrorCodeFirstAppearanceModelTaskNameDuplicatedError          = 170108
	ErrorCodeFirstAppearanceModelDeleteTaskInAthenaFailWithRetry  = 170109
	ErrorCodeFirstAppearanceModelCreateTaskError                  = 170110
	ErrorCodeFirstAppearanceModelTaskNotFoundError                = 170111
	ErrorCodeFirstAppearanceModelUpdateTypeInvalidError           = 170112
	ErrorCodeFirstAppearanceModelTargetTaskNotUserVisibilityError = 170113
	ErrorCodeFirstAppearanceModelCreateTaskPlateTextEmptyError    = 170114

	//昼伏夜出
	ErrorCodeAthenaDayInNightOutTaskModelGotErrorFromAthena                 = 190000
	ErrorCodeAthenaDayInNightOutTaskModelGotTaskNotExistErrorFromAthena     = 190001
	ErrorCodeAthenaDayInNightOutTaskModelGotTaskStateChangedErrorFromAthena = 190002
	ErrorCodeAthenaDayInNightOutTaskModelGotMaxTaskNumErrorFromAthena       = 190003

	ErrorCodeAthenaDayInNightOutModelInsertIntoDBError                                              = 190100
	ErrorCodeAthenaDayInNightOutModelCreateTaskCouldNotHaveEmptySensorsError                        = 190101
	ErrorCodeAthenaDayInNightOutModelCreateTaskSensorCountReachMaxLimitError                        = 190102
	ErrorCodeAthenaDayInNightOutModelCreateTaskTaskNameTooLongError                                 = 190103
	ErrorCodeAthenaDayInNightOutModelCreateTaskEmptyTaskNameError                                   = 190104
	ErrorCodeAthenaDayInNightOutTaskCreateTaskTsSuperposition                                       = 190105
	ErrorCodeAthenaDayInNightOutModelCreateTaskIllegalTsError                                       = 190106
	ErrorCodeAthenaDayInNightOutModelCreateTaskDurationDaysBetweenStartTsAndEndTsReachMaxLimitError = 190107
	ErrorCodeAthenaDayInNightOutModelReachMaxTaskNumInAthena                                        = 190108
	ErrorCodeAthenaDayInNightOutModelReachMaxTaskNumOfOneUser                                       = 190109
	ErrorCodeAthenaDayInNightOutTaskCreateFailWithNoTaskInAthenaAndNoTaskInPg                       = 190110
	ErrorCodeAthenaDayInNightOutModelTaskDeleteTaskInAthenaFailWithRetry                            = 190111
	ErrorCodeAthenaDayInNightOutModelCreateTaskTaskNameError                                        = 190112
	ErrorCodeGetDayInNightOutTaskResultErr                                                          = 190113
	ErrorCodeAthenaDayInNightOutTaskNotUser                                                         = 190114

	// 套牌车 athena
	ErrorCodeFakePlateModelGotErrorFromAthena                 = 180000
	ErrorCodeFakePlateModelGotTaskNotExistErrorFromAthena     = 180001
	ErrorCodeFakePlateModelGotTaskStateChangedErrorFromAthena = 180002
	ErrorCodeFakePlateModelGotMaxTaskNumErrorFromAthena       = 180003
	// 套牌车 loki
	ErrorCodeFakePlateModelCreateTaskEmptyTaskNameError     = 180100
	ErrorCodeFakePlateModelCreateTaskNameTooLongError       = 180101
	ErrorCodeFakePlateModelCreateTaskSensorIDsEmptyError    = 180102
	ErrorCodeFakePlateModelCreateTaskSensorIDsTooManyError  = 180103
	ErrorCodeFakePlateModelCreateTaskSelectedTimeError      = 180104
	ErrorCodeFakePlateModelCreateTaskBacktrackingTimeError  = 180105
	ErrorCodeFakePlateModelReachSystemMaxTaskLimitError     = 180106
	ErrorCodeFakePlateModelReachUserMaxTaskLimitError       = 180107
	ErrorCodeFakePlateModelTaskNameDuplicatedError          = 180108
	ErrorCodeFakePlateModelDeleteTaskInAthenaFailWithRetry  = 180109
	ErrorCodeFakePlateModelCreateTaskError                  = 180110
	ErrorCodeFakePlateModelTaskNotFoundError                = 180111
	ErrorCodeFakePlateModelUpdateTypeInvalidError           = 180112
	ErrorCodeFakePlateModelTargetTaskNotUserVisibilityError = 180113
	// 用户收藏
	ErrorCodeUserFavoriteExceedLimitError = 210000
	ErrorCodeUserFavoriteTargetExistError = 210001
	// 设备预览
	ErrorCodeSensorPreviewNotSupport      = 220001
	ErrorCodeSensorPreviewIllegalSensorID = 220002
	ErrorCodeSensorPreviewSensorNotExist  = 220003

	//隐匿车辆
	ErrorCodeAthenaHiddenTaskModelGotErrorFromAthena                 = 230000
	ErrorCodeAthenaHiddenTaskModelGotTaskNotExistErrorFromAthena     = 230001
	ErrorCodeAthenaHiddenTaskModelGotTaskStateChangedErrorFromAthena = 230002
	ErrorCodeAthenaHiddenTaskModelGotMaxTaskNumErrorFromAthena       = 230003

	ErrorCodeAthenaHiddenModelCreateTaskTaskNameError                  = 230100
	ErrorCodeAthenaHiddenModelReachMaxTaskNumOfOneUser                 = 230101
	ErrorCodeAthenaHiddenModelInsertIntoDBError                        = 230102
	ErrorCodeAthenaHiddenModelTaskDeleteTaskInAthenaFailWithRetry      = 230103
	ErrorCodeAthenaHiddenTaskCreateFailWithNoTaskInAthenaAndNoTaskInPg = 230104
	ErrorCodeAthenaHiddenModelCreateTaskEmptyTaskNameError             = 230105
	ErrorCodeAthenaHiddenModelCreateTaskTaskNameTooLongError           = 230106
	ErrorCodeHiddenModelCreateTaskSensorIDsEmptyError                  = 230107
	ErrorCodeHiddenModelCreateTaskSensorIDsTooManyError                = 230108
	ErrorCodeHiddenModelCreateTaskSelectedTimeError                    = 230109
	ErrorCodeHiddenModelTargetTaskNotUserVisibilityError               = 230110
	ErrorCodeHiddenModelUpdateTypeInvalidError                         = 230111
	ErrorCodeAthenaHiddenTaskNotUser                                   = 230112
	ErrorCodeGetHiddenTaskResultErr                                    = 230113

	//轨迹分析
	ErrorCodeAthenaPathAnalyzeTaskModelGotErrorFromAthena                 = 240000
	ErrorCodeAthenaPathAnalyzeTaskModelGotTaskNotExistErrorFromAthena     = 240001
	ErrorCodeAthenaPathAnalyzeTaskModelGotTaskStateChangedErrorFromAthena = 240002
	ErrorCodeAthenaPathAnalyzeTaskModelGotMaxTaskNumErrorFromAthena       = 250003

	ErrorCodeAthenaPathAnalyzeModelCreateTaskTaskNameError                  = 240100
	ErrorCodeAthenaPathAnalyzeModelReachMaxTaskNumOfOneUser                 = 240101
	ErrorCodeAthenaPathAnalyzeModelInsertIntoDBError                        = 240102
	ErrorCodeAthenaPathAnalyzeModelTaskDeleteTaskInAthenaFailWithRetry      = 240103
	ErrorCodeAthenaPathAnalyzeTaskCreateFailWithNoTaskInAthenaAndNoTaskInPg = 240104
	ErrorCodeAthenaPathAnalyzeModelCreateTaskEmptyTaskNameError             = 240105
	ErrorCodeAthenaPathAnalyzeModelCreateTaskTaskNameTooLongError           = 240106
	ErrorCodePathAnalyzeModelCreateTaskSensorIDsEmptyError                  = 240107
	ErrorCodePathAnalyzeModelCreateTaskSensorIDsTooManyError                = 240108
	ErrorCodePathAnalyzeModelCreateTaskSelectedTimeError                    = 240109
	ErrorCodePathAnalyzeModelTargetTaskNotUserVisibilityError               = 240110
	ErrorCodePathAnalyzeModelUpdateTypeInvalidError                         = 240111
	ErrorCodeAthenaPathAnalyzeTaskNotUser                                   = 240112
	ErrorCodeGetPathAnalyzeTaskResultErr                                    = 240113
	ErrorCodePathAnalyzeModelCreateTaskTypeError                            = 240114
)

var ErrorMessages = map[int]string{
	// 用户
	ErrorCodeUserModelNotFound:                       "User not found",
	ErrorCodeUserModelProhibitOperation:              "You shouldn't operating this user",
	ErrorCodeUserModelUserNameDuplicated:             "UserName duplicated",
	ErrorCodeUserModelInvalidUser:                    "Invalid user",
	ErrorCodeUserModelCouldNotDeleteOrUpdateYourSelf: "You couldn't delete or update yourself",
	ErrorCodeUserModelCouldNotUpdateOrdID:            "OrgId couldn't be update",
	ErrorCodeUserModelUserIDShouldNotEmpty:           "UserID shouldn't be empty",
	ErrorCodeUserModelUserNameShouldNotEmpty:         "UserName shouldn't be empty",
	ErrorCodeUserModelUserPasswordShouldNotEmpty:     "UserPassword shouldn't be empty",
	ErrorCodeUserModelOrgIDShouldNotEmpty:            "OrgID shouldn't be empty",
	ErrorCodeUserModelUserFuncRoleShouldNotEmpty:     "UserFuncRole shouldn't be empty",
	ErrorCodeUserModelFuncRoleIDShouldNotEmpty:       "FuncRoleID shouldn't be empty",
	ErrorCodeUserModelUserDeviceRoleShouldNotEmpty:   "UserDeviceRole shouldn't be empty",
	ErrorCodeUserModelDeviceRoleIDShouldNotEmpty:     "DeviceRoleID shouldn't be empty",
	ErrorCodeUserModelSuperAdminShouldNotBeOperation: "Super admin shouldn't be operation",
	// 功能角色
	ErrorCodeFuncRoleModelNotFound:                   "FuncRole not found",
	ErrorCodeFuncRoleModelProhibitOperation:          "You shouldn't operating this func role",
	ErrorCodeFuncRoleModelFuncRoleNameDuplicated:     "FuncRoleName duplicated",
	ErrorCodeFuncRoleModelCouldNotUpdateOrdID:        "OrgId couldn't be update",
	ErrorCodeFuncRoleModelFuncRoleTypeNotNormal:      "FuncRoleType not normal, not normal type couldn't be edit or delete",
	ErrorCodeFuncRoleModelFuncRoleIDShouldNotEmpty:   "FuncRoleID shouldn't be empty",
	ErrorCodeFuncRoleModelOrgIDShouldNotEmpty:        "OrgID shouldn't be empty",
	ErrorCodeFuncRoleModelFuncRoleNameShouldNotEmpty: "FuncRoleName shouldn't be empty",
	ErrorCodeFuncRoleModelUserCountNotZero:           "UserCount not zero",
	// 设备角色
	ErrorCodeDeviceRoleModelNotFound:                               "DeviceRole not found",
	ErrorCodeDeviceRoleModelProhibitOperation:                      "You shouldn't operating this device role",
	ErrorCodeDeviceRoleModelDeviceRoleNameDuplicatedInSameOrgLevel: "DeviceRoleName duplicated in same org level",
	ErrorCodeDeviceRoleModelCouldNotUpdateOrdID:                    "OrgId couldn't be update",
	ErrorCodeDeviceRoleModelDeviceRoleTypeNotNormal:                "DeviceRoleType not normal, not normal type couldn't be edit or delete",
	ErrorCodeDeviceRoleModelDeviceRoleIDShouldNotEmpty:             "DeviceRoleID shouldn't be empty",
	ErrorCodeDeviceRoleModelOrgIDShouldNotEmpty:                    "OrgID shouldn't be empty",
	ErrorCodeDeviceRoleModelDeviceRoleNameShouldNotEmpty:           "DeviceRoleName shouldn't be empty",
	ErrorCodeDeviceRoleModelUserCountNotZero:                       "UserCount not zero",
	ErrorCodeDeviceRoleModelSensorIsNotVisibleForCurrentUser:       "Sensor is not visible for current user",
	ErrorCodeDeviceRoleModelDeviceRoleTypeInvalid:                  "DeviceRoleType is not 2/3, which is invalid",
	ErrorCodeDeviceRoleModelDeviceRoleIDCouldNotBeEmptyWhenEdit:    "Can not edit empty DeviceRoleID",
	// 组织
	ErrorCodeOrgModelNotFound:          "Org not found",
	ErrorCodeOrgModelProhibitOperation: "You shouldn't operating this org",
	// 设备
	ErrorCodeSensorModelSingleSensorUpdateFail:                    "Single Sensor update fail, 置位数据库'更新中'失败",
	ErrorCodeSensorNormalDeviceTypeHasNoJurisdictionOfAddSensor:   "NormalDeviceType(type=3) Has No Jurisdiction Of Add Sensor",
	ErrorCodeSensorModelCountExceedLimit:                          "Sensor Count Exceed Limit",
	ErrorCodeSensorCouldNotDeleteOrUpdateBecauseBindToTask:        "can not change status of sensor, because there was a task bind to sensor, please delete task first",
	ErrorCodeSensorCouldNotDeleteOrUpdateBecauseBindToMonitorRule: "can not change status of sensor, because there was a monitor rule bind to sensor, please delete monitor rule first",
	// 登陆者鉴权
	ErrorCodeEmptyLoginUserFound:             "Empty Login User Found",
	ErrorCodeEmptyDeviceRoleOfLoginUserFound: "Empty DeviceRole Of Login User Found",
	// 布控
	ErrorCodeMonitorRuleCouldNotDeleteOrEditBecauseOfOrg:            "该用户没有对该布控删除/编辑(非报警声音/等级的属性)的权限! 因为不符合原因1(组织); 当且仅当以下3点共同成立才有权限: 1. 登陆用户为'布控LastUpdateOrgID'的本身或直属上级; 2. 登陆用户可见该布控内的所有设备; 3. 登陆用户可见该布控内的所有库.",
	ErrorCodeMonitorRuleCouldNotDeleteOrEditBecauseOfSensor:         "该用户没有对该布控删除/编辑(非报警声音/等级的属性)的权限! 因为不符合原因2(设备); 当且仅当以下3点共同成立才有权限: 1. 登陆用户为'布控LastUpdateOrgID'的本身或直属上级; 2. 登陆用户可见该布控内的所有设备; 3. 登陆用户可见该布控内的所有库.",
	ErrorCodeMonitorRuleCouldNotDeleteOrEditBecauseOfRepo:           "该用户没有对该布控删除/编辑(非报警声音/等级的属性)的权限! 因为不符合原因3(库);   当且仅当以下3点共同成立才有权限: 1. 登陆用户为'布控LastUpdateOrgID'的本身或直属上级; 2. 登陆用户可见该布控内的所有设备; 3. 登陆用户可见该布控内的所有库.",
	ErrorCodeMonitorRuleNoVisibleSensorOrNoVisibleRepoWhenAddOrEdit: "新建或编辑布控时: 用户传入的设备/库, 与 用户可见的设备/库, 取交集后没有剩余设备/库, 不符合新建/编辑要求, 请至少选择一个可见的设备/库",
	ErrorCodeMonitorRuleNoExistedMonitorRuleByIDFoundWhenEdit:       "编辑布控接口: 查数据库, 根据被编辑的monitorRuleId, 没找到布控信息",
	ErrorCodeMonitorRuleAllDstStatusAreNotSame:                      "批量开关布控接口, 不满足所有布控的dstEditStatus相同, 只能全部开启或全部关闭",
	ErrorCodeMonitorRuleOneSensorCouldNotBeMonitorRuledTwice:        "this sensor already been ruled yet, can not add a new monitor_rule by using this sensor twice",
	ErrorCodeMonitorRuleGotErrorWhenQuerySensorOfMonitorRuleByID:    "鉴权布控时: 在monitorRuleSensorRelation表中找不到此布控id, 无法完成鉴权",
	ErrorCodeMonitorRuleGotErrorWhenQueryRepoOfMonitorRuleByID:      "鉴权布控时: 在monitorRuleRepoRelation表中找不到此布控id, 无法完成鉴权",
	ErrorCodeMonitorRuleNameTooLongError:                            "monitorRule name too long",
	ErrorCodeMonitorRuleNameAlreadyExistError:                       "monitorRule name already exist",
	// 库
	ErrorCodeRepoCouldNotGetIssuedReposByEmptyOrgID:        "查某org下的issued repo时, org_id非法传参: 为空",
	ErrorCodeRepoCouldNotGetIssuedToOrgOfRepoByEmptyRepoID: "查某repo曾被下发给了哪些org时, repo_id非法传参: 为空",
	// 平台同步
	ErrorCodePlatformSyncModelNotFound:           "Target not found",
	ErrorCodePlatformSyncModelWaitForSync:        "Wait for sync",
	ErrorCodePlatformSyncModelAlreadyDistributed: "Target already distributed",
	ErrorCodePlatformSyncModelConfigNotFound:     "Config not found",
	// 导出
	ErrorCodeExportModelFieldNoFound:              "Filed not found",
	ErrorCodeExportModelLimitInvalid:              "Limit invalid",
	ErrorCodeExportModelTargetTaskNotFound:        "Target task not found",
	ErrorCodeExportModelTargetInnerTypeNotFound:   "Target inner type not found",
	ErrorCodeExportModelFeatureIDMapNOExceedLimit: "Feature IDMapNO exceed limit",
	// 频次分析通用(人+车)
	ErrorCodeAthenaFrequencyModelReachMaxFrequencyAnalyzeTaskNumInAthena:  "Athena Frequency Model 系统可开启频次分析的任务个数为10个（包含人员+车辆频次分析）",
	ErrorCodeAthenaFrequencyModelReachMaxFrequencyAnalyzeTaskNumOfOneUser: "Athena Frequency Model 单个用户同时最多拥有频次分析的任务个数为10个（包含人员+车辆频次分析）",
	ErrorCodeAthenaFrequencyTaskCreateFailWithNoTaskInAthenaAndNoTaskInPg: "创建 task(人频次 或 车频次) 失败, 但是athena 与 pg数据一致的情况, 没有残留的 task",
	// 车辆频次分析
	// athena类
	ErrorCodeAthenaFrequencyTaskModelGotErrorFromAthena:                 "Athena Vehicle Frequency Model Got Error From Athena",
	ErrorCodeAthenaFrequencyTaskModelGotTaskNotExistErrorFromAthena:     "Athena Vehicle Frequency Model Got TaskNotExist Error From Athena",
	ErrorCodeAthenaFrequencyTaskModelGotTaskStateChangedErrorFromAthena: "Athena Vehicle Frequency Model Got TaskStateChanged Error From Athena, athena内部状态已切换, 当前状态不支持此接口调用, 请稍后(根据状态)重试",
	ErrorCodeAthenaFrequencyTaskModelGotMaxTaskNumErrorFromAthena:       "Athena Vehicle Frequency Model Got Max Task Num Error From Athena",
	// loki类
	ErrorCodeAthenaVehicleFrequencyModelInsertIntoDBError:                                              "Athena Vehicle Frequency Model loki将此user和task的信息写入数据库的 athena task + vehicle frequency task 表失败",
	ErrorCodeAthenaVehicleFrequencyModelCreateTaskCouldNotHaveEmptySensorsError:                        "Athena Vehicle Frequency Model create task: sensor数量非法: 为空",
	ErrorCodeAthenaVehicleFrequencyModelCreateTaskSensorCountReachMaxLimitError:                        "Athena Vehicle Frequency Model create task: sensor数量非法: 超过上限(500个)",
	ErrorCodeAthenaVehicleFrequencyModelCreateTaskStartTsGreaterEqualThanEndTsError:                    "Athena Vehicle Frequency Model create task: 时间非法: starttime >= endts",
	ErrorCodeAthenaVehicleFrequencyModelCreateTaskDurationDaysBetweenStartTsAndEndTsReachMaxLimitError: "Athena Vehicle Frequency Model create task: 时间非法: starttime 与 endts 之间间隔超过上限(31天)",
	ErrorCodeAthenaVehicleFrequencyModelCreateTaskIllegalTsError:                                       "Athena Vehicle Frequency Model create task: 时间非法: 如负数",
	ErrorCodeAthenaVehicleFrequencyModelCreateTaskEmptyTaskNameError:                                   "Athena Vehicle Frequency Model create task: 任务名称参数非法",
	ErrorCodeAthenaVehicleFrequencyModelCreateTaskTaskNameTooLongError:                                 "Athena Vehicle Frequency Model create task: 任务名称太长",
	ErrorCodeAthenaVehicleFrequencyModelCreateTaskIllegalTimesThresholdError:                           "Athena Vehicle Frequency Model create task: 频次参数非法",
	ErrorCodeAthenaFrequencyModelSomeTaskIDIsNotUserVisibleError:                                       "Athena Frequency Model 部分传来的taskid不是用户可见的 (因为不是此登陆者创建的 task, 任何用户(包括admin)都只能看到自己建立的频次task",
	ErrorCodeAthenaFrequencyModelTaskDeleteTaskInAthenaFailWithRetry:                                   "Athena Frequency Model 在 athena 内删除此任务失败(包括重试机制在内)",
	// 人员频次分析
	ErrorCodeAthenaPersonFrequencyModelInsertIntoDBError: "Athena Person Frequency Model loki将此user和task的信息写入数据库的 athena task + person frequency task 表失败",

	ErrorCodeAthenaCollisionModelReachMaxCollisionAnalyzeTaskNumInAthena:  "Athena Collision Model 系统可开启区域碰撞的任务个数为10个（包含人员+车辆区域碰撞）",
	ErrorCodeAthenaCollisionModelReachMaxCollisionAnalyzeTaskNumOfOneUser: "Athena Collision Model 单个用户同时最多拥有区域碰撞的任务个数为10个（包含人员+车辆区域碰撞）",
	ErrorCodeAthenaCollisionTaskCreateFailWithNoTaskInAthenaAndNoTaskInPg: "创建 task(人员 或 车区域碰撞) 失败, 但是athena 与 pg数据一致的情况, 没有残留的 task",

	//碰撞任务 athena类
	ErrorCodeAthenaCollisionTaskModelGotErrorFromAthena:                 "Athena Vehicle Collision Model Got Error From Athena",
	ErrorCodeAthenaCollisionTaskModelGotTaskNotExistErrorFromAthena:     "Athena Vehicle Collision Model Got TaskNotExist Error From Athena",
	ErrorCodeAthenaCollisionTaskModelGotTaskStateChangedErrorFromAthena: "Athena Vehicle Collision Model Got TaskStateChanged Error From Athena, athena内部状态已切换, 当前状态不支持此接口调用, 请稍后(根据状态)重试",
	ErrorCodeAthenaCollisionTaskModelGotMaxTaskNumErrorFromAthena:       "Athena Vehicle Collision Model Got Max Task Num Error From Athena",

	ErrorCodeAthenaCollisionModelInsertIntoDBError:                                              "Athena  Collision Model loki将此user和task的信息写入数据库的 athena task  Collision task 表失败",
	ErrorCodeAthenaCollisionModelCreateTaskCouldNotHaveEmptyTimeSpacesError:                     "Athena  Collision Model create task: sensor数量非法: 为空",
	ErrorCodeAthenaCollisionModelCreateTaskSensorCountReachMaxLimitError:                        "Athena  Collision Model create task: sensor数量非法: 超过上限(500个)",
	ErrorCodeAthenaCollisionModelCreateTaskTaskNameTooLongError:                                 "Athena  Collision Model create task: 任务名称太长",
	ErrorCodeAthenaCollisionModelCreateTaskEmptyTaskNameError:                                   "Athena  Collision Model create task: 任务名称参数非法",
	ErrorCodeAthenaCollisionTaskCreateTaskTsAndSensorSuperposition:                              "创建任务区域间得时间和设备有同时重叠",
	ErrorCodeAthenaCollisionModelCreateTaskIllegalTsError:                                       "Athena  Collision Model create task: 时间非法: 如负数",
	ErrorCodeAthenaCollisionModelCreateTaskStartTsGreaterEqualThanEndTsError:                    "Athena  Collision Model create task: 时间非法: starttime >= endts",
	ErrorCodeAthenaCollisionModelCreateTaskDurationDaysBetweenStartTsAndEndTsReachMaxLimitError: "Athena  Collision Model create task: 时间非法: starttime 与 endts 之间间隔超过上限(31天)",
	ErrorCodeAthenaCollisionModelSomeTaskIDIsNotUserVisibleError:                                "Athena Collision Model 部分传来的taskid不是用户可见的 (因为不是此登陆者创建的 task, 任何用户(包括admin)都只能看到自己建立的task",
	ErrorCodeAthenaCollisionModelCreateTaskTaskNameError:                                        "Athena  Collision Model create task: sensorName:已经存在 ",
	ErrorCodeAthenaCollisionTaskNotUser:                                                         "Athena  Collision Task 不是此用户得任务",
	// 首次入城 athena
	ErrorCodeFirstAppearanceModelGotErrorFromAthena:                 "First appearance got error from athena",
	ErrorCodeFirstAppearanceModelGotTaskNotExistErrorFromAthena:     "First appearance got taskNotExist Error From Athena",
	ErrorCodeFirstAppearanceModelGotTaskStateChangedErrorFromAthena: "First appearance got taskStateChanged Error From Athena, athena内部状态已切换, 当前状态不支持此接口调用, 请稍后(根据状态)重试",
	ErrorCodeFirstAppearanceModelGotMaxTaskNumErrorFromAthena:       "First appearance got max task num error from athena",
	// 首次入城 loki
	ErrorCodeFirstAppearanceModelCreateTaskEmptyTaskNameError:     "First appearance create task empty name",
	ErrorCodeFirstAppearanceModelCreateTaskNameTooLongError:       "First appearance create task name too long",
	ErrorCodeFirstAppearanceModelCreateTaskSensorIDsEmptyError:    "First appearance create task sensorIDs empty",
	ErrorCodeFirstAppearanceModelCreateTaskSensorIDsTooManyError:  "First appearance create task sensorIDs too many",
	ErrorCodeFirstAppearanceModelCreateTaskSelectedTimeError:      "First appearance create task selected time invalid",
	ErrorCodeFirstAppearanceModelCreateTaskBacktrackingTimeError:  "First appearance create task backtracking time invalid",
	ErrorCodeFirstAppearanceModelReachSystemMaxTaskLimitError:     "First appearance reach system max task limit",
	ErrorCodeFirstAppearanceModelReachUserMaxTaskLimitError:       "First appearance reach user max task limit",
	ErrorCodeFirstAppearanceModelTaskNameDuplicatedError:          "First appearance task name duplicated",
	ErrorCodeFirstAppearanceModelDeleteTaskInAthenaFailWithRetry:  "First appearance delete task in athena fail with retry",
	ErrorCodeFirstAppearanceModelCreateTaskError:                  "First appearance create task error",
	ErrorCodeFirstAppearanceModelTaskNotFoundError:                "First appearance task not found",
	ErrorCodeFirstAppearanceModelUpdateTypeInvalidError:           "First appearance task update type invalid",
	ErrorCodeFirstAppearanceModelTargetTaskNotUserVisibilityError: "First appearance task target task not user visibility",
	ErrorCodeFirstAppearanceModelCreateTaskPlateTextEmptyError:    "First appearance create task plate text empty",
	// 套牌车 athena
	ErrorCodeFakePlateModelGotErrorFromAthena:                 "Fake plate got error from athena",
	ErrorCodeFakePlateModelGotTaskNotExistErrorFromAthena:     "Fake plate got taskNotExist Error From Athena",
	ErrorCodeFakePlateModelGotTaskStateChangedErrorFromAthena: "Fake plate got taskStateChanged Error From Athena, athena内部状态已切换, 当前状态不支持此接口调用, 请稍后(根据状态)重试",
	ErrorCodeFakePlateModelGotMaxTaskNumErrorFromAthena:       "Fake plate got max task num error from athena",
	// 套牌车 loki
	ErrorCodeFakePlateModelCreateTaskEmptyTaskNameError:     "Fake plate create task empty name",
	ErrorCodeFakePlateModelCreateTaskNameTooLongError:       "Fake plate create task name too long",
	ErrorCodeFakePlateModelCreateTaskSensorIDsEmptyError:    "Fake plate create task sensorIDs empty",
	ErrorCodeFakePlateModelCreateTaskSensorIDsTooManyError:  "Fake plate create task sensorIDs too many",
	ErrorCodeFakePlateModelCreateTaskSelectedTimeError:      "Fake plate create task selected time invalid",
	ErrorCodeFakePlateModelCreateTaskBacktrackingTimeError:  "Fake plate create task backtracking time invalid",
	ErrorCodeFakePlateModelReachSystemMaxTaskLimitError:     "Fake plate reach system max task limit",
	ErrorCodeFakePlateModelReachUserMaxTaskLimitError:       "Fake plate reach user max task limit",
	ErrorCodeFakePlateModelTaskNameDuplicatedError:          "Fake plate task name duplicated",
	ErrorCodeFakePlateModelDeleteTaskInAthenaFailWithRetry:  "Fake plate delete task in athena fail with retry",
	ErrorCodeFakePlateModelCreateTaskError:                  "Fake plate create task error",
	ErrorCodeFakePlateModelTaskNotFoundError:                "Fake plate task not found",
	ErrorCodeFakePlateModelUpdateTypeInvalidError:           "Fake plate task update type invalid",
	ErrorCodeFakePlateModelTargetTaskNotUserVisibilityError: "Fake plate task target task not user visibility",

	// 昼伏夜出 athena
	ErrorCodeAthenaDayInNightOutTaskModelGotErrorFromAthena:                 "DayInNightOut got error from athena",
	ErrorCodeAthenaDayInNightOutTaskModelGotTaskNotExistErrorFromAthena:     "DayInNightOut got taskNotExist Error From Athena",
	ErrorCodeAthenaDayInNightOutTaskModelGotTaskStateChangedErrorFromAthena: "DayInNightOut got taskStateChanged Error From Athena, athena内部状态已切换, 当前状态不支持此接口调用, 请稍后(根据状态)重试",
	ErrorCodeAthenaDayInNightOutTaskModelGotMaxTaskNumErrorFromAthena:       "DayInNightOut max task num error from athena",
	//昼伏夜出 loki
	ErrorCodeAthenaDayInNightOutModelInsertIntoDBError:                                              "DayInNightOut insert db err",
	ErrorCodeAthenaDayInNightOutModelCreateTaskCouldNotHaveEmptySensorsError:                        "DayInNightOut create task sensorIDs empty",
	ErrorCodeAthenaDayInNightOutModelCreateTaskSensorCountReachMaxLimitError:                        "DayInNightOut  create task sensorIDs too many",
	ErrorCodeAthenaDayInNightOutModelCreateTaskTaskNameTooLongError:                                 "DayInNightOut create task name too long",
	ErrorCodeAthenaDayInNightOutModelCreateTaskEmptyTaskNameError:                                   "DayInNightOut create task name empty",
	ErrorCodeAthenaDayInNightOutTaskCreateTaskTsSuperposition:                                       "DayInNightOut create task ts superposition",
	ErrorCodeAthenaDayInNightOutModelCreateTaskIllegalTsError:                                       "DayInNightOut create task ts illegal",
	ErrorCodeAthenaDayInNightOutModelCreateTaskDurationDaysBetweenStartTsAndEndTsReachMaxLimitError: "DayInNightOut create task start and stop ts MaxLimit",
	ErrorCodeAthenaDayInNightOutModelReachMaxTaskNumInAthena:                                        "DayInNightOut reach system max task limit",
	ErrorCodeAthenaDayInNightOutModelReachMaxTaskNumOfOneUser:                                       "DayInNightOut reach user max task limit",
	ErrorCodeAthenaDayInNightOutTaskCreateFailWithNoTaskInAthenaAndNoTaskInPg:                       "DayInNightOut task to Pg err",
	ErrorCodeAthenaDayInNightOutModelTaskDeleteTaskInAthenaFailWithRetry:                            "DayInNightOut delete task in athena fail",
	ErrorCodeAthenaDayInNightOutModelCreateTaskTaskNameError:                                        "DayInNightOut task name duplicated",
	ErrorCodeGetDayInNightOutTaskResultErr:                                                          "DayInNightOut query result err",
	ErrorCodeAthenaDayInNightOutTaskNotUser:                                                         "DayInNightOut task not current user",
	// 用户收藏
	ErrorCodeUserFavoriteExceedLimitError: "User favorite exceed limit",
	ErrorCodeUserFavoriteTargetExistError: "User favorite target exist",
	// 设备预览
	ErrorCodeSensorPreviewNotSupport:      "sensor not support preview",
	ErrorCodeSensorPreviewIllegalSensorID: "illegal sensor_id for preview",
	ErrorCodeSensorPreviewSensorNotExist:  "try to start stream by sensor.url, but sensor not found in sensors table by sensor_id, can not preview",
	// 隐匿车辆
	ErrorCodeAthenaHiddenTaskModelGotErrorFromAthena:                 "Hidden got error from athena",
	ErrorCodeAthenaHiddenTaskModelGotTaskNotExistErrorFromAthena:     "Hidden got taskNotExist Error From Athena",
	ErrorCodeAthenaHiddenTaskModelGotTaskStateChangedErrorFromAthena: "Hidden got taskStateChanged Error From Athena, athena内部状态已切换, 当前状态不支持此接口调用, 请稍后(根据状态)重试",
	ErrorCodeAthenaHiddenTaskModelGotMaxTaskNumErrorFromAthena:       "Hidden max task num error from athena",

	ErrorCodeAthenaHiddenModelCreateTaskTaskNameError:                  "Hidden create task name exist",
	ErrorCodeAthenaHiddenModelReachMaxTaskNumOfOneUser:                 "Hidden reach user max task limit",
	ErrorCodeAthenaHiddenModelInsertIntoDBError:                        "Hidden insert db err",
	ErrorCodeAthenaHiddenModelTaskDeleteTaskInAthenaFailWithRetry:      "Hidden delete task in athena fail",
	ErrorCodeAthenaHiddenTaskCreateFailWithNoTaskInAthenaAndNoTaskInPg: "Hidden task to Pg err",
	ErrorCodeAthenaHiddenModelCreateTaskEmptyTaskNameError:             "Hidden create task name empty",
	ErrorCodeAthenaHiddenModelCreateTaskTaskNameTooLongError:           "Hidden create task name too long",
	ErrorCodeHiddenModelCreateTaskSensorIDsEmptyError:                  "Hidden create task sensorids empty",
	ErrorCodeHiddenModelCreateTaskSensorIDsTooManyError:                "Hidden create task sensorids reach max limit ",
	ErrorCodeHiddenModelCreateTaskSelectedTimeError:                    "Hidden time param err",
	ErrorCodeHiddenModelTargetTaskNotUserVisibilityError:               "Hidden task target task not user visibility",
	ErrorCodeHiddenModelUpdateTypeInvalidError:                         "Hidden task update type invalid",
	ErrorCodeAthenaHiddenTaskNotUser:                                   "Hidden task not is user",
	ErrorCodeGetHiddenTaskResultErr:                                    "Hidden get result err",
	//轨迹分析
	ErrorCodeAthenaPathAnalyzeTaskModelGotErrorFromAthena:                 "PathAnalyze got error from athena",
	ErrorCodeAthenaPathAnalyzeTaskModelGotTaskNotExistErrorFromAthena:     "PathAnalyze got taskNotExist Error From Athena",
	ErrorCodeAthenaPathAnalyzeTaskModelGotTaskStateChangedErrorFromAthena: "PathAnalyze got taskStateChanged Error From Athena, athena内部状态已切换, 当前状态不支持此接口调用, 请稍后(根据状态)重试",
	ErrorCodeAthenaPathAnalyzeTaskModelGotMaxTaskNumErrorFromAthena:       "PathAnalyze max task num error from athena",

	ErrorCodeAthenaPathAnalyzeModelCreateTaskTaskNameError:                  "PathAnalyze create task name exist",
	ErrorCodeAthenaPathAnalyzeModelReachMaxTaskNumOfOneUser:                 "PathAnalyze reach user max task limit",
	ErrorCodeAthenaPathAnalyzeModelInsertIntoDBError:                        "PathAnalyze insert db err",
	ErrorCodeAthenaPathAnalyzeModelTaskDeleteTaskInAthenaFailWithRetry:      "PathAnalyze delete task in athena fail",
	ErrorCodeAthenaPathAnalyzeTaskCreateFailWithNoTaskInAthenaAndNoTaskInPg: "PathAnalyze task to Pg err",
	ErrorCodeAthenaPathAnalyzeModelCreateTaskEmptyTaskNameError:             "PathAnalyze create task name empty",
	ErrorCodeAthenaPathAnalyzeModelCreateTaskTaskNameTooLongError:           "PathAnalyze create task name too long",
	ErrorCodePathAnalyzeModelCreateTaskSensorIDsEmptyError:                  "PathAnalyze create task sensorids empty",
	ErrorCodePathAnalyzeModelCreateTaskSensorIDsTooManyError:                "PathAnalyze create task sensorids reach max limit ",
	ErrorCodePathAnalyzeModelCreateTaskSelectedTimeError:                    "PathAnalyze time param err",
	ErrorCodePathAnalyzeModelTargetTaskNotUserVisibilityError:               "PathAnalyze task target task not user visibility",
	ErrorCodePathAnalyzeModelUpdateTypeInvalidError:                         "PathAnalyze task update type invalid",
	ErrorCodeAthenaPathAnalyzeTaskNotUser:                                   "PathAnalyze task not is user",
	ErrorCodeGetPathAnalyzeTaskResultErr:                                    "PathAnalyze get result err",
	ErrorCodePathAnalyzeModelCreateTaskTypeError:                            "PathAnalyze create task type err",
}

type Error struct {
	Code       int
	Message    string
	InnerError error
}

func (err *Error) Error() string {
	return fmt.Sprintf("Code: %d Message: %s InnerError: %v", err.Code, err.Message, err.InnerError)
}

func (err *Error) String() string {
	return err.InnerError.Error()
}

func NewError(code int, innerError error) *Error {
	return &Error{
		Code:       code,
		Message:    ErrorMessages[code],
		InnerError: innerError,
	}
}

func NewErrorf(code int, format string, args ...interface{}) *Error {
	return &Error{
		Code:       code,
		Message:    ErrorMessages[code],
		InnerError: fmt.Errorf(format, args...),
	}
}

func NewErrorVerificationFailed(code int) *Error {
	return &Error{
		Code:       code,
		Message:    ErrorMessages[code],
		InnerError: fmt.Errorf("Verification failed "),
	}
}

func NewErrorTarget(code int, args ...interface{}) *Error {
	return &Error{
		Code:       code,
		Message:    ErrorMessages[code],
		InnerError: fmt.Errorf("Target %v ", args...),
	}
}
