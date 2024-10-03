# Concatenator

Concatenator est un outil CLI qui parcourt un répertoire, collecte les informations des fichiers et les génère sous forme de fichier JSON structuré.

## Fonctionnalités

- Parcours de répertoire (récursif ou non)
- Collecte d'informations sur les fichiers (nom, contenu, taille, date de modification, chemin)
- Génération d'un fichier JSON structuré
- Exclusion de fichiers basée sur des motifs (wildcards)
- Inclusion sélective de fichiers basée sur des motifs (wildcards)

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
- `--exclude`, `-e` : Exclure des fichiers basés sur des motifs (supporte les wildcards)
- `--include`, `-i` : Inclure uniquement les fichiers correspondant aux motifs (supporte les wildcards)

Pour afficher la version :

```bash
concatenator version
```

## Exemples

Générer un fichier JSON en incluant uniquement les fichiers .txt et .log :

```bash
concatenator concatenate --dir=/chemin/vers/dossier --recursive --include="*.txt,*.log" output.json
```

Inclure certains fichiers mais exclure les fichiers .tmp :

```bash
concatenator concatenate --dir=/chemin/vers/dossier --recursive --include="*.txt,*.log" --exclude="*.tmp" output.json
```

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

[chrlesur](https://github.com/chrlesur)
