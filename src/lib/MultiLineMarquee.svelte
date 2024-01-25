<script lang="ts">
	import Marqueeck, { type MarqueeckOptions } from '@arisbh/marqueeck';

	export let title: string;
	export let content: string;

	let halfLength = Math.ceil(content.length / 2);
	let firstLine = content.substring(0, halfLength);
	let secondLine = content.substring(halfLength);

	const options = {
		direction: 'left',
		paddingX: 0,
		speed: 25
	} as MarqueeckOptions;
</script>

<Marqueeck {options} hoverClasses="language-hover" stickyClasses="language-title">
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
		height: 3.5rem;
		flex-direction: column;
		user-select: none;

		& > * {
			height: 2rem;
			font-size: 1.5rem;
			color: var(--color-text-light);
			opacity: 60%;
		}
	}
</style>
