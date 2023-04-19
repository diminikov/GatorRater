Front End:
-integrated front end with backend for signing up and logging in
-Can now sign up which adds a new user to the API.
-Can log in and will be verified that user exists in our data base.
-updated front end to use routing for navigation between pages
-updating overall css styling with bootstrap
-need to add team members pictures and bio to 'About Us' page
-need to now implement user messages with the backend
-implemented routing to individual class discussion boards but need to style and integrate with backend
-need to add logout and customization per user

Unit Tests through Cypress: 
For the Unit Testing for sprint 3 (using Cypress), I wanted to be able to demonstrate the functionality of all the different features that we have added to the project. The first test goes through how a user would sign up and make an account on our platform. Essentially, the user would go on our browser and first see the login page. Assuming they are a new user, they would click on the "New user? Click to signup!!" This would take them to the sign up page where they would type the email address and password they want to use for their account. Then they would click the "Sign up" button. The second is logging in and so, on the login page of our browser, the user would use the same email address and password (from when they signed up) to login to their account. After clicking "Login!" it would take them straight to the "Find a Course to Rate" page. The third test demonstrates the functionality of clicking on the three tabs at the top. The "Find a Course to Rate" tab takes them to a page where they can choose what class they are looking for. The "About Us" tab takes them to a page with descriptions about who we are as the creators. Lastly, the "Login" tab takes them to a page where they can login or sign up. The fourth test shows the functionality of our router. And so, whenever the user clicks anything that is not the three tabs at the tab like the "Manage" dropdown at the top right, it automatically sends them straight to the login page. The fifth test is something that we will continue to work on and this is the discussion board page. On the "Find a Course to Rate" page, the user can go to the discussion board page for the course they are looking for by clicking on the link that is directly under the course name and id. We will continue working on the discussion board page for sprint 4. Lastly, the sixth test would demonstrate the functionality of the search filter on the "Find a Course to Rate" page. And so, the user would simply type in the course name of the class they are looking for and it should filter it to be the only one showing on the page.

Back End:

TestMakeClass - Testing to see if we can create a class in the backend and that this class is successfully uploaded and stored in the databse. TestGetClass - Testing to see if we can access the database to find a lass with a specified name. TestEditClass - Testing to see if we can find a class with a specified name and succesfully edit its data. TestDeleteClass - Testing to see if we can find a class with a specified name and succesfully remove it from the table.

Backend API documentation

POST /class: - posting the new information from the frontend and creating a new class in the backend using that information. GET /classes - Navigates to the backend and gets all the classes from the database GET /class/:name - Navigates to the backend and gets the class from the database if that class exists. Otherwise outputs an error. PUT /class/:name - Navigates to the backend and updates an class's information in the databse. DELETE /class/:name - Navigates to the backend and deletes a class from the database.
