package models

type Even struct {
	Id                  int    `json:"id"`
	QqNum               string `json:"qq_num"`
	EvenContentMoreInfo string `json:"even_content_more_info"`
	EvenContent         string `json:"even_content"`
	VerifyStatus        int    `json:"verify_status"`
}

func (even *Even) QqNumIsExist() (error error, isExist bool) {
	sqlSql := "select 1 from `qqblk_even` where `qq_num` = ? limit 1;"
	if row, err := Db.Query(sqlSql, even.QqNum); err != nil {
		return err, false
	} else if !row.Next() {
		return nil, false
	}
	return nil, true
}

func (even *Even) GetEvenAllDataForQq() error {
	sqlStr := "SELECT  `id`, `qq_num`, `even_content_more_info`, `even_content`, `verify_status` FROM `qqblk_even` WHERE `qq_num` = ?;"
	err := Db.QueryRow(sqlStr, even.QqNum).
		Scan(&even.Id, &even.QqNum, &even.EvenContentMoreInfo, &even.EvenContent, &even.VerifyStatus)
	if err != nil {
		return err
	}
	return nil
}
