package service

import (
	"database/sql"
	"github.com/jboulet/fizzbuzz-go/dto"
	"log"
)

const (
	INSERT = "INSERT INTO statistic (int1, int2, lim, str1, str2, count) VALUES ($1, $2, $3, $4, $5, 1);"
	UPDATE = "UPDATE statistic SET count = (count +1) WHERE int1=$1 AND int2=$2 AND lim=$3 AND str1=$4 AND str2=$5;"
	EXISTS = "SELECT EXISTS(SELECT 1 FROM statistic WHERE int1=$1 AND int2=$2 AND lim=$3 AND str1=$4 AND str2=$5);"
	TOP    = "SELECT int1, int2, lim, str1, str2, count FROM statistic ORDER BY count DESC LIMIT 1;"
)

func Save(database *sql.DB, gameParameter dto.GameParameter) {

	var exists bool
	err := database.QueryRow(EXISTS, gameParameter.Int1, gameParameter.Int2, gameParameter.Limit, gameParameter.Str1, gameParameter.Str2).Scan(&exists)
	if err != nil {
		log.Fatal("Error checking if row exists : ", err)
	}
	if exists {
		_, err := database.Exec(UPDATE, gameParameter.Int1, gameParameter.Int2, gameParameter.Limit, gameParameter.Str1, gameParameter.Str2)
		if err != nil {
			log.Fatal("Database UPDATE failed", err)
		}
	} else {
		_, err := database.Exec(INSERT, gameParameter.Int1, gameParameter.Int2, gameParameter.Limit, gameParameter.Str1, gameParameter.Str2)
		if err != nil {
			log.Fatal("Database INSERT failed", err)
		}
	}
}

func FindMax(database *sql.DB) dto.StatisticResult {
	statisticResult := dto.StatisticResult{}
	err := database.QueryRow(TOP).Scan(&statisticResult.Int1, &statisticResult.Int2, &statisticResult.Limit, &statisticResult.Str1, &statisticResult.Str2, &statisticResult.Count)
	if err != nil {
		log.Fatal("Error checking top result : ", err)
	}
	return statisticResult
}
