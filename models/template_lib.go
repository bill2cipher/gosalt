package models

import (
	log "github.com/Sirupsen/logrus"
	"github.com/jellybean4/gosalt/util"
  "github.com/spf13/viper"
)

import (
	"os"
	"path"
  "fmt"
  "text/template"
)

const (
	CONFIG_TEMPLATE = `
{{ range $k, $v := .Config }}
{{printf "%s = salt['pillar.get']('%s', '%s')" $k $k $v }}
{{ end }}
`
)

var (
	configParser *template.Template
)

func initTemplLib() {
	if templ, err := template.New("config").Parse(CONFIG_TEMPLATE); err != nil {
		log.WithFields(log.Fields{
			"template": "config",
			"reason":   err.Error(),
		}).Fatal(util.TEMPL_PARSE_LOG)
	} else {
		configParser = templ
	}
}

/**
 * SetTemplate edit exits template or add new template into gosalt
 */
func SetTemplate(templ *Template) error {
  if err := TemplateCache.Set(templ.Name, templ); err != nil {
    return err
  } else if err := SyncTemplFile(templ); err != nil {
    return err
  } else {
    log.WithFields(log.Fields{
      "name": templ.Name,
    }).Info("settemplate success")
    return nil
  }
}

/**
 * SyncTemplFile sync template data into config template file
 */
func SyncTemplFile(templ *Template) error {
	templDir := getTemplDir(templ)
	writer, err := util.OpenFile(templDir, templ.Name, true)
	if err != nil {
		return err
	}

	if err := configParser.Execute(writer, templ.Config); err != nil {
		log.WithFields(log.Fields{
			"args":   templ.Config,
			"reason": err.Error(),
		}).Error(util.TEMPL_EXEC_LOG)
		return err
	} else if err := writer.Close(); err != nil {
		log.WithFields(log.Fields{
			"file":   path.Join(templDir, templ.Name),
			"reason": err.Error(),
		}).Error(util.FILE_CLOSE_LOG)
		return err
	} else {
		log.WithFields(log.Fields{
			"file": path.Join(templDir, templ.Name),
			"args": templ.Config,
		}).Info("synctemplfile success")
		return nil
	}
}

/**
 * DelTemplFile delete the corresponding config template file
 */
func DelTemplFile(name string) error {
	filename := viper.GetString(util.CONFIG_TEMPL_DIR) + "/" + name
	if err := os.Remove(filename); err != nil {
		log.WithFields(log.Fields{
			"file":   filename,
			"reason": err.Error(),
		}).Error(util.FILE_DELETE_LOG)
		return err
	} else {
		return nil
	}
}


func getTemplDir(templ *Template) string {
  return fmt.Sprintf("%s/%s/%s", util.GetConfig(util.CONFIG_TEMPL_DIR), templ.Env, templ.Version)
}
