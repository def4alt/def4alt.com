<script lang="ts">
	export let title: string;
</script>

<div class="timeline">
	<div class="timeline-header">
		<h2>{title}</h2>
	</div>
	<div class="timeline-items">
		<slot />
	</div>
</div>

<style>
	.timeline {
		position: relative;
		margin-bottom: 40px;
	}

	.timeline-header {
		margin-bottom: 20px;
	}

	.timeline-items {
		padding-left: 2rem;
		border-left: solid 2px var(--color-accent);
	}
</style>
