Front End:


Unit Tests through Cypress: 


Backend: 
 - We have implemented creating posts and storing them in the database.
 - We have restructured the databse to more efficiently support posts.
 - The user can now post a message to leave a review for a class and can edit it if them like.
 - We used to store three different object containing a reference to the databse but instead we merged it now into one object for a more efficient backend.


Backend tests: We added new tests for the new functionality that we added this sprint which was being able to post messages.

TestMakePost: Checks to see if post has been made correctly

TestGetPostFromUser: Searching for all posts from a specific user

TestGetPostFromClass: Searching for all posts from a specific class

TestIncorrectUserPost: Testing to make sure it will return an error if a post is made from an non-existing user

TestIncorrectClassPost: Testing to make sure it will return an error if a post is made from an non-existing class 

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