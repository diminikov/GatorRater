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