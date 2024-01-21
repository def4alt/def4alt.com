<script lang="ts">
	export let text: string;
	export let startYear: number | undefined = undefined;
	export let endYear: number | undefined = undefined;
	export let inProgress = false;
</script>

<div class="timeline-item">
	<div class="timeline-item-text">
		{text}
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
	}
</style>
