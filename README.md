## Logcheck. A plugin for the golangci linter ##
A modular plugin for checking the formatting rules for log text content

## Plugin installation guide ##
1. Downolad the linter software
#### git repo #####
```sh
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```
#### or apt ####
```sh
sudo apt install golangci-lint
```
2. Clone this repository on your local machine
```sh
git clone https://github.com/Moxcelix/linter
```
3. Create or set up two config files

#### .custom-gcl.yml #### 
defines plugin path
```yaml
version: v2.9.0
plugins:
  - module: 'linters'
    path: .
```

#### .golangci.yml #### 
the main linters config, which includes "settings" section
```yaml
version: "2"
linters:
  default: none
  enable:
    - logcheck

  settings:
    custom:
      logcheck:
        type: module
        description: checks log messages for proper format
        settings:
          rules:
            english-rule:
              enabled: true
            lowercase-rule:
              enabled: true
            secret-rule:
              enabled: true
              words: ["password", "passwd", "token", "secret", "key", "api_key", "apikey"]
            special-rule:
              enabled: true
              chars: ["!", "@", "#", "$", "%", "^", "&", "*"]

          loggers:
            - pkg_name: "go.uber.org/zap"
              logger_obj:
                - name: "Logger"
                  methods: ["Debug", "Info", "Warn", "Error", "DPanic", "Panic", "Fatal",
                            "Debugf", "Infof", "Warnf", "Errorf", "DPanicf", "Panicf", "Fatalf", "Debugw", "Infow",
                            "Warnw", "Errorw", "DPanicw", "Panicw", "Fatalw", "With", "Named", "WithOptions", "Core", "Check", "Sugar"]
                - name: "SugaredLogger"
                  methods: ["Debug", "Info", "Warn", "Error", "DPanic", "Panic", "Fatal", "Debugf", "Infof", "Warnf",
                            "Errorf", "DPanicf", "Panicf", "Fatalf", "Debugw", "Infow", "Warnw", "Errorw", "DPanicw",
                            "Panicw", "Fatalw", "With", "Named", "Desugar"]
              funcs: []

            - pkg_name: "log/slog"
              logger_obj:
                - name: "Logger"
                  methods: ["Debug", "Info", "Warn", "Error", "DebugContext", "InfoContext", "WarnContext",
                            "ErrorContext", "Log", "LogContext", "Enabled", "Handler", "With", "WithGroup"]
              funcs: ["Debug", "Info", "Warn", "Error", "DebugContext", "InfoContext", "WarnContext",
                      "ErrorContext", "Log", "LogContext", "Default", "SetDefault", "New", "NewJSONHandler",
                      "NewTextHandler", "With", "NewRecord", "NewLogLogger"]
```

4. Change config for your task

You can set up four options for lintering:
- english characters validation
- lowercase text format validation
- special characters validation
- sensetive data validation

You can enable and disable all of this options.

Also you can set up a loggers congig, which defines logger matching patterns (section "loggers"). 

5. Build linter
```sh
golangci-lint custom
```

6. Start checking
```
./custom-gcl run ./testdata/test2.go
```
The test data for your best expirience is provided
