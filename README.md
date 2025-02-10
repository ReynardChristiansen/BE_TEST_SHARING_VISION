
## Step to Run Backend
Clone the Repository
    
    git clone https://github.com/ReynardChristiansen/BE_TEST_SHARING_VISION.git

Create .env File

    DB_USER=avnadmin
    DB_PASSWORD=AVNS_TwOjP-9WnrwEC9-Zu9i
    DB_HOST=mysql-160eef6b-reynard-de35.h.aivencloud.com
    DB_PORT=22909
    DB_NAME=article

Install Dependencies

    go mod tidy

Migrate Database (Optional)

    go run migrate.go

Run the Server

    go run main.go

## Step to Run Frontend

https://github.com/ReynardChristiansen/FE_TEST_SHARING_VISION


## Endpoints

- Create Article (POST): http://localhost:8080/article
- Get Article (GET): http://localhost:8080/article/{limit}/{offset}
- Get ArticleById (GET): http://localhost:8080/article/{id}
- Delete ArticleById (DELETE): http://localhost:8080/article/{id}
- Update ArticleById (PUT/PATCH): http://localhost:8080/article/{id}

## Request Body

- Create Article

    To create an article, send a POST request with the following body:

        {
            "title": STRING,
            "content": STRING,
            "category": STRING,
            "status": STRING
        }

- Update ArticleById

    To update an article by id, send a PUT/PATCH request with the following body:

        {
            "title": STRING,
            "content": STRING,
            "category": STRING,
            "status": STRING
        }
