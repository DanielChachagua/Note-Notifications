#!/bin/bash

set -e

echo "🚀 Iniciando la desinstalación de ntn..."

# --- Detección del Sistema Operativo ---
OS="$(uname)"
SERVICE_NAME="ntn.service"

if [ "$OS" == "Linux" ]; then
  echo "🐧 Detectado sistema Linux. Desinstalando servicio systemd --user..."

  SERVICE_PATH="$HOME/.config/systemd/user/$SERVICE_NAME"

  systemctl --user stop "$SERVICE_NAME" 2>/dev/null || echo "❗ No se pudo detener el servicio o no estaba en ejecución."
  systemctl --user disable "$SERVICE_NAME" 2>/dev/null || echo "❗ No se pudo deshabilitar el servicio o no estaba habilitado."

  if [ -f "$SERVICE_PATH" ]; then
    rm -f "$SERVICE_PATH"
    echo "✅ Archivo de servicio eliminado: $SERVICE_PATH"
    systemctl --user daemon-reload
  else
    echo "❗ Archivo de servicio no encontrado en $SERVICE_PATH"
  fi

  # (Opcional) Deshabilitar linger si ya no se necesita
  loginctl disable-linger "$USER" 2>/dev/null || true

elif [ "$OS" == "Darwin" ]; then
  echo "🍎 Detectado sistema macOS. Desinstalando servicio launchd..."
  PLIST_PATH="$HOME/Library/LaunchAgents/com.gemini.ntn.plist"

  if [ -f "$PLIST_PATH" ]; then
    launchctl bootout gui/$(id -u) "$PLIST_PATH" || echo "❗ No se pudo descargar el servicio o ya estaba descargado."
    rm -f "$PLIST_PATH"
    echo "✅ Servicio eliminado: $PLIST_PATH"
  else
    echo "❗ Archivo .plist no encontrado en $PLIST_PATH"
  fi

else
  echo "❌ Sistema operativo no soportado: $OS" >&2
  exit 1
fi

# --- Eliminar binario ---
BIN_PATH="/usr/local/bin/ntn"

if [ -f "$BIN_PATH" ]; then
  echo "🗑️  Eliminando el binario de $BIN_PATH..."
  sudo rm -f "$BIN_PATH"
  echo "✅ Binario eliminado."
else
  echo "ℹ️  El binario no existe en $BIN_PATH. Nada que eliminar."
fi

echo "🎉 ¡Desinstalación completada!"



# #!/bin/bash

# # --- Verificación de Permisos ---
# if [ "$EUID" -ne 0 ]; then
#   echo "❌ Por favor, ejecuta este script con sudo: sudo ./uninstall.sh"
#   exit 1
# fi

# # --- Detección del Sistema Operativo ---
# OS="$(uname)"

# echo "🚀 Iniciando la desinstalación de ntn..."

# # --- Lógica Específica del Sistema Operativo ---
# if [ "$OS" == "Linux" ]; then
#   echo "🐧 Detectado sistema Linux. Desinstalando servicio systemd..."
#   systemctl disable --now ntn.service || echo "El servicio no estaba activo o no existía."
#   rm -f /etc/systemd/system/ntn.service
#   systemctl daemon-reload
#   echo "✅ Servicio systemd de ntn eliminado."

# elif [ "$OS" == "Darwin" ]; then
#   echo "🍎 Detectado sistema macOS. Desinstalando servicio launchd..."
#   SERVICE_USER=${SUDO_USER:-$(whoami)}
#   PLIST_PATH="/Users/$SERVICE_USER/Library/LaunchAgents/com.gemini.ntn.plist"
#   if [ -f "$PLIST_PATH" ]; then
#     launchctl bootout gui/$UID "$PLIST_PATH" || echo "El servicio no estaba cargado."
#     rm -f "$PLIST_PATH"
#     echo "✅ Servicio launchd de ntn eliminado."
#   else
#     echo "El archivo .plist no fue encontrado. Saltando desinstalación del servicio."
#   fi

# else
#   echo "❌ Sistema operativo no soportado: $OS" >&2
#   exit 1
# fi

# # --- Eliminar Binario ---
# echo "🗑️  Eliminando el binario de /usr/local/bin..."
# rm -f /usr/local/bin/ntn

# echo "🎉 ¡Desinstalación completada!"