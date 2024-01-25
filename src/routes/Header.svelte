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
		height: 10rem;
		width: 100vw;
		display: flex;
		flex-direction: row;
		align-items: center;
		justify-content: space-around;
	}

	.logo {
		font-variant-alternates: styleset(symbols);
		user-select: none;
		color: var(--color-background);
		font-size: 8rem;

		@media screen and (max-width: 768px) {
			display: none;
		}
	}

	.brand {
		font-variant-numeric: proportional-nums;
		font-size: 4rem;
		color: var(--color-background);
		font-variant-alternates: styleset(alternate);

		@media screen and (max-width: 768px) {
			font-size: 3rem;
		}
	}

	ul {
		display: flex;
		flex-direction: row;
		padding: 0;
		gap: 30px;

		@media screen and (max-width: 768px) {
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
				font-size: 2rem;
				font-weight: 400;

				@media screen and (max-width: 768px) {
					font-size: 1.5rem;
				}
			}
		}
	}
</style>
