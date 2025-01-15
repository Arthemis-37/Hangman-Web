# Hangman-Web

---

Vous devez créer une version web du jeu classique Hangman en utilisant le langage de programmation Go. Voici une explication détaillée des exigences pour votre dépôt **hangman-web** :

1. **Objectif** :  
   - Créez une application web pour le jeu classique Hangman.  
   - Utilisez un **Go module** pour réutiliser les fonctions et la logique de votre projet **hangman-classic**. Évitez de copier ou réécrire ce code dans le nouveau projet ; au lieu de cela, référencez-le comme une dépendance.  

2. **Dépôt et Configuration** :  
   - Créez un dépôt privé nommé **hangman-web**.
   - Ajoutez un fichier nommé `words.txt`, qui contient une liste de mots (un mot par ligne) à utiliser comme source pour le jeu Hangman.  
   - Votre programme acceptera `words.txt` comme paramètre au démarrage.  

3. **Fonctionnalité principale** :  
   - Le comportement du jeu (règles, sélection de mots, conditions de victoire/perte) doit correspondre à celui du projet **hangman-classic**.
   - Implémentez au moins deux points de terminaison HTTP pour gérer le jeu :
     - **GET /** :  
       - Ce point de terminaison servira l'interface HTML principale du jeu.  
       - Utilisez les **templates** Go pour afficher dynamiquement les données du jeu (par exemple, l'état actuel du mot à deviner, les erreurs, et les tentatives restantes).  
     - **POST /hangman** :  
       - Ce point de terminaison gérera les entrées de l'utilisateur (par exemple, les lettres devinées) envoyées depuis l'interface web.  
       - Utilisez des formulaires HTML ou d'autres méthodes d'entrée pour collecter et envoyer les données à ce point de terminaison.  

4. **Exigences de l'interface web** :  
   - La page principale doit inclure :  
     - Un affichage textuel montrant le mot partiellement révélé (par exemple, `_ _ A _ M _ N`).  
     - Un champ de saisie texte où l'utilisateur peut taper une lettre à deviner.  
     - Un bouton pour soumettre la tentative, déclenchant une requête POST vers `/hangman`.  
     - Après la soumission, la page doit être mise à jour pour afficher l'état du jeu (par exemple, mot mis à jour, lettres incorrectes, tentatives restantes).  

5. **Conseils techniques** :  
   - Utilisez les **templates Go** pour générer dynamiquement l'HTML, permettant ainsi de mettre à jour les données du jeu sur l'interface.  
   - Assurez-vous que la requête POST vers `/hangman` redirige l'utilisateur vers `/` avec l'état du jeu mis à jour.  

6. **Consistance du comportement** :  
   - Suivez les règles et mécaniques de votre projet **hangman-classic** pour la sélection des mots, les conditions de victoire/perte, et le suivi des tentatives.  

Ce projet démontrera votre capacité à créer une application web en Go tout en maintenant de bonnes pratiques telles que la modularité, la réutilisation du code et une interface utilisateur conviviale.
