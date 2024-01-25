<script>
	import andrii from '$lib/images/andrii.webp';
	import LinkedIn from '../icons/linkedin.svelte';
	import Instagram from '../icons/instagram.svelte';
	import Github from '../icons/github.svelte';
	import Timeline from '$lib/Timeline.svelte';
	import TimelineItem from '$lib/TimelineItem.svelte';
	import MultiLineMarquee from '$lib/MultiLineMarquee.svelte';
</script>

<svelte:head>
	<title>Home</title>
	<meta name="description" content="def4alt's blog" />
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
		<a href="https://linkedin.com/in/def4alt"><LinkedIn /></a>
		<a href="https://instagram.com/def4alt"><Instagram /></a>
		<a href="https://github.com/def4alt"><Github /></a>
	</div>
</div>
<Timeline title="My alma mater">
	<TimelineItem
		content="ukrainian physics and mathematics lyceum"
		link="http://upml.knu.ua/about-upml/"
		startYear={2018}
		endYear={2022}
	/>
	<TimelineItem
		content="taras shevchenko national university of kyiv"
		link="https://knu.ua/en"
		startYear={2022}
		endYear={2023}
	/>
	<TimelineItem
		content="studienkolleg münchen"
		link="https://xn--studienkolleg-mnchen-3ec.de/"
		startYear={2023}
		inProgress={true}
	/>
</Timeline>
<div class="work-hobbies-section">
	<div>
		<Timeline title="I worked at">
			<TimelineItem content="home 🏠" startYear={2004} inProgress={true} />
			<TimelineItem
				content="bamberger haus 🍻"
				startYear={2023}
				link="https://www.bambergerhaus.com/"
			/>
			<TimelineItem content="steinheil 16 💁" startYear={2023} link="https://www.steinheil16.de" />
		</Timeline>
	</div>

	<div>
		<Timeline title="My hobbies">
			<TimelineItem content="traveling ✈️️" />
			<TimelineItem content="swimming 🏊" />
			<TimelineItem content="playing guitar 🎸" />
		</Timeline>
	</div>
</div>

<h1>I speak</h1>
<div class="languages">
	<MultiLineMarquee
		title="ukrainian"
		content="Живеш на сьомому поверсі, стіни кімнати завісив козаками й діячами ЗУНР, з вікна бачиш московські дахи, безрадісні тополині алеї, Останкінської телевежі не бачиш — її видно з кімнат, розташованих по другий бік коридору, — але близька її присутність відчувається щохвилини; випромінює щось дуже снодійне, віруси млявості та апатії, тому вранці ніяк не можеш прокинутися, переходиш з одного сновидіння в інше, ніби з країни в країну. Спиш самовіддано, найчастіше до одинадцятої, поки узбек за стіною не увімкне на повний регулятор духмяну орієнтальну музику 'адин палка два струна'. Проклинаючи незлостиво нещасну нашу історію, дружбу народів і союзний договір 22-го року, усвідомлюєш: далі спати не можна."
	/>
	<MultiLineMarquee
		title="english"
		content="It was a bright cold day in April, and the clocks were striking thirteen. Winston Smith, his chin nuzzled into his breast in an effort to escape the vile wind, slipped quickly through the glass doors of Victory Mansions, though not quickly enough to prevent a swirl of gritty dust from entering along with him"
	/>
	<MultiLineMarquee
		title="german"
		content="Ihr naht euch wieder, schwankende Gestalten! Die früh sich einst dem trüben Blick gezeigt. Versuch ich wohl euch diesmal fest zu halten? Fühl ich mein Herz noch jenem Wahn geneigt? Ihr drängt euch zu! nun gut, so mögt ihr walten, 5 Wie ihr aus Dunst und Nebel um mich steigt; Mein Busen fühlt sich jugendlich erschüttert Vom Zauberhauch, der euren Zug umwittert."
	/>
	<MultiLineMarquee
		title="spanish"
		content="Paracéme, Sancho, que no hay refrán que no sea verdadero, porque todos son sentencias sacadas de la misma experiencia, madre de las ciencias todas."
	/>
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
			top: 50%;
			left: 50%;
			width: 100%;
			height: 100%;
			object-fit: cover;
			overflow: clip;
			overflow-clip-margin: content-box;
			display: inline-block;
			vertical-align: baseline;
			transform: translate(-50%, -50%);
			-webkit-transform: translate(-50%, -50%);
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

			& > *:hover {
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

	.languages {
		display: grid;
		gap: 30px;
		grid-template-columns: auto auto;
		grid-template-rows: auto auto;

		& > * {
			border-radius: 15px;
		}
	}
</style>
