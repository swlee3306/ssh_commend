package dblinker

import (
	"log"
	"ssh_commend/internal/sysenv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error){
	var err error
	db, err := gorm.Open(mysql.Open(sysenv.Database.Dsn), &gorm.Config{})
	if err != nil {
			log.Printf("failed to connect database: %v", err)
			return nil, err
	}
	return db, nil
}

func Dbconnection(db *gorm.DB) ([]*sysenv.VmEnv, error){
	fnc := "dbconnection"
	// 공급자 로드
	if Vmlist, err := LoadModule(db); err != nil {
		log.Printf("%s: main_LoadPrvId 실패: %s", fnc, err.Error())
		return nil, err
	} else {
		for i := range Vmlist {
			// 데이터베이스 로드
			if err = LoadDb(db, Vmlist[i]); err != nil {
				log.Printf("%s: main_LoadDb 실패: %s", fnc, err.Error())
				return nil, err
			}
		}
		return Vmlist, nil
	}
}
