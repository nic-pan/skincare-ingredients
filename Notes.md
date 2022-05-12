CRUD Operations:
	GET /ingredient
	GET /ingredient/:id
	POST /ingredient
	PUT /ingredient
	DELETE /ingredient

	GET /skinTypes
	GET /skinType/:id
	POST /skinType
	PUT /skinType
	DELETE /skinType

var skinTypes = []models.SkinType{
		{ID: 1, Name: "Dry"},
		{ID: 2, Name: "Oily"},
		{ID: 3, Name: "Combination"},
		{ID: 4, Name: "Dehydrated"},
		{ID: 5, Name: "Sensitive"},
		{ID: 6, Name: "Acne-Prone"},
		{ID: 7, Name: "Normal"},
	}
	
Setting ingredient compatibility

    POST /
