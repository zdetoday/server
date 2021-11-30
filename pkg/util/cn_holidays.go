package util

// ChinaHolidays 中国法定节假日
var ChinaHolidays = map[string]map[string]bool{
	"2021": ChinaHolidays2021,
	"2022": ChinaHolidays2022,
}

// ChinaHolidays2021 2021年中国法定节假日
var ChinaHolidays2021 = map[string]bool{
	// 周五 ～ 周日	元旦
	"2021-01-01": true,
	"2021-01-02": true,
	"2021-01-03": true,
	// 周四 ～ 周三	春节
	"2021-02-11": true,
	"2021-02-12": true,
	"2021-02-13": true,
	"2021-02-14": true,
	"2021-02-15": true,
	"2021-02-16": true,
	"2021-02-17": true,
	// 周六 ～ 周一	清明节
	"2021-04-03": true,
	"2021-04-04": true,
	"2021-04-05": true,
	// 周六 ～ 周三	劳动节
	"2021-05-03": true,
	"2021-05-04": true,
	"2021-05-05": true,
	// 周六 ～ 周一	端午节
	"2021-06-12": true,
	"2021-06-13": true,
	"2021-06-14": true,
	// 周日 ～ 周二	中秋节
	"2021-09-19": true,
	"2021-09-20": true,
	"2021-09-21": true,
	// 周五 ～ 周四	国庆日
	"2021-10-01": true,
	"2021-10-02": true,
	"2021-10-03": true,
	"2021-10-04": true,
	"2021-10-05": true,
	"2021-10-06": true,
	"2021-10-07": true,
}

// ChinaHolidays2022 2022年中国法定节假日
var ChinaHolidays2022 = map[string]bool{
	// 周六 ～ 周一	元旦
	"2022-01-01": true,
	"2022-01-02": true,
	"2022-01-03": true,
	// 周一 ～ 周日	春节
	"2022-01-31": true,
	"2022-02-01": true,
	"2022-02-02": true,
	"2022-02-03": true,
	"2022-02-04": true,
	"2022-02-05": true,
	"2022-02-06": true,
	// 周日 ～ 周二	清明节
	"2022-04-03": true,
	"2022-04-04": true,
	"2022-04-05": true,
	// 周六 ～ 周三	劳动节
	"2022-04-30": true,
	"2022-05-01": true,
	"2022-05-02": true,
	"2022-05-03": true,
	"2022-05-04": true,
	// 周五 ～ 周日	端午节
	"2022-06-03": true,
	"2022-06-04": true,
	"2022-06-05": true,
	// 周六 ～ 周一	中秋节
	"2022-09-10": true,
	"2022-09-11": true,
	"2022-09-12": true,
	// 周六 ～ 周五	国庆日
	"2022-10-01": true,
	"2022-10-02": true,
	"2022-10-03": true,
	"2022-10-04": true,
	"2022-10-05": true,
	"2022-10-06": true,
	"2022-10-07": true,
}
