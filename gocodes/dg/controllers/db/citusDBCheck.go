package db

func dbCheck() {
	// citus扩容 - 20200915
	// 建分区表
	// 灌数据
	// 校验扩容后是否还能用(通过是不是还能查的相等来确定)
	/*
		sourceDBConnStr := "postgres://postgres:Zstvgcs@9102@192.168.2.130/deepface_v7?sslmode=disable"
		targetDBConnStr := "postgres://postgres:Zstvgcs@9102@192.168.2.47/deepface_v6?sslmode=disable"
		tableNames := []string{
			entities.TableNameFaces,
			entities.TableNameFacesIndex,
			entities.TableNameVehicleCaptureEntity,
			entities.TableNameVehicleCaptureIndexEntity,
			entities.TableNameNonmotorCapture,
			entities.TableNameNonmotorIndexCapture,
			entities.TableNamePedestrianCapture,
			entities.TableNamePedestrianIndexCapture,
			//entities.TableNamePetrolVehicles,
		}
		startTs := utils.ParseLocalTimeStr2Ts("2020-01-06 00:00:00")
		endTs := utils.ParseLocalTimeStr2Ts("2020-05-04 00:00:00")

		AutoCreatePartitionTable(startTs, endTs, targetDBConnStr)
		QueryDataFromDB1AndInsertIntoDB2(sourceDBConnStr, targetDBConnStr, tableNames, startTs, endTs)
		CheckIsDBStatisticsEqual(FirstStatisticsType, targetDBConnStr, tableNames, startTs, endTs)
		CheckIsDBStatisticsEqual(SecondStatisticsType, targetDBConnStr, tableNames, startTs, endTs)
	*/
}
