#show heading.where(level: 1): set text(size: 20pt, weight: "bold", fill: rgb("#17324D"))
#show heading.where(level: 2): set text(size: 16pt, weight: "bold", fill: rgb("#245B82"))
#show heading.where(level: 3): set text(size: 12pt, weight: "bold", fill: rgb("#4B6F86"))

= Analyse des besoins

== Contexte et objectifs

Le projet consiste à créer un site web permettant de jouer au Skyjo en ligne,
seul ou avec d'autres joueurs.

== Périmètre

Le projet couvre les fonctionnalités suivantes :

- la gestion des joueurs invités et des joueurs inscrits ;
- la création et la participation à des lobbies ;
- le déroulement d'une partie de Skyjo en ligne ;
- la synchronisation de la partie entre les joueurs ;
- une messagerie interne pour les joueurs d'un même lobby.

== Acteurs et systèmes externes

- *Joueur invité* : participe à une partie sans créer de compte ;
- *Joueur inscrit* : possède un compte et peut gérer son profil ;
- *Créateur du lobby* : crée un lobby et invite d'autres joueurs ;
- *Hôte du lobby* : gère le lobby et peut lancer la partie (par défaut, le créateur du lobby est l'hôte, mais ce rôle peut être transféré) ;
- *Système de jeu* : gère les lobbies, les cartes et le déroulement des parties ;
- *Service d'e-mail* : envoie les liens nécessaires à la gestion des comptes ;
- *Service de connexion Google* : permet aux joueurs de se connecter avec leur compte Google.

== Règles métier du jeu

- La partie doit respecter les règles officielles du Skyjo.
- Chaque joueur doit retourner deux cartes avant le début effectif de la
  partie. La partie ne peut pas commencer tant que tous les joueurs n'ont pas effectué
  cette action.
- En cas de déconnexion prolongée d'un joueur, les joueurs restants peuvent
  voter pour son exclusion.

== Besoins fonctionnels

=== Gestion du joueur

- *BF-01* : Le système doit permettre à un joueur de participer à une partie
  en tant qu'invité.
- *BF-02* : Un joueur invité doit pouvoir renseigner un pseudo.
- *BF-03* : Un joueur invité doit pouvoir ajouter une photo de profil, de façon
  optionnelle.
- *BF-04* : Le système doit permettre à un joueur de créer un compte et de
  s'y connecter.
- *BF-05* : Un compte doit comporter un pseudo, une adresse e-mail, un mot de
  passe et, éventuellement, une photo de profil.
- *BF-06* : Un joueur inscrit doit pouvoir modifier son pseudo, sa photo de
  profil et son mot de passe.
- *BF-07* : Le système doit utiliser un lien envoyé par e-mail pour les actions
  nécessitant une vérification de l'adresse e-mail ou une modification du mot
  de passe.
- *BF-08* : Le système doit permettre à un joueur de supprimer son compte, ce qui entraîne la suppression de toutes les données le concernant.
- *BF-09* : Le système doit permettre à un joueur de se déconnecter de son compte.
- *BF-10* : Le système doit permettre à un joueur de se connecter avec son compte Google.

=== Gestion du lobby

- *BF-11* : Un joueur doit pouvoir créer un lobby.
- *BF-12* : Un lobby doit pouvoir accueillir jusqu'à 8 joueurs ou IA, auxquels peuvent s'ajouter un nombre illimité de spectateurs.
- *BF-13* : Un lobby doit posséder un code permettant à d'autres joueurs de le
  rejoindre.
- *BF-14* : Le système doit permettre de partager un lien d'invitation vers un
  lobby.
- *BF-15* : Une partie doit pouvoir être lancée uniquement lorsque le lobby
  contient entre 2 et 8 joueurs.
- *BF-16* : L'hôte du lobby doit pouvoir exclure un joueur quand il le souhaite.
- *BF-17* : L'hôte du lobby doit pouvoir transférer son rôle à un autre joueur.
- *BF-18* : Le système doit permettre à un joueur de quitter un lobby.
- *BF-19* : L'hôte du lobby doit pouvoir fermer le lobby, ce qui entraîne l'exclusion de tous les joueurs et spectateurs.
- *BF-20* : L'hôte du lobby doit pouvoir ajouter des IA pour compléter le nombre de joueurs et modifier leur niveau de difficulté.
- *BF-21* : Le système doit permettre aux joueurs d'un lobby de communiquer entre
  eux via une messagerie interne.
- *BF-22* : Le lobby doit pouvoir être public ou privé. Dans le cas d'un lobby privé, seuls les joueurs ayant reçu le lien d'invitation ou ayant le code peuvent le rejoindre.
- *BF-23* : Le système doit permettre à un joueur qui n'est pas dans un lobby de voir la liste des lobbys publics et de les rejoindre.
- *BF-24* : Le système doit permettre à un joueur de voir la liste des joueurs et spectateurs présents dans le lobby.
- *BF-25* : Le système doit permettre à l'hôte du lobby de déplacer un spectateur vers la liste des joueurs, et inversement.

=== Gestion de la partie

- *BF-26* : Le système doit distribuer automatiquement les cartes au début de
  la partie.
- *BF-27* : Le système doit appliquer les règles du Skyjo définies dans la
  section "Règles métier du jeu".
- *BF-28* : Le système doit synchroniser l'état de la partie entre les joueurs.
- *BF-29* : Le système doit gérer la déconnexion d'un joueur pendant une
  partie.
- *BF-30* : Le système doit permettre à un joueur de quitter une partie en cours,
  auquel cas il est remplacé par une IA si nécessaire.

=== Pages principales
- *BF-31* : Le système doit proposer une page d'accueil présentant le jeu et
  permettant de se connecter, de créer un compte ou de continuer en tant qu'invité.
- *BF-32* : Le système doit proposer une page de profil permettant à un joueur
  inscrit de modifier ses informations personnelles.
- *BF-33* : Après s'être connecté ou avoir continué en tant qu'invité, le système doit proposer un menu principal permettant de créer ou rejoindre un lobby.
- *BF-34* : Le système doit proposer une page de lobby permettant aux joueurs de voir les autres joueurs et spectateurs, de communiquer via la messagerie interne et de lancer la partie.
- *BF-35* : Le système doit proposer une page permettant de voir la liste des lobbys publics et de les rejoindre.
- *BF-36* : Le système doit proposer une page de jeu.

== Besoins non fonctionnels

- Le site doit être utilisable sur ordinateur, tablette et téléphone.
- Les actions réalisées par un joueur doivent être transmises aux autres
  joueurs dans un délai suffisamment court pour permettre une partie fluide.
- Les données des comptes et des parties doivent être protégées.
- Le site doit signaler clairement les erreurs, notamment lorsqu'un lobby est
  complet, inexistant ou fermé.

== Architecture technique

=== Interface utilisateur
- *Framework* : SvelteKit

=== Serveur et API
- *Langage* : Go (Golang)
- *API web* : Gin (`gin-gonic/gin`)
- *WebSocket* : Gorilla WebSocket (`gorilla/websocket`)

=== Données
- *Base de données* : PostgreSQL
- *ORM* : GORM (`gorm.io/gorm`)

=== Intelligence artificielle
- *Entraînement* : PyTorch (Python)
- *Exécution du modèle* : ONNX Runtime

=== Connexion et sécurité
- *Connexion avec Google* : Google Identity Services (OAuth 2.0)
- *Vérification des identifiants* : `google.golang.org/api/idtoken`
- *Sessions* : `golang-jwt/jwt`

=== Outils de développement
- *Environnement local* : Docker & Docker Compose
- *Tests Go* : package standard `testing` + `stretchr/testify`
- *Tests de l'interface* : Vitest

