package models

type IdType int32

const (
	IdType_Id_Type_Unknown   IdType = 0
	IdType_Id_Type_IdNo      IdType = 1
	IdType_Id_Type_Passport  IdType = 2
	IdType_Id_Type_Officer   IdType = 3
	IdType_Id_Type_Sergeant  IdType = 4
	IdType_Id_Type_Temporary IdType = 5
)

var (
	IdTypeMap = make(map[IdType]string, 0)
)

func init() {
	IdTypeMap[IdType_Id_Type_Unknown] = "未知"
	IdTypeMap[IdType_Id_Type_IdNo] = "身份证"
	IdTypeMap[IdType_Id_Type_Passport] = "护照"
	IdTypeMap[IdType_Id_Type_Officer] = "警官证"
	IdTypeMap[IdType_Id_Type_Sergeant] = "军官证"
	IdTypeMap[IdType_Id_Type_Temporary] = "临时身份证"
}
