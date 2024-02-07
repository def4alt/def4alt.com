<script lang="ts">
	export let title: string;
</script>

<div class="timeline">
	<h2>{title}</h2>
	<div class="timeline-items">
		<slot />
	</div>
</div>

<style>
	.timeline {
		position: relative;
	}

	.timeline-items {
		padding-left: 2rem;
		border-left: solid 2px var(--color-accent);
	}
</style>
