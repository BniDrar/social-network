#!/bin/bash
# 📦 Docker Compose Commands Utiles pour le Développement

echo "📦 Docker Compose Helper 🚀"

echo "1️⃣  Up (démarrage en arrière-plan)"
echo "2️⃣  Up avec logs"
echo "3️⃣  Build/Rebuild"
echo "4️⃣  Up & Build"
echo "5️⃣  Status"
echo "6️⃣  Logs (tous)"
echo "7️⃣  Logs (backend)"
echo "8️⃣  Logs -f (live)"
echo "9️⃣  Down (garder volumes)"
echo "1️⃣ 0️⃣  Down (supprimer volumes)"
echo "1️⃣ 1️⃣  Restart (all)"
echo "1️⃣ 2️⃣  Restart (backend)"
echo "1️⃣ 3️⃣  Exec shell (backend)"
echo "1️⃣ 4️⃣  Resources (top)"
echo "1️⃣ 5️⃣  Clean docker system ⚠️"
echo "0️⃣  Quitter"
echo ""

read -p "🧠 Choix: " choix

case $choix in
  1) docker compose up -d ;;
  2) docker compose up ;;
  3) docker compose build ;;
  4) docker compose up --build ;;
  5) docker compose ps ;;
  6) docker compose logs ;;
  7) docker compose logs backend ;;
  8) docker compose logs -f ;;
  9) docker compose down ;;
  10) docker compose down -v ;;
  11) docker compose restart ;;
  12) docker compose restart backend ;;
  13) docker compose exec backend sh ;;
  14) docker compose top ;;
  15) 
    echo "⚠️ Tu es sur le point de tout nettoyer ! (images, conteneurs, volumes)"
    read -p "Continuer ? (y/N): " confirm
    if [[ "$confirm" == "y" || "$confirm" == "Y" ]]; then
      docker system prune -a --volumes -f
    else
      echo "✅ Opération annulée."
    fi
    ;;
  0) echo "👋 Bye!" ;;
  *) echo "❌ Choix invalide." ;;
esac
