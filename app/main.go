package main

import (
	"database/sql"
	"log"
	"net/url"
	"os"
	"strconv"
	"time"

	colors "github.com/Aritiaya50217/Test-Assignment-ANC/color"
	mysqlRepo "github.com/Aritiaya50217/Test-Assignment-ANC/internal/repository/mysql"
	"github.com/Aritiaya50217/Test-Assignment-ANC/internal/rest"
	"github.com/Aritiaya50217/Test-Assignment-ANC/internal/rest/middleware"
	product "github.com/Aritiaya50217/Test-Assignment-ANC/product"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

const (
	defaultTimeout = 30
	defaultAddress = ":8080"
)

func main() {

	// dbHost := os.Getenv("DATABASE_HOST")
	// dbPort := os.Getenv("DATABASE_PORT")
	// dbUser := os.Getenv("DATABASE_USER")
	// dbPass := os.Getenv("DATABASE_PASS")
	// dbName := os.Getenv("DATABASE_NAME")
	// connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Bangkok")

	dsn := "root:test@tcp(127.0.0.1:3306)/ecommerce"
	dbConn, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("failed to open connection to database", err)
	}

	err = dbConn.Ping()
	if err != nil {
		log.Fatal("failed to ping database ", err)
	}
	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal("got error when closing the DB connection", err)
		}
	}()

	e := echo.New()
	e.Use(middleware.CORS)
	timeoutStr := os.Getenv("CONTEXT_TIMEOUT")
	timeout, err := strconv.Atoi(timeoutStr)
	if err != nil {
		log.Println("failed to parse timeout,  using default timeout")
		timeout = defaultTimeout
	}
	timeoutContext := time.Duration(timeout) * time.Second
	e.Use(middleware.SetRequestContextWithTimeout(timeoutContext))

	// Prepare service layer
	colorRepo := mysqlRepo.NewColorRepository(dbConn)
	productRepo := mysqlRepo.NewProductRepository(dbConn)

	// Build service layer
	colorService := colors.NewService(colorRepo)
	productService := product.NewService(productRepo)

	rest.NewColorHandler(e, colorService)
	rest.NewProductHandler(e, productService)

	// Start Server
	address := os.Getenv("PORT")
	if address == "" {
		address = defaultAddress
	}
	log.Fatal(e.Start(address))

}
