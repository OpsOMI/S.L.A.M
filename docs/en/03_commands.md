# Commands

Below is the list of commands that can be used by both users and administrators in the S.L.A.M client.

## General Commands (For All Users)

| Command                         | Description                                                                                         | Example                    |
| ------------------------------- | --------------------------------------------------------------------------------------------------- | -------------------------- |
| `/login`                        | Used for user login.                                                                                | `/login`                   |
| `/room/create [true/false]`     | Creates a new room. If `true`, creates a password-protected room; if `false`, creates an open room. | `/room/create true`        |
| `/room/list`                    | Lists all rooms owned by the user.                                                                  | `/room/list`               |
| `/room/join [code] [password?]` | Joins the specified room. If it's password-protected, enter the password; leave blank if not.       | `/room/join publicroom123` |
| `/room/clean`                   | Instantly deletes all messages in a room owned by the user.                                         | `/room/clean`              |
| `/me`                           | Displays the user's own information (nickname, username).                                           | `/me`                      |
| `/reconnect`                    | Attempts to reconnect to the server if the connection is lost.                                      | `/reconnect`               |
| `/clear`                        | Clears the terminal screen.                                                                         | `/clear`                   |
| `/logout`                       | Logs the user out.                                                                                  | `/logout`                  |
| `/exit`, `/quit`                | Completely exits the application.                                                                   | `/exit` or `/quit`         |

💬 **Note:** If you are inside a room, you can send a message directly by typing plain text without using a command.

## Administrator Commands (Requires Owner Privileges)

| Command     | Description                                                                                     | Example     |
| ----------- | ----------------------------------------------------------------------------------------------- | ----------- |
| `/register` | Creates a new user. A client binary file is automatically generated (`client/[username]/main`). | `/register` |
| `/online`   | Lists the number of currently connected users on the server.                                    | `/online`   |

[← Back](./02_features.md)
