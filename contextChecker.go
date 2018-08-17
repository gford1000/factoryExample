package factory

// Hidden type
type contextChecker struct {
	ctx     string
	onChg   func()
	factory FactoryContextChecker
}

func (c *contextChecker) Check() {
	if ctxNew := c.factory.GetExecutionContextHash(); ctxNew != c.ctx {
		c.ctx = ctxNew
		if c.onChg != nil {
			c.onChg()
		}
	}
}
