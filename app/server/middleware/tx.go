package middleware

import (
	"log"
	"net/http"

	"github.com/iqbalrestu07/datting-apps-api/common"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

// DBTransactionMiddleware sets up the database transaction middleware for Echo
func DBTransactionMiddleware(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			txHandle := db.Begin()
			if txHandle.Error != nil {
				return txHandle.Error
			}
			log.Print("beginning database transaction")

			defer func() {
				if r := recover(); r != nil {
					_ = txHandle.Rollback()
				}
			}()
			c.Set("db_trx", txHandle)

			if err := next(c); err != nil {
				err = txHandle.Rollback().Error
				common.LogErrorWithLine(err)
				log.Print("rolling back transaction due to error: ", err)
				return err
			}

			if StatusInList(c.Response().Status, []int{http.StatusOK, http.StatusCreated}) {
				log.Print("committing transactions")
				if err := txHandle.Commit().Error; err != nil {
					common.LogErrorWithLine(err)
					log.Print("trx commit error: ", err)
					return err
				}
			} else {
				log.Print("rolling back transaction due to status code: ", c.Response().Status)
				_ = txHandle.Rollback()
			}

			return nil
		}
	}
}

// StatusInList checks if a given status code is in the list
func StatusInList(status int, statusList []int) bool {
	for _, i := range statusList {
		if i == status {
			return true
		}
	}
	return false
}
