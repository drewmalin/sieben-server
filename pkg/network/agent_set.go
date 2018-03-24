package network

// AgentSet is a collection of unique AgentSets.
type AgentSet struct {
	Agents map[Agent]bool
}

// NewAgentSet instantiates a new agent set.
func NewAgentSet() *AgentSet {
	set := make(map[Agent]bool)
	return &AgentSet{
		Agents: set,
	}
}

// Add attempts to add the specified Agent to the set if it is not already present.
// Returns true if this set did not already contain the specified Agent.
func (agentSet *AgentSet) Add(agent Agent) bool {
	exists := agentSet.Contains(agent)
	if !exists {
		agentSet.Agents[agent] = true
	}
	return !exists
}

// Remove attempts to remove the specified Agent from the set if it is present.
// Returns true if this set contained the specified Agent.
func (agentSet *AgentSet) Remove(agent Agent) bool {
	exists := agentSet.Contains(agent)
	if exists {
		delete(agentSet.Agents, agent)
	}
	return exists
}

// Size returns the number of Agents in this set.
func (agentSet *AgentSet) Size() int {
	return len(agentSet.Agents)
}

// Contains returns true if this set contains the specified Agent.
func (agentSet *AgentSet) Contains(agent Agent) bool {
	_, exists := agentSet.Agents[agent]
	return exists
}
