<script>
	import { page } from '$app/stores';
</script>

<header>
	<nav>
		<ul>
			<li aria-current={$page.url.pathname === '/posts' ? 'page' : undefined}>
				<a href="/posts">posts</a>
			</li>
			<li aria-current={$page.url.pathname === '/projects' ? 'page' : undefined}>
				<a href="/projects">projects</a>
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
	}

	.brand {
		font-variant-numeric: proportional-nums;
		font-size: 4rem;
		color: var(--color-background);
		font-variant-alternates: styleset(alternate);
	}

	ul {
		display: flex;
		flex-direction: row;
		padding: 0;
		gap: 30px;
		& li {
			list-style-type: none;
			& a {
				color: var(--color-background);
				font-variant-alternates: styleset(alternate);
				text-transform: lowercase;
				font-size: 2rem;
				font-weight: 400;
			}
		}
	}
</style>
