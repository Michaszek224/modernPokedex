package database

import (
	"database/sql"
	"fmt"
)

func SeedData(db *sql.DB) error {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM pokemons").Scan(&count)
	if err != nil {
		return fmt.Errorf("error checking existing data: %v", err)
	}

	if count > 0 {
		fmt.Printf("Database already contains %d pokemon, skipping seed\n", count)
		return nil
	}

	fmt.Println("Seeding database with initial Pokemon data...")

	pokemons := []Pokemon{
		{Name: "Bulbasaur", Type1: "Grass", Type2: "Poison"},
		{Name: "Ivysaur", Type1: "Grass", Type2: "Poison"},
		{Name: "Venusaur", Type1: "Grass", Type2: "Poison"},
		{Name: "Charmander", Type1: "Fire", Type2: ""},
		{Name: "Charmeleon", Type1: "Fire", Type2: ""},
		{Name: "Charizard", Type1: "Fire", Type2: "Flying"},
		{Name: "Squirtle", Type1: "Water", Type2: ""},
		{Name: "Wartortle", Type1: "Water", Type2: ""},
		{Name: "Blastoise", Type1: "Water", Type2: ""},
		{Name: "Caterpie", Type1: "Bug", Type2: ""},
		{Name: "Metapod", Type1: "Bug", Type2: ""},
		{Name: "Butterfree", Type1: "Bug", Type2: "Flying"},
		{Name: "Weedle", Type1: "Bug", Type2: "Poison"},
		{Name: "Kakuna", Type1: "Bug", Type2: "Poison"},
		{Name: "Beedrill", Type1: "Bug", Type2: "Poison"},
		{Name: "Pidgey", Type1: "Normal", Type2: "Flying"},
		{Name: "Pidgeotto", Type1: "Normal", Type2: "Flying"},
		{Name: "Pidgeot", Type1: "Normal", Type2: "Flying"},
		{Name: "Rattata", Type1: "Normal", Type2: ""},
		{Name: "Raticate", Type1: "Normal", Type2: ""},
		{Name: "Spearow", Type1: "Normal", Type2: "Flying"},
		{Name: "Fearow", Type1: "Normal", Type2: "Flying"},
		{Name: "Ekans", Type1: "Poison", Type2: ""},
		{Name: "Arbok", Type1: "Poison", Type2: ""},
		{Name: "Pikachu", Type1: "Electric", Type2: ""},
		{Name: "Raichu", Type1: "Electric", Type2: ""},
		{Name: "Sandshrew", Type1: "Ground", Type2: ""},
		{Name: "Sandslash", Type1: "Ground", Type2: ""},
		{Name: "Nidoran♀", Type1: "Poison", Type2: ""},
		{Name: "Nidorina", Type1: "Poison", Type2: ""},
		{Name: "Nidoqueen", Type1: "Poison", Type2: "Ground"},
		{Name: "Nidoran♂", Type1: "Poison", Type2: ""},
		{Name: "Nidorino", Type1: "Poison", Type2: ""},
		{Name: "Nidoking", Type1: "Poison", Type2: "Ground"},
		{Name: "Clefairy", Type1: "Fairy", Type2: ""},
		{Name: "Clefable", Type1: "Fairy", Type2: ""},
		{Name: "Vulpix", Type1: "Fire", Type2: ""},
		{Name: "Ninetales", Type1: "Fire", Type2: ""},
		{Name: "Jigglypuff", Type1: "Normal", Type2: "Fairy"},
		{Name: "Wigglytuff", Type1: "Normal", Type2: "Fairy"},
		{Name: "Zubat", Type1: "Poison", Type2: "Flying"},
		{Name: "Golbat", Type1: "Poison", Type2: "Flying"},
		{Name: "Oddish", Type1: "Grass", Type2: "Poison"},
		{Name: "Gloom", Type1: "Grass", Type2: "Poison"},
		{Name: "Vileplume", Type1: "Grass", Type2: "Poison"},
		{Name: "Paras", Type1: "Bug", Type2: "Grass"},
		{Name: "Parasect", Type1: "Bug", Type2: "Grass"},
		{Name: "Venonat", Type1: "Bug", Type2: "Poison"},
		{Name: "Venomoth", Type1: "Bug", Type2: "Poison"},
		{Name: "Diglett", Type1: "Ground", Type2: ""},
		{Name: "Dugtrio", Type1: "Ground", Type2: ""},
		{Name: "Meowth", Type1: "Normal", Type2: ""},
		{Name: "Persian", Type1: "Normal", Type2: ""},
		{Name: "Psyduck", Type1: "Water", Type2: ""},
		{Name: "Golduck", Type1: "Water", Type2: ""},
		{Name: "Mankey", Type1: "Fighting", Type2: ""},
		{Name: "Primeape", Type1: "Fighting", Type2: ""},
		{Name: "Growlithe", Type1: "Fire", Type2: ""},
		{Name: "Arcanine", Type1: "Fire", Type2: ""},
		{Name: "Poliwag", Type1: "Water", Type2: ""},
		{Name: "Poliwhirl", Type1: "Water", Type2: ""},
		{Name: "Poliwrath", Type1: "Water", Type2: "Fighting"},
		{Name: "Abra", Type1: "Psychic", Type2: ""},
		{Name: "Kadabra", Type1: "Psychic", Type2: ""},
		{Name: "Alakazam", Type1: "Psychic", Type2: ""},
		{Name: "Machop", Type1: "Fighting", Type2: ""},
		{Name: "Machoke", Type1: "Fighting", Type2: ""},
		{Name: "Machamp", Type1: "Fighting", Type2: ""},
		{Name: "Bellsprout", Type1: "Grass", Type2: "Poison"},
		{Name: "Weepinbell", Type1: "Grass", Type2: "Poison"},
		{Name: "Victreebel", Type1: "Grass", Type2: "Poison"},
		{Name: "Tentacool", Type1: "Water", Type2: "Poison"},
		{Name: "Tentacruel", Type1: "Water", Type2: "Poison"},
		{Name: "Geodude", Type1: "Rock", Type2: "Ground"},
		{Name: "Graveler", Type1: "Rock", Type2: "Ground"},
		{Name: "Golem", Type1: "Rock", Type2: "Ground"},
		{Name: "Ponyta", Type1: "Fire", Type2: ""},
		{Name: "Rapidash", Type1: "Fire", Type2: ""},
		{Name: "Slowpoke", Type1: "Water", Type2: "Psychic"},
		{Name: "Slowbro", Type1: "Water", Type2: "Psychic"},
		{Name: "Magnemite", Type1: "Electric", Type2: "Steel"},
		{Name: "Magneton", Type1: "Electric", Type2: "Steel"},
		{Name: "Farfetchd", Type1: "Normal", Type2: "Flying"},
		{Name: "Doduo", Type1: "Normal", Type2: "Flying"},
		{Name: "Dodrio", Type1: "Normal", Type2: "Flying"},
		{Name: "Seel", Type1: "Water", Type2: ""},
		{Name: "Dewgong", Type1: "Water", Type2: "Ice"},
		{Name: "Grimer", Type1: "Poison", Type2: ""},
		{Name: "Muk", Type1: "Poison", Type2: ""},
		{Name: "Shellder", Type1: "Water", Type2: ""},
		{Name: "Cloyster", Type1: "Water", Type2: "Ice"},
		{Name: "Gastly", Type1: "Ghost", Type2: "Poison"},
		{Name: "Haunter", Type1: "Ghost", Type2: "Poison"},
		{Name: "Gengar", Type1: "Ghost", Type2: "Poison"},
		{Name: "Onix", Type1: "Rock", Type2: "Ground"},
		{Name: "Drowzee", Type1: "Psychic", Type2: ""},
		{Name: "Hypno", Type1: "Psychic", Type2: ""},
		{Name: "Krabby", Type1: "Water", Type2: ""},
		{Name: "Kingler", Type1: "Water", Type2: ""},
		{Name: "Voltorb", Type1: "Electric", Type2: ""},
		{Name: "Electrode", Type1: "Electric", Type2: ""},
		{Name: "Exeggcute", Type1: "Grass", Type2: "Psychic"},
		{Name: "Exeggutor", Type1: "Grass", Type2: "Psychic"},
		{Name: "Cubone", Type1: "Ground", Type2: ""},
		{Name: "Marowak", Type1: "Ground", Type2: ""},
		{Name: "Hitmonlee", Type1: "Fighting", Type2: ""},
		{Name: "Hitmonchan", Type1: "Fighting", Type2: ""},
		{Name: "Lickitung", Type1: "Normal", Type2: ""},
		{Name: "Koffing", Type1: "Poison", Type2: ""},
		{Name: "Weezing", Type1: "Poison", Type2: ""},
		{Name: "Rhyhorn", Type1: "Ground", Type2: "Rock"},
		{Name: "Rhydon", Type1: "Ground", Type2: "Rock"},
		{Name: "Chansey", Type1: "Normal", Type2: ""},
		{Name: "Tangela", Type1: "Grass", Type2: ""},
		{Name: "Kangaskhan", Type1: "Normal", Type2: ""},
		{Name: "Horsea", Type1: "Water", Type2: ""},
		{Name: "Seadra", Type1: "Water", Type2: ""},
		{Name: "Goldeen", Type1: "Water", Type2: ""},
		{Name: "Seaking", Type1: "Water", Type2: ""},
		{Name: "Staryu", Type1: "Water", Type2: ""},
		{Name: "Starmie", Type1: "Water", Type2: "Psychic"},
		{Name: "Mr. Mime", Type1: "Psychic", Type2: "Fairy"},
		{Name: "Scyther", Type1: "Bug", Type2: "Flying"},
		{Name: "Jynx", Type1: "Ice", Type2: "Psychic"},
		{Name: "Electabuzz", Type1: "Electric", Type2: ""},
		{Name: "Magmar", Type1: "Fire", Type2: ""},
		{Name: "Pinsir", Type1: "Bug", Type2: ""},
		{Name: "Tauros", Type1: "Normal", Type2: ""},
		{Name: "Magikarp", Type1: "Water", Type2: ""},
		{Name: "Gyarados", Type1: "Water", Type2: "Flying"},
		{Name: "Lapras", Type1: "Water", Type2: "Ice"},
		{Name: "Ditto", Type1: "Normal", Type2: ""},
		{Name: "Eevee", Type1: "Normal", Type2: ""},
		{Name: "Vaporeon", Type1: "Water", Type2: ""},
		{Name: "Jolteon", Type1: "Electric", Type2: ""},
		{Name: "Flareon", Type1: "Fire", Type2: ""},
		{Name: "Porygon", Type1: "Normal", Type2: ""},
		{Name: "Omanyte", Type1: "Rock", Type2: "Water"},
		{Name: "Omastar", Type1: "Rock", Type2: "Water"},
		{Name: "Kabuto", Type1: "Rock", Type2: "Water"},
		{Name: "Kabutops", Type1: "Rock", Type2: "Water"},
		{Name: "Aerodactyl", Type1: "Rock", Type2: "Flying"},
		{Name: "Snorlax", Type1: "Normal", Type2: ""},
		{Name: "Articuno", Type1: "Ice", Type2: "Flying"},
		{Name: "Zapdos", Type1: "Electric", Type2: "Flying"},
		{Name: "Moltres", Type1: "Fire", Type2: "Flying"},
		{Name: "Dratini", Type1: "Dragon", Type2: ""},
		{Name: "Dragonair", Type1: "Dragon", Type2: ""},
		{Name: "Dragonite", Type1: "Dragon", Type2: "Flying"},
		{Name: "Mewtwo", Type1: "Psychic", Type2: ""},
		{Name: "Mew", Type1: "Psychic", Type2: ""},
	}

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("error starting transaction: %v", err)
	}
	defer func() {
		if err := tx.Rollback(); err != nil && err != sql.ErrTxDone {
			fmt.Printf("rollback failed: %v\n", err)
		}
	}()

	stmt, err := tx.Prepare("INSERT INTO pokemons (name, type1, type2) VALUES ($1, $2, $3)")
	if err != nil {
		return fmt.Errorf("error preparing statement: %v", err)
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			fmt.Printf("error closing statement: %v\n", err)
		}
	}()

	for _, pokemon := range pokemons {
		_, err = stmt.Exec(pokemon.Name, pokemon.Type1, pokemon.Type2)
		if err != nil {
			return fmt.Errorf("error inserting pokemon %s: %v", pokemon.Name, err)
		}
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %v", err)
	}

	fmt.Printf("Successfully seeded database with %d pokemon!\n", len(pokemons))
	return nil
}
