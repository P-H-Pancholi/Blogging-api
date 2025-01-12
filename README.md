# Blogging-api
RESTful API for a Blogging Platform

## API Endpoints

| Endpoint       | HTTP Method | Description                |
|----------------|-------------|----------------------------|
| `/posts`       | GET         | Fetch all blog posts       |
| `/posts/{id}`  | GET         | Fetch a single blog post   |
| `/posts`       | POST        | Create a new blog post     |
| `/posts/{id}`  | PUT         | Update an existing blog post |
| `/posts/{id}`  | DELETE      | Delete a blog post         |
| `/users`       | POST        | Register a new user        |
| `/auth/login`  | POST        | Login to get a token       |