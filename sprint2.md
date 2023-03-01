Backend API documentation

POST /users: - posting the new information from the frontend and creating a new user in the backend using that information.
GET /users - Navigates to the backend and gets all the users from the database
GET /users/:username - Navigates to the backend and gets the user from the databse if that user exists. Otherwise outputs an error.
PUT /users/:username - Navigates to the backend and updates an user's information in the databse.
DELETE /users/:username - Navigates to the backend and deletes a user from the database.


BACKEND UNIT tests
TestpingRoute - Testing to see if we are successfully connecting to the backend.
TestGetuser - Testing to see if we can create an user in the backend and that this user is successfully uploaded and stored in the databse.