package org

import (
	"bytes"
	login2 "codes/gocodes/dg/service/login"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"codes/gocodes/dg/dbengine"
	"codes/gocodes/dg/entities"
	"codes/gocodes/dg/models"
	"codes/gocodes/dg/utils/math"

	"github.com/golang/glog"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"xorm.io/builder"
)

// 目标: 自动化建组织, 建3级组织, 每级10个兄弟组织
// orgLevelToCreate: 3
// orgNumInEachLevelToCreate: 10 // 每级兄弟org从level2-no0到level2-no9

// 读org: db
// 写org: 接口
func BatchCreate(orgLevelToCreate, orgNumInEachLevelToCreate int) {
	token, err := login2.Login()
	if err != nil {
		glog.Fatal(err)
	}

	for orgLevel := 1; orgLevel < orgLevelToCreate; orgLevel++ { // orgLevel: 待添加的子org的name(若为1, 则上级orgLevel为0)
		for orgNum := 0; orgNum < math.PowInt(orgNumInEachLevelToCreate, orgLevel); orgNum++ { // orgNum: 兄弟组织的number序号, level1有10^1个兄弟组织, level2有10^2个兄弟组织, level3有10^3个兄弟组织

			orgNameToCreate := GenCreateOrgName(orgLevel, orgNum) // level2-no0

			var superOrgID string
			if orgLevel == 1 { // 2
				superOrgID = "0000"
			} else {
				superOrgName := GenCreateOrgName(orgLevel-1, orgNum/orgNumInEachLevelToCreate) // 根据本机orgLevel+orgNum, 确定上级orgName GenCreateOrgName(2-1, 0/10)=>superOrgName=level1-no0
				id, err := GetIDByOrgName(superOrgName)
				if err != nil {
					glog.Fatal(err)
				}
				superOrgID = id
			}

			orgJs, err := json.Marshal(
				&models.Org{
					SuperiorOrgId: superOrgID,
					OrgLevel:      int64(orgLevel),
					OrgName:       orgNameToCreate,
					Ts:            time.Now().UnixNano() / 1e6,
				})
			r := bytes.NewReader(orgJs)
			req, err := http.NewRequest(http.MethodPost, models.OrgCreateURL, r)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("authorization", fmt.Sprintf("Bearer %v", token))
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				glog.Fatal(err)
			}
			if resp.StatusCode/100 != 2 { // should be 204
				glog.Fatal(resp.StatusCode)
			}
			glog.Infof("create org %v done", orgNameToCreate)
		}
	}
	glog.Infof("create org ok, orgLevelToCreate:%v, orgNumInEachLevelToCreate:%v", orgLevelToCreate, orgNumInEachLevelToCreate)
}

func GenCreateOrgName(orgLevel, orgNum int) string {
	return fmt.Sprintf("level%v-no%v", orgLevel, orgNum)
}

func GetByOrgName(orgName string) (*entities.OrgStructure, error) {
	ret := make([]*entities.OrgStructure, 0)
	err := dbengine.GetV5Instance().
		Table(entities.TableNameOrgStructure).
		Where("org_name = ?", orgName).
		OrderBy("uts desc").
		Limit(1).
		Find(&ret)
	if err != nil {
		return nil, err
	}
	if len(ret) == 0 {
		return nil, errors.New("len(ret) == 0")
	}

	return ret[0], err
}

func GetIDByOrgName(orgName string) (string, error) {
	b := builder.Postgres().
		Select("org_id").
		From(entities.TableNameOrgStructure).
		Where(builder.Eq{"org_name": orgName})
	sql, err := b.ToBoundSQL()
	if err != nil {
		return "", err
	}
	ret, err := dbengine.GetV5Instance().QueryString(sql)
	if err != nil {
		return "", err
	}
	if len(ret) == 0 {
		return "", errors.New("len(ret) == 0")
	}

	return ret[0]["org_id"], err
}
