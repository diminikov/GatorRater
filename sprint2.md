Backend API documentation

POST /users: - posting the new information from the frontend and creating a new user in the backend using that information.
GET /users - Navigates to the backend and gets all the users from the database
GET /users/:username - Navigates to the backend and gets the user from the databse if that user exists. Otherwise outputs an error.
PUT /users/:username - Navigates to the backend and updates an user's information in the databse.
DELETE /users/:username - Navigates to the backend and deletes a user from the database.


Backend Unit Tests

TestPingRoute - Testing to see if we are successfully connecting to the backend.
TestMakeUser - Testing to see if we can create an user in the backend and that this user is successfully uploaded and stored in the databse.
TestGetUser - Testing to see if we can access the database to find a user with a specified username.
TestEditUser -  Testing to see if we can find a user with a specified username and succesfully edit its data.
TestDeleteUser - Testing to see if we can find a user with a specified username and succesfully remove it from the table.

Front End Additions
Added an about us page, giving a simple background to each of the team members
Added a course section using a search bar, so the user can search for desired course
Created buttons as part of the header using event binding to load different pages based on user input




Front End Cypress Tests

For the Cypress and Unit Testing, the first test essentially visits the local browser and finds the three tabs "About us", "Login" and "Find a Course to rate" and clicks on them, resulting in the user going to different pages. The second clicks on one of the courses in the list of computer science courses. The third test implements the search bar in that it finds the search bar, types the name of a course (e.g. "Introduction to Computer Organization") and through a search filter pipe, filters the list of courses until it finds the course the user is looking for and clicks on the user. Lastly, the fourth test, it goes to the login page by clicking on login on the tabs and then inputs the username and password of some account and clicks submit. The code returns "Login Successful".




Front End Work to be done next sprint:

Integrating front and back end- we were unable to do this because of a problem creating a proxy in angular. We plan to meet with a TA to resolve this issue as soon as possible.
Adding images to About us page and to the course search page
modifying the login page to allow a user to make a new account.
being able to click a class and it route to a new message board for the class.
storing messages and user id in back end.

