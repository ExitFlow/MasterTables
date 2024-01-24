package model

import (
	"github.com/josephspurrier/gowebapp/app/shared/database"
)

// *****************************************************************************
// Opciones
// *****************************************************************************

// Opciones table contains the information for each Option
type Opciones struct {
	ID           uint32 `db:"id" json:"id,omitempty"`
	Name         string `db:"name" json:"name"`
	Url          string `db:"URL" json:"URL"`
	Menu         string `db:"menu" json:"menu"`
	Comment      string `db:"comment" json:"comment,omitempty"`
	Exec         string `db:"exec" json:"exec"`
	InternalName string `db:"internalname" json:"internalname"`
	Order        uint8  `db:"order" json:"order"`
}

// OptionsByUserID gets all notes for a user
func OpcionesByUserID(userID string) ([]Opciones, error) {
	var err error

	var result []Opciones

	err = database.SQL.Select(&result, "SELECT  optn.`id`, optn.`name`, optn.`URL`, optn.`menu`, optn.`comment`, optn.`exec`, optn.`internalname`, optn.`order`"+
		" FROM `sec_options` optn"+
		" join `sec_profiles_options` prop on prop.id_option = optn.id and prop.deleted = 0 "+
		" join `sec_profiles` prof	on prop.id_profile = prof.id and prof.deleted = 0 "+
		" join `sec_user_profiles` uspr on uspr.id_profile = prof.id "+
		" join `sec_user`				usrs	on uspr.id_user		= usrs.id and usrs.deleted = 0 "+
		" where  optn.`deleted` = 0 "+
		" and    usrs.id = ? "+
		" and    optn.`application` = 'MASTERTABLES' "+
		" group by optn.`id`, optn.`name`, optn.`URL`, optn.`menu`, optn.`comment`, optn.`exec`, optn.`internalname`, optn.`order` "+
		" Order by optn.`menu`, `order`", userID)

	return result, standardizeError(err)
}
