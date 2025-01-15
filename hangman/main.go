package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
)

type GameState struct {
	Word              string
	MaskedWord        []string
	RemainingAttempts int
	Ascii             string
}

func main() {
	var statusjeu GameState
	var err error
	lutil := []string{}
	position := gettxt("hangman")
	lettreascii := []string{}
	d := 0
	a := 0
	ascii := ' '
	if len(os.Args) > 1 {
		if os.Args[1] == "save" {
			statusjeu, err = chargeJeu()
			if err != nil {
				fmt.Print("Erreur avec la sauvegarde")
				return
			}
			d++
			a++
			if statusjeu.Ascii == "maj" {
				lettreascii = gettxt("maj")
				ascii = 'M'
			} else if statusjeu.Ascii == "min" {
				lettreascii = gettxt("min")
				ascii = 'm'
			} else {
				ascii = 'n'
			}
		} else {
			fmt.Print("Trop d'arguments ou arguments invalide !")
			return
		}
	}
	difficulté := ' '
	for d == 0 {
		fmt.Print("Choisir une difficulté :")
		fmt.Scanf("%c\n", &difficulté)
		if difficulté == 'f' {
			d++
			statusjeu.Word = choimot("facile.txt")
			statusjeu.MaskedWord = motcache(statusjeu.Word)
			nbrand := (len(statusjeu.Word) / 2) - 1
			for nbrand != 0 {
				n := rand.IntN(len(statusjeu.Word) - 1)
				if statusjeu.MaskedWord[n] == "_" {
					statusjeu.MaskedWord[n] = ToUpper(string(statusjeu.Word[n]))
					nbrand--
				}
			}
			statusjeu.RemainingAttempts = 10
		} else if difficulté == 'm' {
			d++
			statusjeu.Word = choimot("moyen.txt")
			statusjeu.MaskedWord = motcache(statusjeu.Word)
			statusjeu.MaskedWord[0] = ToUpper(string(statusjeu.Word[0]))
			statusjeu.RemainingAttempts = 8
		} else if difficulté == 'd' {
			d++
			statusjeu.Word = choimot("difficile.txt")
			statusjeu.MaskedWord = motcache(statusjeu.Word)
			statusjeu.RemainingAttempts = 5
		} else {
			fmt.Print("Caractère invalide\n")
		}
	}
	for a == 0 {
		fmt.Print("Jouer en ascii ? : ")
		fmt.Scanf("%c\n", &ascii)
		if ascii == 'n' {
			a++
		} else if ascii == 'y' {
			a++
			for a == 1 {
				fmt.Print("Majuscule ou minuscule ? : ")
				fmt.Scanf("%c\n", &ascii)
				if ascii == 'M' {
					a++
					lettreascii = gettxt("maj")
					statusjeu.Ascii = "maj"
				} else if ascii == 'm' {
					a = 3
					lettreascii = gettxt("min")
					statusjeu.Ascii = "min"
				} else {
					fmt.Print("Caractère invalide\n")
				}
			}
		} else {
			fmt.Print("Caractère invalide\n")
		}
	}

	fmt.Printf("\n \n Tu as %d essais pour trouver le bon mot\n", statusjeu.RemainingAttempts)
	fmt.Print("\t \tBonne chance\n \n")
	affichemot(statusjeu.MaskedWord, lettreascii, ascii)
	for statusjeu.RemainingAttempts != 0 {
		if MotFini(statusjeu.MaskedWord) {
			break
		}
		danslemot := 0
		l := ""
		fmt.Print("\n Choisir une lettre :")
		fmt.Scanf("%s\n", &l)
		if !simplelettre(l) {
			fmt.Print("Caractère invalide\n")
			continue
		}
		if l == "STOP" {
			sauvegarde(statusjeu)
			return
		}
		if len(l) > 1 {
			if l == statusjeu.Word {
				break
			} else {
				if statusjeu.RemainingAttempts == 1 {
					statusjeu.RemainingAttempts--
				} else {
					statusjeu.RemainingAttempts -= 2
				}
				if statusjeu.RemainingAttempts != 0 {
					fmt.Printf("Il te reste %d essais pour trouver le bon mot\n", statusjeu.RemainingAttempts)
					fmt.Print(string(position[10-statusjeu.RemainingAttempts-1]) + "\n")
					affichemot(statusjeu.MaskedWord, lettreascii, ascii)
					continue
				}
			}
		}
		if InTab(lutil, l) {
			fmt.Print("Lettre déjà utiliser réessayer\n")
			continue
		}
		lutil = append(lutil, l)
		for i := 0; i < len(statusjeu.Word); i++ {
			if string(statusjeu.Word[i]) == l && string(statusjeu.MaskedWord[i]) == "_" {
				statusjeu.MaskedWord[i] = ToUpper(l)
				danslemot++
			}
		}
		if !InTab(statusjeu.MaskedWord, "_") {
			continue
		} else if danslemot == 0 {
			statusjeu.RemainingAttempts--
		}
		if statusjeu.RemainingAttempts == 0 {
			continue
		} else {
			if statusjeu.RemainingAttempts != 10 {
				fmt.Print(string(position[10-statusjeu.RemainingAttempts-1]) + "\n")
			}
			fmt.Printf("Il te reste %d essais pour trouver le bon mot\n", statusjeu.RemainingAttempts)
			affichemot(statusjeu.MaskedWord, lettreascii, ascii)
		}
	}
	if statusjeu.RemainingAttempts == 0 {
		fmt.Print(string(position[9]))
		fmt.Print("Perdu !! Le mot était\n")
		for i := 0; i < len(statusjeu.Word); i++ {
			statusjeu.MaskedWord[i] = ToUpper(string(statusjeu.Word[i]))
		}
		affichemot(statusjeu.MaskedWord, lettreascii, ascii)
	} else {
		fmt.Print("bravo vous avez trouvez le mot cacher\n")
		for i := 0; i < len(statusjeu.Word); i++ {
			statusjeu.MaskedWord[i] = ToUpper(string(statusjeu.Word[i]))
		}
		affichemot(statusjeu.MaskedWord, lettreascii, ascii)
	}
}

func choimot(fich string) string {
	fichier, err := os.Open(fich)
	if err != nil {
		fmt.Print(err)
	}
	fileScanner := bufio.NewScanner(fichier)
	fileScanner.Split(bufio.ScanLines)
	mots := []string{}
	for fileScanner.Scan() {
		mots = append(mots, fileScanner.Text())
	}
	fichier.Close()
	mot := mots[rand.IntN(len(mots)-1)]
	return mot
}

func motcache(mot string) []string {
	motcacher := []string{}
	for i := 0; i < len(mot); i++ {
		motcacher = append(motcacher, "_")
	}
	return motcacher
}

func ToUpper(s string) string {
	h := []rune(s)
	result := ""
	for i := 0; i < len(h); i++ {
		if (h[i] >= 'a') && (h[i] <= 'z') {
			h[i] = h[i] - 32
		}
		result += string(h[i])
	}
	return result
}

func MotFini(tab []string) bool {
	for _, i := range tab {
		if i == "_" {
			return false
		}
	}
	return true
}

func InTab(tab []string, lettre string) bool {
	for _, i := range tab {
		if i == lettre {
			return true
		}
	}
	return false
}

func affichemot(mot []string, tabascii []string, ascii rune) {
	if ascii == 'n' {
		affichemotnormal(mot)
	} else {
		affichasciimot(tabascii, mot, ascii)
	}
}

func affichemotnormal(tab []string) {
	affiche := ""
	for _, i := range tab {
		affiche += i + " "
	}
	fmt.Print(affiche + "\n")
}

func affichasciimot(tab []string, mot []string, ascii rune) {
	lignes := make([]string, 8)
	charac := ""
	for _, lettre := range mot {
		if lettre == "_" {
			charac = tab[0]
		} else {
			if ascii == 'm' {
				charac = tab[lettre[0]-64]
			} else {
				charac = tab[lettre[0]-64]
			}
		}
		ligneascii := strings.Split(charac, "\n")
		for i, line := range ligneascii {
			if i < len(lignes) {
				lignes[i] += line
			}
		}
	}
	for _, ligne := range lignes {
		fmt.Printf(ligne + "\n")
	}
}

func simplelettre(l string) bool {
	for _, i := range l {
		if (i < 'a' || i > 'z') && (i < 'A' || i > 'Z') {
			return false
		}
	}
	return true
}

func gettxt(taille string) []string {
	tab := []string{}
	fichier, err := os.Open(taille + ".txt")
	if err != nil {
		fmt.Print(err)
	}
	fileScanner := bufio.NewScanner(fichier)
	fileScanner.Split(bufio.ScanLines)
	lettre := ""
	lscan := 0
	for fileScanner.Scan() {
		if lscan < 0 {
			lscan++
			continue
		}
		lettre += fileScanner.Text() + "\n"
		lscan++
		if lscan%7 == 0 {
			tab = append(tab, lettre)
			lettre = ""
			lscan -= 8
		}
	}
	fichier.Close()
	return tab
}

func sauvegarde(state GameState) error {
	fichier, err := os.Create("save.txt")
	if err != nil {
		return err
	}
	defer fichier.Close()

	_, err = fichier.WriteString(fmt.Sprintf("%s\n%s\n%d\n%s\n", state.Word, convertmotenstr(state.MaskedWord), state.RemainingAttempts, state.Ascii))
	if err != nil {
		return err
	}
	fmt.Println("Partie sauvegardée dans save.txt.")
	return nil
}

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
		state.MaskedWord = convertmotentab(scanner.Text())
	}
	if scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%d", &state.RemainingAttempts)
	}
	if scanner.Scan() {
		state.Ascii = scanner.Text()
	}
	return state, nil
}

func convertmotentab(mot string) []string {
	motf := []string{}
	for _, i := range mot {
		motf = append(motf, string(i))
	}
	return motf
}

func convertmotenstr(mot []string) string {
	motf := ""
	for _, i := range mot {
		motf += i
	}
	return motf
}
