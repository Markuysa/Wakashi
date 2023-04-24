package messages

const (
	GreetingMessage = `
		Hi! It's a role game where u can choose your role and play with other people
		The list of roles:
			⭐ Administrator - can create entities, bind other entities with each other, create cards etc.
			⭐ Shogun - can output the data about his slaves, can create cards and bind them to Daimyo etc.
			⭐ Daimyo = can output the data about cards, create card requests
			⭐ Samurai - binds to one Daimyo and performs some useful actions
			⭐ Collector - handles daimyo card requests
		🔐 To start the game you should register then authorize 
		To register use the endpoint: /register username=.. password=.. role=..
		To login user the endpoint: /login username=.. password=..
		Sya!👋
`
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
	To bind card to dimyo you should know their usernames.
	To bind use command: /admin_bindSlave master_username=... slave_username=...
	For example, /admin_bindSlave master_username=Markuysa slave_username=Makaroni1234
	Notice that you can bind only HIGHER-LEVEL role with LOWER-LEVEL. For example, Shogun with Daimyo,
	Daimyo with Samurai etc.
`
)
