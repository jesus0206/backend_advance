
drop DATABASE if EXISTS bd_yofio;
CREATE DATABASE bd_yofio;
COMMENT on DATABASE bd_yofio is 'base de datos para empresa yofio';


-- credit_asignacion(monto,type_300,type_500,type_700,exito)
drop table if exists credit_asignacion;
create table credit_asignacion(
  id integer GENERATED ALWAYS AS IDENTITY PRIMARY key, 
  monto INTEGER,
  type_300 INTEGER,
  type_500 INTEGER,
  type_700 INTEGER,
  exito BOOLEAN
);

