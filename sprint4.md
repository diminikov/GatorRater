Front End:


Unit Tests through Cypress: 

Back End Tests:

TestMakePost -

TestGetPostFromUser - 

TestGetPostFromClass - 

TestIncorrectUserPost - 

TestIncorrectClassPost - 

Backend API documentation

POST /post - this request recieves a json in the format:
{
		"Username": "",
		"Class": "",
		"Body": ""
}
and will add the post into the database finding the correct user id and class id to insert into the database

GET /post/class/:classname - This does not recieve a json file, it just uses the classname parameter from the
url. It will return a json filled with all posts for that class in the form of:
{
    	"Username": "",
		"Body": ""
}

GET /post/username/:username - This does not recieve a json file, it just uses the username parameter from the
url. It will return a json filled with all posts made by that user in the form of:
{
    	"Class": "",
		"Body": ""
}