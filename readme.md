# Concatenator

Concatenator est un outil CLI qui parcourt un répertoire, collecte les informations des fichiers et les génère sous forme de fichier JSON structuré.

## Fonctionnalités

- Parcours de répertoire (récursif ou non)
- Collecte d'informations sur les fichiers (nom, contenu, taille, date de modification, chemin)
- Génération d'un fichier JSON structuré

## Installation

Assurez-vous d'avoir Go installé sur votre système, puis exécutez :

```bash
go get github.com/chrlesur/concatenator
```

## Utilisation

```bash
concatenator concatenate [output_file]
```

Options :
- `--dir`, `-d` : Répertoire d'entrée (par défaut : répertoire courant)
- `--recursive`, `-r` : Parcourir le répertoire de manière récursive

Pour afficher la version :

```bash
concatenator version
```

## Exemple

```bash
concatenator concatenate --dir=/chemin/vers/dossier --recursive output.json
```

Cela générera un fichier `output.json` contenant les informations de tous les fichiers du dossier spécifié et de ses sous-dossiers.

## Développement

### Prérequis

- Go 1.22.0 ou supérieur

### Dépendances

- github.com/spf13/cobra
- github.com/spf13/pflag

### Construction

```bash
go build
```

## Licence

Ce projet est sous licence [GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html).

## Auteur

[[chrlesur](https://github.com/chrlesur)](https://github.com/chrlesur)
