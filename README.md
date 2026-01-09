# VoIPrax 🚀

<p align="center">
  <img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go">
  <img src="https://img.shields.io/badge/License-MIT-green?style=for-the-badge" alt="License">
  <img src="https://img.shields.io/badge/Version-1.0.0-blue?style=for-the-badge" alt="Version">
</p>

**VoIPrax** (VoIP Penetration & Analysis eXtreme) is a modern, modular, and high-performance toolkit designed for VoIP security professionals. Written in Go, it offers a robust foundation for SIP protocol analysis, fuzzing, and automated security testing.

## 🎯 Vision

VoIPrax aims to modernize VoIP security testing by providing a portable, extensible, and developer-friendly alternative to legacy tools. It focuses on modern protocol compliance, rich user experience, and automation-first design.

## ✨ Key Features

- 🛡️ **SIP Protocol Engine**: Full RFC 3261 compliance with deep packet inspection and manipulation.
- ⚡ **SIP Fuzzing**: Advanced fuzzing modules for headers, methods, and payloads to discover zero-day vulnerabilities.
- 🔑 **Brute-force & Enumeration**: High-speed REGISTER and INVITE brute-force capabilities.
- 📟 **Stunning CLI**: A premium terminal experience with stylized banners, colorized logs, and real-time progress indicators.
- 🔌 **Modular Architecture**: Easy to extend with custom plugins and security modules.
- 🌐 **REST API**: Built-in API for seamless integration into CI/CD pipelines and security orchestration.

## 🛠️ Installation

### Prerequisites
- Go 1.21 or higher

### Build from Source
```bash
git clone https://github.com/ismailtsdln/VoIPrax.git
cd VoIPrax
go build -o voiprax ./cmd/voiprax
```

## 🚀 Quick Start

### Display Help
```bash
./voiprax --help
```

### Start a Fuzzing Session
```bash
./voiprax fuzz --target 192.168.1.50:5060 --count 1000 --verbose
```

### Run the API Server
```bash
# Future implementation for API control
go run cmd/voiprax/main.go server
```

## 📂 Project Structure

```text
VoIPrax/
├── api/             # REST API implementation
├── cmd/             # CLI entry points and commands
├── internal/
│   ├── sip/         # core SIP protocol parser & stack
│   ├── fuzz/        # fuzzing logic and templates
│   ├── exploit/     # security testing modules
│   ├── ui/          # CLI UX and styling
│   └── logger/      # structured logging
├── docs/            # technical documentation
└── tests/           # unit and integration tests
```

## 🤝 Roadmap

- [ ] Interactive SIP/TLS Proxy
- [ ] Automated HTML/JSON Reporting
- [ ] Support for H.323 and MGCP
- [ ] Web-based Dashboard for Real-time Monitoring

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<p align="center">
Developed with ❤️ by <a href="https://github.com/ismailtsdln">Ismail Tasdelen</a>
</p>
