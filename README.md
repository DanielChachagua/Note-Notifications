# NoteNotification CLI (`ntn`)

`ntn` is a command-line tool (CLI) built in Go to manage events in your Google Calendar. It allows you to create, list, update, and delete events directly from your terminal.

This project uses Cobra for the command structure and Tablewriter to display lists in an orderly manner.

## Main Features

- **Create events**: Add new events with summary, description, location, date, and time.
- **List events**: View all your upcoming events in a well-formatted table in the console.
- **Update events**: Modify the details of an existing event using its ID.
- **Delete events**: Remove one or more events from your calendar.
- **Google Calendar Integration**: The CLI securely authenticates with your Google account to manage events and notifications.
- **Background Service**: Automatically runs in the background to handle notifications and keep events updated.

## Installation

Choose the command for your operating system. The installer will download the latest version of `ntn`, place it in the appropriate system directory, and set it up as a background service.

**Note:** You will need to download the appropriate binary (`ntn` or `ntn.exe`) and the installer script from the [latest release on GitHub](https://github.com/DanielChachagua/NoteNotification/releases/latest) and run them from the same directory.

### Linux & macOS

Open your terminal and run the following command. It will ask for your password to install the application system-wide.

```bash
# Make sure install.sh and the 'ntn' binary are in the same directory
chmod +x install.sh
sudo ./install.sh
```

### Windows

Open **PowerShell as Administrator** and run the following command:

```powershell
# Make sure install.ps1 and 'ntn.exe' are in the same directory
Set-ExecutionPolicy Bypass -Scope Process -Force; .\install.ps1
```

## Usage

Once installed, the `ntn` command will be available globally in your terminal.

The general structure of the commands is `ntn calendar <subcommand> [flags]`.

### Authentication (First Use)

The first time you run a command that requires access to Google Calendar, an authentication process will be initiated:
1.  A tab will open in your browser asking you to log in and authorize access to your calendar.
2.  After authorizing, you will be redirected, and the CLI will save an access token in your user directory for future sessions.

### Commands

- `ntn calendar login`: Authenticates with your Google account.
- `ntn calendar logout`: Removes the stored authentication token.
- `ntn calendar add`: Creates a new event.
- `ntn calendar list`: Lists all upcoming events.
- `ntn calendar put`: Updates an existing event.
- `ntn calendar rm`: Deletes one or more events.
- `ntn calendar update`: Forces an update of the local event cache.

For detailed flags for each command, you can use `ntn calendar <subcommand> --help`.

## Uninstallation

To completely remove `ntn` and its background service from your system, use the corresponding uninstaller script.

### Linux & macOS

```bash
# Make sure uninstall.sh is in the current directory
chmod +x uninstall.sh
sudo ./uninstall.sh
```

### Windows

*A dedicated uninstaller script for Windows (`uninstall.ps1`) can be created to remove the application, PATH entries, and the scheduled task.*