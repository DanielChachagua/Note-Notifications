function Wait-ForInternet {
    $url = "https://www.google.com"
    $timeout = 30
    $elapsed = 0
    $delay = 3

    Write-Host "🌐 Verificando conexión a Internet..."

    while ($true) {
        try {
            $response = Invoke-WebRequest -Uri $url -UseBasicParsing -TimeoutSec 5
            if ($response.StatusCode -eq 200) {
                Write-Host "✅ Conexión a Internet detectada."
                break
            }
        } catch {
            Write-Host "⏳ Sin conexión aún. Esperando..."
        }

        Start-Sleep -Seconds $delay
        $elapsed += $delay
        if ($elapsed -ge $timeout) {
            Write-Warning "⚠️  No se detectó conexión a Internet luego de $timeout segundos."
            break
        }
    }
}

# --- Verificación de Permisos ---
if (-NOT ([Security.Principal.WindowsPrincipal] [Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole(`
    [Security.Principal.WindowsBuiltInRole]::Administrator)) {
    Write-Error "❌ Este script debe ser ejecutado como Administrador."
    exit 1
}

# --- Configuración de Rutas ---
$installDir = Join-Path $env:ProgramFiles "ntn"
$exePath = Join-Path $installDir "ntn.exe"

# --- Verificación del Binario ---
if (-NOT (Test-Path ".\ntn.exe")) {
    Write-Error "❌ El archivo 'ntn.exe' no se encontró en el directorio actual."
    exit 1
}

Write-Host "🚀 Iniciando la instalación de ntn para Windows..."

# --- Crear Directorio de Instalación ---
if (-NOT (Test-Path $installDir)) {
    New-Item -ItemType Directory -Path $installDir | Out-Null
}

# --- Copiar el Binario ---
Write-Host "📂 Copiando 'ntn.exe' a $installDir..."
Copy-Item -Path ".\ntn.exe" -Destination $exePath -Force

# --- Añadir al PATH del Sistema ---
Write-Host "🔧 Añadiendo $installDir al PATH del sistema..."
$currentPath = [Environment]::GetEnvironmentVariable("Path", "Machine")
if ($currentPath -notlike "*$installDir*") {
    $newPath = "$currentPath;$installDir"
    [Environment]::SetEnvironmentVariable("Path", $newPath, "Machine")
    Write-Host "✅ Ruta añadida al PATH. Reiniciá tu terminal para aplicar los cambios."
} else {
    Write-Host "ℹ️ El PATH ya incluye $installDir."
}

# --- Esperar conexión a Internet (opcional) ---
Wait-ForInternet

# --- Crear tarea programada ---
Write-Host "⚙️  Configurando tarea programada para ejecutar 'ntn serve'..."

$taskName = "NoteNotificationService"
$taskAction = New-ScheduledTaskAction -Execute $exePath -Argument "serve"
$taskTrigger = New-ScheduledTaskTrigger -AtLogOn
$taskPrincipal = New-ScheduledTaskPrincipal -UserId "SYSTEM" -LogonType ServiceAccount -RunLevel Highest

# Eliminar tarea anterior si ya existe
if (Get-ScheduledTask -TaskName $taskName -ErrorAction SilentlyContinue) {
    Unregister-ScheduledTask -TaskName $taskName -Confirm:$false
}

# Registrar nueva tarea
Register-ScheduledTask -TaskName $taskName -Action $taskAction -Trigger $taskTrigger -Principal $taskPrincipal

Write-Host "🎉 ¡Instalación completada correctamente!"


# function Wait-ForInternet {
#     $url = "https://www.google.com"
#     $timeout = 30
#     $elapsed = 0
#     $delay = 3

#     Write-Host "🌐 Verificando conexión a Internet..."

#     while ($true) {
#         try {
#             $response = Invoke-WebRequest -Uri $url -UseBasicParsing -TimeoutSec 5
#             if ($response.StatusCode -eq 200) {
#                 Write-Host "✅ Conexión a Internet detectada."
#                 break
#             }
#         } catch {
#             Write-Host "⏳ Sin conexión aún. Esperando..."
#         }

#         Start-Sleep -Seconds $delay
#         $elapsed += $delay
#         if ($elapsed -ge $timeout) {
#             Write-Warning "⚠️  No se detectó conexión a Internet luego de $timeout segundos."
#             break
#         }
#     }
# }


# # Requiere ejecución como Administrador

# # --- Verificación de Permisos ---
# if (-NOT ([Security.Principal.WindowsPrincipal][Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole] "Administrator")) {
#     Write-Error "Este script debe ser ejecutado como Administrador."
#     exit 1
# }

# # --- Configuración de Rutas ---
# $installDir = "$env:ProgramFiles\ntn"
# $exePath = Join-Path $installDir "ntn.exe"

# # --- Verificación del Binario ---
# if (-NOT (Test-Path ".\ntn.exe")) {
#     Write-Error "El archivo 'ntn.exe' no se encontró en el directorio actual."
#     exit 1
# }

# Write-Host "🚀 Iniciando la instalación de ntn para Windows..."

# # --- Crear Directorio de Instalación ---
# if (-NOT (Test-Path $installDir)) {
#     New-Item -ItemType Directory -Path $installDir
# }

# # --- Copiar el Binario ---
# Write-Host "📂 Copiando ntn.exe a $installDir..."
# Copy-Item -Path ".\ntn.exe" -Destination $exePath -Force

# # --- Añadir al PATH del Sistema ---
# Write-Host "🔧 Añadiendo $installDir al PATH del sistema..."
# $currentPath = [Environment]::GetEnvironmentVariable("Path", "Machine")
# if ($currentPath -notlike "*$installDir*") {
#     [Environment]::SetEnvironmentVariable("Path", "$currentPath;$installDir", "Machine")
#     Write-Host "✅ ntn ha sido añadido al PATH. Por favor, reinicia tu terminal para usar el comando."
# } else {
#     Write-Host "El PATH ya está configurado."
# }

# # --- Configurar Tarea Programada para Ejecución en Segundo Plano ---
# Wait-ForInternet

# Write-Host "⚙️  Configurando una Tarea Programada para mantener ntn activo..."

# $taskName = "NoteNotificationService"
# $taskAction = New-ScheduledTaskAction -Execute $exePath -Argument "serve"
# $taskTrigger = New-ScheduledTaskTrigger -AtLogOn
# $taskPrincipal = New-ScheduledTaskPrincipal -UserId "SYSTEM" -LogonType ServiceAccount -RunLevel Highest

# # Eliminar la tarea si ya existe
# Get-ScheduledTask -TaskName $taskName -ErrorAction SilentlyContinue | Unregister-ScheduledTask -Confirm:$false

# # Registrar la nueva tarea
# Register-ScheduledTask -TaskName $taskName -Action $taskAction -Trigger $taskTrigger -Principal $taskPrincipal

# Write-Host "🎉 ¡Instalación completada!"
