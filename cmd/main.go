package main

import (
	_ "github.com/lib/pq"
	"library/config"
	"library/internal/db"
	"library/internal/server"
	"log"
)

const (
	ConnectStatus = "Connect status:Succesful"
	PingStatus    = "Ping status:Succesful"
)

func main() {
	cfg, err := config.StartConfig()
	if err != nil {
		log.Fatal(err)
	}

	//Database
	DB, err := db.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	defer DB.Close()

	server := server.StartServer(DB)

	if err := server.Router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

	//server := gin.Default()

	//server.GET("/", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"status": ConnectStatus,
	//	})
	//})
	//
	//server.Run(":8080")

	//	DB, err := db.Connect(cfg.DatabaseURL)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	fmt.Println(ConnectStatus)
	//
	//	if err := db.Ping(DB); err != nil {
	//		log.Fatal(err)
	//	}
	//	fmt.Println(PingStatus)

	//	fmt.Println("Enter the operation\n1.Book\n2.Books")
	//	var operation int
	//	fmt.Scan(&operation)
	//	switch operation {
	//	case 1:
	//		fmt.Println("Enter:")
	//		reader := bufio.NewReader(os.Stdin)
	//
	//		BookName, _ := reader.ReadString('\n')
	//		BookName = strings.TrimSpace(BookName)
	//
	//		Author, _ := reader.ReadString('\n')
	//		Author = strings.TrimSpace(BookName)
	//
	//		YearRealease, _ := reader.ReadString('\n')
	//		YearRealease = strings.TrimSpace(BookName)
	//
	//		Book := &internal.Book{
	//			BookName:     BookName,
	//			Author:       Author,
	//			YearRealease: YearRealease,
	//		}
	//
	//		if err := db.InsertBook(DB, Book); err != nil {
	//			log.Fatal(err)
	//		}
	//
	//		fmt.Println("Insert status:succesful")
	//	}
	//
	//info, err := db.InfoForID(DB, 3)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//if err := db.DeleteFromID(DB, 7); err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("Delete is succesful!")

}
