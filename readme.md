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
    environement:
      - TARGET_DIR_PATH: /git_folder
      - GIT_URL: https://github.com/profile/reposity

```

## Variables d'environement

- TARGET_DIR_PATH : Lien vers le dossier à synchroniser dans le conteneur. A adapter selon votre volume. Par défaut : /git_folder
- GIT_URL : le lien vers le repertoire Git à synchroniser. Par défaut: vide
- PERIOD : en secondes, durée entre deux pull. Par défaut : 60