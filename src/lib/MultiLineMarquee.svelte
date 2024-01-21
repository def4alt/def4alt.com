<script lang="ts">
	import Marqueeck, { type MarqueeckOptions, type MarqueeckHoverEvent } from '@arisbh/marqueeck';

	export let title: string;
	export let content: string;

	let hovered = false;

	let halfLength = Math.ceil(content.length / 2);
	let firstLine = content.substring(0, halfLength);
	let secondLine = content.substring(halfLength);

	const handleHover = (e: MarqueeckHoverEvent) => {
		hovered = e.detail.valueOf();
	};

	const options = {
		direction: 'left',
		paddingX: 0
	} as MarqueeckOptions;
</script>

<Marqueeck
	{options}
	hoverClasses="language-hover"
	stickyClasses="language-title"
	on:hover={handleHover}
>
	<div class="text-container">
		<span>{firstLine}</span>
		<span>{secondLine}</span>
	</div>
	<svelte:fragment slot="stickyStart">{title}</svelte:fragment>
</Marqueeck>

<style>
	.text-container {
		display: flex;
		width: 100%;
		height: 5rem;
		flex-direction: column;

		& > * {
			height: 2.5rem;
			font-size: 2rem;
			color: var(--color-text-light);
			opacity: 70%;
		}
	}
</style>
