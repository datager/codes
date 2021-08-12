package smartquery

import "codes/gocodes/dg/utils/log"

// case1: 输入(offset, limit) = (0, 100), 输出 explainCnt = 1, lenList = 0, ===>修正为 fixedExplainCnt = offset(0) + lenList(0) = 0
// case2: 输入(offset, limit) = (300, 400), 输出 explainCnt = 5, lenList = 1, ===>修正为 fixedExplainCnt = offset(0) + lenList(0) = 301
func FixSillyExplainCnt(explainCnt, offset, limit, lenList int64) (fixedExplainCnt int64) {
	log.Debugf("start FixSillyExplainCnt(), explainCnt: %d, offset: %d, limit: %d, lenList: %d", explainCnt, offset, limit, lenList)
	offsetPlusLenList := offset + lenList

	if lenList < limit { // 如果按 limit 去查, 但是查出的lenList没查满limit, 则说明pg 里只有 offset + lenList 的总数, 返回此数
		log.Debugln("because lenList < limit, so return offsetPlusLenList")
		return offsetPlusLenList
	} else if lenList == limit && explainCnt <= offsetPlusLenList { // 如果按 limit 去查, 而且查出的lenList查满了limit, 此时需要判断 "explainCnt" 与 "offset + limit" 的大小 (此时offset + lenList = offset + limit)
		log.Debugln("because lenList == limit && explainCnt <= offsetPlusLenList, so return offsetPlusLenList")
		return offsetPlusLenList
	}

	log.Debugln("no need to fix silly explainCnt, it is ok")
	return explainCnt
}

/*
type CntAndErr struct {chr
	cnt int64
	err error
}

type ListAndErr struct {
	list []*entities.VehiclesIndex
	err  error
}

func NewListAndErr() ListAndErr {
	return ListAndErr{
		list: make([]*entities.VehiclesIndex, 0),
		err:  nil,
	}
}

// todo: later
func SearchAndCountAsync(condition models.PaginationInterface, repositoryFactory repositories.RepositoryFactory,
	countAsync func(interface{}, chan CntAndErr, repositories.TransactionScope),
	searchAsync func(interface{}, chan ListAndErr, repositories.TransactionScope),
	searchSync func(interface{}, chan ListAndErr, repositories.TransactionScope)) (result *models.SearchResult, err error) {
	//var vehicleCondition *models.VehicleConditionRequest
	//var pedestrainCondition *models.PedestrianConditionRequest
	//var nonmotorCondition *models.NonmotorConditionRequest
	//var faceCondition *models.FaceConditionRequestConditionRequest
	//
	//
	//
	//switch condition.(type) {
	//case *models.VehicleConditionRequest:
	//	condition
	//}


	st := time.Now()
	log.Infoln("[--Capture--Query--Vehicle--] start VehicleService.SearchAndCountAsync()")
	defer log.Infof("[--Capture--Query--Vehicle--] finish VehicleService.SearchAndCountAsync(), took: %s", time.Since(st))

	//listResult = make([]*entities.VehiclesIndex, 0)

	// tx
	txCnt, err := repositoryFactory.NewTransactionScope()
	if err != nil {
		return nil, err
	}

	txList, err := repositoryFactory.NewTransactionScope()
	if err != nil {
		return nil, err
	}

	// step1: count + list
	chanCnt := make(chan CntAndErr)
	chanList := make(chan ListAndErr)
	language countAsync(condition, chanCnt, txCnt)
	language searchAsync(condition, chanList, txList)

	//defer func() {
	//	log.Infof("[--Capture--Query--Vehicle--] 查完了, cntResult: %d", cntResult)
	//	log.Infof("[--Capture--Query--Vehicle--] 查完了, len(listResult): %d", len(listResult))
	//}()
	isCntOK := false
	for {
		select {
		case cnt := <-chanCnt:
			log.Infoln("[--Capture--Query--Vehicle--] chanCnt 比 chanList 先收到, 取`已经得到且准确的cnt`, 不用去取`explainCnt了")

			if cnt.err != nil {
				log.Errorf("[--Capture--Query--Vehicle--] 接收到的cnt有错: %v", cnt.err)
				return nil, err
			}

			cntResult = cnt.cnt
			language func(tx repositories.TransactionScope) { // 既然已收到 cnt, 可关 txCnt
				_ = tx.Commit() // we do not need to handle err of tx.Commit() which is readonly of pg
				tx.Close()
			}(txCnt)

			// cnt 与 condition 的 3 种关系
			if cntResult <= int64(condition.Offset) { // 关系1: 查出来的 cnt(如 5), 小于 offset(如 300), 理论上前端不会出现此调用情况, 但是为防止 postman 等方式, 后端以返回空的方式保护
				log.Infof("[--Capture--Query--Vehicle--] 关系1成立, 返回空")
				language txList.Close()
				listResult = []*entities.VehiclesIndex{}
			} else if cntResult > int64(condition.Offset) && cntResult < int64(condition.Offset+condition.Limit) { // 关系2: 查出来的 cnt(如 5) 介于 [offset(如 300), offset+limit(如 400)]之间, 需要停止 listTx, 重新开始 listTx2
				log.Infof("[--Capture--Query--Vehicle--] 关系2成立, 停止 listTx, 重新开始 listTx2")

				language txList.Close()
				log.Infof("[--Capture--Query--Vehicle--] txList已停止")

				// 依据 cnt, 重新修正limit参数
				condition.GetLimit() = int(cntResult) - condition.GetOffset()
				log.Infof("[--Capture--Query--Vehicle--] 依据 cnt, 已重新修正limit参数, 修正后 offset:%d, limit: %d, 准备重新按修正的参数再查一次list2", condition.Offset, condition.Limit)

				st := time.Now()
				log.Infof("[--Capture--Query--Vehicle--] 开始接收list2的结果")
				ssList2 := vs.repositoryFactory.NewSimpleSessionScope()
				list2, err := vs.SearchSync(condition, ssList2)
				ssList2.Close()
				log.Infof("[--Capture--Query--Vehicle--] 结束接收list2的结果, took: %s", time.Since(st))
				if err != nil {
					log.Errorf("[--Capture--Query--Vehicle--] 接收到的list2有错: %v", err)
					return 0, nil, err
				}

				listResult = list2
				return cntResult, listResult, nil
			} else { // 关系3: 查出来的 cnt(如 401) 大于 offset+limit(如 400), 继续等待 listTx 的结果即可
				log.Infof("[--Capture--Query--Vehicle--] 关系3成立, 继续等待 listTx 的结果即可")
				isCntOK = true
				continue
			}

		case list := <-chanList:
			log.Infoln("[--Capture--Query--Vehicle--] chanList 比 chanCnt 先收到, 终止掉`还在运行的cnt`, 准备去取`explainCnt`")

			if list.err != nil {
				log.Errorf("[--Capture--Query--Vehicle--] 接收到的list有错: %v", list.err)
				return nil, err
			}

			if !isCntOK {
				// step2: explain count
				explainCnt, err := vs.ExplainCount(condition, txList)
				if err != nil {
					return nil, err
				}
				cntResult = explainCnt

				language func(tx repositories.TransactionScope) { // 既然已收到 list, 可关 txList
					_ = tx.Commit()
					tx.Close()
				}(txList)

				language txCnt.Close()
			}
			listResult = list.list

			return cntResult, listResult, nil
		}
	}

	getSensorNameByIDs()
}

func getSensorNameByIDs([]string) []string {
	return nil
}
*/
