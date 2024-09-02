package handlers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    logrus "github.com/sirupsen/logrus"
)

func SetupRoutes(router *gin.Engine, facade Facade, logger *logrus.Entry) {
    	router.GET("/getFNOPosition/:UCCID", func(c *gin.Context) {
		UCCID := c.Param("UCCID")

		// Log incoming request
		logger.WithFields(logrus.Fields{
			"endpoint": "/getFNOPosition",
			"UCCID":    UCCID,
		}).Info("Received request")

		resp, err := facade.GetFNOPosition(UCCID)
		if err != nil {
			logger.WithFields(logrus.Fields{
				"endpoint": "/getFNOPosition",
				"UCCID":    UCCID,
				"error":    err.Error(),
			}).Error("Error fetching FNO positions")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch FNO positions"})
			return
		}

		// Log successful response
		logger.WithFields(logrus.Fields{
			"endpoint": "/getFNOPosition",
			"UCCID":    UCCID,
		}).Info("Successfully fetched FNO positions")

		c.JSON(http.StatusOK, gin.H{"positions": resp})
	})

    	router.GET("/getOrderDetails/:OrderID", func(c *gin.Context) {
		OrderID := c.Param("OrderID")

		// Log incoming request
		logger.WithFields(logrus.Fields{
			"endpoint": "/getOrderDetails",
			"OrderID":  OrderID,
		}).Info("Received request")

		resp, err := facade.GetOrderDetails(OrderID)
		if err != nil {
			logger.WithFields(logrus.Fields{
				"endpoint": "/getOrderDetails",
				"OrderID":  OrderID,
				"error":    err.Error(),
			}).Error("Error fetching order details")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch order details"})
			return
		}

		// Log successful response
		logger.WithFields(logrus.Fields{
			"endpoint": "/getOrderDetails",
			"OrderID":  OrderID,
		}).Info("Successfully fetched order details")

		c.JSON(http.StatusOK, gin.H{"orderDetails": resp})
	})


    router.GET("/GetCommodityPositions/:account", func(c *gin.Context) {
        account := c.Param("account")
        resp, err := facade.GetCommodityPositions(account)
        if err != nil {
            logger.Errorf("Error fetching commodity positions: %v", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, resp)
    })

	router.GET("/GetMtfPosition/:account",func(c *gin.Context) {
		account := c.Param("account")
		resp,err := facade.GetMtfPositions(account)
		if err !=nil{
			logger.Error("Error fetching MTF Positions: %s",err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK,resp)
	})
}


// package handlers

// import (
// 	//"context"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	log "github.com/sirupsen/logrus"
// )

// func SetupRoutes(router *gin.Engine, facade Facade, logger *log.Entry) {
// 	router.GET("/getFNOPosition/:UCCID", func(c *gin.Context) {
// 		UCCID := c.Param("UCCID")

// 		// Log incoming request
// 		logger.WithFields(log.Fields{
// 			"endpoint": "/getFNOPosition",
// 			"UCCID":    UCCID,
// 		}).Info("Received request")

// 		resp, err := facade.GetFNOPosition(UCCID)
// 		if err != nil {
// 			logger.WithFields(log.Fields{
// 				"endpoint": "/getFNOPosition",
// 				"UCCID":    UCCID,
// 				"error":    err.Error(),
// 			}).Error("Error fetching FNO positions")
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch FNO positions"})
// 			return
// 		}

// 		// Log successful response
// 		logger.WithFields(log.Fields{
// 			"endpoint": "/getFNOPosition",
// 			"UCCID":    UCCID,
// 		}).Info("Successfully fetched FNO positions")

// 		c.JSON(http.StatusOK, gin.H{"positions": resp})
// 	})

// 	router.GET("/getOrderDetails/:OrderID", func(c *gin.Context) {
// 		OrderID := c.Param("OrderID")

// 		// Log incoming request
// 		logger.WithFields(log.Fields{
// 			"endpoint": "/getOrderDetails",
// 			"OrderID":  OrderID,
// 		}).Info("Received request")

// 		resp, err := facade.GetOrderDetails(OrderID)
// 		if err != nil {
// 			logger.WithFields(log.Fields{
// 				"endpoint": "/getOrderDetails",
// 				"OrderID":  OrderID,
// 				"error":    err.Error(),
// 			}).Error("Error fetching order details")
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch order details"})
// 			return
// 		}

// 		// Log successful response
// 		logger.WithFields(log.Fields{
// 			"endpoint": "/getOrderDetails",
// 			"OrderID":  OrderID,
// 		}).Info("Successfully fetched order details")

// 		c.JSON(http.StatusOK, gin.H{"orderDetails": resp})
// 	})
// }