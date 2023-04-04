Table of Contents
Getting Started
Prerequisites
Installation
Usage
Resources
User
Recipe
Contributing
License
Getting Started
These instructions will help you get a copy of the project up and running on your local machine for development and testing purposes.

Prerequisites
Before you begin, ensure that you have the following software installed on your system:

Go (version 1.19 or higher)
Git
Installation
To set up the project on your local machine, follow these steps:

Clone the repository:
bash
Copy code
git clone https://github.com/datudou/test.git
Change to the project directory:
bash
Copy code
cd test
Install the required dependencies:
bash
Copy code
go mod download
Build the project:
bash
Copy code
go build
Run the application:
bash
Copy code
./test
Now, the server should be running on localhost:8080.


d
Usage
Once the server is up and running, you can interact with the available resources using a REST client or your browser. The following are the supported endpoints for each resource:

Resources
User
List all users

Method: GET
Endpoint: /users
Description: Retrieves a list of all users.
Create a new user

Method: POST

Endpoint: /users

Description: Creates a new user.

Payload:

```
{
  "name": "string",
  "email": "string"
}
```

Method: GET
Endpoint: /users/:id
Description: Retrieves a specific user by ID.
Update a specific user

Method: PUT

Endpoint: /users/:id

Description: Updates a specific user by ID.

Payload:

json
Copy code
{
  "name": "string",
  "email": "string"
}
Delete a specific user

Method: DELETE
Endpoint: /users/:id
Description: Deletes a specific user by ID.
Recipe
List all recipes

Method: GET
Endpoint: /recipes
Description: Retrieves a list of all recipes.
Create a new recipe

Method: POST

Endpoint: /recipes

Description: Creates a new recipe.

Payload:

```
{
  "title": "string",
  "description": "string",
  "ingredients": "string",
  "instructions": "string"
}
```
3. Get a specific recipe

Method: GET
Endpoint: /recipes/:id
Description: Retrieves a specific recipe by ID.
Update a specific recipe

Method: PUT

Endpoint: /recipes/:id

Description: Updates a specific recipe by ID.

Payload:

```
{
  "title": "string",
  "description": "string",
  "ingredients": "string",
  "instructions": "string"
}
```
Delete a specific recipe

Method: DELETE
Endpoint: /recipes/:id
Description: Deletes a specific recipe by ID.
Contributing
Contributions to this project are welcome! To contribute, please follow these steps:

Fork the repository on GitHub.
Create a new branch for your feature or bugfix.
Commit your changes to the new branch.
Push your changes to your forked repository.
Create a pull request, describing the changes you made and their purpose.
Before submitting your pull request, please make sure your code adheres to the project's coding style and that you've tested your changes.

License
This project is licensed under the MIT License. See the LICENSE file for more details.
