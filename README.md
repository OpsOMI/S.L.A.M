# S.L.A.M — Secure Link Anonymous Messaging

S.L.A.M is a TCP-based communication system developed for privacy-focused, anonymous messaging. The system consists of a dedicated TCP server running on a remote host and clients that securely connect to this server.

For each user, a **unique and dedicated client** is generated. Even if someone else has both their own login credentials and another client, they cannot access the system, because the client-user mapping is unique.

When a new user is added by the system administrator, a client executable specific to that user is automatically created. This client is the only way to access the server.

S.L.A.M is designed to run from **USB drives** or other portable devices and can be used especially in environments requiring high security.

## Purpose

The goal of S.L.A.M is to enable users to communicate securely, anonymously, and quickly. It is designed for sensitive operations where communication must happen without leaving traces and under centralized control.

## Documentation

- [Installation](docs/en/01_installation.md)
- [Features](docs/en/02_features.md)
- [Commands](docs/en/03_commands.md)

## Support Us

If you like our project, don't forget to give it a ⭐️ **Star**!. This motivates us and helps the project reach more people. Thank you! 🙌

## Languages

[Turkish](README.tr.md)
