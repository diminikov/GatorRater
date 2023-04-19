Front End:
-created a data service to inject data from the back end into any component on our site.
-created a new discussion component to display each class's posts
-used class objects to store users and posts for each class.
-used get requests to load users and posts from our API for each classes discussiom page
-used variable routing so that we could load different class information on the same discussion component.
-stored the current user in our data service.
-added a logout component to switch users.
-updated styling and formatting on the course and about us page.

Unit Tests through Cypress: 
For the Unit Testing for sprint 3 (using Cypress), I wanted to be able to demonstrate the functionality of all the different features that we have added to the project. The first test goes through how a user would sign up and make an account on our platform. Essentially, the user would go on our browser and first see the login page. Assuming they are a new user, they would click on the "New user? Click to signup!!" This would take them to the sign up page where they would type the email address and password they want to use for their account. Then they would click the "Sign up" button. The second is logging in and so, on the login page of our browser, the user would use the same email address and password (from when they signed up) to login to their account. After clicking "Login!" it would take them straight to the "Find a Course to Rate" page. The third test demonstrates the functionality of clicking on the three tabs at the top. The "Find a Course to Rate" tab takes them to a page where they can choose what class they are looking for. The "About Us" tab takes them to a page with descriptions about who we are as the creators. Lastly, the "Login" tab takes them to a page where they can login or sign up. The fourth test shows the functionality of our router. And so, whenever the user clicks anything that is not the three tabs at the tab like the "Manage" dropdown at the top right, it automatically sends them straight to the login page. The fifth test is something that we will continue to work on and this is the discussion board page. On the "Find a Course to Rate" page, the user can go to the discussion board page for the course they are looking for by clicking on the link that is directly under the course name and id. We will continue working on the discussion board page for sprint 4. Lastly, the sixth test would demonstrate the functionality of the search filter on the "Find a Course to Rate" page. And so, the user would simply type in the course name of the class they are looking for and it should filter it to be the only one showing on the page.


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
