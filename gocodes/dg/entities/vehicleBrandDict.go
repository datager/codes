package entities

type VehicleBrandDict struct {
	BrandID    int64  `xorm:"brand_id not null VARCHAR(1024)"`
	SubBrandID int64  `xorm:"sub_brand_id not null VARCHAR(1024)"`
	YearID     int64  `xorm:"year_id not null VARCHAR(1024)"`
	Desc       string `xorm:"desc not null VARCHAR(1024)"`
}

func (VehicleBrandDict) TableName() string {
	return TableNameVehicleBrandDict
}

//func ModelToEntities(vehicleBrandDict *models.VehicleBrandDict) []*VehicleBrandDict {
//	var result []*VehicleBrandDict
//	for brandID, brandFloor := range vehicleBrandDict.Data["brand"] {
//		for subBrandID, subBrandFloor := range brandFloor.SubBrand {
//			for yearID, yearFloor := range subBrandFloor.Year {
//				brandIDInt, _ := strconv.ParseInt(brandID, 0, 0)
//				subBrandIDInt, _ := strconv.ParseInt(subBrandID, 0, 0)
//				yearIDInt, _ := strconv.ParseInt(yearID, 0, 0)
//				result = append(result, &VehicleBrandDict{
//					BrandID:    brandIDInt,
//					SubBrandID: subBrandIDInt,
//					YearID:     yearIDInt,
//					Desc:       fmt.Sprintf("%s_%s_%s", brandFloor.Name.Chinese, subBrandFloor.Name.Chinese, yearFloor.Name.Chinese),
//				})
//			}
//		}
//	}
//	return result
//}
