package barrier

import "Final_Assessment/common/models"

type BarrierModel struct {
	models.Model
	TransType        string
	Gid              string
	BranchID         string
	Op               string
	BarrierID        int
	DBType           string // DBTypeMysql | DBTypePostgres
	BarrierTableName string
}
