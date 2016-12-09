package util

const (
	ROOT_DIR   = "basic.root_dir"
	CODE_DIR   = "basic.code_dir"

  DB_NAME = "db.name"
  DB_USER = "db.user"
  DB_HOST = "db.host"
  DB_PASS = "db.pass"

	WEB_DIR    = "web.dir"

	RELEASE_DIR    = "release.dir"
	RELEASE_SCRIPT = "release.script"


	DEPLOY_SCRIPT  = "deploy.deploy_script"
	INIT_SCRIPT    = "deploy.init_script"
	SYNC_SCRIPT    = "deploy.sync_script"

	MASTER           = "saltstack.master"
	MASTER_PORT      = "saltstack.master_port"
	MINION_USER      = "saltstack.minion_user"
	MINION_ROOT      = "saltstack.minion_root"
	MINION_DIR       = "saltstack.minion_dir"
	CONFIG_TEMPL_DIR = "saltstack.templ_dir"
)

const (
	CONFIG_INIT          = "config"
	LOG_INIT             = "log"
	DB_INIT              = "db"
	UTIL_INIT            = "util"
	MODEL_SVR_INIT       = "model.server"
	MODEL_SVR_LIB_INIT   = "models.server_lib"
	MODEL_TEMPL_INIT     = "model.template"
	MODEL_TEMPL_LIB_INIT = "model.template_lib"
	CMD_INIT             = "cmd"
)

const (
	MODE = 0755
)
