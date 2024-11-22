TODO:

Small app microservice which pulls the 1000 posts with API and store into db.
- Generate random 1000 posts json file using www.mockaroo.com [done]
- Create a simple API to return the posts from the json with pagination [done]
- Service to fetch the posts from the API and store into db


1) Create .env from .env.example
2) docker compose up -d --build
3) `make fetch-posts` to fetch posts
