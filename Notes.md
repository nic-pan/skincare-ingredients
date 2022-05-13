CRUD Operations:
	GET /ingredient
	GET /ingredient/:id
	POST /ingredient
	PUT /ingredient/:id
	DELETE /ingredient/:id

	GET /skinTypes
	GET /skinType/:id
	POST /skinType
	PUT /skinType/:id
	DELETE /skinType/:id

[
	{"name": "Dry", "characteristics":"small pores;feels tight in the morning;dull, rough skin;prone to redness and flaking"},
	{"name": "Oily", "characteristics":"larger pores;shiny/greasy skin;may have blemishes or blackheads"},
	{"name": "Combination", "characteristics":"medium sized pores in T-zone;oily in T-zone;may have blemishes or blackheads"},
	{"name": "Dehydrated", "characteristics":"dry skin;may have rough or flakey spots but produces excessive sebum;should be treated as oily skin, despite dry appearance"},
	{"name": "Sensitive", "characteristics":"fine/larger pores;redness, itching and dryness of skin;prone to irritation"},
	{"name": "Acne-Prone", "characteristics":"usually oily or combination skin;products easily irritate;caused by sensitivity or hormonal imbalances"},
	{"name": "Normal", "characteristics":"pores barely visible;even skin tone;minimal skin sensitivities and blemishes"}
]

Separate each characteristic with a ;

Setting ingredient compatibility

    POST /combination
	
