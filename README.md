# PinBoard Display

## Functionalities

### Editor

Editor is the default screen. Every new editor will have a unique identified at
the end for identifying a new board after saving its title.

URLs:

* http://localhost:8080/#/edit/?boardId=5aaab0321def2d41da352ee9

### Kiosk

Kiosk mode is for displaying the post. There can be more than one post in a
kiosk board.

* http://localhost:8080/#/?boardId=5aaab0321def2d41da352ee9

It has functionality of refreshing board automatically at specific interval.

## Architecture

### Storage

* Long term storage MongoDB

### Server

* Golang Gorilla Mux


## Future Tasks

- [ ] Document invalid responses
- [ ] Export API Documentation from Swagger JSON/YAML
- [ ] Add test cases
- [ ] Login and Authentication

## Notes

* This project is still under development.

## API Docs

API Docs uses [sourcey/spectacle](https://github.com/sourcey/spectacle) for generating static documentation from [swagger](https://swagger.io) json

API Docs are available [here](docs/api-docs).
