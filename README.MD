<p align="left"><h1 align="left">Authentication API (gRPC/HTTP) </h1></p>

An authentication API for user management and authorization supporting both **gRPC** and **HTTP** protocols

<p align="left">
	<img src="https://img.shields.io/github/license/fsobh/auth?style=default&logo=opensourceinitiative&logoColor=white&color=00fff2" alt="license">
	<img src="https://img.shields.io/github/last-commit/fsobh/auth?style=default&logo=git&logoColor=white&color=00fff2" alt="last-commit">
	<img src="https://img.shields.io/github/languages/top/fsobh/auth?style=default&color=00fff2" alt="repo-top-language">
	<img src="https://img.shields.io/github/languages/count/fsobh/auth?style=default&color=00fff2" alt="repo-language-count">
</p>
<p align="left"><!-- default option, no dependency badges. -->
</p>
<p align="left">
	<!-- default option, no dependency badges. -->
</p>
<br>

## 🔗 Table of Contents

- [📍 Overview](#-overview)
- [👾 Features](#-features)
- [📁 Project Structure](#-project-structure)
    - [📂 Project Index](#-project-index)
- [🚀 Getting Started](#-getting-started)
    - [☑️ Prerequisites](#-prerequisites)
    - [⚙️ Installation](#-installation)
    - [🤖 Usage](#🤖-usage)
    - [🧪 Testing](#🧪-testing)
- [📌 Project Roadmap](#-project-roadmap)
- [🔰 Contributing](#-contributing)
- [🎗 License](#-license)
- [🙌 Acknowledgments](#-acknowledgments)

---

## 📍 Overview

<code>❯ REPLACE-ME</code>

---

## 👾 Features

<code>❯ REPLACE-ME</code>

---

## 📁 Project Structure

```sh
└── auth/
    ├── db
    │   ├── migration
    │   ├── query
    │   └── sqlc
    ├── doc
    │   ├── db.dbml
    │   ├── statik
    │   └── swagger
    ├── gapi
    │   ├── authorization.go
    │   ├── converter.go
    │   ├── error.go
    │   ├── logger.go
    │   ├── metadata.go
    │   ├── rpc_create_user.go
    │   ├── rpc_login_user.go
    │   ├── rpc_update_user.go
    │   ├── rpc_verify_email.go
    │   └── server.go
    ├── go.mod
    ├── go.sum
    ├── main.go
    ├── makefile
    ├── pb
    │   ├── rpc_create_user.pb.go
    │   ├── rpc_login_user.pb.go
    │   ├── rpc_update_user.pb.go
    │   ├── rpc_verify_email.pb.go
    │   ├── service_auth.pb.go
    │   ├── service_auth.pb.gw.go
    │   ├── service_auth_grpc.pb.go
    │   └── user.pb.go
    ├── proto
    │   ├── google
    │   ├── protoc-gen-openapiv2
    │   ├── rpc_create_user.proto
    │   ├── rpc_login_user.proto
    │   ├── rpc_update_user.proto
    │   ├── rpc_verify_email.proto
    │   ├── service_auth.proto
    │   └── user.proto
    ├── sqlc.yaml
    ├── util
    │   ├── config.go
    │   ├── password.go
    │   └── random.go
    ├── val
    │   └── validator.go
    └── worker
        ├── distributor.go
        ├── logger.go
        ├── processor.go
        └── task_send_verify_email.go
```


### 📂 Project Index
<details open>
	<summary><b><code>AUTH/</code></b></summary>
	<details> <!-- __root__ Submodule -->
		<summary><b>__root__</b></summary>
		<blockquote>
			<table>
			<tr>
				<td><b><a href='https://github.com/fsobh/auth/blob/master/makefile'>makefile</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/fsobh/auth/blob/master/main.go'>main.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/fsobh/auth/blob/master/go.mod'>go.mod</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/fsobh/auth/blob/master/go.sum'>go.sum</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/fsobh/auth/blob/master/sqlc.yaml'>sqlc.yaml</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			</table>
		</blockquote>
	</details>
	<details> <!-- worker Submodule -->
		<summary><b>worker</b></summary>
		<blockquote>
			<table>
			<tr>
				<td><b><a href='https://github.com/fsobh/auth/blob/master/worker/processor.go'>processor.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/fsobh/auth/blob/master/worker/task_send_verify_email.go'>task_send_verify_email.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/fsobh/auth/blob/master/worker/logger.go'>logger.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/fsobh/auth/blob/master/worker/distributor.go'>distributor.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			</table>
		</blockquote>
	</details>
	<details> <!-- proto Submodule -->
		<summary><b>proto</b></summary>
		<blockquote>
			<table>
			<tr>
				<td><b><a href='https://github.com/fsobh/auth/blob/master/proto/user.proto'>user.proto</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/fsobh/auth/blob/master/proto/rpc_update_user.proto'>rpc_update_user.proto</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/fsobh/auth/blob/master/proto/rpc_create_user.proto'>rpc_create_user.proto</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/fsobh/auth/blob/master/proto/service_auth.proto'>service_auth.proto</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/fsobh/auth/blob/master/proto/rpc_verify_email.proto'>rpc_verify_email.proto</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/fsobh/auth/blob/master/proto/rpc_login_user.proto'>rpc_login_user.proto</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			</table>
			<details>
				<summary><b>protoc-gen-openapiv2</b></summary>
				<blockquote>
					<details>
						<summary><b>options</b></summary>
						<blockquote>
							<table>
							<tr>
								<td><b><a href='https://github.com/fsobh/auth/blob/master/proto/protoc-gen-openapiv2/options/openapiv2.proto'>openapiv2.proto</a></b></td>
								<td><code>❯ REPLACE-ME</code></td>
							</tr>
							<tr>
								<td><b><a href='https://github.com/fsobh/auth/blob/master/proto/protoc-gen-openapiv2/options/annotations.proto'>annotations.proto</a></b></td>
								<td><code>❯ REPLACE-ME</code></td>
							</tr>
							</table>
						</blockquote>
					</details>
				</blockquote>
			</details>
			<details>
				<summary><b>google</b></summary>
				<blockquote>
					<details>
						<summary><b>api</b></summary>
						<blockquote>
							<table>
							<tr>
								<td><b><a href='https://github.com/fsobh/auth/blob/master/proto/google/api/httpbody.proto'>httpbody.proto</a></b></td>
								<td><code>❯ REPLACE-ME</code></td>
							</tr>
							<tr>
								<td><b><a href='https://github.com/fsobh/auth/blob/master/proto/google/api/field_behavior.proto'>field_behavior.proto</a></b></td>
								<td><code>❯ REPLACE-ME</code></td>
							</tr>
							<tr>
								<td><b><a href='https://github.com/fsobh/auth/blob/master/proto/google/api/http.proto'>http.proto</a></b></td>
								<td><code>❯ REPLACE-ME</code></td>
							</tr>
							<tr>
								<td><b><a href='https://github.com/fsobh/auth/blob/master/proto/google/api/annotations.proto'>annotations.proto</a></b></td>
								<td><code>❯ REPLACE-ME</code></td>
							</tr>
							</table>
						</blockquote>
					</details>
				</blockquote>
			</details>
		</blockquote>
	</details>
	<details> <!-- doc Submodule -->
		<summary><b>doc</b></summary>
		<blockquote>
			<table>
			<tr>
				<td><b><a href='https://github.com/fsobh/auth/blob/master/doc/db.dbml'>db.dbml</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			</table>
			<details>
				<summary><b>statik</b></summary>
				<blockquote>
					<table>
					<tr>
						<td><b><a href='https://github.com/fsobh/auth/blob/master/doc/statik/statik.go'>statik.go</a></b></td>
						<td><code>❯ REPLACE-ME</code></td>
					</tr>
					</table>
				</blockquote>
			</details>
			<details>
				<summary><b>swagger</b></summary>
				<blockquote>
					<table>
					<tr>
						<td><b><a href='https://github.com/fsobh/auth/blob/master/doc/swagger/swagger-initializer.js'>swagger-initializer.js</a></b></td>
						<td><code>❯ REPLACE-ME</code></td>
					</tr>
					<tr>
						<td><b><a href='https://github.com/fsobh/auth/blob/master/doc/swagger/swagger-ui-standalone-preset.js'>swagger-ui-standalone-preset.js</a></b></td>
						<td><code>❯ REPLACE-ME</code></td>
					</tr>
					<tr>
						<td><b><a href='https://github.com/fsobh/auth/blob/master/doc/swagger/auth.swagger.json'>auth.swagger.json</a></b></td>
						<td><code>❯ REPLACE-ME</code></td>
					</tr>
					<tr>
						<td><b><a href='https://github.com/fsobh/auth/blob/master/doc/swagger/swagger-ui-es-bundle.js'>swagger-ui-es-bundle.js</a></b></td>
						<td><code>❯ REPLACE-ME</code></td>
					</tr>
					<tr>
						<td><b><a href='https://github.com/fsobh/auth/blob/master/doc/swagger/swagger-ui-bundle.js'>swagger-ui-bundle.js</a></b></td>
						<td><code>❯ REPLACE-ME</code></td>
					</tr>
					<tr>
						<td><b><a href='https://github.com/fsobh/auth/blob/master/doc/swagger/index.css'>index.css</a></b></td>
						<td><code>❯ REPLACE-ME</code></td>
					</tr>
					<tr>
						<td><b><a href='https://github.com/fsobh/auth/blob/master/doc/swagger/swagger-ui-es-bundle-core.js'>swagger-ui-es-bundle-core.js</a></b></td>
						<td><code>❯ REPLACE-ME</code></td>
					</tr>
					<tr>
						<td><b><a href='https://github.com/fsobh/auth/blob/master/doc/swagger/swagger-ui.js'>swagger-ui.js</a></b></td>
						<td><code>❯ REPLACE-ME</code></td>
					</tr>
					<tr>
						<td><b><a href='https://github.com/fsobh/auth/blob/master/doc/swagger/swagger-ui.css'>swagger-ui.css</a></b></td>
						<td><code>❯ REPLACE-ME</code></td>
					</tr>
					<tr>
						<td><b><a href='https://github.com/fsobh/auth/blob/master/doc/swagger/index.html'>index.html</a></b></td>
						<td><code>❯ REPLACE-ME</code></td>
					</tr>
					<tr>
						<td><b><a href='https://github.com/fsobh/auth/blob/master/doc/swagger/oauth2-redirect.html'>oauth2-redirect.html</a></b></td>
						<td><code>❯ REPLACE-ME</code></td>
					</tr>
					</table>
				</blockquote>
			</details>
		</blockquote>
	</details>
	<details> <!-- gapi Submodule -->
		<summary><b>gapi</b></summary>
		<blockquote>
			<table>
			<tr>
				<td><b><a href='https://github.com/fsobh/auth/blob/master/gapi/metadata.go'>metadata.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/fsobh/auth/blob/master/gapi/logger.go'>logger.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/fsobh/auth/blob/master/gapi/authorization.go'>authorization.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/fsobh/auth/blob/master/gapi/rpc_update_user.go'>rpc_update_user.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/fsobh/auth/blob/master/gapi/rpc_verify_email.go'>rpc_verify_email.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/fsobh/auth/blob/master/gapi/error.go'>error.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/fsobh/auth/blob/master/gapi/rpc_login_user.go'>rpc_login_user.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/fsobh/auth/blob/master/gapi/converter.go'>converter.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/fsobh/auth/blob/master/gapi/rpc_create_user.go'>rpc_create_user.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/fsobh/auth/blob/master/gapi/server.go'>server.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			</table>
		</blockquote>
	</details>
	<details> <!-- val Submodule -->
		<summary><b>val</b></summary>
		<blockquote>
			<table>
			<tr>
				<td><b><a href='https://github.com/fsobh/auth/blob/master/val/validator.go'>validator.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			</table>
		</blockquote>
	</details>
	<details> <!-- pb Submodule -->
		<summary><b>pb</b></summary>
		<blockquote>
			<table>
			<tr>
				<td><b><a href='https://github.com/fsobh/auth/blob/master/pb/user.pb.go'>user.pb.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/fsobh/auth/blob/master/pb/service_auth_grpc.pb.go'>service_auth_grpc.pb.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/fsobh/auth/blob/master/pb/rpc_login_user.pb.go'>rpc_login_user.pb.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/fsobh/auth/blob/master/pb/rpc_update_user.pb.go'>rpc_update_user.pb.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/fsobh/auth/blob/master/pb/rpc_create_user.pb.go'>rpc_create_user.pb.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/fsobh/auth/blob/master/pb/service_auth.pb.go'>service_auth.pb.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/fsobh/auth/blob/master/pb/service_auth.pb.gw.go'>service_auth.pb.gw.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/fsobh/auth/blob/master/pb/rpc_verify_email.pb.go'>rpc_verify_email.pb.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			</table>
		</blockquote>
	</details>
	<details> <!-- util Submodule -->
		<summary><b>util</b></summary>
		<blockquote>
			<table>
			<tr>
				<td><b><a href='https://github.com/fsobh/auth/blob/master/util/password.go'>password.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/fsobh/auth/blob/master/util/config.go'>config.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			<tr>
				<td><b><a href='https://github.com/fsobh/auth/blob/master/util/random.go'>random.go</a></b></td>
				<td><code>❯ REPLACE-ME</code></td>
			</tr>
			</table>
		</blockquote>
	</details>
	<details> <!-- db Submodule -->
		<summary><b>db</b></summary>
		<blockquote>
			<details>
				<summary><b>sqlc</b></summary>
				<blockquote>
					<table>
					<tr>
						<td><b><a href='https://github.com/fsobh/auth/blob/master/db/sqlc/models.go'>models.go</a></b></td>
						<td><code>❯ REPLACE-ME</code></td>
					</tr>
					<tr>
						<td><b><a href='https://github.com/fsobh/auth/blob/master/db/sqlc/db.go'>db.go</a></b></td>
						<td><code>❯ REPLACE-ME</code></td>
					</tr>
					<tr>
						<td><b><a href='https://github.com/fsobh/auth/blob/master/db/sqlc/verify_email.sql.go'>verify_email.sql.go</a></b></td>
						<td><code>❯ REPLACE-ME</code></td>
					</tr>
					<tr>
						<td><b><a href='https://github.com/fsobh/auth/blob/master/db/sqlc/user.sql.go'>user.sql.go</a></b></td>
						<td><code>❯ REPLACE-ME</code></td>
					</tr>
					<tr>
						<td><b><a href='https://github.com/fsobh/auth/blob/master/db/sqlc/querier.go'>querier.go</a></b></td>
						<td><code>❯ REPLACE-ME</code></td>
					</tr>
					<tr>
						<td><b><a href='https://github.com/fsobh/auth/blob/master/db/sqlc/tx_verify_email.go'>tx_verify_email.go</a></b></td>
						<td><code>❯ REPLACE-ME</code></td>
					</tr>
					<tr>
						<td><b><a href='https://github.com/fsobh/auth/blob/master/db/sqlc/store.go'>store.go</a></b></td>
						<td><code>❯ REPLACE-ME</code></td>
					</tr>
					<tr>
						<td><b><a href='https://github.com/fsobh/auth/blob/master/db/sqlc/tx_create_user.go'>tx_create_user.go</a></b></td>
						<td><code>❯ REPLACE-ME</code></td>
					</tr>
					<tr>
						<td><b><a href='https://github.com/fsobh/auth/blob/master/db/sqlc/sessions.sql.go'>sessions.sql.go</a></b></td>
						<td><code>❯ REPLACE-ME</code></td>
					</tr>
					</table>
				</blockquote>
			</details>
			<details>
				<summary><b>query</b></summary>
				<blockquote>
					<table>
					<tr>
						<td><b><a href='https://github.com/fsobh/auth/blob/master/db/query/user.sql'>user.sql</a></b></td>
						<td><code>❯ REPLACE-ME</code></td>
					</tr>
					<tr>
						<td><b><a href='https://github.com/fsobh/auth/blob/master/db/query/sessions.sql'>sessions.sql</a></b></td>
						<td><code>❯ REPLACE-ME</code></td>
					</tr>
					<tr>
						<td><b><a href='https://github.com/fsobh/auth/blob/master/db/query/verify_email.sql'>verify_email.sql</a></b></td>
						<td><code>❯ REPLACE-ME</code></td>
					</tr>
					</table>
				</blockquote>
			</details>
			<details>
				<summary><b>migration</b></summary>
				<blockquote>
					<table>
					<tr>
						<td><b><a href='https://github.com/fsobh/auth/blob/master/db/migration/000001_init_schema.up.sql'>000001_init_schema.up.sql</a></b></td>
						<td><code>❯ REPLACE-ME</code></td>
					</tr>
					<tr>
						<td><b><a href='https://github.com/fsobh/auth/blob/master/db/migration/000001_init_schema.down.sql'>000001_init_schema.down.sql</a></b></td>
						<td><code>❯ REPLACE-ME</code></td>
					</tr>
					</table>
				</blockquote>
			</details>
		</blockquote>
	</details>
</details>

---
## 🚀 Getting Started

### ☑️ Prerequisites

Before getting started with auth, ensure your runtime environment meets the following requirements:

- **Programming Language:** Go
- **Package Manager:** Go modules


### ⚙️ Installation

Install auth using one of the following methods:

**Build from source:**

1. Clone the auth repository:
```sh
❯ git clone https://github.com/fsobh/auth
```

2. Navigate to the project directory:
```sh
❯ cd auth
```

3. Install the project dependencies:


**Using `go modules`** &nbsp; [<img align="center" src="https://img.shields.io/badge/Go-00ADD8.svg?style={badge_style}&logo=go&logoColor=white" />](https://golang.org/)

```sh
❯ go build
```




### 🤖 Usage
Run auth using the following command:
**Using `go modules`** &nbsp; [<img align="center" src="https://img.shields.io/badge/Go-00ADD8.svg?style={badge_style}&logo=go&logoColor=white" />](https://golang.org/)

```sh
❯ go run {entrypoint}
```


### 🧪 Testing
Run the test suite using the following command:
**Using `go modules`** &nbsp; [<img align="center" src="https://img.shields.io/badge/Go-00ADD8.svg?style={badge_style}&logo=go&logoColor=white" />](https://golang.org/)

```sh
❯ go test ./...
```


---
## 📌 Project Roadmap

- [X] **`Task 1`**: <strike>Implement feature one.</strike>
- [ ] **`Task 2`**: Implement feature two.
- [ ] **`Task 3`**: Implement feature three.

---

## 🔰 Contributing

- **💬 [Join the Discussions](https://github.com/fsobh/auth/discussions)**: Share your insights, provide feedback, or ask questions.
- **🐛 [Report Issues](https://github.com/fsobh/auth/issues)**: Submit bugs found or log feature requests for the `auth` project.
- **💡 [Submit Pull Requests](https://github.com/fsobh/auth/blob/main/CONTRIBUTING.md)**: Review open PRs, and submit your own PRs.

<details closed>
<summary>Contributing Guidelines</summary>

1. **Fork the Repository**: Start by forking the project repository to your github account.
2. **Clone Locally**: Clone the forked repository to your local machine using a git client.
   ```sh
   git clone https://github.com/fsobh/auth
   ```
3. **Create a New Branch**: Always work on a new branch, giving it a descriptive name.
   ```sh
   git checkout -b new-feature-x
   ```
4. **Make Your Changes**: Develop and test your changes locally.
5. **Commit Your Changes**: Commit with a clear message describing your updates.
   ```sh
   git commit -m 'Implemented new feature x.'
   ```
6. **Push to github**: Push the changes to your forked repository.
   ```sh
   git push origin new-feature-x
   ```
7. **Submit a Pull Request**: Create a PR against the original project repository. Clearly describe the changes and their motivations.
8. **Review**: Once your PR is reviewed and approved, it will be merged into the main branch. Congratulations on your contribution!
</details>

<details closed>
<summary>Contributor Graph</summary>
<br>
<p align="left">
   <a href="https://github.com{/fsobh/auth/}graphs/contributors">
      <img src="https://contrib.rocks/image?repo=fsobh/auth">
   </a>
</p>
</details>

---

## 🎗 License

This project is protected under the [SELECT-A-LICENSE](https://choosealicense.com/licenses) License. For more details, refer to the [LICENSE](https://choosealicense.com/licenses/) file.

---

## 🙌 Acknowledgments

- List any resources, contributors, inspiration, etc. here.

---