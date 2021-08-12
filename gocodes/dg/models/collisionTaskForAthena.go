package models

//-----------------------------------------------------------------------
// 车辆区域碰撞
//-----------------------------------------------------------------------
//车辆区域碰撞任务状态类型
type VehicleAreaCollisionTaskStateType int32

const (
	VehicleAreaCollisionTaskStateTypeRunning VehicleAreaCollisionTaskStateType = iota
	VehicleAreaCollisionTaskStateTypeError
	VehicleAreaCollisionTaskStateTypeFinish
	VehicleAreaCollisionTaskStateTypeRollback
	VehicleAreaCollisionTaskStateTypeQueueUp
	VehicleAreaCollisionTaskStateTypeStop
)

type VehicleAreaCollisionTaskConfigInfo struct {
	TaskID     string
	TaskName   string
	TimeSpaces []*TimeSpace
}

type VehicleAreaCollisionTaskStatusInfo struct {
	//任务状态类型
	StateType VehicleAreaCollisionTaskStateType

	// 任务排队信息
	PosInQueue  int
	QueueLength int

	//任务进度信息
	ProcessBar int

	Err string
}

// 由clients.StatusResponse 解析后的 车辆碰撞的Athena 结构体
type AthenaUnmarshalVehicleCollisionTaskStatusResp struct {
	ID         string                              // 任务ID
	Type       int                                 // 任务类别
	ConfigInfo *VehicleAreaCollisionTaskConfigInfo // 任务配置信息
	StatusInfo *VehicleAreaCollisionTaskStatusInfo // 任务状态信息
}

//-----------------------------------------------------------------------
// 人员区域碰撞
//-----------------------------------------------------------------------
//车辆区域碰撞任务状态类型
type PersonAreaCollisionTaskStateType int32

const (
	PersonAreaCollisionTaskStateTypeRunning PersonAreaCollisionTaskStateType = iota
	PersonAreaCollisionTaskStateTypeError
	PersonAreaCollisionTaskStateTypeFinish
	PersonAreaCollisionTaskStateTypeRollback
	PersonAreaCollisionTaskStateTypeQueueUp
	PersonAreaCollisionTaskStateTypeStop
)

type PersonAreaCollisionTaskConfigInfo struct {
	TaskID     string
	TaskName   string
	TimeSpaces []*TimeSpace
}

type PersonAreaCollisionTaskStatusInfo struct {
	//任务状态类型
	StateType PersonAreaCollisionTaskStateType

	// 任务排队信息
	PosInQueue  int
	QueueLength int

	//任务进度信息
	ProcessBar int

	Err string
}

// 由clients.StatusResponse 解析后的 车辆碰撞的Athena 结构体
type AthenaUnmarshalPersonCollisionTaskStatusResp struct {
	ID         string                             // 任务ID
	Type       int                                // 任务类别
	ConfigInfo *PersonAreaCollisionTaskConfigInfo // 任务配置信息
	StatusInfo *PersonAreaCollisionTaskStatusInfo // 任务状态信息
}
