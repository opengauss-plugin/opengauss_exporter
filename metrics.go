package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var openGaussClient *sql.DB

func initOpenGaussClient() (db *sql.DB, err error) {
	defer func() {
		//捕获func抛出的panic，防止主程序崩溃
		if err := recover(); err != nil {
			Error("initOpenGaussClient发生错误：", err)
		}
	}()

	if openGaussClient != nil {
		return openGaussClient, nil
	} else {
		connStr := DATA_SOURCE_NAME
		openGaussClient, err := sql.Open("postgres", connStr)
		if err != nil {
			log.Fatal(err)
		}
		return openGaussClient, nil
	}
}

/* 慢查询语句数量统计

	性能很差的SQL被大量执行
    xact_stay 执行持续秒数 超过60秒就是慢查询
*/
func og_slow_select_count() (count string, err error) {
	defer func() {
		//捕获func抛出的panic，防止主程序崩溃
		if err := recover(); err != nil {
			Error("og_slow_select_count发生错误：", err)
		}
	}()

	db, _ := initOpenGaussClient()
	rows, err := db.Query(`select 
	datname,usename,client_addr,application_name,state,backend_start,xact_start,xact_stay,query_start,query_stay,
	replace(query,chr(10),' ') as query 
  from 
	(
	  select 
		pgsa.datname as datname, 
		pgsa.usename as usename, 
		pgsa.client_addr client_addr, 
		pgsa.application_name as application_name, 
		pgsa.state as state, 
		pgsa.backend_start as backend_start, 
		pgsa.xact_start as xact_start, 
		extract(
		  epoch 
		  from 
			(now() - pgsa.xact_start)
		) as xact_stay, 
		pgsa.query_start as query_start, 
		extract(
		  epoch 
		  from 
			(now() - pgsa.query_start)
		) as query_stay, 
		pgsa.query as query 
	  from 
		pg_stat_activity as pgsa 
	  where 
		pgsa.state != 'idle' 
		and pgsa.state != 'idle in transaction' 
		and pgsa.state != 'idle in transaction (aborted)'
	) idleconnections 
  order by 
	query_stay desc`)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var counter int
	for rows.Next() {
		counter++
		var xact_stay int
		err = rows.Scan(&xact_stay)
		if err := rows.Scan(&xact_stay); err != nil {
		}
	}

	return fmt.Sprintf("%d", counter), err
}
