package entities

type PGExplain struct {
	Plan PGPlan `json:"Plan"`
}

type PGPlan struct {
	NodeType      string  `json:"Node Type"`
	ParallelAware bool    `json:"Parallel Aware"`
	StartupCost   float64 `json:"Startup Cost"`
	TotalCost     float64 `json:"Total Cost"`
	PlanRows      int64   `json:"Plan Rows"`
	PlanWidth     float64 `json:"Plan Width"`
}

type PGQueryPlan struct {
	QueryPlan []byte `xorm:"QUERY PLAN"`
}
