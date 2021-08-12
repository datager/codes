package models

type TaskStatus int32

const (
	TaskStatus_Task_Status_Unknown    TaskStatus = 0
	TaskStatus_Task_Status_Created    TaskStatus = 1
	TaskStatus_Task_Status_Processing TaskStatus = 2
	TaskStatus_Task_Status_Finished   TaskStatus = 3
	TaskStatus_Task_Status_Deleted    TaskStatus = 4
	TaskStatus_Task_Status_Outdate    TaskStatus = 5
	TaskStatus_Task_Status_Paused     TaskStatus = 6
	TaskStatus_Task_Status_Stopped    TaskStatus = 7
)
