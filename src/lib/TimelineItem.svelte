<script lang="ts">
	import ArrowUp from '../icons/arrowup.svelte';

	export let content: string;
	export let link: string | undefined = undefined;
	export let startYear: number | undefined = undefined;
	export let endYear: number | undefined = undefined;
	export let inProgress = false;
</script>

<div class="timeline-item">
	<div class="timeline-item-text">
		{#if link}
			<a href={link}>
				{content}
				<span class="timeline-item-link-arrow"><ArrowUp /></span>
			</a>
		{:else}
			{content}
		{/if}
		{#if startYear}
			<span class="year-range">
				{#if startYear && endYear}
					{startYear}-{endYear}
				{:else if startYear && inProgress}
					{startYear}-
				{:else if startYear}
					{startYear}
				{/if}
			</span>
		{/if}
	</div>
</div>

<style>
	.timeline-item {
		margin-bottom: 20px;
	}

	.year-range {
		font-size: 12px;
		color: var(--color-secondary-dark);
	}

	.timeline-item-text {
		font-size: 1.5rem;

		font-weight: 300;

		& a {
			color: inherit;
		}

		& .timeline-item-link-arrow > * {
			width: 1.5rem;
			height: 1.5rem;
			stroke: var(--color-text);
		}
	}
</style>
