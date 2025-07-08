#!/bin/bash

set -e

# --- Verificación del Binario ---
if [ -f "ntn" ]; then
  BIN="ntn"
elif [ -f "ntn-darwin" ]; then
  BIN="ntn-darwin"
else
  echo "❌ No se encontró ni 'ntn' ni 'ntn-darwin' en el directorio actual."
  exit 1
fi

echo "🚀 Iniciando la instalación de ntn..."

# --- Instalación del Binario (requiere sudo) ---
echo "📂 Moviendo el binario a /usr/local/bin..."
sudo mv ntn /usr/local/bin/

# --- Detección del Sistema Operativo ---
OS="$(uname)"

if [ "$OS" == "Linux" ]; then
  echo "🐧 Detectado sistema Linux. Configurando servicio systemd de usuario..."

  SERVICE_NAME="ntn.service"
  SYSTEMD_USER_DIR="$HOME/.config/systemd/user"
  mkdir -p "$SYSTEMD_USER_DIR"

  cat <<EOF > "$SYSTEMD_USER_DIR/$SERVICE_NAME"
[Unit]
Description=Note Notification Service
After=graphical-session.target network-online.target default.target
Wants=network-online.target

[Service]
ExecStart=/bin/bash -c "sleep 15 && /usr/local/bin/ntn serve"
Restart=always
RestartSec=5s
Environment=DISPLAY=:0
Environment=XAUTHORITY=$HOME/.Xauthority

[Install]
WantedBy=default.target
EOF

  echo "✅ Archivo $SERVICE_NAME creado en $SYSTEMD_USER_DIR"

  # Recargar y habilitar el servicio de usuario
  systemctl --user daemon-reload
  systemctl --user enable --now "$SERVICE_NAME"

  # Para que funcione tras reinicio incluso sin login explícito (opcional)
  loginctl enable-linger "$USER" || true

  echo "✅ Servicio de usuario $SERVICE_NAME habilitado y ejecutándose."

elif [ "$OS" == "Darwin" ]; then
  echo "🍎 Detectado macOS. Configurando con launchd..."

  WRAPPER_PATH="$HOME/.local/bin/ntn-wrapper.sh"
  PLIST_PATH="$HOME/Library/LaunchAgents/com.gemini.ntn.plist"

  cat <<EOF > "$WRAPPER_PATH"
#!/bin/bash
sleep 30
/usr/local/bin/ntn serve
EOF

  chmod +x "$WRAPPER_PATH"


  mkdir -p "$(dirname "$PLIST_PATH")"

  cat <<EOF > "$PLIST_PATH"
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>Label</key>
    <string>com.gemini.ntn</string>

    <key>ProgramArguments</key>
    <array>
        <string>/usr/local/bin/ntn</string>
        <string>serve</string>
    </array>

    <key>RunAtLoad</key>
    <true/>

    <key>KeepAlive</key>
    <true/>

    <!-- Para que funcione en sesiones gráficas -->
    <key>EnvironmentVariables</key>
    <dict>
        <key>PATH</key>
        <string>/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin</string>
        <key>DISPLAY</key>
        <string>:0</string>
    </dict>

    <!-- Retardo para asegurar que el sistema esté listo -->
    <key>StartInterval</key>
    <integer>0</integer>

</dict>
</plist>
EOF

  echo "✅ Archivo .plist creado en $PLIST_PATH"
  echo "✅ Archivo .plist creado en $PLIST_PATH"

  launchctl unload "$PLIST_PATH" 2>/dev/null || true
  launchctl load "$PLIST_PATH"
  
  echo "✅ Servicio ntn cargado con launchd."

else
  echo "❌ Sistema operativo no soportado: $OS" >&2
  exit 1
fi

echo "🎉 ¡Instalación completada!"


# #!/bin/bash

# # Salir inmediatamente si un comando falla
# set -e

# # --- Verificación de Permisos ---
# if [ "$EUID" -ne 0 ]; then
#   echo "❌ Ejecuta este script con sudo: sudo ./install.sh"
#   exit 1
# fi

# # --- Detección del Sistema Operativo ---
# OS="$(uname)"

# # --- Verificación del Binario ---
# if [ ! -f "ntn" ]; then
#   echo "❌ El archivo binario 'ntn' no se encontró. Asegúrate de que esté en el mismo directorio que este script."
#   exit 1
# fi

# echo "🚀 Iniciando la instalación de ntn..."

# # --- Instalación del Binario ---
# echo "📂 Moviendo el binario a /usr/local/bin..."
# mv ntn /usr/local/bin/

# # --- Lógica Específica del Sistema Operativo ---
# if [ "$OS" == "Linux" ]; then
#   echo "🐧 Detectado sistema Linux. Configurando con systemd..."

#   # Usamos $SUDO_USER para obtener el usuario que ejecutó sudo
#   SERVICE_USER=${SUDO_USER:-$(whoami)}
  
#   cat <<EOF > /etc/systemd/system/ntn.service
# [Unit]
# Description=Note Notification Service
# After=network-online.target
# Wants=network-online.target

# [Service]
# User=$SERVICE_USER
# Group=$(id -gn $SERVICE_USER)
# Environment="HOME=/home/$SERVICE_USER"
# Environment=DISPLAY=:0
# Environment=XAUTHORITY=/home/$SERVICE_USER/.Xauthority
# ExecStart=/usr/local/bin/ntn serve
# Restart=always
# RestartSec=5s

# [Install]
# WantedBy=multi-user.target
# EOF

#   echo "✅ Archivo ntn.service creado."
#   systemctl daemon-reload
#   systemctl enable --now ntn.service
#   echo "✅ Servicio ntn habilitado y iniciado con systemd."

# elif [ "$OS" == "Darwin" ]; then
#   echo "🍎 Detectado sistema macOS. Configurando con launchd..."

#   # Usamos $SUDO_USER para obtener el usuario que ejecutó sudo
#   SERVICE_USER=${SUDO_USER:-$(whoami)}
#   PLIST_PATH="/Users/$SERVICE_USER/Library/LaunchAgents/com.gemini.ntn.plist"

#   cat <<EOF > "$PLIST_PATH"
# <?xml version="1.0" encoding="UTF-8"?>
# <!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
# <plist version="1.0">
# <dict>
#     <key>Label</key>
#     <string>com.gemini.ntn</string>
#     <key>ProgramArguments</key>
#     <array>
#         <string>/usr/local/bin/ntn</string>
#         <string>serve</string>
#     </array>
#     <key>RunAtLoad</key>
#     <true/>
#     <key>KeepAlive</key>
#     <true/>
# </dict>
# </plist>
# EOF

#   echo "✅ Archivo .plist creado en $PLIST_PATH"
#   # Cambiar el propietario al usuario para que pueda cargarlo
#   chown $SERVICE_USER "$PLIST_PATH"
#   # Cargar el servicio
#   launchctl bootstrap gui/$UID "$PLIST_PATH"
#   echo "✅ Servicio ntn cargado con launchd."

# else
#   echo "❌ Sistema operativo no soportado: $OS" >&2
#   exit 1
# fi

# echo "🎉 ¡Instalación completada!"