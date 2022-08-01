package discord

import "github.com/bwmarrin/discordgo"

type Router struct {
	actions  []Handler
	branches []*Branch
}

type Branch struct {
	requirement BranchRequirement
	subRouter   *Router
}

func NewRouter() *Router {
	return &Router{
		actions:  []Handler{},
		branches: []*Branch{},
	}
}

func NewBranch(requirement BranchRequirement, subRouter *Router) *Branch {
	return &Branch{
		requirement: requirement,
		subRouter:   subRouter,
	}
}

func (r *Router) Serve(session *discordgo.Session, message any) {
	for _, action := range r.actions {
		action(session, message)
	}
	for _, branch := range r.branches {
		if branch.requirement(session, message) {
			branch.subRouter.Serve(session, message)
			break
		}
	}
}

func (r *Router) AddAction(handler Handler) *Router {
	r.actions = append(r.actions, handler)
	return r
}

func (r *Router) AddBranch(branch *Branch) *Router {
	r.branches = append(r.branches, branch)
	return r
}
