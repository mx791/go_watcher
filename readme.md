# Go_Watcher

Un outil dockerisé écrit en go pour synchroniser automatiquement un dossier avec un répertoire Git.

## Lancement
Avec docker-compose
```
version: "3"

services:

  web:
    image: nginx
    volumes:
      - ./data:/data/www/html
    ports:
      - 80:80

  watcher:
    image: gowatcher
    volumes:
      - ./data:/git_folder
    environment:
      - TARGET_DIR_PATH=/git_folder
      - GIT_URL=https://github.com/....git
      - POST_UPDATE=echo updating... > out.txt
      - PERIOD=60

```

## Variables d'environement

- TARGET_DIR_PATH : Lien vers le dossier à synchroniser dans le conteneur. A adapter selon votre volume. Par défaut : /git_folder
- GIT_URL : le lien vers le repertoire Git à synchroniser. Par défaut: vide
- PERIOD : en secondes, durée entre deux pull. Par défaut : 60
- POST_UPDATE : Commandes bash à executer en cas de modifications sur le repertoire.