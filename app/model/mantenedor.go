package model

import (
	"database/sql"

	"github.com/josephspurrier/gowebapp/app/shared/database"
)

var (
	QRYListaTotalTablas string = "CALL GPR_CNS_TAB_PAR(?);"
	QRYListaTablas      string = "CALL MET_SEL_Lista_Tablas();"
)

func ListaTablas(IdTabla string) ([]map[string]interface{}, error) {
	var err error
	var rows *sql.Rows

	if IdTabla == "" || IdTabla == "*" {
		// No hacer nada si IdTabla es vacío o "*"
		return []map[string]interface{}{}, nil
	} else {
		rows, err = database.SQL.Query(QRYListaTotalTablas, IdTabla)
		if err != nil {
			return nil, standardizeError(err)
		}
		defer rows.Close()
	}

	// Slice para almacenar los mapas de resultados
	var results []map[string]interface{}

	cols, err := rows.Columns()
	if err != nil {
		return nil, standardizeError(err)
	}

	// Preparar un slice de interfaces para cada columna
	values := make([]interface{}, len(cols))
	for i := range values {
		values[i] = new(interface{})
	}

	for rows.Next() {
		err = rows.Scan(values...)
		if err != nil {
			return nil, standardizeError(err)
		}

		// Crear un mapa para almacenar cada fila
		result := make(map[string]interface{})
		for i, colName := range cols {
			val := *(values[i].(*interface{}))
			if b, ok := val.([]byte); ok {
				result[colName] = string(b)
			} else {
				result[colName] = val
			}
		}

		results = append(results, result)
	}
	return results, nil
}

type MET_Lista_Tablas struct {
	Message     string `db:"Message" bson:"Message"`
	IdTabla     int32  `db:"IdTabla" bson:"IdTabla"`
	NombreTabla string `db:"NombreTabla" bson:"NombreTabla"`
}

func CargaCbxListaTabla() ([]MET_Lista_Tablas, error) {

	var err error
	var result []MET_Lista_Tablas
	err = database.SQL.Select(&result, QRYListaTablas)
	return result, standardizeError(err)

}
