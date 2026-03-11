import { gql } from 'urql';
import { client } from './graphqlClient';

export interface AgentAction {
	label: string;
	prompt: string;
}

const AGENT_ACTIONS_QUERY = gql`
	query AgentActions {
		agentActions {
			label
			prompt
		}
	}
`;

class AgentActionsStore {
	actions = $state<AgentAction[]>([]);

	async fetch(): Promise<void> {
		const result = await client.query(AGENT_ACTIONS_QUERY, {}).toPromise();

		if (result.error) {
			console.error('Failed to fetch agent actions:', result.error);
			return;
		}

		this.actions = result.data?.agentActions ?? [];
	}
}

export const agentActionsStore = new AgentActionsStore();
