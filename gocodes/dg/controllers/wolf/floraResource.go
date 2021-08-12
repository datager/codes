package wolf

import (
	"codes/gocodes/dg/dbengine"
	"codes/gocodes/dg/entities"
	"codes/gocodes/dg/utils/json"
	"context"
	"github.com/golang/glog"
	"golang.org/x/sync/errgroup"
	"strconv"
	"sync"
	"time"
)

/*
结论: v3最快, 即用xormQueryInterface(), 然后直接类型强转, 且用最小的entity(v31和v3的区别).
I0618 09:47:44.533896  181720 floraResource.language:38] tststsv1:1.823557034s // 直接查QueryInterface()的性能
I0618 09:47:50.171193  181720 floraResource.language:69] tststsv2:5.637262435s // QueryByte() + 大entity + marshal + unmarshal
I0618 09:47:52.449834  181720 floraResource.language:93] tststsv3:2.278606783s // QueryInterface() + 大entity + 类型强转
I0618 09:47:54.355559  181720 floraResource.language:218] tststsv31:1.905684259s // QueryInterface() + 小entity + 类型强转
I0618 09:47:57.119168  181720 floraResource.language:145] ts9259s // 10000个gouroutine, QueryByte() + 大entity + 类型强转, 比不开gouroutine快(可能是因为本身耗时就少, 没必要开这么多gouroutine)
I0618 09:47:59.865393  181720 floraResource.language:176] tststsv5:2.746188841s
I0618 09:48:07.701373  181720 floraResource.language:193] tststsv6:7.835814076s
*/

func GetResourceV1() {
	ss := dbengine.GetV5Instance().NewSession()
	defer ss.Close()
	st := time.Now()

	for i := 0; i < 100; i++ {
		ss.Table("sensors").Select("sensor_id, org_id, org_code")
		iface, err := ss.QueryInterface()
		if err != nil {
			panic("111")
		}
		_ = iface
	}
	tm := time.Since(st)
	glog.Infof("tststsv1:%v", tm)
}

func GetResourceV2() {
	ss := dbengine.GetV5Instance().NewSession()
	defer ss.Close()
	st := time.Now()

	for i := 0; i < 100; i++ {
		ss.Table("sensors").Select("sensor_id, org_id, org_code")
		iface, err := ss.QueryInterface()
		if err != nil {
			panic("111")
		}

		sensors := make([]*entities.Sensors, 0)

		for _, data := range iface {
			byte, err := json.Marshal(data)
			if err != nil {
				panic("err")
			}

			s := &entities.Sensors{}
			err = json.Unmarshal(byte, s)
			if err != nil {
				panic("err")
			}
			sensors = append(sensors, s)
		}
	}
	tm := time.Since(st)
	glog.Infof("tststsv2:%v", tm)
}

func GetResourceV3() {
	ss := dbengine.GetV5Instance().NewSession()
	defer ss.Close()
	st := time.Now()

	for i := 0; i < 100; i++ {
		sensors := make([]*entities.Sensors, 0)
		ss.Table("sensors").Select("sensor_id, org_id, org_code")
		iface, err := ss.QueryInterface()
		if err != nil {
			panic("111")
		}
		for _, data := range iface {
			sensors = append(sensors, &entities.Sensors{
				SensorID: data["sensor_id"].(string),
				OrgCode:  data["org_code"].(int64),
				OrgId:    data["org_id"].(string),
			})
		}
	}
	tm := time.Since(st)
	glog.Infof("tststsv3:%v", tm)
}

// v3之上的v31, 精简model
func GetResourceV31() {
	ss := dbengine.GetV5Instance().NewSession()
	defer ss.Close()
	st := time.Now()

	for i := 0; i < 100; i++ {
		sensors := make([]*entities.FloraMinSensors, 0)
		ss.Table("sensors").Select("sensor_id, org_id, org_code")
		iface, err := ss.QueryInterface()
		if err != nil {
			panic("111")
		}
		for _, data := range iface {
			sensors = append(sensors, &entities.FloraMinSensors{
				SensorID: data["sensor_id"].(string),
				OrgCode:  data["org_code"].(int64),
				OrgId:    data["org_id"].(string),
			})
		}
	}
	tm := time.Since(st)
	glog.Infof("tststsv31:%v", tm)
}
func GetResourceV4() {
	ss := dbengine.GetV5Instance().NewSession()
	defer ss.Close()
	st := time.Now()

	for i := 0; i < 100; i++ {
		ss.Table("sensors").Select("sensor_id, org_id, org_code")
		iface, err := ss.QueryInterface()
		if err != nil {
			panic("111")
		}

		threadCh := make(chan struct{}, 100)
		g, _ := errgroup.WithContext(context.TODO())
		lk := sync.Mutex{}
		sensors := make([]*entities.FloraMinSensors, 0)

		for _, data := range iface {
			data := data
			threadCh <- struct{}{}

			g.Go(func() error {
				defer func() {
					<-threadCh
				}()

				s := &entities.FloraMinSensors{
					SensorID: data["sensor_id"].(string),
					OrgCode:  data["org_code"].(int64),
					OrgId:    data["org_id"].(string),
				}

				lk.Lock()
				sensors = append(sensors, s)
				lk.Unlock()
				return nil
			})
		}

		if err := g.Wait(); err != nil {
			panic("wait")
		}
	}

	tm := time.Since(st)
	glog.Infof("tststsv4:%v", tm)
}

func GetResourceV41() {
	ss := dbengine.GetV5Instance().NewSession()
	defer ss.Close()
	st := time.Now()

	for i := 0; i < 100; i++ {
		ss.Table("sensors").Select("sensor_id, org_id, org_code")
		iface, err := ss.QueryInterface()
		if err != nil {
			panic("111")
		}

		sensors := make([]*entities.FloraMinSensors, 0)
		threadCh := make(chan struct{}, 100)

		g, _ := errgroup.WithContext(context.TODO())
		retCh := make(chan *entities.FloraMinSensors, len(iface))

		for i := range iface {
			data := iface[i]
			threadCh <- struct{}{}

			g.Go(func() error {
				defer func() {
					<-threadCh
				}()

				s := &entities.FloraMinSensors{
					SensorID: data["sensor_id"].(string),
					OrgCode:  data["org_code"].(int64),
					OrgId:    data["org_id"].(string),
				}

				retCh <- s
				return nil
			})
		}
		if err := g.Wait(); err != nil {
			panic("wait")
		}

		//for s := range retCh {
		//	sensors = append(sensors, s)
		//}
		for i := 0; i < len(iface); i++ {
			s := <-retCh
			sensors = append(sensors, s)
		}
		glog.Infof("ok%V", 1)
	}

	tm := time.Since(st)
	glog.Infof("tststsv41:%v", tm)
}

func GetResourceV5() {
	ss := dbengine.GetV5Instance().NewSession()
	defer ss.Close()
	st := time.Now()

	for i := 0; i < 100; i++ {
		ss.Table("sensors").Select("sensor_id, org_id, org_code")
		iface, err := ss.Query()
		if err != nil {
			panic("111")
		}

		sensors := make([]*entities.Sensors, 0)
		for _, data := range iface {
			oc, err := strconv.ParseInt(string(data["org_code"]), 10, 64)
			if err != nil {
				panic("int")
			}

			s := &entities.Sensors{
				SensorID: string(data["sensor_id"]),
				OrgCode:  oc,
				OrgId:    string(data["org_id"]),
			}
			sensors = append(sensors, s)
		}
	}
	tm := time.Since(st)
	glog.Infof("tststsv5:%v", tm)
}

func GetResourceV6() {
	ss := dbengine.GetV5Instance().NewSession()
	defer ss.Close()
	st := time.Now()

	for i := 0; i < 100; i++ {
		sensors := make([]*entities.Sensors, 0)
		ss.Table("sensors").Select("sensor_id, org_id, org_code")
		err := ss.Find(&sensors)
		if err != nil {
			panic("111")
		}
	}
	tm := time.Since(st)
	glog.Infof("tststsv6:%v", tm)
}
