# Simple TODO CLI for a developer

## Install
Create a config ~/.todo in YAML format and specify the storage location (sqlite):
```yaml
todos_file: <your_sqlite_db_path>
github_token: <your_github_token>
```
```bash
go install .
```
For help, type:
```bash
todo --help
```
