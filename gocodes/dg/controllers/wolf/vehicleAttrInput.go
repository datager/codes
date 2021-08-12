package wolf

import (
	"time"

	"codes/gocodes/dg/dbengine"
	"codes/gocodes/dg/entities"
	"codes/gocodes/dg/utils"
	"codes/gocodes/dg/utils/uuid"

	"github.com/golang/glog"
)

func InputVehicleAttr() {
	plates, err := getPlateTexts(1591286400000, 1591347600000)
	if err != nil {
		glog.Fatal(err)
	}

	err = inputIntoVehicleAttrTable("0000", "d80952df-26fc-413e-b991-4d1c4848d077", plates)
	if err != nil {
		glog.Fatal(err)
	}
}

func getPlateTexts(st, et int64) ([]string, error) {
	ret := make([]string, 0)
	err := dbengine.GetV5Instance().
		Table(entities.TableNameVehicleCaptureIndexEntity).
		Select("distinct(plate_text)").
		Where("plate_text != ''").
		Where("ts >= ?", st).
		Where("ts < ?", et).
		Limit(300).
		Find(&ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func inputIntoVehicleAttrTable(orgID, repoID string, plateTexts []string) error {
	attrs := make([]*entities.VehicleAttrs, 0)
	for _, p := range plateTexts {
		attrs = append(attrs,
			&entities.VehicleAttrs{
				ID:        uuid.NewV4().String(),
				RepoID:    repoID,
				OrgID:     orgID,
				PlateText: p,
				Comment:   "车辆",
				Ts:        utils.GetNowTs(),
				Uts:       time.Now(),
			})
	}

	sess := dbengine.GetV5Instance().NewSession()
	_, err := sess.Table(entities.TableNameVehicleAttrs).InsertMulti(attrs)
	if err != nil {
		return err
	}
	return nil
}
