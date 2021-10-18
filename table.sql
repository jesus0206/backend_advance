
drop DATABASE if EXISTS bd_yofio;
CREATE DATABASE bd_yofio;
COMMENT on DATABASE bd_yofio is 'base de datos para empresa yofio';


-- credit_asignacion(monto,type_300,type_500,type_700,exito)
drop table if exists credit_asignacion;
create table credit_asignacion(
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY key, 
  monto INTEGER not null,
  type_300 INTEGER,
  type_500 INTEGER,
  type_700 INTEGER,
  exito BOOLEAN not null
);


select*from credit_asignacion;


select
  count(*) solicitudes,
  count(case when exito='True' THEN 1 end) as exito_total,
  count(case when exito='False' THEN 1 end) as no_exito_total,
  avg(case when exito='True' THEN monto end) as inversion_exitosa,
  avg(case when exito='False' THEN monto end) as inversion_no_exitosa
from credit_asignacion
;