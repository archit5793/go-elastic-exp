package main


func notfound(c *gin.Context){
	c.JSON(404,gin.H{"code":"PAGE_NOT_FOUND","message":"Page not found"})
}

func createIndex(c *gin.Context){
	exists, err := client.IndexExists("users").Do(context.Background())

	if err != nil{
		panic(err)
	}

	if !exists {
	mapping :=`
				{
					"settings":{
						"number_of_shards":1,
						"number_of_replicas":0
					},
					"mappings":{
						"_default_":{
							"_all":{
								"enabled":true
							}
						},
						"teacher":{
							"properties":{
								"first_name":{
									"type":"keyword"
								},
								"last_name":{
									"type":"keyword"
								},
								"institute_name":{
									"type":"text"
								},
								"subjects":{

								},
								"classes":{

								},
								"street":{
									"type":"text"
								},
								"landmark":{
									"type":"text"
								},
								"sector":{
									"type":"text"
								},
								"city":{
									"type":"text"
								},
								"location":{
									"type":"geo_point"
								}

							}
						}
					}
				}
				`
	createIndex, err := client.CreateIndex("users").Body(mapping).Do(context.Background())

	if err != nil {
		panic(err)
	}

	if !createIndex.Acknowledged {
		//No ack
	}
}

}
