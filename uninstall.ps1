# Requiere ejecución como Administrador

# --- Verificación de Permisos ---
if (-NOT ([Security.Principal.WindowsPrincipal][Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole] "Administrator")) {
    Write-Error "Este script debe ser ejecutado como Administrador."
    exit 1
}

Write-Host "🚀 Iniciando la desinstalación de ntn..."

# --- Configuración de Rutas y Nombres ---
$installDir = "$env:ProgramFiles\ntn"
$taskName = "NoteNotificationService"

# --- Eliminar la Tarea Programada ---
Write-Host "⚙️  Eliminando la Tarea Programada..."
Get-ScheduledTask -TaskName $taskName -ErrorAction SilentlyContinue | Unregister-ScheduledTask -Confirm:$false

# --- Eliminar del PATH del Sistema ---
Write-Host "🔧 Eliminando la ruta del PATH del sistema..."
$currentPath = [Environment]::GetEnvironmentVariable("Path", "Machine")
if ($currentPath -like "*$installDir*") {
    $newPath = ($currentPath.Split(';') | Where-Object { $_ -ne $installDir }) -join ';'
    [Environment]::SetEnvironmentVariable("Path", $newPath, "Machine")
    Write-Host "✅ La ruta ha sido eliminada del PATH."
} else {
    Write-Host "La ruta no se encontró en el PATH."
}

# --- Eliminar el Directorio de Instalación ---
if (Test-Path $installDir) {
    Write-Host "🗑️  Eliminando el directorio de instalación..."
    Remove-Item -Path $installDir -Recurse -Force
}

Write-Host "🎉 ¡Desinstalación completada!"

