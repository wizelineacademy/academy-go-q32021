# Golang Bootcamp course application

Instructions regarding the challenge on GOLANGBOOTCAMP.md file. This file is regarding the application itself (how to run it, tests, etc.)

## Technologies

This project runs with Go. In order to run the project we need to install [Go](https://golang.org/doc/install)

## Run the project

We need to go into the root of the project and run

    go run main.go

## Functionality

The project contains one endpoint as of now _/pokemons/{pokemon_id}_

We can hit the endpoint with the following link http://localhost:8000/pokemons/155

### Pokemons endpoint

Example when pinging the previous link

    {
        "id": 155,
        "name": "Cyndaquil",
        "type_1": "Fire",
        "type_2": "",
        "total_points": 309,
        "hp": 39,
        "attack": 52,
        "defense": 43,
        "speed_attack": 60,
        "speed_defense": 50,
        "speed": 65,
        "generation": 2,
        "legendary": "False"
    }
