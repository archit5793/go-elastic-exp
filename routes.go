package main 

func initializeRoutes() {
	router.GET("/search",func(c *gin.Context){
		c.JSON(200,gin.H{"code":"OK","message":"found data"})
	})

	router.NoRoute(notfound())
}