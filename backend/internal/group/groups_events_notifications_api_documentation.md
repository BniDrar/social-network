
# Groups, Events & Notifications Module API Documentation

This document outlines the API routes and handlers available for the **Groups, Events, and Notifications** module.

### Overview

The Groups module allows users to manage groups, send and respond to invitations, and interact with events. Events can be created, voted on, and viewed. Notifications are handled for group invitations, membership requests, and event interactions.

This README includes the following sections:
1. **Groups Endpoints**
2. **Events Endpoints**
3. **Notifications Endpoints**

Each section contains a table of routes, HTTP methods, and the corresponding handler functions.

## Groups Endpoints

| **Route**                         | **Handler**             | **HTTP Method** | **Role** | **Description**                          |
|------------------------------------|-------------------------|-----------------|----------|------------------------------------------|
| `/api/user/groups`                | `GetUserGroups`         | `GET`           | `User`   | Get a list of groups for a user.         |
| `/api/group`                      | `GetGroupById`          | `GET`           | `User`   | Get a specific group by ID.              |
| `/api/group/create`               | `CreateGroup`           | `POST`          | `User`   | Create a new group.                      |
| `/api/groups`                     | `GetAllGroups`          | `GET`           | `User`   | Get all available groups.                |
| `/api/group/members`              | `GetGroupMembers`       | `GET`           | `User`   | Get members of a specific group.         |
| `/api/group/invite/response`      | `InvitationResponse`    | `POST`          | `User`   | Respond to a group invitation.           |
| `/api/group/invite/request`       | `InviteToJoinGroup`     | `POST`          | `User`   | Invite a user to join a group.           |
| `/api/group/join/request`         | `RequestToJoinGroup`    | `POST`          | `User`   | Request to join a group.                 |
| `/api/group/join/response`        | `RequestToJoinResponse` | `POST`          | `User`   | Respond to a group join request.         |

## Events Endpoints

| **Route**                         | **Handler**             | **HTTP Method** | **Role** | **Description**                          |
|------------------------------------|-------------------------|-----------------|----------|------------------------------------------|
| `/api/event/create`               | `CreateEvent`           | `POST`          | `User`   | Create a new event.                      |
| `/api/event/vote`                 | `VoteEvent`             | `POST`          | `User`   | Vote for an event.                       |
| `/api/event/get`                  | `GetEvent`              | `GET`           | `User`   | Get details of an event.                 |

## Notifications Endpoints

| **Route**                         | **Handler**             | **HTTP Method** | **Role** | **Description**                          |
|------------------------------------|-------------------------|-----------------|----------|------------------------------------------|
| `/api/group/invite/response`      | `InvitationResponse`    | `POST`          | `User`   | Respond to group invitation notifications.|
| `/api/group/join/response`        | `RequestToJoinResponse` | `POST`          | `User`   | Respond to a group join request.         |

### Conclusion

This module covers the core functionalities for managing groups, handling events, and processing notifications. The routes outlined here should help frontend developers integrate the backend logic with the user interface.
