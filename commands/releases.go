package commands

import (
	"github.com/Sirupsen/logrus"
	"github.com/hobeone/gonab/config"
	"github.com/hobeone/gonab/db"
	"gopkg.in/alecthomas/kingpin.v2"
)

// ReleasesCommand comment
type ReleasesCommand struct{}

func (r *ReleasesCommand) run(c *kingpin.ParseContext) error {
	if *debug {
		logrus.SetLevel(logrus.DebugLevel)
	}
	logrus.Infof("Reading config %s\n", *configfile)
	cfg := config.NewConfig()
	err := cfg.ReadConfig(*configfile)
	if err != nil {
		return err
	}

	dbh := db.NewDBHandle(cfg.DB.Path, cfg.DB.Verbose)

	err = dbh.MakeReleases()
	return err
}