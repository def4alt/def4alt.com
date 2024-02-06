<script>
	import Header from './Header.svelte';
	import './styles.css';
</script>

<div class="app">
	<Header />

	<main>
		<slot />
	</main>

	<footer>©2024 Andrii Olkhovych</footer>
</div>

<style>
	.app {
		display: flex;
		flex-direction: column;
		min-height: 100vh;
	}

	main {
		flex: 1;
		display: flex;
		flex-direction: column;
		width: 100%;
		margin: 0;
		box-sizing: border-box;
	}

	footer {
		display: flex;
		line-height: 1.5;
		font-size: 1.5rem;
		color: var(--color-text-light);
		opacity: 20%;
		margin-right: 10px;
		flex-direction: column;
		align-items: end;
		font-weight: bold;
	}
</style>
