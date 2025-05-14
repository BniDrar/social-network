#!/bin/bash
# Commandes Docker Compose essentielles

# Démarrer tous les services en arrière-plan
docker compose up -d

# Démarrer les services et voir les logs
docker compose up

# Construire ou reconstruire les services
docker compose build

# Démarrer et reconstruire les services
docker compose up --build

# Voir l'état des services
docker compose ps

# Voir les logs de tous les services
docker compose logs

# Voir les logs d'un service spécifique (ex: backend)
docker compose logs backend

# Voir les logs et les suivre en temps réel
docker compose logs -f

# Arrêter les services mais conserver les volumes
docker compose down

# Arrêter les services et supprimer les volumes
docker compose down -v

# Redémarrer tous les services
docker compose restart

# Redémarrer un service spécifique
docker compose restart backend

# Exécuter une commande dans un conteneur en cours d'exécution
docker compose exec backend sh

# Afficher les ressources utilisées (CPU, mémoire)
docker compose top