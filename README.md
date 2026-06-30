# \# Focal-Agent Inventory Plugin

# 

# > A read-only system inventory plugin for Focal-Agent written in Go.

# 

# \*\*Status:\*\* 🚧 Active Development

# 

# Focal-Agent Inventory Plugin collects hardware, operating system, storage, networking, service, and process information from a local Windows system and exposes it through Focal-Agent's JSON plugin protocol.

# 

# The project is designed to provide AI agents with structured, read-only access to system information while maintaining a secure and language-independent plugin architecture.

# 

# \---

# 

# \## Features

# 

# Current and planned capabilities include:

# 

# \- System inventory

# \- CPU information

# \- Memory information

# \- Operating system details

# \- Disk inventory

# \- Network adapter information

# \- Running processes

# \- Windows Services

# \- JSON-based communication

# \- Read-only operation

# 

# \---

# 

# \## Plugin Protocol

# 

# The plugin communicates with Focal-Agent using JSON over standard input and output.

# 

# Example request:

# 

# ```json

# {

# &#x20; "method": "inventory.collect"

# }

# ```

# 

# Example response:

# 

# ```json

# {

# &#x20; "success": true,

# &#x20; "data": {

# &#x20;   "hostname": "DESKTOP-12345",

# &#x20;   "os": "Windows 11",

# &#x20;   "cpu": "Intel Core i9-13900K",

# &#x20;   "memory\_gb": 64

# &#x20; }

# }

# ```

# 

# \---

# 

# \## Supported Commands

# 

# | Command | Description |

# |----------|-------------|

# | `inventory.collect` | Collect complete system inventory |

# | `disks.list` | List local disks and usage |

# | `network.list` | List network adapters and addresses |

# | `processes.list` | List running processes |

# | `services.list` | List Windows services |

# 

# Additional commands will be added as the project evolves.

# 

# \---

# 

# \## Goals

# 

# \- Read-only by default

# \- Fast startup

# \- Minimal dependencies

# \- Simple JSON protocol

# \- Easy to extend

# \- Cross-language compatibility

# \- Cross-platform architecture

# 

# \---

# 

# \## Project Structure

# 

# ```

# cmd/

# internal/

# docs/

# README.md

# LICENSE

# go.mod

# ```

# 

# \---

# 

# \## Building

# 

# ```bash

# go build

# ```

# 

# or

# 

# ```bash

# go build -o focal-agent-inventory.exe

# ```

# 

# \---

# 

# \## Running

# 

# Example:

# 

# ```bash

# echo '{"method":"inventory.collect"}' | focal-agent-inventory.exe

# ```

# 

# \---

# 

# \## Future Enhancements

# 

# \- Linux support

# \- macOS support

# \- Hardware sensors

# \- Event log access

# \- Installed software inventory

# \- GPU information

# \- Security configuration

# \- Windows Updates

# \- Performance counters

# \- Plugin version negotiation

# 

# \---

# 

# \## Relationship to Focal-Agent

# 

# This repository is one component of the Focal-Agent ecosystem.

# 

# Focal-Agent launches the plugin as an external process, sends JSON requests through standard input, and receives structured JSON responses through standard output.

# 

# This language-independent protocol allows plugins to be implemented in Go, Rust, C#, Python, or other languages while presenting a consistent interface to the AI agent.

# 

# \---

# 

# \## Contributing

# 

# Contributions, bug reports, and feature requests are welcome.

# 

# \---

# 

# \## License

# 

# MIT License

