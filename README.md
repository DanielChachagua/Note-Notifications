# NoteNotification CLI (`ntn`)

`ntn` es una herramienta de línea de comandos (CLI) construida en Go para gestionar notas y programar notificaciones. Te permite crear, listar y actualizar notas directamente desde tu terminal.

Este proyecto utiliza Cobra para la estructura de los comandos y Tablewriter para mostrar las listas de forma ordenada.

## Características Principales

- **Crear notas**: Añade nuevas notas con título, descripción, fecha, hora y una URL opcional.
- **Listar notas**: Visualiza todas tus notas en una tabla bien formateada en la consola.
- **Actualizar notas**: Modifica los detalles de una nota existente usando su ID.
- **Integración con Google Calendar**: La CLI se autentica de forma segura con tu cuenta de Google para poder gestionar eventos y notificaciones (basado en el código de `jobs/calendar.go`).

## Requisitos Previos

1.  **Go**: Asegúrate de tener Go 1.18 o una versión superior instalada.
2.  **Credenciales de Google API**: Para la funcionalidad de calendario, necesitas credenciales de la API de Google.

    - Ve a la Google Cloud Console.
    - Crea un nuevo proyecto o selecciona uno existente.
    - En el menú de navegación, ve a **APIs & Services > Library**.
    - Busca y habilita la **Google Calendar API**.
    - Ve a **APIs & Services > Credentials**.
    - Haz clic en **Create Credentials > OAuth client ID**.
    - Selecciona **Desktop app** como tipo de aplicación.
    - Descarga el archivo JSON con las credenciales. Es importante que lo renombres a `credentials.json` y lo coloques en la raíz del proyecto para que el job de calendario pueda encontrarlo.

## Instalación

1.  **Clona el repositorio:**
    ```sh
    git clone <URL_DEL_REPOSITORIO>
    cd <NOMBRE_DEL_DIRECTORIO>
    ```

2.  **Construye el ejecutable:**
    Puedes construir el binario y nombrarlo `ntn` para un uso más fácil.
    ```sh
    go build -o ntn ./cmd/note_notification
    ```

3.  **Mueve el binario a tu PATH (Opcional):**
    Para poder ejecutar `ntn` desde cualquier lugar, mueve el binario a un directorio incluido en tu `PATH`.
    ```sh
    # Para Linux/macOS
    sudo mv ntn /usr/local/bin/

    # Para Windows, muévelo a una carpeta que esté en tus variables de entorno.
    ```

## Uso

La estructura general de los comandos es `ntn note <subcomando> [flags]`.

### Autenticación (Primer Uso)

La primera vez que ejecutes un comando que requiera acceso a Google Calendar, se iniciará un proceso de autenticación:
1.  Se abrirá una pestaña en tu navegador pidiéndote que inicies sesión y autorices el acceso a tu calendario.
2.  Después de autorizar, serás redirigido y la CLI guardará un token de acceso en tu directorio de usuario (`~/.credentials/calendar-go.json`) para futuras sesiones.

### Comandos y Ejemplos

#### `add` - Crear una nueva nota

Crea una nota con los detalles especificados.

```sh
# Ejemplo básico
ntn note add --title "Reunión de equipo" --description "Discutir el sprint actual" --date "28-11-2024" --time "10:30"

# Ejemplo con URL y aviso desactivado
ntn note add -n "Comprar dominio" -b "Buscar y comprar el dominio para el proyecto" -d "29-11-2024" -t "15:00" -u "https://dominios.com" -w "false"
```

**Flags:**
- `-n`, `--title` (string, **requerido**): Título de la nota.
- `-b`, `--description` (string, **requerido**): Descripción de la nota.
- `-d`, `--date` (string, **requerido**): Fecha en formato `dd-mm-yyyy`.
- `-t`, `--time` (string, **requerido**): Hora en formato `hh:mm`.
- `-u`, `--url` (string, opcional): URL asociada.
- `-w`, `--warn` (string, opcional): Activar/desactivar aviso (`true`/`false`). Por defecto es `true`.

---

#### `list` - Listar todas las notas

Muestra una tabla con todas las notas guardadas.

```sh
ntn note list
```

---

#### `update` - Actualizar una nota existente

Actualiza los campos de una nota existente, identificada por su ID.

```sh
# Para actualizar solo el título de una nota
ntn note update --id "ID_DE_LA_NOTA" --title "Nuevo título para la reunión"
```

**Flags:**
- `-i`, `--id` (string, **requerido**): ID de la nota a actualizar.
- `-n`, `--title` (string, opcional): Nuevo título.
- `-b`, `--description` (string, opcional): Nueva descripción.
- `-d`, `--date` (string, opcional): Nueva fecha (`dd-mm-yyyy`).
- `-t`, `--time` (string, opcional): Nueva hora (`hh:mm`).
- `-u`, `--url` (string, opcional): Nueva URL.
- `-w`, `--warn` (string, opcional): Nuevo estado del aviso (`true`/`false`).

---