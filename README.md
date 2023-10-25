# ApisOps

ApisOps is an effective operations automation system tailored for SRE use, adept at handling massive concurrent tasks.

## Introduction

ApisOps is an effective, open-source, high-performance operations automation system tailored for SRE use, adept at handling massive concurrent tasks. Lightweight and extensible, ApisOps offers a practical tool for streamlined integration, user-friendly operation, and extensive adaptability.

ApisOps is composed of these components:

- **Concourse**: A back-end server acting as an entry point of operations.
- **Interchange**: A high-performance execution service for submitting tasks such as script execution, requesting third-party systems, and file distribution. It also retrieves task logs, converts task status, and persists execution records.
- **Station**: A lightweight agent that runs on target servers.

## Architecture

![Architecture of ApisOps](docs/pictures/architecture.png)

## License

[GNU General Public License v2](https://www.gnu.org/licenses/old-licenses/gpl-2.0.en.html).
