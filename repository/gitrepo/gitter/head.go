package gitter

func (g *git) GetHead() string {
	cmd := g.command("git", "rev-parse", "HEAD")

	v, err := cmd.Output()
	if err != nil {
		return ""
	}

	return string(v)
}
