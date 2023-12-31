![Logo](docs/pictures/logo.png)

[![License: GPLv3](https://img.shields.io/badge/License-GPLv3-green.svg)](https://www.gnu.org/licenses/gpl-3.0.en.html)

## Notification: Slowdown in the development process

Dear ApisOps Community,

We apologise for the current slowdown in the development of the ApisOps project. The lead developer (me, Ada), who is a student pursuing a Master's degree, is navigating a particularly demanding schedule at the moment, impacting the project's progress. 

We appreciate your understanding, patience, and continued support as we work through this phase. For questions or additional contributions, please don’t hesitate to reach out.

Best regards,

Ada

## Introduction

ApisOps is an effective, open-source, high-performance operations automation system tailored for SRE use, adept at handling massive concurrent tasks. Lightweight and extensible, ApisOps offers a practical tool for streamlined integration, user-friendly operation, and extensive adaptability.

## Architecture

![Architecture of ApisOps](docs/pictures/architecture.png)

ApisOps is developed in multiple languages including Go and TypeScript with a simple, asynchronized, distributed, Event Loop-based architecture.

### Front-end

Built by **React.js**, **TypeScript**, **Ant Design Pro** and **ProComponents**.

### Back-end

Written in **Go**, the back-end services are composed of two separate services: **Concourse** and **Interchange**.

**Concourse**: The entry point of ApisOps, a simple web server based on the MVC model, providing RESTful APIs to manage workflows, scripts, tasks, and cron jobs.

![Architecture of Concourse](docs/pictures/architecture_of_Concourse.png)

**Interchange**: A high-performance execution service for submitting tasks such as script execution, requesting third-party systems, and file distribution. It also retrieves task logs, converts task status, and persists execution records.

![Architecture of Interchange](docs/pictures/architecture_of_Interchange.png)

### Client on remote servers

**Station** is a lightweight client that runs on target servers for dispatching files and executing scripts.

### Design considerations

- High-performance
- Lightweight
- Efficient
- Easy to extend
- Keep it simple
- Easy to use
