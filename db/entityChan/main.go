package entityChan

import (
	"github.com/goatcms/goat-core/app"
	"github.com/goatcms/goat-core/db"
)

// Factory create entity intance
type Factory func() interface{}

// EntityChan is entity channel
type EntityChan chan interface{}

// ChanCorverter is channel for entities
type ChanCorverter struct {
	Rows    db.Rows
	Factory Factory
	Chan    EntityChan
	Scope   app.Scope
	inited  bool
	kill    bool
}

// NewChanCorverter create new instance of ChanCorverter
func NewChanCorverter(s app.Scope, r db.Rows, f Factory) *ChanCorverter {
	c := &ChanCorverter{
		Rows:    r,
		Factory: f,
		Scope:   s,
	}
	c.Init()
	return c
}

// Init prepare struct to run
func (c *ChanCorverter) Init() {
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
func (c *ChanCorverter) Go() {
	c.Init()
	//var entities = []*models.ArticleEntity{}
	for c.Rows.Next() && !c.kill {
		entity := c.Factory()
		if err := c.Rows.StructScan(entity); err != nil {
			c.Scope.Set(app.Error, err)
			c.Scope.Trigger(app.ErrorEvent)
			c.Scope.Trigger(app.KillEvent)
			c.close()
			return
		}
		c.Chan <- entity
	}
	c.close()
}

// Close close converter
func (c *ChanCorverter) close() error {
	if err := c.Rows.Close(); err != nil {
		return err
	}
	close(c.Chan)
	return nil
}

// Kill thread
func (c *ChanCorverter) Kill() error {
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
