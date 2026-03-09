<script lang="ts">
	import type { Bean } from '$lib/beans.svelte';
	import { beansStore } from '$lib/beans.svelte';
	import { backlogDrag } from '$lib/backlogDrag.svelte';
	import BeanCard from './BeanCard.svelte';
	import BeanItem from './BeanItem.svelte';

	interface Props {
		bean: Bean;
		/** Parent ID of this bean's sibling group (null = top-level) */
		parentId?: string | null;
		index?: number;
		depth?: number;
		selectedId?: string | null;
		onSelect?: (bean: Bean) => void;
	}

	let { bean, parentId = null, index = 0, depth = 0, selectedId = null, onSelect }: Props =
		$props();

	const children = $derived(beansStore.children(bean.id));

	function handleClick(e: MouseEvent) {
		e.stopPropagation();
		onSelect?.(bean);
	}
</script>

<div class="bean-item" data-bean-id={bean.id}>
	<!-- Drop indicator before this card -->
	<div
		class={[
			'mx-1 h-0.5 rounded-full transition-colors',
			backlogDrag.showIndicator(parentId, index, bean.id) ? 'bg-accent' : 'bg-transparent'
		]}
	></div>

	<!-- svelte-ignore a11y_click_events_have_key_events -->
	<!-- svelte-ignore a11y_no_static_element_interactions -->
	<div
		class={[
			'rounded transition-all',
			backlogDrag.draggedBeanId === bean.id && 'opacity-40',
			backlogDrag.isReparentTarget(bean.id) && 'ring-2 ring-accent ring-offset-1'
		]}
		draggable="true"
		ondragstart={(e) => backlogDrag.startDrag(e, bean)}
		ondragend={() => backlogDrag.endDrag()}
		ondragover={(e) => backlogDrag.hoverCard(e, parentId, index, bean.id)}
		onclick={handleClick}
	>
		<BeanCard
			{bean}
			variant="list"
			selected={selectedId === bean.id}
			onclick={() => onSelect?.(bean)}
		/>
	</div>

	{#if children.length > 0}
		<div
			class="mt-1 ml-4 border-l border-border pl-2"
			ondragover={(e) => backlogDrag.hoverList(e, bean.id, children.length)}
			ondragleave={(e) => backlogDrag.leaveList(e, e.currentTarget, bean.id)}
			ondrop={(e) => backlogDrag.drop(e, bean.id, children)}
			role="list"
		>
			{#each children as child, i (child.id)}
				<BeanItem
					bean={child}
					parentId={bean.id}
					index={i}
					depth={depth + 1}
					{selectedId}
					{onSelect}
				/>
			{/each}

			<!-- Drop indicator at end of children -->
			<div
				class={[
					'mx-1 h-0.5 rounded-full transition-colors',
					backlogDrag.showEndIndicator(bean.id, children.length)
						? 'bg-accent'
						: 'bg-transparent'
				]}
			></div>
		</div>
	{/if}
</div>
