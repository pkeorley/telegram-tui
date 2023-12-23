# telegram-tui

![Static Badge](https://img.shields.io/badge/Go-1.21-blue?logo=go)
![GitHub License](https://img.shields.io/github/license/pkeorley/telegram-tui)
![GitHub issues](https://img.shields.io/github/issues/pkeorley/telegram-tui)
![GitHub commit activity (branch)](https://img.shields.io/github/commit-activity/y/pkeorley/telegram-tui)
![GitHub contributors](https://img.shields.io/github/contributors/pkeorley/telegram-tui)

-----

**Table of Contents**

- [Installation](#installation)
- [License](#license)
- [Architecture](#architecture)

## Installation

```console
go build .
```

## License

`telegram-tui` is distributed under the terms of the [MIT](https://spdx.org/licenses/MIT.html) license.

## Architecture

`telegram-tui` has the following architecture (for visualisation, [mermaid](https://github.com/mermaid-js/mermaid) is used)

```mermaid
---
title: Log-in
---
flowchart LR

    subgraph fa:fa-file Log-in environment
        L[Log in]
        S{"`**Successful**`"}
    end

    C[(database.go)]
    L --> |login.go| S

    S --> |fa:fa-check true| C
    S --> |fa:fa-ban false| L
```

```mermaid
---
title: Interaction
---
flowchart LR
    T[fa:fa-telegram telegram.go]
    C[(database.go)]
    
    subgraph Login
        T --> |Get login data| C
        C --> |Login data| T
    end
    
    subgraph fa:fa-file Work environment
        T --> |fa:fa-check| X
    end
```
    