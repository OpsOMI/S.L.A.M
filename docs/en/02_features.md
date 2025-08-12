# Features

The S.L.A.M (Secure Link Anonymous Messaging) project is developed to meet the need for privacy- and security-focused anonymous messaging. Below are the key features of the project:

## General Features

- **Secure TCP-based communication:**
  Connects to a remote server over TCP. Communication is encrypted using TLS (Transport Layer Security).

- **User-specific client binary:**
  A unique client executable, dedicated to each user, is automatically generated.

- **Client-user pairing:**
  Since the client binary belongs only to the related user, even if others have user credentials, they cannot log in without the client.

- **USB and portable device support:**
  The client application can easily run from portable media like USB drives without any additional installation.

- **Administrator-controlled client binary creation:**
  Only the administrator can create new clients when adding new users.

- **Room-based communication:**
  Users are organized into rooms, which can be public or private. The private room password is set by the system administrator.

- **24-hour message storage and automatic deletion:**
  Messages are stored in the database for 24 hours and then automatically deleted, enabling communication without leaving traces.

## Security Features

- **End-to-end encryption with TLS:**
  All network communication is encrypted via TLS, protecting against eavesdropping and tampering.

- **JWT-based authentication:**
  User authentication is secured using JWT tokens.

- **Encrypted messages in the database:**
  Messages are stored encrypted within the database.

## Usage and Administration

- **Automatic admin and client creation:**
  Upon server startup, an "owner" role admin user and its associated client are automatically created based on admin credentials.

- **Automatic public and private rooms for admin:**
  Two rooms (public and private) are automatically created for the admin account. The private room’s password is set by the admin.

- **Easy installation and deployment:**
  Can be easily installed and run using Docker and Docker Compose, and operates well on portable environments.

[← Back](./01_installation.md)   |   [Next →](./03_features.md)
