# Golang Webapp Skeleton

This project is a quickstart to developing a golang webapp.

## Local Setup

- Install Go version 1.19
- Clone this repository to your local machine
- Copy database config file `cp database-sample.yml database.yml` and fill in your values
- Find and replace `github.com/jdonahue135/golang-web-skeleton` with your project repository
- Install `pop v6.1.1`
- run migrations `soda migrate`
- Edit `run.sh` file and add in your env variables
- Run `chmod +x run.sh` to register command
- Run `./run.sh` in a terminal to launch application

## Testing with Coverage

- First time setup: run `chmod +x coverage.sh` to register coverage command (edit default browser in command if needed)
- Run `./coverage.sh`

## Dependencies

- Built in Go version 1.19
- Uses the [chi router](https://github.com/go-chi/chi)
- Uses [alex edwards SCS](https://github.com/alexedwards/scs) session management
- Uses [nosurf](https://github.com/justinas/nosurf)

## References

- Adapted from [Trevor Sawler's bookings app](https://github.com/tsawler/bookings)
