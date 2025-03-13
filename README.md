# Blog Aggregator

A simple command-line application for aggregating blog feeds. This application allows users to register, log in, add feeds, follow feeds, and browse posts from their followed feeds.

## Features

- **User Registration and Login**: Create and manage user accounts.
- **Feed Management**: Add and follow blog feeds.
- **Post Browsing**: View posts from followed feeds.
- **Scheduled Aggregation**: Collect feeds at specified intervals.
- **Database Reset**: Clear all data and start fresh.

## Technologies Used

- **PostgreSQL**: Database for storing blog data.
- **Goose**: Database migration tool ([GitHub](https://github.com/pressly/goose)).
- **sqlc**: Compile-time query generation tool for Go ([GitHub](https://github.com/sqlc-dev/sqlc)).

## Setup

### 1. Run the Database

Ensure you have Docker installed, then start the PostgreSQL database with:

```sh
docker compose up -d
```

### 2. Apply Migrations

Run the following command to apply database migrations:

```sh
goose -dir sql/schema postgres "postgres://postgres:postgres@localhost:5432/gator" up
```

### 3. Build the Project

Compile the application:

```sh
go build 
```

### 4. Update Configuration

Update the connection string in the `.gatorconfig.json` file if necessary:

```json
{
  "db_url": "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}
```

### Usage

After starting the application, you can use the following commands:

- **Register a new user**:
  ```bash
  ./blog-aggregator register <name>
  ```

- **Log in as a user**:
  ```bash
  ./blog-aggregator login <name>
  ```

- **Reset the database**:
  ```bash
  ./blog-aggregator reset
  ```

- **List all users**:
  ```bash
  ./blog-aggregator users
  ```

- **Add a new feed**:
  ```bash
  ./blog-aggregator addfeed <name> <url>
  ```

- **List all feeds**:
  ```bash
  ./blog-aggregator feeds
  ```

- **Follow a feed**:
  ```bash
  ./blog-aggregator follow <url>
  ```

- **List followed feeds**:
  ```bash
  ./blog-aggregator following
  ```

- **Unfollow a feed**:
  ```bash
  ./blog-aggregator unfollow <url>
  ```

- **Browse posts from followed feeds**:
  ```bash
  ./blog-aggregator browse [limit]
  ```

- **Aggregate feeds at specified intervals**:
  ```bash
  ./blog-aggregator agg <time_between_reqs>
  ```

### Feed Examples

Here’s a list of RSS feed URLs you can use for testing:

1. **TechCrunch**: `https://feeds.feedburner.com/TechCrunch/`
2. **Wired**: `https://www.wired.com/feed/rss`
3. **Hacker News**: `https://news.ycombinator.com/rss`
4. **BBC News**: `http://feeds.bbci.co.uk/news/rss.xml`
5. **Reuters**: `http://feeds.reuters.com/reuters/topNews`
6. **NPR News**: `https://www.npr.org/rss/rss.php?id=1001`
7. **Lifehacker**: `https://lifehacker.com/rss`
8. **The Verge**: `https://www.theverge.com/rss/index.xml`
9. **Mental Floss**: `https://www.mentalfloss.com/rss`
10. **ESPN**: `https://www.espn.com/espn/rss/news`
11. **Bleacher Report**: `https://bleacherreport.com/articles/feed`
12. **Serious Eats**: `https://www.seriouseats.com/rss`

You can use any of these links to test the blog aggregator application!

### Configuration

The application uses a configuration file named `.gatorconfig.json` to store the database URL and the current user name. Make sure to update this file with your database connection details.

### Database Schema

The application uses the following tables:

- **users**: Stores user information.
- **feeds**: Stores feed information.
- **feed_follows**: Stores the relationship between users and feeds.
- **posts**: Stores posts fetched from feeds.

### Acknowledgments

- [SQLC](https://sqlc.dev/) for generating type-safe Go code from SQL queries.
- [PostgreSQL](https://www.postgresql.org/) for the database.
- [Go](https://golang.org/) for the programming language.

