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
)
