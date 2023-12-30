# SnowCamp 2024 ⛄️
> handson

## 🇫🇷 Créer des plugins en Rust, TinyGo, ... pour vos applications Go avec WebAssembly et Wazero et Extism

[WebAssembly](https://webassembly.org/) (ou Wasm, son petit nom) était initialement prévu pour s'exécuter dans le navigateur. Grâce à ses qualités (polyglotte, efficacité, rapidité, ...) , très vite il est apparu indispensable de pouvoir exécuter des programmes Wasm en dehors du navigateur (ce qui est maintenant possible avec la spécification [Wasi](https://wasi.dev/)).
L'objectif de ce hands-on est de vous faire développer une application en Go (un serveur web), capable d'exécuter des plugins développés (par vous) avec des langages différents et compilé en Wasm.

1ère partie: ce sera théorique
- Wasm, Wasi, qu'est-ce que c'est ?
- Les possibilités et les limites (et comment contourner ces limites)
- Les outils existants
- Présentation de [Wazero](https://wazero.io/) (un runtime Wasm sans dépendance)
- Présentation d'[Extism](https://extism.org/) (un framework pour réaliser des applications et des plugins)

2ème partie: écrire des plugins Wasm avec Extism pour la CLI Extism
- 2ème plugin en Go
- 1er plugin en Rust
- Écrire votre propre "lanceur" de plugin en Go

3ème partie: écrire un serveur d'application (http) avec le SDK Extism qui va appeler/exécuter les plugins Wasm de la 2ème partie
- un 1er serveur http qui va fonctionner (mais pas toujours)
- un 2ème serveur http qui va fonctionner (tout le temps)
- un 3ème serveur http qui va encore mieux fonctionner

Et pour finir nous verrons comment dockeriser tout ça pour le déployer facilement avec la plus petite image possible

PS: Pré-requis (**avant de venir**), vous devrez avoir un compte Docker et un compte Gitpod (https://www.gitpod.io/) qui est un VSCode dans le cloud (il existe un plan free). Si pour x raisons vous ne souhaitez pas utiliser Gitpod et utiliser vos propres outils, c'est tout à fait possible, mais vous devrez avoir un environnement fonctionnel dès le début du workshop, et nous ne pourrons pas vous aider si votre environnement ne fonctionne pas)


## 🇬🇧 Create Rust, TinyGo plugins for your Golang applications thanks to WebAssembly, Wazero and Extism

### Use this repository with Gitpod (for the handson at SnowCamp)

#### Read-only mode

- [🍊 Open it with Gitpod](https://gitpod.io/#https://github.com/bots-garden/snowcamp-2024-wasm)

#### Read-write mode

- Fork this repository
- Then use this link to open it with Gitpod: https://gitpod.io/#https://github.com/YOUR-GITHUB-HANDLE/snowcamp-2024-wasm
