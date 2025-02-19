package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

func getEnvOrDefault(key, defVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defVal
}

func getCommandName() string {
	if strings.HasPrefix(filepath.Base(os.Args[0]), "kubectl-") {
		// cobra will split on " " and take the first element
		return "kubectl\u2002kopilot"
	}
	return "kopilot"
}

func main() {
	lang := getEnvOrDefault(envKopilotLang, langEN)
	typ := getEnvOrDefault(envKopilotType, typeChatGPT)
	opt := option{
		lang:  lang,
		typ:   typ,
		token: os.Getenv(envKopilotToken),
	}
	cmd := cobra.Command{
		Use: getCommandName(),
		Long: fmt.Sprintf(`
You need three ENVs to run Kopilot.
Set %s to specify your token.
Set %s to specify your token type, current type is: %s.
Set %s to specify the language, current language is: %s. Valid options like Chinese, French, Spain, etc.
`, envKopilotToken, envKopilotType, typ, envKopilotLang, lang),
	}
	cmd.AddCommand(newDiagnoseCommand(opt))
	cmd.AddCommand(newAuditCommand(opt))
	_ = cmd.Execute()
}
