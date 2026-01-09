# VoIPrax 🚀

**VoIP Penetration & Analysis eXtreme toolkit**

VoIPrax is a modern, high-performance VoIP penetration testing and analysis toolkit written in Go. It is designed to be modular, extensible, and suitable for both PoC and real-world security audits.

## ✨ Features

- **SIP Protocol Analysis**: RFC 3261 compliant message parsing and validation.
- **Advanced Fuzzing**: Modulal SIP message fuzzing to discover vulnerabilities.
- **Brute-force & Enumeration**: Efficient REGISTER and INVITE brute-force modules.
- **Modern Architecture**: Clean, modular Go code with structured logging.
- **CLI & REST API**: Flexible interfaces for manual testing and automation.
- **Reporting**: JSON and HTML reporting outputs (WIP).

## 🛠️ Installation

```bash
git clone https://github.com/ismailtsdln/VoIPrax.git
cd VoIPrax
go build -o voiprax ./cmd/voiprax
```

## 🚀 Usage

### CLI Help
```bash
./voiprax --help
```

### SIP Fuzzing
```bash
./voiprax fuzz --target 192.168.1.1:5060 --count 1000
```

## 📂 Project Structure

- `cmd/voiprax`: CLI entry point and commands.
- `internal/sip`: SIP protocol parser and stack.
- `internal/fuzz`: Fuzzing logic and templates.
- `internal/exploit`: Exploit modules (Brute-force, etc.).
- `api/`: REST API implementation.

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## ⚖️ License

MIT License - see LICENSE for details.
