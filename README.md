# datting-apps-api

## Introduction
This is datting-apps-api, This project is create for a simple datting apps  similar like tinder or bumble with basic functionality. 
## System Proccess
- Sign up and Sign in to the App
  - Implement auth for user
  - User must be logged in to have access to the apps

- User Interaction
   - User Able to only view, swipe left (pass) and swipe right (like) 10 other dating profiles in
total (pass + like) in 1 day.
  - Same profiles can’t appear twice in the same day.
  - Implement feature to purchase premium packages that unlocks one premium feature:
    - No swipe quota for user
    - Verified label for user
  - Default apps will have default user interest in the database by seeder when initiate the project

## Tech Used
- [Golang] - https://golang.org/
- [Echo] - https://echo.labstack.com/
- [PostgreSQL] - https://www.postgresql.org/
- [Gorm] - https://gorm.io/
- [Docker] - https://www.docker.com/
- [gosec] - https://github.com/securego/gosec

## Clean Code Architecture
This project architecture is follows clean code architecture by https://github.com/bxcodec/go-clean-arch with some custom layer. The reason of using this architecture is every module is independent and not depending to specifik framework, testable, not depending on the specific database used and easy to understand.

## Get Started, How to run this project

### Clone the repository:

```bash
git clone https://github.com/iqbalrestu07/datting-apps-api.git
```

```bash
$ docker-compose up -d --build
```
### Clone the repository:

```bash
$ go test ./...
```

## API Documentation
The API documentation can found in [Link] - https://elements.getpostman.com/redirect?entityId=6069561-6f80b4fb-fa94-48c8-9f15-ba1201555461&entityType=collection

Thank you, if any question about the project we can keep in touch by email.
