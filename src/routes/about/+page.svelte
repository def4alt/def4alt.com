<script>
	import andrii from '$lib/images/andrii.png';
	import linkedin from '$lib/images/linkedin.png';
	import Instagram from '../../icons/instagram.svelte';
	import Github from '../../icons/github.svelte';
	import Timeline from '$lib/Timeline.svelte';
	import TimelineItem from '$lib/TimelineItem.svelte';
</script>

<svelte:head>
	<title>About</title>
	<meta name="description" content="About me" />
</svelte:head>

<div class="video-background">
	<video src="https://static.def4alt.com/mountain_range.webm" loop muted autoplay playsinline />
	<div class="video-overlay" />
</div>
<div class="profile">
	<div class="profile-cover">
		<p class="profile-name">Andrii Olkhovych</p>
		<a
			class="transcription-link"
			href="http://ipa-reader.xyz/?text=%C9%91n%CB%88drij%20%C9%94l%CB%88x%C9%94v%C9%AAt%CA%83&voice=Ewa"
			>/andrij olhovich/</a
		>
		<img src={andrii} alt="Andrii Olkhovych" />
	</div>
	<div class="profile-bottom-line" />
</div>
<div class="personal-info">
	<p>programmer, student</p>
	<div class="social-links">
		<a href="https://linkedin.com/in/def4alt"><img src={linkedin} alt="linkedin" /></a>
		<a href="https://instagram.com/def4alt"><Instagram /></a>
		<a href="https://github.com/def4alt"><Github /></a>
	</div>
</div>
<Timeline title="education">
	<TimelineItem text="upml" startYear={2018} endYear={2022} />
	<TimelineItem
		text="taras shevchenko national university of kyiv"
		startYear={2022}
		endYear={2023}
	/>
	<TimelineItem text="studienkolleg münchen" startYear={2023} inProgress={true} />
</Timeline>
<div class="work-hobbies-section">
	<div>
		<Timeline title="workplaces">
			<TimelineItem text="home 🏠" startYear={2004} inProgress={true} />
			<TimelineItem text="bamberger haus 🍻" startYear={2023} />
			<TimelineItem text="steinheil 16 💁" startYear={2023} inProgress={true} />
		</Timeline>
	</div>

	<div>
		<Timeline title="hobbies">
			<TimelineItem text="traveling ✈️️" />
			<TimelineItem text="swimming 🏊" />
			<TimelineItem text="playing guitar 🎸" />
		</Timeline>
	</div>
</div>

<style>
	.video-background {
		position: absolute;
		top: 0;
		left: 0;
		width: 100vw;
		height: 100vh;
		z-index: -1;
		overflow: hidden;
		background: var(--color-text);

		& video {
			position: absolute;
			z-index: -2;
			top: 0;
			left: 0;
			width: 100%;
			height: 100%;
			object-fit: cover;
		}

		& .video-overlay {
			height: 30vh;
			width: 100vw;
			position: absolute;
			top: 0;
			left: 0;
			background: linear-gradient(
				180deg,
				rgba(0, 0, 0, 0.86) 0%,
				rgba(0, 0, 0, 0.65) 35%,
				rgba(0, 0, 0, 0) 100%
			);
		}
	}

	.profile {
		top: 11rem;
		left: 0;
		padding: 0;
		width: 100%;
		height: calc(100vh - 10rem - 1rem);
		display: flex;
		flex-direction: column;
		justify-content: end;
		align-items: center;
		overflow-x: hidden;

		& .profile-cover {
			width: 100%;
			margin-bottom: 5rem;

			& .profile-name {
				font-size: 8rem;
				font-weight: 700;
				width: 40rem;
				margin: 0;
				line-height: 1.1;
				font-variant-alternates: styleset(alternate);
				color: var(--color-background);
			}

			& .transcription-link {
				font-size: 3rem;
				margin: 0;
				padding: 0;
				color: var(--color-secondary-light);
			}

			& img {
				position: absolute;
				bottom: 6rem;
				object-fit: contain;
				width: 40rem;
				right: auto;
				border-radius: 15px;
				z-index: -1;
			}
		}

		& .profile-bottom-line {
			width: 5rem;
			height: 0px;
			border: solid;
			border-width: 3px;
			border-color: var(--color-primary);
			border-radius: 15px;
			margin-bottom: 10px;
		}
	}

	.personal-info {
		display: flex;
		flex-direction: row;
		justify-content: space-between;

		& p {
			font-size: 24px;
		}

		& .social-links {
			display: flex;
			flex-direction: row;
			align-items: center;
			justify-content: space-between;
			width: 20rem;

			& * {
				width: 2.5rem;
				height: 2.5rem;
				opacity: 90%;
				transition: transform 0.1s ease-in-out;
			}

			& img {
				filter: grayscale(100%) contrast(100%) brightness(0);
			}

			& *:hover {
				opacity: 1;
				transform: scale(1.1);
			}
		}
	}

	.work-hobbies-section {
		display: flex;
		flex-direction: row;

		& > * {
			width: 50%;
		}
	}
</style>
