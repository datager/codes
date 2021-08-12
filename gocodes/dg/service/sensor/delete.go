package sensor

import (
	"bytes"
	"codes/gocodes/dg/dbengine"
	"codes/gocodes/dg/models"
	login2 "codes/gocodes/dg/service/login"
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"io/ioutil"
	"net/http"
	"sync"
)

func DelAllSensors() {
	token, err := login2.Login()
	if err != nil {
		glog.Fatal(err)
	}

	//err = delSensorsByIDs([]string{"45b31f52-0b96-4794-9a1a-2fdb1fc252f0"}, token)
	//if err != nil {
	//	glog.Fatal(err)
	//}

	sensorIDs, err := GetAllSensorIDs()
	if err != nil {
		glog.Fatal(err)
	}

	var lk sync.RWMutex
	var g errgroup.Group

	section := len(sensorIDs) / BatchDelSizeForSensor
	for i := 0; i < section; i++ {
		i := i
		start := i
		end := i + BatchDelSizeForSensor
		if end > len(sensorIDs) {
			end = len(sensorIDs)
		}

		g.Go(func() error {
			lk.RLock()
			sIDs := sensorIDs[start:end]
			lk.RUnlock()
			return delSensorsByIDs(sIDs, token)
		})
	}

	if err := g.Wait(); err != nil {
		glog.Fatal(err)
	}

	glog.Info("done")
}

func delSensorsByIDs(sIDs []string, token string) error {
	sIDsJs, err := json.Marshal(sIDs)
	if err != nil {
		return err
	}

	r := bytes.NewReader(sIDsJs)
	req, err := http.NewRequest("DELETE", models.SensorDelURL, r)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("authorization", fmt.Sprintf("Bearer %v", token))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode/100 != 2 {
		return errors.New("StatusCode not 2xx")
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	glog.Info(string(respBytes))
	return nil
}

func DelSensorsByIDsInDB(sIDs []string) error {
	x, err := dbengine.GetV5Instance().Delete(sIDs)
	_ = x
	if err != nil {
		return err
	}
	return nil
}
