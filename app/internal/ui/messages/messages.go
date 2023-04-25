package messages

const (
	InfoMessage = `
	Role game "Wakashi" - play with your friends and plunge into the atmosphere of Japanese culture
`
	LoginMessage = `
	To login you should use command:
	/login password=your_password
	You can also use our web-site to authorize, just follow the link:
	<a href="http://127.0.0.1:8080/login">Login</a>
`
	RegisterMessage = `
	To register you should use command:
	/register password=your_password role=your_role
`
	ResetMessage = `
	If you forgot your password you should use the token, that 
	was given to you after login.
	/reset_password token=your_token new_password=your_password
`
	ExitMessage = `
	To exit from session use: /exit command
`
	GreetingMessage = `
		Hi! It's a role game where u can choose your role and play with other people
		The list of roles:
			⭐ Administrator - can create entities, bind other entities with each other, create cards etc.
			⭐ Shogun - can output the data about his slaves, can create cards and bind them to Daimyo etc.
			⭐ Daimyo = can output the data about cards, create card requests
			⭐ Samurai - binds to one Daimyo and performs some useful actions
			⭐ Collector - handles daimyo card requests
		🔐 To start the game you should register then authorize 
		To register use the endpoint: /register password=.. role=..
		To login user the endpoint: /login  password=..
		Sya!👋
`
	// Admin
	AdminDoc_createEntity = `
	To create entity you should know the username, password and role.
	Example of command, that creates new shogun user: /admin_createEntity username=tgUsername password=qwerty role=shogun
	Notice the order of parameters and the username should be a username of real telegram user, otherwise u will
	not be able to get access to user data.
`
	AdminDoc_createCard = `
	To create card you should know the card number, bank unique id, owner username and cvv-code.
	Example of command, that creates new card: /admin_createCard number=4567344598433456 bank_id=1481 owner=tgUsername cvv=123
	Notice the order of parameters and the username should be a username of real telegram user.
	Also you should use one of the available bank id's:
	1481 - Sberbank
	1326 - Alpha
	2673 - Tinkoff
	3292 - Raiffeisenbank
`
	AdminDoc_bindSlave = `
	To bind slave you should know their usernames.
	To bind use command: /admin_bindSlave master_username=... slave_username=...
	For example, /admin_bindSlave master_username=Markuysa slave_username=Makaroni1234
	Notice that you can bind only HIGHER-LEVEL role with LOWER-LEVEL. For example, Shogun with Daimyo,
	Daimyo with Samurai etc.
`
	AdminDoc_bindCardToDaimyo = `
	To bind card to dimyo you should know the card number and the username of daimyo.
	To bind use command: /admin_bindCard cardNumber=... username=...
	For example, /admin_bindCard card=1234123412341234 daimyo=Markuysa
	The card number is 16-digit length integer.
`
	AdminDoc_getEntityData = `
	To get entity data you should know the username of that person.
	To form report use: /admin_entityData username=...
`

	// Shogun
	ShodunDoc_getSlavesList = `
	To get all your slaves list use command:
	/shogun_getSlavesList
`
	ShodunDoc_createCard = `
	To create card you should know the card number, bank unique id, owner username and cvv-code.
	Example of command, that creates new card: /shogun_createCard number=4567344598433456 bank_id=1481 owner=tgUsername cvv=123
	Notice the order of parameters and the username should be a username of real telegram user.
	Also you should use one of the available bank id's:
	1481 - Sberbank
	1326 - Alpha
	2673 - Tinkoff
	3292 - Raiffeisenbank
`
	ShodunDoc_getSlavesData = `
	To get the slave data you should know his/her username, for example:
	/shogun_getSlavesData username=someone
`
	// Daimyo
	DaimyoDoc_getCards = `
	To get all your cards list use command:
	/daimyo_getCards
`
	DaimyoDoc_increase = `
	To create increment card total request to collectors you 
	should know the card nubmer and type increment value.
	For example: /daimyo_increase card=1234123412341234 value=123.12
	Notice, that float values should separate with dot, not with comma
	The card number is 16-digit length integer.
`
	DaimyoDoc_getSamurai = `
	To get all your samurai list use command:
	/daimyo_getSamurai
`
	DaimyoDoc_getTurnover = `
	To get samurai's turnover you should know his/her username
	For example: /daimyo_getTurnover username=Markuysa
`
	DaimyoDoc_getCardsTotal = `
	To get all your cards total at the end of the working shift you should use:
	/daimyo_getCardsTotal 
`
	DaimyoDoc_bindShogun = `
	To bind yourself to shogun you should know his username:
	/daimyo_bindShogun username=Markuysa
`

	// Samurai
	SamuraiDoc_getTurnover = `
	To get your turnover in the working shift you should use:
	/samurai_getTurnover 
`
	SamuraiDoc_bindDaimyo = `
	To bind yourself to daimyo you should know the username of daimyo: 
	/samurai_bindDaimyo username=Markuysa
`

	// Collector
	CollectorDoc_performInc = `
	To handle daimyo increment request you should know the unique id of 
	transaction.
	To get it you can use /collector_showTranasctions method and see	
	all active transactions.
	Example: /collector_performInc id=3
	id is a positive number.
`
	CollectorDoc_showTranasctions = `
	To get it you can use /collector_showTranasctions method and see	
	all active transactions.
`
)
