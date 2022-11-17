package main

// "context"

// "github.com/vilchykau/golangtest/internal/db/model"
// "github.com/vilchykau/golangtest/internal/db/price"
// "github.com/vilchykau/golangtest/internal/drivers"

import "github.com/vilchykau/golangtest/internal/app"

func main() {
	// ctx := context.Background()

	// driver :=
	// 	drivers.NewPgDriver(drivers.PgConfig{DnsString: "User ID=habrpguser;Password=pgpwd4habr;Host=127.0.0.1;Port=5432;Database=habrdb;Pooling=true;Min Pool Size=0;Max Pool Size=100;Connection Lifetime=0;"})
	// driver.Connect(ctx)

	// url := "SomeUrl"
	// var price1 float64 = 1233.4

	// priceMan := price.NewPrice(driver.GetDb())
	// priceMan.InsertPrice(ctx, &model.Price{Url: &url, Price: &price1})

	app := &app.App{}
	app.Start()
}
