package global

import (
	"github.com/panda8z/istone/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	// ISA_CONFIG: iStone-Admin-config
	ISA_CONFIG config.Server
	// ISA_DB: iStone-Admin-database
	ISA_DB *gorm.DB
	// ISA_VP: iStone-Admin-viper
	ISA_VP *viper.Viper
	// ISA_LOG: iStone-Admin-zap
	ISA_LOG *zap.Logger
)
