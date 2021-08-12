package sensor

import (
	"fmt"
	"time"

	"codes/gocodes/dg/dbengine"
	"codes/gocodes/dg/entities"
	"codes/gocodes/dg/service/org"
	"codes/gocodes/dg/utils/math"
	"codes/gocodes/dg/utils/uuid"

	"github.com/golang/glog"
	"golang.org/x/sync/errgroup"
)

// 目标: 每个org下x个sensor
// orgLevelToCreate: 3 // 有几级org
// sensorNumInEachOrgToCreate: 10 // 每级兄弟org从level2-no0到level2-no9

// 读org: db
// 写org: 接口
func BatchCreate(orgLevelToCreate, orgNumInEachLevelToCreate, sensorNumInEachOrgToCreate int) {
	//token, err := login.Login()
	//if err != nil {
	//	glog.Fatal(err)
	//}

	g := errgroup.Group{}
	tokenPool := make(chan struct{}, 100)

	// 从db读org(按orgName的命名规则)
	// 对每个org下写sensor
	for orgLevel := 1; orgLevel < orgLevelToCreate; orgLevel++ { // orgLevel: 待添加的子org的name(若为1, 则上级orgLevel为0)
		for orgNum := 0; orgNum < math.PowInt(orgNumInEachLevelToCreate, orgLevel); orgNum++ { // orgNum: 兄弟组织的number序号, level1有10^1个兄弟组织, level2有10^2个兄弟组织, level3有10^3个兄弟组织

			orgNameCreated := org.GenCreateOrgName(orgLevel, orgNum) // level2-no0
			org, err := org.GetByOrgName(orgNameCreated)
			if err != nil {
				glog.Fatal(err)
			}

			/*
				// by api http
				for sensorNum := 0; sensorNum < sensorNumInEachOrgToCreate; sensorNum++ {

					tokenPool <- struct{}{}
					g.Go(func() error {
						defer func() {
							<-tokenPool
						}()

						sensor := &models.Sensor{
							OrgId:      orgID,
							//OrgCode: 	org.OrgCode // 新建设备用api时, loki会查并写入orgCode
							SensorName: GenCreateSensorName(orgLevel, orgNum, sensorNum),
							Url:        "rtsp://192.168.2.123/live/t1009",
							Type:       3,
							Longitude:  -200,
							Latitude:   -200,
							Status:     1,
							Timestamp:  time.Now().UnixNano() / 1e6,
						}

						sensorJs, err := json.Marshal(sensor)
						r := bytes.NewReader(sensorJs)
						req, err := http.NewRequest(http.MethodPost, models.SensorCreateURL, r)
						req.Header.Set("Content-Type", "application/json")
						req.Header.Set("authorization", fmt.Sprintf("Bearer %v", token))
						resp, err := http.DefaultClient.Do(req)
						if err != nil {
							glog.Fatal(err)
						}
						if resp.StatusCode/100 != 2 { // should be 204
							glog.Fatal(resp.StatusCode)
						}

						glog.Infof("create sensor %v done", GenCreateSensorName(orgLevel, orgNum, sensorNum))
						return nil
					})
				}

				if err := g.Wait(); err != nil {
					glog.Fatal("wait fail", err)
				}
			*/

			// by db engine, 可以batch
			tokenPool <- struct{}{}
			g.Go(func() error {
				defer func() {
					<-tokenPool
				}()

				tx := dbengine.GetV5Instance().NewSession()
				if err := tx.Begin(); err != nil {
					return err
				}

				sensors := make([]*entities.Sensors, 0)
				for sensorNum := 0; sensorNum < sensorNumInEachOrgToCreate; sensorNum++ {
					sensors = append(sensors,
						&entities.Sensors{
							OrgId:      org.OrgId,
							OrgCode:    org.OrgCode,
							SensorID:   uuid.NewV4().String(),
							SensorName: GenCreateSensorName(orgLevel, orgNum, sensorNum),
							Url:        "rtsp://192.168.2.123/live/t1009",
							Type:       3,
							Longitude:  -200,
							Latitude:   -200,
							Status:     1,
							Ts:         time.Now().UnixNano() / 1e6,
							Uts:        time.Now(),
						},
					)
				}

				_, err = dbengine.GetV5Instance().Table(entities.TableNameSensors).InsertMulti(&sensors)
				glog.Infof("create sensor %v done", GenCreateSensorName(orgLevel, orgNum, -123456789))
				if err != nil {
					return err
				}

				if err := tx.Commit(); err != nil {
					return err
				}
				return nil
			})
		}
	}

	if err := g.Wait(); err != nil {
		glog.Fatal("wait fail", err)
	}
	glog.Info("create sensor ok, orgLevelToCreate:%v, orgNumInEachLevelToCreate:%v, sensorNumInEachOrgToCreate:%v", orgLevelToCreate, orgNumInEachLevelToCreate, sensorNumInEachOrgToCreate)

}

func GenCreateSensorName(orgLevel, orgNum int, sensorNum int) string {
	return fmt.Sprintf("level%v-no%v-s%v", orgLevel, orgNum, sensorNum)
}
