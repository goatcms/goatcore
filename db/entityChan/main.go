package entitychan

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/db"
)

// Factory create entity intance
type Factory func() interface{}

// EntityChan is entity channel
type EntityChan chan interface{}

// ChanConverter is channel for entities
type ChanConverter struct {
	Rows    db.Rows
	Factory Factory
	Chan    EntityChan
	Scope   app.Scope
	inited  bool
	kill    bool
}

// NewChanConverter create new instance of ChanConverter
func NewChanConverter(s app.Scope, r db.Rows, f Factory) *ChanConverter {
	c := &ChanConverter{
		Rows:    r,
		Factory: f,
		Scope:   s,
	}
	c.Init()
	return c
}

// Init prepare struct to run
func (c *ChanConverter) Init() {
	if c.inited {
		return
	}
	if c.Scope != nil {
		c.Scope.On(app.KillEvent, c.Kill)
	}
	if c.Chan == nil {
		c.Chan = make(EntityChan, 30)
	}
}

// Go convert entities and add to channel
func (c *ChanConverter) Go() {
	c.Init()
	//var entities = []*models.ArticleEntity{}
	for c.Rows.Next() && !c.kill {
		entity := c.Factory()
		if err := c.Rows.StructScan(entity); err != nil {
			c.close()
			c.Scope.Set(app.Error, err)
			c.Scope.Trigger(app.ErrorEvent, err)
			c.Scope.Trigger(app.KillEvent, err)
			return
		}
		c.Chan <- entity
	}
	c.close()
}

// Close close converter
func (c *ChanConverter) close() error {
	c.kill = true
	close(c.Chan)
	if err := c.Rows.Close(); err != nil {
		return err
	}
	return nil
}

// Kill thread
func (c *ChanConverter) Kill(interface{}) error {
	// select & case is fix to get element without deadlock
	select {
	case _, ok := <-c.Chan:
		if ok {
			c.kill = true
		} else {
			c.kill = true
		}
	default:
		c.kill = true
	}
	return nil
}
