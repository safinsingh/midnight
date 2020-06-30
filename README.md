<p align="center">
  <img src="./assets/logo.png" style="
    border-radius: 10px;
  "/>
</p>

<h3 align="center">
  <i>
    An extensible security auditing tool
  </i>
</h3>

<p align="center">
  <img alt="Version" src="https://img.shields.io/badge/Version-1.0-red.svg" />
  <img alt="License: MIT" src="https://img.shields.io/badge/License-GPLv3-orange.svg" />
  <img alt="Made with Go" src="https://img.shields.io/badge/Made%20with-Go-yellow.svg" />
  <img alt="PRs Welcome" src="https://img.shields.io/badge/PRs-Welcome-brightgreen.svg">
  <img alt="Maintained" src="https://img.shields.io/badge/Maintained-Yes-blue">
  <img alt="awesome" src="https://img.shields.io/badge/Awesome-Yes-blueviolet">
</p>

<hr>

## 💡 Features

- System audit to report vulnerabilities
- Based on STIGs and (in the future) CIS Controls
- User can specify custom controls in JSON
- Ability to run on virtual machines with Docker

## ✨ In Action

<p align="center">
  <img src="./assets/screen.png" />
</p>

## 💻 Try it out

```sh
git clone https://github.com/safinsingh/midnight.git
bash install.sh
go build .
./midnight -file checks/u16stig.json
```

## 🔮 Usage

```sh
./midnight -h

# output

Usage of ./midnight:
  -file string
        Configuration file to use (mandatory)
  -mode string
        Mode to run midnight in. Possible modes: audit, enforce, docker (default "audit")

```

## 👨‍💻 Author

Linkedin: [Safin Singh](https://www.linkedin.com/in/safin-singh-b2630918a/) <br>
GitHub: [safinsingh](https://github.com/safinsingh) <br>
Dribbble: [Safin Singh](https://dribbble.com/safinsingh/) <br>
YouTube: [Safin Singh](https://www.youtube.com/channel/UCvb01sUdAgcPAG1j0SLxAtA)

## 🤝 Contributing

Contributions, PRs, issues and feature requests are welcome! Feel free to check out my [issues page](https://github.com/safinsingh/midnight/issues).

## ❤️ Show your support

Give a ⭐️ if this project helped you!
Hope you enjoy it!
