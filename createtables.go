package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "12345678"
	dbname   = "dbdbdb"
)

func main() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS orders (
		id SERIAL PRIMARY KEY,
		order_uid VARCHAR(255) NOT NULL,
		track_number VARCHAR(255) NOT NULL,
		entry VARCHAR(255) NOT NULL,
		delivery_name VARCHAR(255) NOT NULL,
		delivery_phone VARCHAR(20) NOT NULL,
		delivery_zip VARCHAR(10) NOT NULL,
		delivery_city VARCHAR(255) NOT NULL,
		delivery_address VARCHAR(255) NOT NULL,
		delivery_region VARCHAR(255) NOT NULL,
		delivery_email VARCHAR(255) NOT NULL,
		transaction VARCHAR(255) NOT NULL,
		request_id VARCHAR(255) NOT NULL,
		currency VARCHAR(10) NOT NULL,
		provider VARCHAR(255) NOT NULL,
		amount INT NOT NULL,
		payment_dt INT NOT NULL,
		bank VARCHAR(255) NOT NULL,
		delivery_cost INT NOT NULL,
		goods_total INT NOT NULL,
		custom_fee INT NOT NULL,
		item_chrt_id INT NOT NULL,
		item_track_number VARCHAR(255) NOT NULL,
		item_price INT NOT NULL,
		item_rid VARCHAR(255) NOT NULL,
		item_name VARCHAR(255) NOT NULL,
		item_sale INT NOT NULL,
		item_size VARCHAR(255) NOT NULL,
		item_total_price INT NOT NULL,
		item_nm_id INT NOT NULL,
		item_brand VARCHAR(255) NOT NULL,
		item_status INT NOT NULL,
		locale VARCHAR(10) NOT NULL,
		internal_signature VARCHAR(255) NOT NULL,
		customer_id VARCHAR(255) NOT NULL,
		delivery_service VARCHAR(255) NOT NULL,
		shardkey VARCHAR(10) NOT NULL,
		sm_id INT NOT NULL,
		date_created TIMESTAMP NOT NULL,
		oof_shard VARCHAR(10) NOT NULL
	);`

	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Table 'orders' created successfully!")

	//код для вставки данных в таблицу

	insertData := `
	INSERT INTO orders (
		order_uid, track_number, entry, delivery_name, delivery_phone, delivery_zip,
		delivery_city, delivery_address, delivery_region, delivery_email, transaction,
		request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total,
		custom_fee, item_chrt_id, item_track_number, item_price, item_rid, item_name,
		item_sale, item_size, item_total_price, item_nm_id, item_brand, item_status,
		locale, internal_signature, customer_id, delivery_service, shardkey, sm_id,
		date_created, oof_shard
	) VALUES (
		'b563feb7b2b84b6test', 'WBILMTESTTRACK', 'WBIL', 'Test Testov', '+9720000000',
		'2639809', 'Kiryat Mozkin', 'Ploshad Mira 15', 'Kraiot', 'test@gmail.com',
		'b563feb7b2b84b6test', '', 'USD', 'wbpay', 1817, 1637907727, 'alpha', 1500,
		317, 0, 9934930, 'WBILMTESTTRACK', 453, 'ab4219087a764ae0btest', 'Mascaras',
		30, '0', 317, 2389212, 'Vivienne Sabo', 202, 'en', '', 'test', 'meest', '9',
		99, '2021-11-26T06:22:19Z', '1'
	);
	`

	_, err = db.Exec(insertData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data inserted successfully!")
	fmt.Println("Tables created...")

}
