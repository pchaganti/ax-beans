<script lang="ts">
  import { gql } from 'urql';
  import { onMount, onDestroy } from 'svelte';
  import { changesStore } from '$lib/changes.svelte';
  import { ui } from '$lib/uiState.svelte';
  import { client } from '$lib/graphqlClient';
  import PaneHeader from '$lib/components/PaneHeader.svelte';

  interface AgentAction {
    id: string;
    label: string;
    description: string | null;
  }

  interface Props {
    path?: string;
    beanId?: string;
    agentBusy?: boolean;
  }

  let { path, beanId, agentBusy = false }: Props = $props();

  const AGENT_ACTIONS_QUERY = gql`
    query AgentActions($beanId: ID!) {
      agentActions(beanId: $beanId) {
        id
        label
        description
      }
    }
  `;

  const EXECUTE_AGENT_ACTION = gql`
    mutation ExecuteAgentAction($beanId: ID!, $actionId: ID!) {
      executeAgentAction(beanId: $beanId, actionId: $actionId)
    }
  `;

  let actions = $state<AgentAction[]>([]);
  let executingAction = $state<string | null>(null);

  async function fetchActions() {
    if (!beanId) return;
    const result = await client.query(AGENT_ACTIONS_QUERY, { beanId }).toPromise();
    if (result.error) {
      console.error('Failed to fetch agent actions:', result.error);
      return;
    }
    if (result.data?.agentActions) {
      actions = result.data.agentActions;
    }
  }

  async function executeAction(actionId: string) {
    if (!beanId || agentBusy) return;
    executingAction = actionId;
    try {
      await client.mutation(EXECUTE_AGENT_ACTION, { beanId, actionId }).toPromise();
    } finally {
      executingAction = null;
    }
  }

  // Re-fetch actions when beanId changes
  $effect(() => {
    if (beanId) {
      fetchActions();
    }
  });

  // Re-fetch actions when agent transitions to idle
  let wasAgentBusy = $state(false);
  $effect(() => {
    if (wasAgentBusy && !agentBusy) {
      fetchActions();
    }
    wasAgentBusy = agentBusy;
  });

  const stagedChanges = $derived(changesStore.changes.filter((c) => c.staged));
  const unstagedChanges = $derived(changesStore.changes.filter((c) => !c.staged));
  const totalCount = $derived(changesStore.changes.length);

  onMount(() => {
    changesStore.startPolling(path);
  });

  onDestroy(() => {
    changesStore.stopPolling();
  });

  function statusColor(status: string): string {
    switch (status) {
      case 'added':
      case 'untracked':
        return 'text-success';
      case 'deleted':
        return 'text-danger';
      case 'renamed':
        return 'text-accent';
      default:
        return 'text-warning';
    }
  }

  function statusLabel(status: string): string {
    switch (status) {
      case 'modified':
        return 'M';
      case 'added':
        return 'A';
      case 'deleted':
        return 'D';
      case 'untracked':
        return '?';
      case 'renamed':
        return 'R';
      default:
        return '?';
    }
  }

  function fileName(path: string): string {
    return path.split('/').pop() ?? path;
  }

  function dirName(path: string): string {
    const parts = path.split('/');
    if (parts.length <= 1) return '';
    return parts.slice(0, -1).join('/') + '/';
  }
</script>

<div class="flex h-full flex-col border-l border-border bg-surface">
  <PaneHeader title="Status" onClose={() => ui.toggleChanges()}>
    {#snippet extra()}
      {#if totalCount > 0}
        <span class="ml-1 text-sm text-text-muted">({totalCount})</span>
      {/if}
    {/snippet}
  </PaneHeader>

  <div class="flex-1 overflow-auto">
    {#if totalCount === 0}
      <p class="px-3 py-4 text-center text-text-muted">No changes</p>
    {:else}
      {#if stagedChanges.length > 0}
        <div class="px-3 pt-2 pb-1 font-medium text-text-muted">Staged</div>
        {#each stagedChanges as change (change.path + ':staged')}
          <div class="flex items-center gap-1.5 px-3 py-0.5 hover:bg-surface-alt">
            <span class={['w-3 shrink-0 text-center font-mono font-bold', statusColor(change.status)]}>
              {statusLabel(change.status)}
            </span>
            <span class="min-w-0 flex-1 truncate" title={change.path}>
              <span class="text-text-muted">{dirName(change.path)}</span><span class="text-text">{fileName(change.path)}</span>
            </span>
            {#if change.additions > 0 || change.deletions > 0}
              <span class="shrink-0 font-mono">
                {#if change.additions > 0}<span class="text-success">+{change.additions}</span>{/if}
                {#if change.deletions > 0}<span class={[change.additions > 0 && 'ml-1', 'text-danger']}>-{change.deletions}</span>{/if}
              </span>
            {/if}
          </div>
        {/each}
      {/if}

      {#if unstagedChanges.length > 0}
        {#if stagedChanges.length > 0}
          <div class="px-3 pt-2 pb-1 font-medium text-text-muted">Unstaged</div>
        {/if}
        {#each unstagedChanges as change (change.path + ':unstaged')}
          <div class="flex items-center gap-1.5 px-3 py-0.5 hover:bg-surface-alt">
            <span class={['w-3 shrink-0 text-center font-mono font-bold', statusColor(change.status)]}>
              {statusLabel(change.status)}
            </span>
            <span class="min-w-0 flex-1 truncate" title={change.path}>
              <span class="text-text-muted">{dirName(change.path)}</span><span class="text-text">{fileName(change.path)}</span>
            </span>
            {#if change.additions > 0 || change.deletions > 0}
              <span class="shrink-0 font-mono">
                {#if change.additions > 0}<span class="text-success">+{change.additions}</span>{/if}
                {#if change.deletions > 0}<span class={[change.additions > 0 && 'ml-1', 'text-danger']}>-{change.deletions}</span>{/if}
              </span>
            {/if}
          </div>
        {/each}
      {/if}
    {/if}
  </div>

  {#if beanId && actions.length > 0}
    <div class="flex gap-2 border-t border-border px-3 py-2">
      {#each actions as action (action.id)}
        <button
          class={[
            'flex-1 rounded border border-border px-3 py-1.5 text-sm font-medium transition-colors',
            agentBusy || executingAction
              ? 'cursor-not-allowed text-text-faint'
              : 'cursor-pointer text-text-muted hover:bg-surface-alt hover:text-text'
          ]}
          disabled={agentBusy || !!executingAction}
          title={action.description ?? undefined}
          onclick={() => executeAction(action.id)}
        >
          {action.label}
        </button>
      {/each}
    </div>
  {/if}
</div>
