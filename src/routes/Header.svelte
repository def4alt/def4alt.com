<script>
	import { page } from '$app/stores';
</script>

<header>
	<nav>
		<ul>
			<li aria-current={$page.url.pathname === '/posts' ? 'page' : undefined}>
				<a href="/">posts</a>
			</li>
			<li aria-current={$page.url.pathname === '/projects' ? 'page' : undefined}>
				<a href="/">projects</a>
			</li>
		</ul>
		<div class="logo">J</div>
		<a class="brand" href="/">def4alt</a>
	</nav>
</header>

<style>
	nav {
		height: 150px;
		width: 100vw;
		display: flex;
		flex-direction: row;
		align-items: center;
		justify-content: space-around;
		font-size: calc(1rem + 2vw);
		line-height: calc(1.4rem + 2.4vw);
	}

	.logo {
		font-variant-alternates: styleset(symbols);
		user-select: none;
		color: var(--color-background);
		font-size: 8rem;

		@media screen and (orientation: portrait) {
			display: none;
		}
	}

	.brand {
		font-variant-numeric: proportional-nums;
		font-size: 2em;
		line-height: calc(1.4rem + 4.8vw);
		color: var(--color-background);
		font-variant-alternates: styleset(alternate);
	}

	ul {
		display: flex;
		flex-direction: row;
		padding: 0;
		margin: 0;
		gap: 30px;

		@media screen and (orientation: portrait) {
			flex-direction: column;
			gap: 15px;
			margin: 0;
		}

		& li {
			list-style-type: none;
			& a {
				color: var(--color-background);
				font-variant-alternates: styleset(alternate);
				text-transform: lowercase;
				font-weight: 400;
			}
		}
	}
</style>
