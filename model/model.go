package model

import (
    // "os"
    // "log"
    //"github.com/jinzhu/gorm"
    "gorm.io/gorm"
    // "io/ioutil"
    // "time"
)

var (
    db              *gorm.DB
    // OEE_SQL         string
    // LOCATION        *time.Location
)

// func init() {
//     // set credential path
//     pwd, err := os.Getwd()
//     if err != nil {
//         log.Println("Error getting directory: %v\n", err)
//         return
//     }

//     // get oee sql from file
//     oeeSqlByte, err := ioutil.ReadFile(pwd+"/sql-query/oee.sql")
//     if err != nil {
//         log.Println("Cannot load oee sql.")
//         return
//     }

//     OEE_SQL = string(oeeSqlByte)

//     // location
//     loc, err := time.LoadLocation("Asia/Taipei")
//     if err != nil {
//         log.Println("Cannot load timezone location.")
//         return
//     }

//     LOCATION = loc
// }

func Init(_db *gorm.DB) {
    db = _db
} // Init()

