package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	game()
}

/* Structure pour stocker l'état de jeu */
type GameState struct {
	Word              string
	MaskedWord        string
	RemainingAttempts int
}

/* Fonction pour sauvegarder l'état de jeu dans save.txt */
func sauvegarde(state GameState) error {
	fichier, err := os.Create("save.txt")
	if err != nil {
		return err
	}
	defer fichier.Close()

	_, err = fichier.WriteString(fmt.Sprintf("%s\n%s\n%d\n", state.Word, state.MaskedWord, state.RemainingAttempts))
	if err != nil {
		return err
	}
	fmt.Println("Partie sauvegardée dans save.txt.")
	return nil
}

/* Fonction pour charger l'état de jeu depuis save.txt */
func chargeJeu() (GameState, error) {
	fichier, err := os.Open("save.txt")
	if err != nil {
		return GameState{}, err
	}
	defer fichier.Close()

	var state GameState
	scanner := bufio.NewScanner(fichier)

	if scanner.Scan() {
		state.Word = scanner.Text()
	}
	if scanner.Scan() {
		state.MaskedWord = scanner.Text()
	}
	if scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%d", &state.RemainingAttempts)
	}
	state.RemainingAttempts++
	return state, nil
}

/* Lit un mot aléatoire dans le fichier words.txt */
func lectureWord() (string, error) {
	fichier, err := os.Open("words.txt")
	if err != nil {
		return "", err
	}
	defer fichier.Close()

	var lines []string
	scanner := bufio.NewScanner(fichier)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if len(lines) == 0 {
		return "", fmt.Errorf("le fichier words.txt est vide")
	}
	randomIndex := rand.Intn(len(lines))
	return lines[randomIndex], nil
}

/* Cache certaines lettres du mot */
func motmaque(word string) []rune {
	runeWord := []rune(word)
	longueur := len(word)
	nbLettresMasquees := len(word)/2 - 1
	if longueur <= 3 {
		nbLettresMasquees = 1
	}

	for i := 1; i <= longueur-nbLettresMasquees; i++ {
		index := rand.Intn(len(word))
		if runeWord[index] == '_' {
			i--
			continue
		}
		runeWord[index] = '_'
	}
	return runeWord
}

/* Affiche le pendu en fonction du nombre d'essais restants */
func afficheHangman(nbrEssai int) {
	fichier, err := os.Open("hangman.txt")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier hangman.txt:", err)
		return
	}
	defer fichier.Close()

	scanner := bufio.NewScanner(fichier)
	lineStart := (9 - nbrEssai) * 8
	lineEnd := lineStart + 8
	lineNumber := 0

	for scanner.Scan() {
		if lineNumber >= lineStart && lineNumber < lineEnd {
			fmt.Println(scanner.Text())
		}
		lineNumber++
	}
}

/* Met à jour le mot masqué avec les lettres devinées */
func updateMaskedWord(motmasque []rune, originalWord string, guess string) {
	for i := 0; i < len(originalWord); i++ {
		if string(originalWord[i]) == guess {
			motmasque[i] = rune(guess[0])
		}
	}
}

var gameState GameState // on utilise la variable gameState a chaque fois que la fonction game est réutilisé pour pouvoir garder les stats du jeu (nbr essai, mot, mot a trou+progression)
var err error

/* Fonction principale du jeu */
func game() {

	gameState.Word, err = lectureWord() // Récupération d'un mot aléatoire
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}
	gameState.MaskedWord = string(motmaque(gameState.Word))
	gameState.RemainingAttempts = 10

	for gameState.RemainingAttempts > 0 {
		fmt.Println("Mot à deviner :", gameState.MaskedWord)
		afficheHangman(gameState.RemainingAttempts)

		fmt.Print("Entrez une lettre : ")
		var guess string
		fmt.Scan(&guess)

		/* Mise à jour du mot masqué si la lettre est correcte */
		motMasqueRune := []rune(gameState.MaskedWord)
		updateMaskedWord(motMasqueRune, gameState.Word, guess)
		gameState.MaskedWord = string(motMasqueRune)

		/* Vérification de la victoire */
		if gameState.MaskedWord == gameState.Word {
			fmt.Println("Félicitations ! Vous avez deviné le mot :", gameState.Word)
			os.Remove("save.txt") // Supprime la sauvegarde après la victoire
			return
		}

		/* Réduction du nombre d'essais si la lettre n'est pas trouvée */
		if !contienlettre(gameState.Word, guess) {
			gameState.RemainingAttempts--
			fmt.Println("Lettre incorrecte, il vous reste :", gameState.RemainingAttempts, "essais")
		}
		if guess == "STOP" {
			Stop()
		}
		if len(guess) > 1 && guess != "STOP" {
			gameState.RemainingAttempts -= 2
			fmt.Println("ne proposez seulent qu'une seule lettre ")

		}
	}

	if gameState.RemainingAttempts <= 0 {
		afficheHangman(gameState.RemainingAttempts)
		fmt.Println("Dommage, vous avez perdu. Le mot était :", gameState.Word)
	}
}

func Stop() {

	if err := sauvegarde(gameState); err != nil {
		fmt.Println("Erreur lors de la sauvegarde :", err)
		return
	}
	/* Option pour charger une partie sauvegardée */
	fmt.Print("Voulez-vous charger une partie sauvegardée ? (o/n): ")
	var choix string
	fmt.Scan(&choix)
	if choix == "o" {
		gameState, err = chargeJeu()
		if err != nil {
			fmt.Println("Erreur lors du chargement de la sauvegarde:", err)
			return
		}
		fmt.Println("Partie chargée depuis save.txt")
	} else {
		game()
	}
}

/* Vérifie si le mot contient une lettre donnée */
func contienlettre(word, letter string) bool {
	for _, char := range word {
		if string(char) == letter {
			return true
		}
	}
	return false
}
