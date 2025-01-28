# pokedexcli
## Overview
`pokedexcli` is a command-line tool that allows you to interact with the Pokémon API to explore locations, catch Pokémon, and manage your Pokédex.

## Installation
To install `pokedexcli`, clone the repository and build the project:
```sh
git clone https://github.com/bzelaznicki/pokedexcli.git
cd pokedexcli
go build
```

## Usage
Run the `pokedexcli` executable to start the application:
```sh
./pokedexcli
```

## Commands
- `help`: Displays a help message with available commands.
- `exit`: Exit the Pokedex.
- `map`: Lists the next 20 map locations.
- `mapb`: Lists the previous 20 map locations.
- `explore <location_name>`: Shows details about a specific location.
- `catch <pokemon_name>`: Attempts to catch a Pokémon.
- `inspect <pokemon_name>`: Inspects a Pokémon in your Pokédex.
- `pokedex`: Displays the Pokémon in your Pokédex.

## Examples
- To list the next 20 map locations:
    ```sh
    map
    ```
- To explore a specific location:
    ```sh
    explore bell-tower-2f
    ```
- To catch a Pokémon:
    ```sh
    catch Pikachu
    ```
- To inspect a Pokémon in your Pokédex:
    ```sh
    inspect Pikachu
    ```
- To display all Pokémon in your Pokédex:
    ```sh
    pokedex
    ```
