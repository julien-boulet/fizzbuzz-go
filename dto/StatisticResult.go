package dto

type StatisticResult struct {
	Int1  int    `schema:"int1"`
	Int2  int    `schema:"int2"`
	Limit int    `schema:"limit"`
	Str1  string `schema:"str1"`
	Str2  string `schema:"str2"`
	Count int    `schema:"count"`
}
