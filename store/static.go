package store

import "yofio/avanzado/store/model"

func (repo *Repository) Static() (*model.CreditAssignment, error) {
	static := &model.CreditAssignment{}
	rows, err := repo.db.Raw(
		`select
		count(*) solicitudes,
		count(case when exito='True' THEN 1 end) as exito_total,
		count(case when exito='False' THEN 1 end) as no_exito_total,
		avg(case when exito='True' THEN monto end) as inversion_exitosa,
		avg(case when exito='False' THEN monto end) as inversion_no_exitosa
		from credit_asignacion`,
	).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&static.Total,
			&static.SuccessTotal,
			&static.SuccessTotalNot,
			&static.InvestmentSuccess,
			&static.InvestmentSuccessNot,
		)
		if err != nil {
			return nil, err
		}
	}
	return static, nil
}
