package store

import (
	"yofio/avanzado/store/model"

	"gorm.io/gorm"
)

// A Repository connection db
type Repository struct {
	db *gorm.DB
}

// NewRepository exported
func NewRepository(db *gorm.DB) IRepository {
	return &Repository{db: db}
}

// IDataBase interface
type IRepository interface {
	CreditAssignment(investment int32) (int32, int32, int32, error)
	Static() (*model.CreditAssignment, error)
}

// select
//   count(*) solicitudes,
//   count(case when exito='True' THEN 1 end) as exito_total,
//   count(case when exito='False' THEN 1 end) as no_exito_total,
//   avg(case when exito='True' THEN monto end) as inversion_exitosa,
//   avg(case when exito='False' THEN monto end) as inversion_no_exitosa
// from credit_asignacion
