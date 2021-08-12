package clients

import (
	"codes/gocodes/dg/models"
	"context"
	"github.com/olivere/elastic"
)

type ElasticsearchClientImpl struct {
	elastic.Client
}

//func NewElasticsearchClient(con lokiconfig.Config) ElasticsearchClient {
//	f := false
//	conf := config.Config{
//		Sniff: &f,
//		URL:   con.GetString("elasticsearch.url"),
//	}
//	client, err := elastic.NewClientFromConfig(&conf)
//	if err != nil {
//		return nil
//	}
//	return &ElasticsearchClientImpl{
//		Client: *client,
//	}
//}

func (esClient ElasticsearchClientImpl) GetFaceCount(request models.FaceConditionRequest) (int64, error) {
	query := elastic.NewBoolQuery().Filter(elastic.NewRangeQuery("Timestamp").Gte(request.Time.StartTimestamp).Lte(request.Time.EndTimestamp))
	if len(request.SensorID) > 0 {
		q := make([]interface{}, len(request.SensorID))
		for k, v := range request.SensorID {
			q[k] = v
		}
		query = query.Filter(elastic.NewTermsQuery("SensorID.keyword", q...))
	}

	if len(request.NationID) > 0 {
		q := make([]interface{}, len(request.NationID))
		for k, v := range request.NationID {
			q[k] = v
		}
		query = query.Filter(elastic.NewTermsQuery("NationID", q...))
	}

	if len(request.GenderID) > 0 {
		q := make([]interface{}, len(request.GenderID))
		for k, v := range request.GenderID {
			q[k] = v
		}
		query = query.Filter(elastic.NewTermsQuery("GenderID", q...))
	}

	if len(request.GlassID) > 0 {
		q := make([]interface{}, len(request.GlassID))
		for k, v := range request.GlassID {
			q[k] = v
		}
		query = query.Filter(elastic.NewTermsQuery("GlassID", q...))
	}

	if len(request.HatID) > 0 {
		q := make([]interface{}, len(request.HatID))
		for k, v := range request.HatID {
			q[k] = v
		}
		query = query.Filter(elastic.NewTermsQuery("HatID", q...))
	}

	if len(request.MaskID) > 0 {
		q := make([]interface{}, len(request.MaskID))
		for k, v := range request.MaskID {
			q[k] = v
		}
		query = query.Filter(elastic.NewTermsQuery("MaskID", q...))
	}
	return esClient.Count("face").Query(query).Do(context.Background())
}
func (esClient ElasticsearchClientImpl) GetPedestrianCount(request models.PedestrianConditionRequest) (int64, error) {
	query := elastic.NewBoolQuery().Filter(elastic.NewRangeQuery("Timestamp").Gte(request.Time.StartTimestamp).Lte(request.Time.EndTimestamp))
	if len(request.SensorID) > 0 {
		q := make([]interface{}, len(request.SensorID))
		for k, v := range request.SensorID {
			q[k] = v
		}
		query = query.Filter(elastic.NewTermsQuery("SensorID.keyword", q...))
	}
	if len(request.GenderID) > 0 {
		q := make([]interface{}, len(request.GenderID))
		for k, v := range request.GenderID {
			q[k] = v
		}
		query = query.Filter(elastic.NewTermsQuery("GenderID", q...))
	}
	if len(request.NationID) > 0 {
		q := make([]interface{}, len(request.NationID))
		for k, v := range request.NationID {
			q[k] = v
		}
		query = query.Filter(elastic.NewTermsQuery("NationID", q...))
	}
	//hasface入库的时候，就写成0,1,或者2 ,然后，直接查就好了
	if len(request.HasFace) > 0 {
		q := make([]interface{}, len(request.HasFace))
		for k, v := range request.HasFace {
			q[k] = v
		}
		query = query.Filter(elastic.NewTermsQuery("HasFace", q...))
	}
	if len(request.Speed) > 0 {
		q := make([]interface{}, len(request.Speed))
		for k, v := range request.Speed {
			q[k] = v
		}
		query = query.Filter(elastic.NewTermsQuery("Speed", q...))
	}
	if len(request.Direction) > 0 {
		q := make([]interface{}, len(request.Direction))
		for k, v := range request.Direction {
			q[k] = v
		}
		query = query.Filter(elastic.NewTermsQuery("Direction", q...))
	}
	if len(request.UpperStyle) > 0 {
		q := make([]interface{}, len(request.UpperStyle))
		for k, v := range request.UpperStyle {
			q[k] = v
		}
		query = query.Filter(elastic.NewTermsQuery("UpperStyle", q...))
	}
	if len(request.UpColor) > 0 {
		q := make([]interface{}, len(request.UpColor))
		for k, v := range request.UpColor {
			q[k] = v
		}
		query = query.Filter(elastic.NewTermsQuery("UpColor", q...))
	}
	if request.WithBackpack != 0 {
		query = query.Filter(elastic.NewTermQuery("WithBackpack", 1))
	}
	if request.WithShoulderBag != 0 {
		query = query.Filter(elastic.NewTermQuery("WithShoulderBag", 1))
	}
	if request.DescMouth != 0 {
		query = query.Filter(elastic.NewTermQuery("DescMouth", 1))
	}
	if request.DescEye != 0 {
		query = query.Filter(elastic.NewTermQuery("DescEye", 1))
	}
	if request.WithHandbag != 0 {
		query = query.Filter(elastic.NewTermQuery("WithHandbag", 1))
	}
	if request.WithHandCarry != 0 {
		query = query.Filter(elastic.NewTermQuery("WithHandCarry", 1))
	}
	if request.WithPram != 0 {
		query = query.Filter(elastic.NewTermQuery("WithPram", 1))
	}
	if request.WithLuggage != 0 {
		query = query.Filter(elastic.NewTermQuery("WithLuggage", 1))
	}
	if request.WithTrolley != 0 {
		query = query.Filter(elastic.NewTermQuery("WithTrolley", 1))
	}
	if request.WithUmbrella != 0 {
		query = query.Filter(elastic.NewTermQuery("WithUmbrella", 1))
	}
	if request.WithHoldBaby != 0 {
		query = query.Filter(elastic.NewTermQuery("WithHoldBaby", 1))
	}
	if request.WithScarf != 0 {
		query = query.Filter(elastic.NewTermQuery("WithScarf", 1))
	}
	if len(request.DescHead) > 0 {
		q := make([]interface{}, len(request.DescHead))
		for k, v := range request.DescHead {
			q[k] = v
		}
		query = query.Filter(elastic.NewTermsQuery("DescHead", q...))
	}
	if len(request.HairStyle) > 0 {
		q := make([]interface{}, len(request.HairStyle))
		for k, v := range request.HairStyle {
			q[k] = v
		}
		query = query.Filter(elastic.NewTermsQuery("HairStyle", q...))
	}
	if len(request.AgeID) > 0 {
		q := make([]interface{}, len(request.AgeID))
		for k, v := range request.AgeID {
			q[k] = v
		}
		query = query.Filter(elastic.NewTermsQuery("AgeID", q...))
	}
	if len(request.UpperTexture) > 0 {
		q := make([]interface{}, len(request.UpperTexture))
		for k, v := range request.UpperTexture {
			q[k] = v
		}
		query = query.Filter(elastic.NewTermsQuery("UpperTexture", q...))
	}
	if len(request.LowerColor) > 0 {
		q := make([]interface{}, len(request.LowerColor))
		for k, v := range request.LowerColor {
			q[k] = v
		}
		query = query.Filter(elastic.NewTermsQuery("LowerColor", q...))
	}
	if len(request.LowerStyle) > 0 {
		q := make([]interface{}, len(request.LowerStyle))
		for k, v := range request.LowerStyle {
			q[k] = v
		}
		query = query.Filter(elastic.NewTermsQuery("LowerStyle", q...))
	}
	if len(request.ShoesColor) > 0 {
		q := make([]interface{}, len(request.ShoesColor))
		for k, v := range request.ShoesColor {
			q[k] = v
		}
		query = query.Filter(elastic.NewTermsQuery("ShoesColor", q...))
	}
	if len(request.ShoesStyle) > 0 {
		q := make([]interface{}, len(request.ShoesStyle))
		for k, v := range request.ShoesStyle {
			q[k] = v
		}
		query = query.Filter(elastic.NewTermsQuery("ShoesStyle", q...))
	}

	return esClient.Count("pedestrian").Query(query).Do(context.Background())
}

func (esClient ElasticsearchClientImpl) GetVehicleCount(request models.VehicleConditionRequest) (int64, error) {
	query := elastic.NewBoolQuery().Filter(elastic.NewRangeQuery("Timestamp").Gte(request.Time.StartTimestamp).Lte(request.Time.EndTimestamp))
	if len(request.SensorID) > 0 {
		q := make([]interface{}, len(request.SensorID))
		for k, v := range request.SensorID {
			q[k] = v
		}
		query = query.Filter(elastic.NewTermsQuery("SensorID.keyword", q...))
	}

	if len(request.TypeID) > 0 {
		q := make([]interface{}, len(request.TypeID))
		for k, v := range request.TypeID {
			q[k] = v
		}
		query = query.Filter(elastic.NewTermsQuery("TypeID", q...))
	}

	if len(request.ColorID) > 0 {
		q := make([]interface{}, len(request.ColorID))
		for k, v := range request.ColorID {
			q[k] = v
		}
		query = query.Filter(elastic.NewTermsQuery("ColorID", q...))
	}

	if len(request.PlateTypeID) > 0 {
		q := make([]interface{}, len(request.PlateTypeID))
		for k, v := range request.PlateTypeID {
			q[k] = v
		}
		query = query.Filter(elastic.NewTermsQuery("PlateTypeID", q...))
	}

	if len(request.PlateColorID) > 0 {
		q := make([]interface{}, len(request.PlateColorID))
		for k, v := range request.PlateColorID {
			q[k] = v
		}
		query = query.Filter(elastic.NewTermsQuery("PlateColorID", q...))
	}

	if len(request.Direction) > 0 {
		q := make([]interface{}, len(request.Direction))
		for k, v := range request.Direction {
			q[k] = v
		}
		query = query.Filter(elastic.NewTermsQuery("Direction", q...))
	}

	if len(request.Speed) > 0 {
		q := make([]interface{}, len(request.Speed))
		for k, v := range request.Speed {
			q[k] = v
		}
		query = query.Filter(elastic.NewTermsQuery("Speed", q...))
	}

	if len(request.Side) > 0 {
		q := make([]interface{}, len(request.Side))
		for k, v := range request.Side {
			q[k] = v
		}
		query = query.Filter(elastic.NewTermsQuery("Side", q...))
	}

	if len(request.Specials) > 0 {
		for _, v := range request.Specials {
			query = query.Filter(elastic.NewTermQuery("Specials", v))
		}
	}

	if len(request.Symbols) > 0 {
		for _, v := range request.Symbols {
			query = query.Filter(elastic.NewTermQuery("Symbols", v))
		}
	}

	if len(request.HasFace) > 0 {
		q := make([]interface{}, len(request.HasFace))
		for k, v := range request.HasFace {
			q[k] = v
		}
		query = query.Filter(elastic.NewTermsQuery("HasFace", q...))
	}

	if len(request.Brands) > 0 {
		tmpQ := elastic.NewBoolQuery()
		for _, v := range request.Brands {
			tmpQuery := elastic.NewBoolQuery()
			if v.Brand != "" {
				tmpQuery = tmpQuery.Filter(elastic.NewTermQuery("Brand", v.Brand))
			} else {
				continue
			}
			if v.SubBrand != "" {
				tmpQuery = tmpQuery.Filter(elastic.NewTermQuery("SubBrand", v.SubBrand))
			}
			if v.Year != "" {
				tmpQuery = tmpQuery.Filter(elastic.NewTermQuery("Year", v.Year))
			}
			tmpQ = tmpQ.Should(tmpQuery)
		}
		query = query.Filter(tmpQ)
	}

	if request.Plate != "" {
		query = query.Filter(elastic.NewTermQuery("Plate.keyword", request.Plate))
	}
	count, err := esClient.Count("vehicle").Query(query).Do(context.Background())

	return count, err
}

func (esClient ElasticsearchClientImpl) GetNonmotorCount(request models.NonmotorConditionRequest) (int64, error) {
	query := elastic.NewBoolQuery().Filter(elastic.NewRangeQuery("Timestamp").Gte(request.Time.StartTimestamp).Lte(request.Time.EndTimestamp))
	if len(request.SensorID) > 0 {
		q := make([]interface{}, len(request.SensorID))
		for k, v := range request.SensorID {
			q[k] = v
		}
		query = query.Filter(elastic.NewTermsQuery("SensorID.keyword", q...))
	}
	if len(request.PlateColorID) > 0 {
		q := make([]interface{}, len(request.PlateColorID))
		for k, v := range request.PlateColorID {
			q[k] = v
		}
		query = query.Filter(elastic.NewTermsQuery("PlateColorID.keyword", q...))
	}
	if len(request.NonmotorColorID) > 0 {
		q := make([]interface{}, len(request.NonmotorColorID))
		for k, v := range request.NonmotorColorID {
			q[k] = v
		}
		query = query.Filter(elastic.NewTermsQuery("NonmotorColorID", q...))
	}
	if len(request.NonmotorType) > 0 {
		q := make([]interface{}, len(request.NonmotorType))
		for k, v := range request.NonmotorType {
			q[k] = v
		}
		query = query.Filter(elastic.NewTermsQuery("NonmotorType", q...))
	}
	if len(request.NonmotorGesture) > 0 {
		q := make([]interface{}, len(request.NonmotorGesture))
		for k, v := range request.NonmotorGesture {
			q[k] = v
		}
		query = query.Filter(elastic.NewTermsQuery("NonmotorGesture", q...))
	}
	if len(request.DescHead) > 0 {
		q := make([]interface{}, len(request.DescHead))
		for k, v := range request.DescHead {
			q[k] = v
		}
		query = query.Filter(elastic.NewTermsQuery("DescHead", q...))
	}
	if len(request.GenderID) > 0 {
		q := make([]interface{}, len(request.GenderID))
		for k, v := range request.GenderID {
			q[k] = v
		}
		query = query.Filter(elastic.NewTermsQuery("GenderID", q...))
	}
	if len(request.NationID) > 0 {
		q := make([]interface{}, len(request.NationID))
		for k, v := range request.NationID {
			q[k] = v
		}
		query = query.Filter(elastic.NewTermsQuery("NationID", q...))
	}
	if len(request.UpColor) > 0 {
		q := make([]interface{}, len(request.UpColor))
		for k, v := range request.UpColor {
			q[k] = v
		}
		query = query.Filter(elastic.NewTermsQuery("UpColor", q...))
	}
	if len(request.UpperStyle) > 0 {
		q := make([]interface{}, len(request.UpperStyle))
		for k, v := range request.UpperStyle {
			q[k] = v
		}
		query = query.Filter(elastic.NewTermsQuery("UpperStyle", q...))
	}
	if len(request.Speed) > 0 {
		q := make([]interface{}, len(request.Speed))
		for k, v := range request.Speed {
			q[k] = v
		}
		query = query.Filter(elastic.NewTermsQuery("Speed", q...))
	}
	if len(request.Direction) > 0 {
		q := make([]interface{}, len(request.Direction))
		for k, v := range request.Direction {
			q[k] = v
		}
		query = query.Filter(elastic.NewTermsQuery("Direction", q...))
	}
	if request.DescEye == 1 {
		query = query.Filter(elastic.NewTermQuery("DescEye", request.DescEye))
	}
	if request.DescMouth == 1 {
		query = query.Filter(elastic.NewTermQuery("DescMouth", request.DescMouth))
	}
	if request.WithShoulderBag == 1 {
		query = query.Filter(elastic.NewTermQuery("WithShoulderBag", request.WithShoulderBag))
	}
	if request.WithBackpack == 1 {
		query = query.Filter(elastic.NewTermQuery("WithBackpack", request.WithBackpack))
	}

	//hasface入库的时候，就写成0,1,或者2 ,然后，直接查就好了
	if len(request.HasFace) > 0 {
		q := make([]interface{}, len(request.HasFace))
		for k, v := range request.HasFace {
			q[k] = v
		}
		query = query.Filter(elastic.NewTermsQuery("HasFace", q...))
	}

	if request.Plate != "" {
		query = query.Filter(elastic.NewTermQuery("Plate.keyword", request.Plate))
	}
	return esClient.Count("nonmotor").Query(query).Do(context.Background())
}
