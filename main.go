package main

import (
    "phone-book-api/routes"
    "phone-book-api/database"
    "phone-book-api/models"
)

func main() {
   // db connection
   db, err := database.ConnectDB()
   if err != nil {
        panic("Failed to connect to the database")
       return
   }

   // db migration
   if err := db.AutoMigrate(&models.Contact{}); err != nil {
       panic("Failed to migrate to the database")
       return
   }

   // running on default port 8080
   r := routes.SetupRouter(db)
   r.Run()
}
